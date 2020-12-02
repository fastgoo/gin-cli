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


CREATE TABLE `wk_company` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '公司名称',
  `contact_name` varchar(255) NOT NULL DEFAULT '' COMMENT '公司联系人',
  `contact_mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '公司联系电话',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '公司地址',
  `address_pos` varchar(255) NOT NULL DEFAULT '' COMMENT '公司坐标',
  `license_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '营业执照图片',
  `status` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '-1未激活 0正常 1锁定',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公司信息表'

JSON Sample
-------------------------------------
{    "id": 87,    "name": "whbXMYICVsbdibFssnEcGNcdA",    "contact_name": "VtvPMAozoCOeOHsgZjjQzfFuT",    "create_time": "2194-04-09T04:05:59.243701678+08:00",    "contact_mobile": "mNQXSUZzCRKeGfcnLwKMVCiMM",    "address": "yZnpJyjxoXYWMTuHhEpZOyxXC",    "address_pos": "xnnfzZOxHhQDBivzhCoBlGdtA",    "license_pic": "OcXUgVkxxAyQnqYbEmpazyffq",    "status": 27,    "update_time": "2186-12-14T16:09:23.065019266+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WkCompany struct is a row record of the wk_company table in the we-work database
type WkCompany struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string // 公司名称
	//[ 2] contact_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactName string // 公司联系人
	//[ 3] contact_mobile                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactMobile string // 公司联系电话
	//[ 4] address                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Address string // 公司地址
	//[ 5] address_pos                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AddressPos string // 公司坐标
	//[ 6] license_pic                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	LicensePic string // 营业执照图片
	//[ 7] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Status int32 // -1未激活 0正常 1锁定
	//[ 8] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time
	//[ 9] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time
}

var wk_companyTableInfo = &TableInfo{
	Name: "wk_company",
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
			Comment:            `公司名称`,
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
			Name:               "contact_name",
			Comment:            `公司联系人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ContactName",
			GoFieldType:        "string",
			JSONFieldName:      "contact_name",
			ProtobufFieldName:  "contact_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "contact_mobile",
			Comment:            `公司联系电话`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ContactMobile",
			GoFieldType:        "string",
			JSONFieldName:      "contact_mobile",
			ProtobufFieldName:  "contact_mobile",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "address",
			Comment:            `公司地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "address_pos",
			Comment:            `公司坐标`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AddressPos",
			GoFieldType:        "string",
			JSONFieldName:      "address_pos",
			ProtobufFieldName:  "address_pos",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "license_pic",
			Comment:            `营业执照图片`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "LicensePic",
			GoFieldType:        "string",
			JSONFieldName:      "license_pic",
			ProtobufFieldName:  "license_pic",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkCompany) TableName() string {
	return "wk_company"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WkCompany) BeforeSave() error {
	return nil
}

/*func (w *WkCompany) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkCompany) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkCompany) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkCompany) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkCompany) TableInfo() *TableInfo {
	return wk_companyTableInfo
}

// Get One
func (w *WkCompany) Get(fields string, query interface{}, args ...interface{}) (ret WkCompany, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkCompany) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkCompany, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkCompany) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkCompany{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkCompany) Insert(data WkCompany) uint32 {
	err := DB.Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkCompany) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkCompany) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkCompany{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkCompany) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
