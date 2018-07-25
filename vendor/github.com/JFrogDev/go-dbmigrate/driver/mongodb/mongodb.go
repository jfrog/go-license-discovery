package mongodb

import (
	"errors"
	"fmt"
	"github.com/JFrogDev/go-dbmigrate/driver"
	"github.com/JFrogDev/go-dbmigrate/driver/mongodb/gomethods"
	"github.com/JFrogDev/go-dbmigrate/file"
	"github.com/JFrogDev/go-dbmigrate/migrate/direction"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
)

type UnregisteredMethodsReceiverError string

func (e UnregisteredMethodsReceiverError) Error() string {
	return "Unregistered methods receiver for driver: " + string(e)
}

type WrongMethodsReceiverTypeError string

func (e WrongMethodsReceiverTypeError) Error() string {
	return "Wrong methods receiver type for driver: " + string(e)
}

const MIGRATE_C = "db_migrations"
const DRIVER_NAME = "gomethods.mongodb"

type Driver struct {
	Session *mgo.Session

	methodsReceiver MethodsReceiver
	migrator        gomethods.Migrator
	url             string
}

var _ gomethods.GoMethodsDriver = (*Driver)(nil)

type MethodsReceiver interface {
	DbName() string
}

func (d *Driver) MethodsReceiver() interface{} {
	return d.methodsReceiver
}

func (d *Driver) SetMethodsReceiver(r interface{}) error {
	r1, ok := r.(MethodsReceiver)
	if !ok {
		return WrongMethodsReceiverTypeError(DRIVER_NAME)
	}

	d.methodsReceiver = r1
	return nil
}

func init() {
	driver.RegisterDriver("mongodb", &Driver{})
}

type DbMigration struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Version uint64        `bson:"version"`
}

func (driver *Driver) Initialize(url string) error {
	if driver.methodsReceiver == nil {
		return UnregisteredMethodsReceiverError(DRIVER_NAME)
	}

	urlWithoutScheme := strings.SplitN(url, "mongodb://", 2)
	if len(urlWithoutScheme) != 2 {
		return errors.New("invalid mongodb:// scheme")
	}

	driver.url = url
	if err := driver.reconnectToMasterSession(); err != nil {
		return fmt.Errorf("failed to connect to session: %v", err)
	}
	driver.migrator = gomethods.Migrator{MethodInvoker: driver}

	return nil
}

func (driver *Driver) reconnectToMasterSession() error {
	session, err := mgo.Dial(driver.url)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	if driver.Session != nil {
		driver.Session.Close()
	}
	driver.Session = session
	return nil
}

func (driver *Driver) ensureSessionNotClosed() (retErr error) {
	defer func() {
		if r := recover(); r != nil {
			if err := driver.reconnectToMasterSession(); err != nil {
				retErr = fmt.Errorf("recovering from error: '%v'. failed to re-connect to master session: %v", r, err)
				return
			}
		}
	}()

	//Ping panics if session is closed
	if pingErr := driver.Session.Ping(); pingErr != nil {
		if reconnectErr := driver.reconnectToMasterSession(); reconnectErr != nil {
			retErr = fmt.Errorf("Session ping has failed: %v. Failed to re-connect to master session: %v", pingErr, reconnectErr)
			return
		}
	}
	return nil
}

func (driver *Driver) getNewSession() (*mgo.Session, error) {
	if err := driver.ensureSessionNotClosed(); err != nil {
		return nil, fmt.Errorf("failed to ensure master session is not closed: %v", err)
	}
	session := driver.Session.Clone()
	return session, nil
}

func (driver *Driver) Close() error {
	if driver.Session != nil {
		driver.Session.Close()
	}
	return nil
}

func (driver *Driver) FilenameExtension() string {
	return "mgo"
}

func (driver *Driver) Version() (uint64, error) {
	var latestMigration DbMigration

	session, err := driver.getNewSession()
	if err != nil {
		return 0, fmt.Errorf("failed to get new session: %v", err)
	}
	defer session.Close()
	c := session.DB(driver.methodsReceiver.DbName()).C(MIGRATE_C)

	err = c.Find(bson.M{}).Sort("-version").One(&latestMigration)
	switch {
	case err == mgo.ErrNotFound:
		return 0, nil
	case err != nil:
		return 0, err
	default:
		return latestMigration.Version, nil
	}
}
func (driver *Driver) Migrate(f file.File, pipe chan interface{}) {
	defer close(pipe)
	pipe <- f

	err := driver.migrator.Migrate(f, pipe)
	if err != nil {
		return
	}

	session, err := driver.getNewSession()
	if err != nil {
		pipe <- fmt.Errorf("Migrate failed to get new session: %v", err)
		return
	}
	defer session.Close()
	migrate_c := session.DB(driver.methodsReceiver.DbName()).C(MIGRATE_C)

	if f.Direction == direction.Up {
		id := bson.NewObjectId()
		dbMigration := DbMigration{Id: id, Version: f.Version}

		err := migrate_c.Insert(dbMigration)
		if err != nil {
			pipe <- err
			return
		}

	} else if f.Direction == direction.Down {
		err := migrate_c.Remove(bson.M{"version": f.Version})
		if err != nil {
			pipe <- err
			return
		}
	}
}

func (driver *Driver) Validate(methodName string) error {
	methodWithReceiver, ok := reflect.TypeOf(driver.methodsReceiver).MethodByName(methodName)
	if !ok {
		return gomethods.MissingMethodError(methodName)
	}
	if methodWithReceiver.PkgPath != "" {
		return gomethods.MethodNotExportedError(methodName)
	}

	methodFunc := reflect.ValueOf(driver.methodsReceiver).MethodByName(methodName)
	methodTemplate := func(*mgo.Session) error { return nil }

	if methodFunc.Type() != reflect.TypeOf(methodTemplate) {
		return gomethods.WrongMethodSignatureError(methodName)
	}

	return nil
}

func (driver *Driver) Invoke(methodName string) error {
	name := methodName
	migrateMethod := reflect.ValueOf(driver.methodsReceiver).MethodByName(name)
	if !migrateMethod.IsValid() {
		return gomethods.MissingMethodError(methodName)
	}

	session, err := driver.getNewSession()
	if err != nil {
		return fmt.Errorf("Migrate failed to get new session: %v", err)
	}
	defer session.Close()
	retValues := migrateMethod.Call([]reflect.Value{reflect.ValueOf(session)})
	if len(retValues) != 1 {
		return gomethods.WrongMethodSignatureError(name)
	}

	if !retValues[0].IsNil() {
		err, ok := retValues[0].Interface().(error)
		if !ok {
			return gomethods.WrongMethodSignatureError(name)
		}
		return &gomethods.MethodInvocationFailedError{MethodName: name, Err: err}
	}

	return nil
}
