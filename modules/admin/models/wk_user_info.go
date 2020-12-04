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


CREATE TABLE `wk_user_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `create_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `username` char(30) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(2) NOT NULL DEFAULT '-1' COMMENT '-1未激活 0正常 1锁定',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表'

JSON Sample
-------------------------------------
{    "avatar": "EhzgebFsAFxEyeJNbSFizYyba",    "status": 6,    "create_time": "2118-01-04T12:54:04.687881929+08:00",    "update_time": "2247-07-09T04:04:07.523449518+08:00",    "id": 93,    "create_user_id": 20,    "username": "ivdExrKwBuPHISVFjwdNQBjNl",    "password": "rTCOpOaYSdwVnnViWOhMZmQSO",    "name": "hArltWmTGjJzlyDNajhWuLWdV"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WkUserInfo struct is a row record of the wk_user_info table in the we-work database
type WkUserInfo struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `json:"id"`
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `json:"name"` // 昵称
	//[ 2] avatar                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Avatar string `json:"avatar"` // 头像
	//[ 3] create_user_id                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CreateUserID int32 `json:"create_user_id"` // 创建人id
	//[ 4] username                                       char(30)             null: false  primary: false  isArray: false  auto: false  col: char            len: 30      default: []
	Username string `json:"username"` // 用户名
	//[ 5] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `json:"password"` // 密码
	//[ 6] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Status int32 `json:"status"` // -1未激活 0正常 1锁定
	//[ 7] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time `json:"create_time"`
	//[ 8] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `json:"update_time"`
}

type wkuserinfoPages struct {
	Rows        []WkUserInfo `json:"rows"`
	Count       int          `json:"count"`
	CurrentPage int          `json:"current_page"`
	MaxPage     int          `json:"max_page"`
}

var wk_user_infoTableInfo = &TableInfo{
	Name: "wk_user_info",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "avatar",
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
			GoFieldName:        "Avatar",
			GoFieldType:        "string",
			JSONFieldName:      "avatar",
			ProtobufFieldName:  "avatar",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "create_user_id",
			Comment:            `创建人id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CreateUserID",
			GoFieldType:        "int32",
			JSONFieldName:      "create_user_id",
			ProtobufFieldName:  "create_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkUserInfo) TableName() string {
	return "wk_user_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*func (w *WkUserInfo) BeforeSave() error {
	return nil
}

func (w *WkUserInfo) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserInfo) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkUserInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkUserInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkUserInfo) TableInfo() *TableInfo {
	return wk_user_infoTableInfo
}

// Get One
func (w *WkUserInfo) Get(fields string, query interface{}, args ...interface{}) (ret WkUserInfo, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkUserInfo) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkUserInfo, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Get Page
func (w *WkUserInfo) Pages(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (pages wkuserinfoPages, has bool) {
	var ret []WkUserInfo
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
func (w *WkUserInfo) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkUserInfo{}).Omit("CreateTime", "UpdateTime").Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkUserInfo) Insert(data WkUserInfo) uint32 {
	err := DB.Omit("CreateTime", "UpdateTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkUserInfo) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkUserInfo) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkUserInfo{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkUserInfo) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
