// generated by xgen -- DO NOT EDIT
package xsql_test

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	USER          = schema.TABLE("users", "USER")
	USER_ID       = USER.PRIMARY("id", "ID")
	USER_CODE     = USER.COLUMN("code", "CODE")
	USER_NAME     = USER.COLUMN("name", "NAME")
	USER_PASSWORD = USER.COLUMN("password", "PASSWORD")
	USER_MEMO     = USER.COLUMN("memo", "MEMO")
	USER_GENRE    = USER.COLUMN("genre", "GENRE")
	USER_STATUS   = USER.COLUMN("status", "STATUS")
	USER_ROLES    = USER.COLUMN("roles", "ROLES")
	USER_POSTS    = USER.COLUMN("posts", "POSTS")
	USER_ORG      = USER.COLUMN("org", "ORG")
	USER_AREA     = USER.COLUMN("area", "AREA")
	USER_CREATER  = USER.CREATER("creater", "CREATER")
	USER_CREATED  = USER.CREATED("created", "CREATED")
	USER_MODIFIER = USER.MODIFIER("modifier", "MODIFIER")
	USER_MODIFIED = USER.MODIFIED("modified", "MODIFIED")
	USER_VERSION  = USER.VERSION("version", "VERSION")
	USER_DELETION = USER.DELETION("deletion", "DELETION")
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

func (me *User) Code() string {
	return me.code.Value()
}

func (me *User) SetCode(v string) {
	me.code.SetValue(v)
}

func (me *User) Name() string {
	return me.name.Value()
}

func (me *User) SetName(v string) {
	me.name.SetValue(v)
}

func (me *User) Password() string {
	return me.password.Value()
}

func (me *User) SetPassword(v string) {
	me.password.SetValue(v)
}

func (me *User) Memo() string {
	return me.memo.Value()
}

func (me *User) SetMemo(v string) {
	me.memo.SetValue(v)
}

func (me *User) Genre() string {
	return me.genre.Value()
}

func (me *User) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *User) Status() string {
	return me.status.Value()
}

func (me *User) SetStatus(v string) {
	me.status.SetValue(v)
}

func (me *User) Roles() string {
	return me.roles.Value()
}

func (me *User) SetRoles(v string) {
	me.roles.SetValue(v)
}

func (me *User) Posts() string {
	return me.posts.Value()
}

func (me *User) SetPosts(v string) {
	me.posts.SetValue(v)
}

func (me *User) Org() string {
	return me.org.Value()
}

func (me *User) SetOrg(v string) {
	me.org.SetValue(v)
}

func (me *User) Area() string {
	return me.area.Value()
}

func (me *User) SetArea(v string) {
	me.area.SetValue(v)
}

func (me *User) Creater() string {
	return me.creater.Value()
}

func (me *User) SetCreater(v string) {
	me.creater.SetValue(v)
}

func (me *User) Created() int64 {
	return me.created.Value()
}

func (me *User) SetCreated(v int64) {
	me.created.SetValue(v)
}

func (me *User) Modifier() string {
	return me.modifier.Value()
}

func (me *User) SetModifier(v string) {
	me.modifier.SetValue(v)
}

func (me *User) Modified() int64 {
	return me.modified.Value()
}

func (me *User) SetModified(v int64) {
	me.modified.SetValue(v)
}

func (me *User) Version() int {
	return me.version.Value()
}

func (me *User) SetVersion(v int) {
	me.version.SetValue(v)
}

func (me *User) Deletion() int {
	return me.deletion.Value()
}

func (me *User) SetDeletion(v int) {
	me.deletion.SetValue(v)
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
	me.code.SetColumn(USER_CODE)
	me.name.SetColumn(USER_NAME)
	me.password.SetColumn(USER_PASSWORD)
	me.memo.SetColumn(USER_MEMO)
	me.genre.SetColumn(USER_GENRE)
	me.status.SetColumn(USER_STATUS)
	me.roles.SetColumn(USER_ROLES)
	me.posts.SetColumn(USER_POSTS)
	me.org.SetColumn(USER_ORG)
	me.area.SetColumn(USER_AREA)
	me.creater.SetColumn(USER_CREATER)
	me.created.SetColumn(USER_CREATED)
	me.modifier.SetColumn(USER_MODIFIER)
	me.modified.SetColumn(USER_MODIFIED)
	me.version.SetColumn(USER_VERSION)
	me.deletion.SetColumn(USER_DELETION)
}

func (me *User) initSetDefault() {
}

func (me *User) initSetField() {
	me.id.SetField(entity.DefaultField())
	me.code.SetField(entity.DefaultField())
	me.name.SetField(entity.DefaultField())
	me.password.SetField(entity.DefaultField())
	me.memo.SetField(entity.DefaultField())
	me.genre.SetField(entity.DefaultField())
	me.status.SetField(entity.DefaultField())
	me.roles.SetField(entity.DefaultField())
	me.posts.SetField(entity.DefaultField())
	me.org.SetField(entity.DefaultField())
	me.area.SetField(entity.DefaultField())
	me.creater.SetField(entity.DefaultField())
	me.created.SetField(entity.DefaultField())
	me.modifier.SetField(entity.DefaultField())
	me.modified.SetField(entity.DefaultField())
	me.version.SetField(entity.DefaultField())
	me.deletion.SetField(entity.DefaultField())
}

func (me *User) initSetExcel() {
}

func (me *User) initSetJson() {
	me.id.Field().SetJson(entity.NewJsonBy("id"))
	me.code.Field().SetJson(entity.NewJsonBy("code"))
	me.name.Field().SetJson(entity.NewJsonBy("name"))
	me.password.Field().SetJson(entity.NewJsonBy("password"))
	me.memo.Field().SetJson(entity.NewJsonBy("memo"))
	me.genre.Field().SetJson(entity.NewJsonBy("genre"))
	me.status.Field().SetJson(entity.NewJsonBy("status"))
	me.roles.Field().SetJson(entity.NewJsonBy("roles"))
	me.posts.Field().SetJson(entity.NewJsonBy("posts"))
	me.org.Field().SetJson(entity.NewJsonBy("org"))
	me.area.Field().SetJson(entity.NewJsonBy("area"))
	me.creater.Field().SetJson(entity.NewJsonBy("creater"))
	me.created.Field().SetJson(entity.NewJsonBy("created"))
	me.modifier.Field().SetJson(entity.NewJsonBy("modifier"))
	me.modified.Field().SetJson(entity.NewJsonBy("modified"))
	me.version.Field().SetJson(entity.NewJsonBy("version"))
	me.deletion.Field().SetJson(entity.NewJsonBy("deletion"))
}

func (me *User) initSetXml() {
	me.id.Field().SetXml(entity.NewXmlBy("id"))
	me.code.Field().SetXml(entity.NewXmlBy("code"))
	me.name.Field().SetXml(entity.NewXmlBy("name"))
	me.password.Field().SetXml(entity.NewXmlBy("password"))
	me.memo.Field().SetXml(entity.NewXmlBy("memo"))
	me.genre.Field().SetXml(entity.NewXmlBy("genre"))
	me.status.Field().SetXml(entity.NewXmlBy("status"))
	me.roles.Field().SetXml(entity.NewXmlBy("roles"))
	me.posts.Field().SetXml(entity.NewXmlBy("posts"))
	me.org.Field().SetXml(entity.NewXmlBy("org"))
	me.area.Field().SetXml(entity.NewXmlBy("area"))
	me.creater.Field().SetXml(entity.NewXmlBy("creater"))
	me.created.Field().SetXml(entity.NewXmlBy("created"))
	me.modifier.Field().SetXml(entity.NewXmlBy("modifier"))
	me.modified.Field().SetXml(entity.NewXmlBy("modified"))
	me.version.Field().SetXml(entity.NewXmlBy("version"))
	me.deletion.Field().SetXml(entity.NewXmlBy("deletion"))
}

func (me User) New() entity.Interface {
	return NewUser()
}

func (me *User) Get(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.Value()
	case USER_CODE.Name():
		return me.code.Value()
	case USER_NAME.Name():
		return me.name.Value()
	case USER_PASSWORD.Name():
		return me.password.Value()
	case USER_MEMO.Name():
		return me.memo.Value()
	case USER_GENRE.Name():
		return me.genre.Value()
	case USER_STATUS.Name():
		return me.status.Value()
	case USER_ROLES.Name():
		return me.roles.Value()
	case USER_POSTS.Name():
		return me.posts.Value()
	case USER_ORG.Name():
		return me.org.Value()
	case USER_AREA.Name():
		return me.area.Value()
	case USER_CREATER.Name():
		return me.creater.Value()
	case USER_CREATED.Name():
		return me.created.Value()
	case USER_MODIFIER.Name():
		return me.modifier.Value()
	case USER_MODIFIED.Name():
		return me.modified.Value()
	case USER_VERSION.Name():
		return me.version.Value()
	case USER_DELETION.Name():
		return me.deletion.Value()
	}
	return nil
}

func (me *User) GetPtr(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.ValuePtr()
	case USER_CODE.Name():
		return me.code.ValuePtr()
	case USER_NAME.Name():
		return me.name.ValuePtr()
	case USER_PASSWORD.Name():
		return me.password.ValuePtr()
	case USER_MEMO.Name():
		return me.memo.ValuePtr()
	case USER_GENRE.Name():
		return me.genre.ValuePtr()
	case USER_STATUS.Name():
		return me.status.ValuePtr()
	case USER_ROLES.Name():
		return me.roles.ValuePtr()
	case USER_POSTS.Name():
		return me.posts.ValuePtr()
	case USER_ORG.Name():
		return me.org.ValuePtr()
	case USER_AREA.Name():
		return me.area.ValuePtr()
	case USER_CREATER.Name():
		return me.creater.ValuePtr()
	case USER_CREATED.Name():
		return me.created.ValuePtr()
	case USER_MODIFIER.Name():
		return me.modifier.ValuePtr()
	case USER_MODIFIED.Name():
		return me.modified.ValuePtr()
	case USER_VERSION.Name():
		return me.version.ValuePtr()
	case USER_DELETION.Name():
		return me.deletion.ValuePtr()
	}
	return nil
}

func (me *User) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.String()
	case "code":
		return me.code.String()
	case "name":
		return me.name.String()
	case "password":
		return me.password.String()
	case "memo":
		return me.memo.String()
	case "genre":
		return me.genre.String()
	case "status":
		return me.status.String()
	case "roles":
		return me.roles.String()
	case "posts":
		return me.posts.String()
	case "org":
		return me.org.String()
	case "area":
		return me.area.String()
	case "creater":
		return me.creater.String()
	case "created":
		return me.created.String()
	case "modifier":
		return me.modifier.String()
	case "modified":
		return me.modified.String()
	case "version":
		return me.version.String()
	case "deletion":
		return me.deletion.String()
	}
	return ""
}

func (me *User) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	case "code":
		return me.code.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "password":
		return me.password.SetString(value)
	case "memo":
		return me.memo.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "status":
		return me.status.SetString(value)
	case "roles":
		return me.roles.SetString(value)
	case "posts":
		return me.posts.SetString(value)
	case "org":
		return me.org.SetString(value)
	case "area":
		return me.area.SetString(value)
	case "creater":
		return me.creater.SetString(value)
	case "created":
		return me.created.SetString(value)
	case "modifier":
		return me.modifier.SetString(value)
	case "modified":
		return me.modified.SetString(value)
	case "version":
		return me.version.SetString(value)
	case "deletion":
		return me.deletion.SetString(value)
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
	case USER_CODE.Name():
		return &me.code, true
	case USER_NAME.Name():
		return &me.name, true
	case USER_PASSWORD.Name():
		return &me.password, true
	case USER_MEMO.Name():
		return &me.memo, true
	case USER_GENRE.Name():
		return &me.genre, true
	case USER_STATUS.Name():
		return &me.status, true
	case USER_ROLES.Name():
		return &me.roles, true
	case USER_POSTS.Name():
		return &me.posts, true
	case USER_ORG.Name():
		return &me.org, true
	case USER_AREA.Name():
		return &me.area, true
	case USER_CREATER.Name():
		return &me.creater, true
	case USER_CREATED.Name():
		return &me.created, true
	case USER_MODIFIER.Name():
		return &me.modifier, true
	case USER_MODIFIED.Name():
		return &me.modified, true
	case USER_VERSION.Name():
		return &me.version, true
	case USER_DELETION.Name():
		return &me.deletion, true
	}
	return nil, false
}

func (me *User) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return USER_ID, true
	case "code":
		return USER_CODE, true
	case "name":
		return USER_NAME, true
	case "password":
		return USER_PASSWORD, true
	case "memo":
		return USER_MEMO, true
	case "genre":
		return USER_GENRE, true
	case "status":
		return USER_STATUS, true
	case "roles":
		return USER_ROLES, true
	case "posts":
		return USER_POSTS, true
	case "org":
		return USER_ORG, true
	case "area":
		return USER_AREA, true
	case "creater":
		return USER_CREATER, true
	case "created":
		return USER_CREATED, true
	case "modifier":
		return USER_MODIFIER, true
	case "modified":
		return USER_MODIFIED, true
	case "version":
		return USER_VERSION, true
	case "deletion":
		return USER_DELETION, true
	}
	return nil, false
}

func (me *User) Columns() []schema.Column {
	return []schema.Column{
		USER_ID,
		USER_CODE,
		USER_NAME,
		USER_PASSWORD,
		USER_MEMO,
		USER_GENRE,
		USER_STATUS,
		USER_ROLES,
		USER_POSTS,
		USER_ORG,
		USER_AREA,
		USER_CREATER,
		USER_CREATED,
		USER_MODIFIER,
		USER_MODIFIED,
		USER_VERSION,
		USER_DELETION,
	}
}

func (me *User) Names() []string {
	return []string{
		"id",
		"code",
		"name",
		"password",
		"memo",
		"genre",
		"status",
		"roles",
		"posts",
		"org",
		"area",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
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
	b.WriteString(`,"code":"` + jsons.Format(me.GetString("code")) + `"`)
	b.WriteString(`,"name":"` + jsons.Format(me.GetString("name")) + `"`)
	b.WriteString(`,"password":"` + jsons.Format(me.GetString("password")) + `"`)
	b.WriteString(`,"memo":"` + jsons.Format(me.GetString("memo")) + `"`)
	b.WriteString(`,"genre":"` + jsons.Format(me.GetString("genre")) + `"`)
	b.WriteString(`,"status":"` + jsons.Format(me.GetString("status")) + `"`)
	b.WriteString(`,"roles":"` + jsons.Format(me.GetString("roles")) + `"`)
	b.WriteString(`,"posts":"` + jsons.Format(me.GetString("posts")) + `"`)
	b.WriteString(`,"org":"` + jsons.Format(me.GetString("org")) + `"`)
	b.WriteString(`,"area":"` + jsons.Format(me.GetString("area")) + `"`)
	b.WriteString(`,"creater":"` + jsons.Format(me.GetString("creater")) + `"`)
	b.WriteString(`,"created":` + me.GetString("created"))
	b.WriteString(`,"modifier":"` + jsons.Format(me.GetString("modifier")) + `"`)
	b.WriteString(`,"modified":` + me.GetString("modified"))
	b.WriteString(`,"version":` + me.GetString("version"))
	b.WriteString(`,"deletion":` + me.GetString("deletion"))
	b.WriteString("}")
	return b.String()
}

func (me *User) ExcelColumns() []string {
	return nil
}
