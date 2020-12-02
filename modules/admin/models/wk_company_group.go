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


CREATE TABLE `wk_company_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `company_id` int(10) NOT NULL COMMENT '公司id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '群组名称',
  `add_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '添加该群组的用户id',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '群组备注信息',
  `status` tinyint(2) NOT NULL DEFAULT '-1' COMMENT '-1关闭 0正常 1锁定',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`company_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表'

JSON Sample
-------------------------------------
{    "name": "ZeQQkjJntuQcRElcVZFHCicps",    "add_user_id": 58,    "remark": "XgpfEAeJgGJPqvNrZghtzWlNh",    "status": 24,    "create_time": "2187-04-03T08:00:00.582100782+08:00",    "update_time": "2144-04-14T12:25:26.746035042+08:00",    "id": 24,    "company_id": 25}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WkCompanyGroup struct is a row record of the wk_company_group table in the we-work database
type WkCompanyGroup struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] company_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CompanyID int32 // 公司id
	//[ 2] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string // 群组名称
	//[ 3] add_user_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddUserID int32 // 添加该群组的用户id
	//[ 4] remark                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark string // 群组备注信息
	//[ 5] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Status int32 // -1关闭 0正常 1锁定
	//[ 6] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time
	//[ 7] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time
}

var wk_company_groupTableInfo = &TableInfo{
	Name: "wk_company_group",
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
			Name:               "company_id",
			Comment:            `公司id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CompanyID",
			GoFieldType:        "int32",
			JSONFieldName:      "company_id",
			ProtobufFieldName:  "company_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            `群组名称`,
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
			Name:               "add_user_id",
			Comment:            `添加该群组的用户id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddUserID",
			GoFieldType:        "int32",
			JSONFieldName:      "add_user_id",
			ProtobufFieldName:  "add_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "remark",
			Comment:            `群组备注信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `-1关闭 0正常 1锁定`,
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
func (w *WkCompanyGroup) TableName() string {
	return "wk_company_group"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WkCompanyGroup) BeforeSave() error {
	return nil
}

/*func (w *WkCompanyGroup) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompanyGroup) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkCompanyGroup) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkCompanyGroup) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkCompanyGroup) TableInfo() *TableInfo {
	return wk_company_groupTableInfo
}

// Get One
func (w *WkCompanyGroup) Get(fields string, query interface{}, args ...interface{}) (ret WkCompanyGroup, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkCompanyGroup) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkCompanyGroup, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkCompanyGroup) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkCompanyGroup{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkCompanyGroup) Insert(data WkCompanyGroup) uint32 {
	err := DB.Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkCompanyGroup) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkCompanyGroup) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkCompanyGroup{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkCompanyGroup) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
