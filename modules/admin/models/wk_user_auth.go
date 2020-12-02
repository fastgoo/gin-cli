package models

import (
	"database/sql"
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


CREATE TABLE `wk_user_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '授权类型：1企业微信',
  `mark` char(30) NOT NULL DEFAULT '' COMMENT '授权标识',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '绑定的用户id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户第三方授权绑定关系表'

JSON Sample
-------------------------------------
{    "mark": "gmemzdOuGEZyfJxXGtrncopOO",    "user_id": 73,    "create_time": "2159-09-16T04:10:48.161823996+08:00",    "id": 43,    "type": 70}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// WkUserAuth struct is a row record of the wk_user_auth table in the we-work database
type WkUserAuth struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] type                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Type uint32 // 授权类型：1企业微信
	//[ 2] mark                                           char(30)             null: false  primary: false  isArray: false  auto: false  col: char            len: 30      default: []
	Mark string // 授权标识
	//[ 3] user_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UserID int32 // 绑定的用户id
	//[ 4] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time
}

var wk_user_authTableInfo = &TableInfo{
	Name: "wk_user_auth",
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
			Name:               "type",
			Comment:            `授权类型：1企业微信`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mark",
			Comment:            `授权标识`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       30,
			GoFieldName:        "Mark",
			GoFieldType:        "string",
			JSONFieldName:      "mark",
			ProtobufFieldName:  "mark",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "user_id",
			Comment:            `绑定的用户id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkUserAuth) TableName() string {
	return "wk_user_auth"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WkUserAuth) BeforeSave() error {
	return nil
}

/*func (w *WkUserAuth) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserAuth) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkUserAuth) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkUserAuth) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkUserAuth) TableInfo() *TableInfo {
	return wk_user_authTableInfo
}

// Get One
func (w *WkUserAuth) Get(fields string, query interface{}, args ...interface{}) (ret WkUserAuth, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkUserAuth) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkUserAuth, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkUserAuth) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkUserAuth{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkUserAuth) Insert(data WkUserAuth) uint32 {
	err := DB.Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkUserAuth) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkUserAuth) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkUserAuth{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkUserAuth) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
