// generated by xgen -- DO NOT EDIT
package result_test

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	USER         = schema.TABLE("sys_user", "USER")
	USER_ID      = USER.PRIMARY("id", "ID")
	USER_NAME    = USER.COLUMN("name", "NAME")
	USER_PASSWD  = USER.COLUMN("passwd", "PASSWD")
	USER_AGE     = USER.COLUMN("age", "AGE")
	USER_EMAIL   = USER.COLUMN("email", "EMAIL")
	USER_VERSION = USER.COLUMN("version", "VERSION")
)

func NewUser() *User {
	e := &User{}
	e.init()
	return e
}

func (me *User) Id() string {
	return me.id.Value()
}

func (me *User) SetId(v string) {
	me.id.SetValue(v)
}

func (me *User) Name() string {
	return me.name.Value()
}

func (me *User) SetName(v string) {
	me.name.SetValue(v)
}

func (me *User) Passwd() string {
	return me.passwd.Value()
}

func (me *User) SetPasswd(v string) {
	me.passwd.SetValue(v)
}

func (me *User) Age() int {
	return me.age.Value()
}

func (me *User) SetAge(v int) {
	me.age.SetValue(v)
}

func (me *User) Email() string {
	return me.email.Value()
}

func (me *User) SetEmail(v string) {
	me.email.SetValue(v)
}

func (me *User) Version() int {
	return me.version.Value()
}

func (me *User) SetVersion(v int) {
	me.version.SetValue(v)
}

func (me *User) init() {
	me.table = USER
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
	me.initSetJson()
	me.initSetXml()
}

func (me *User) initSetDict() {
}

func (me *User) initSetColumn() {
	me.id.SetColumn(USER_ID)
	me.name.SetColumn(USER_NAME)
	me.passwd.SetColumn(USER_PASSWD)
	me.age.SetColumn(USER_AGE)
	me.email.SetColumn(USER_EMAIL)
	me.version.SetColumn(USER_VERSION)
}

func (me *User) initSetDefault() {
}

func (me *User) initSetField() {
	me.id.SetField(entity.DefaultField())
	me.name.SetField(entity.DefaultField())
	me.passwd.SetField(entity.DefaultField())
	me.age.SetField(entity.DefaultField())
	me.email.SetField(entity.DefaultField())
	me.version.SetField(entity.DefaultField())
}

func (me *User) initSetExcel() {
}

func (me *User) initSetJson() {
	me.id.Field().SetJson(entity.NewJsonBy("id"))
	me.name.Field().SetJson(entity.NewJsonBy("name"))
	me.passwd.Field().SetJson(entity.NewJsonBy("passwd"))
	me.age.Field().SetJson(entity.NewJsonBy("age"))
	me.email.Field().SetJson(entity.NewJsonBy("email"))
	me.version.Field().SetJson(entity.NewJsonBy("version"))
}

func (me *User) initSetXml() {
	me.id.Field().SetXml(entity.NewXmlBy("id"))
	me.name.Field().SetXml(entity.NewXmlBy("name"))
	me.passwd.Field().SetXml(entity.NewXmlBy("passwd"))
	me.age.Field().SetXml(entity.NewXmlBy("age"))
	me.email.Field().SetXml(entity.NewXmlBy("email"))
	me.version.Field().SetXml(entity.NewXmlBy("version"))
}

func (me User) New() entity.Interface {
	return NewUser()
}

func (me *User) Get(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.Value()
	case USER_NAME.Name():
		return me.name.Value()
	case USER_PASSWD.Name():
		return me.passwd.Value()
	case USER_AGE.Name():
		return me.age.Value()
	case USER_EMAIL.Name():
		return me.email.Value()
	case USER_VERSION.Name():
		return me.version.Value()
	}
	return nil
}

func (me *User) GetPtr(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.ValuePtr()
	case USER_NAME.Name():
		return me.name.ValuePtr()
	case USER_PASSWD.Name():
		return me.passwd.ValuePtr()
	case USER_AGE.Name():
		return me.age.ValuePtr()
	case USER_EMAIL.Name():
		return me.email.ValuePtr()
	case USER_VERSION.Name():
		return me.version.ValuePtr()
	}
	return nil
}

func (me *User) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.String()
	case "name":
		return me.name.String()
	case "passwd":
		return me.passwd.String()
	case "age":
		return me.age.String()
	case "email":
		return me.email.String()
	case "version":
		return me.version.String()
	}
	return ""
}

func (me *User) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "passwd":
		return me.passwd.SetString(value)
	case "age":
		return me.age.SetString(value)
	case "email":
		return me.email.SetString(value)
	case "version":
		return me.version.SetString(value)
	}
	return nil
}

func (me *User) Table() schema.Table {
	return me.table
}

func (me *User) Type(column string) (entity.Type, bool) {
	switch column {
	case USER_ID.Name():
		return &me.id, true
	case USER_NAME.Name():
		return &me.name, true
	case USER_PASSWD.Name():
		return &me.passwd, true
	case USER_AGE.Name():
		return &me.age, true
	case USER_EMAIL.Name():
		return &me.email, true
	case USER_VERSION.Name():
		return &me.version, true
	}
	return nil, false
}

func (me *User) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return USER_ID, true
	case "name":
		return USER_NAME, true
	case "passwd":
		return USER_PASSWD, true
	case "age":
		return USER_AGE, true
	case "email":
		return USER_EMAIL, true
	case "version":
		return USER_VERSION, true
	}
	return nil, false
}

func (me *User) Columns() []schema.Column {
	return []schema.Column{
		USER_ID,
		USER_NAME,
		USER_PASSWD,
		USER_AGE,
		USER_EMAIL,
		USER_VERSION,
	}
}

func (me *User) Names() []string {
	return []string{
		"id",
		"name",
		"passwd",
		"age",
		"email",
		"version",
	}
}

func (me *User) Value() *User {
	return me
}

func (me *User) Validate() error {
	return nil
}

func (me *User) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(`,"id":"` + jsons.Format(me.GetString("id")) + `"`)
	b.WriteString(`,"name":"` + jsons.Format(me.GetString("name")) + `"`)
	b.WriteString(`,"passwd":"` + jsons.Format(me.GetString("passwd")) + `"`)
	b.WriteString(`,"age":` + me.GetString("age"))
	b.WriteString(`,"email":"` + jsons.Format(me.GetString("email")) + `"`)
	b.WriteString(`,"version":` + me.GetString("version"))
	b.WriteString("}")
	return b.String()
}

func (me *User) ExcelColumns() []string {
	return nil
}