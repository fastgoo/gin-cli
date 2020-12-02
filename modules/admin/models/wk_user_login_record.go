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


CREATE TABLE `wk_user_login_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `ip` varchar(255) NOT NULL DEFAULT '' COMMENT '登录的ip地址',
  `area` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录记录表'

JSON Sample
-------------------------------------
{    "id": 81,    "user_id": 28,    "ip": "yPYEEcHVHSssuzWWnPAHSGAvp",    "area": "DKfmvArcODlAucNoksKSticfX",    "create_time": "2281-10-28T17:36:29.597923937+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// WkUserLoginRecord struct is a row record of the wk_user_login_record table in the we-work database
type WkUserLoginRecord struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	UserID uint32 // 用户id
	//[ 2] ip                                             varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	IP string // 登录的ip地址
	//[ 3] area                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Area string // ip地址
	//[ 4] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time
}

var wk_user_login_recordTableInfo = &TableInfo{
	Name: "wk_user_login_record",
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
			Name:               "user_id",
			Comment:            `用户id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip",
			Comment:            `登录的ip地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "area",
			Comment:            `ip地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Area",
			GoFieldType:        "string",
			JSONFieldName:      "area",
			ProtobufFieldName:  "area",
			ProtobufType:       "string",
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
func (w *WkUserLoginRecord) TableName() string {
	return "wk_user_login_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*
func (w *WkUserLoginRecord) BeforeSave() error {
	return nil
}

func (w *WkUserLoginRecord) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkUserLoginRecord) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkUserLoginRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkUserLoginRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkUserLoginRecord) TableInfo() *TableInfo {
	return wk_user_login_recordTableInfo
}

// Get One
func (w *WkUserLoginRecord) Get(fields string, query interface{}, args ...interface{}) (ret WkUserLoginRecord, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkUserLoginRecord) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkUserLoginRecord, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkUserLoginRecord) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkUserLoginRecord{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkUserLoginRecord) Insert(data WkUserLoginRecord) uint32 {
	err := DB.Omit("CreateTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkUserLoginRecord) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkUserLoginRecord) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkUserLoginRecord{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkUserLoginRecord) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
