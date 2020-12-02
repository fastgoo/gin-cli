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


CREATE TABLE `wk_qiniu_file_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名称',
  `size` bigint(20) NOT NULL DEFAULT '0' COMMENT '文件大小',
  `key` varchar(255) DEFAULT '' COMMENT '七牛云key',
  `hash` varchar(255) NOT NULL DEFAULT '' COMMENT '七牛云hash',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `size_name` (`size`,`name`) USING BTREE COMMENT '查询文件是否重复'
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='七牛云上传图片记录，主要用于缓存文件的信息，方便后期删除，避免重复上传什么的'

JSON Sample
-------------------------------------
{    "id": 92,    "user_id": 99,    "name": "QrRxxUDNPexhTtAMSltYiglyL",    "size": 95,    "key": "qcyRDlPGmZGorTAUfaAdtNeMX",    "hash": "prLTgwjwbHaeRIxEwhBumwucw",    "create_time": "2156-03-18T14:03:39.861021239+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// WkQiniuFileRecord struct is a row record of the wk_qiniu_file_record table in the we-work database
type WkQiniuFileRecord struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	UserID uint32 // 用户id
	//[ 2] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string // 文件名称
	//[ 3] size                                           bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	Size int64 // 文件大小
	//[ 4] key                                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Key string // 七牛云key
	//[ 5] hash                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Hash string // 七牛云hash
	//[ 6] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time
}

var wk_qiniu_file_recordTableInfo = &TableInfo{
	Name: "wk_qiniu_file_record",
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
			Name:               "name",
			Comment:            `文件名称`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "size",
			Comment:            `文件大小`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "Size",
			GoFieldType:        "int64",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "key",
			Comment:            `七牛云key`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Key",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "key",
			ProtobufFieldName:  "key",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "hash",
			Comment:            `七牛云hash`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Hash",
			GoFieldType:        "string",
			JSONFieldName:      "hash",
			ProtobufFieldName:  "hash",
			ProtobufType:       "string",
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
	},
}

// TableName sets the insert table name for this struct type
func (w *WkQiniuFileRecord) TableName() string {
	return "wk_qiniu_file_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*func (w *WkQiniuFileRecord) BeforeSave() error {
	return nil
}

func (w *WkQiniuFileRecord) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkQiniuFileRecord) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkQiniuFileRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkQiniuFileRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkQiniuFileRecord) TableInfo() *TableInfo {
	return wk_qiniu_file_recordTableInfo
}

// Get One
func (w *WkQiniuFileRecord) Get(fields string, query interface{}, args ...interface{}) (ret WkQiniuFileRecord, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkQiniuFileRecord) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkQiniuFileRecord, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkQiniuFileRecord) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkQiniuFileRecord{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkQiniuFileRecord) Insert(data WkQiniuFileRecord) uint32 {
	err := DB.Omit("CreateTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkQiniuFileRecord) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkQiniuFileRecord) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkQiniuFileRecord{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkQiniuFileRecord) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
