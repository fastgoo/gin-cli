package models

import (
	"database/sql"
	"math"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `wk_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(30) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `head_img` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` tinyint(2) NOT NULL DEFAULT '-1' COMMENT '-1未激活 0正常 1锁定',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统权限菜单表'

JSON Sample
-------------------------------------
{    "username": "OXTPiFyJRgbHZWvNdOqAgMKBj",    "password": "WTKBPEKFjrmTkxUTUteCBFkzV",    "head_img": "hMYrBdAXjIAlZaHqBpkjGPjuA",    "nickname": "ksQDqHwdCXuhrCzWvGxTEXMuk",    "status": 14,    "create_time": "2032-04-12T08:11:29.481748142+08:00",    "update_time": "2256-09-29T13:19:26.033200857+08:00",    "id": 97}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WkRoles struct is a row record of the wk_roles table in the we-work database
type WkRoles struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `json:"id"`
	//[ 1] username                                       char(30)             null: false  primary: false  isArray: false  auto: false  col: char            len: 30      default: []
	Username string `json:"username"` // 用户名
	//[ 2] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `json:"password"` // 密码
	//[ 3] head_img                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	HeadImg string `json:"head_img"` // 头像
	//[ 4] nickname                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Nickname string `json:"nickname"` // 昵称
	//[ 5] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Status int32 `json:"status"` // -1未激活 0正常 1锁定
	//[ 6] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time `json:"create_time"`
	//[ 7] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `json:"update_time"`
}

type wkrolesPages struct {
	Rows        []WkRoles `json:"rows"`
	Count       int       `json:"count"`
	CurrentPage int       `json:"current_page"`
	MaxPage     int       `json:"max_page"`
}

var wk_rolesTableInfo = &TableInfo{
	Name: "wk_roles",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "username",
			Comment:            `用户名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       30,
			GoFieldName:        "Username",
			GoFieldType:        "string",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
			Comment:            `密码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "head_img",
			Comment:            `头像`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "HeadImg",
			GoFieldType:        "string",
			JSONFieldName:      "head_img",
			ProtobufFieldName:  "head_img",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "nickname",
			Comment:            `昵称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Nickname",
			GoFieldType:        "string",
			JSONFieldName:      "nickname",
			ProtobufFieldName:  "nickname",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `-1未激活 0正常 1锁定`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "create_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkRoles) TableName() string {
	return "wk_roles"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*func (w *WkRoles) BeforeSave() error {
	return nil
}

func (w *WkRoles) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkRoles) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkRoles) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkRoles) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkRoles) TableInfo() *TableInfo {
	return wk_rolesTableInfo
}

// Get One
func (w *WkRoles) Get(fields string, query interface{}, args ...interface{}) (ret WkRoles, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkRoles) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkRoles, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Get Page
func (w *WkRoles) Pages(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (pages wkrolesPages, has bool) {
	var ret []WkRoles
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	pages.Rows = ret
	pages.Count = int(w.Count(query, args...))
	pages.CurrentPage = page
	pages.MaxPage = int(math.Ceil(float64(pages.Count) / float64(nums)))
	return
}

// Update
func (w *WkRoles) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkRoles{}).Omit("CreateTime", "UpdateTime").Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkRoles) Insert(data WkRoles) uint32 {
	err := DB.Omit("CreateTime", "UpdateTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkRoles) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkRoles) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkRoles{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkRoles) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
