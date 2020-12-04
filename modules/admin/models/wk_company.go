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


CREATE TABLE `wk_company` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '公司名称',
  `contact_name` varchar(255) NOT NULL DEFAULT '' COMMENT '公司联系人',
  `contact_mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '公司联系电话',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '公司地址',
  `address_pos` varchar(255) NOT NULL DEFAULT '' COMMENT '公司坐标',
  `license_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '营业执照图片',
  `verify_remark` varchar(255) NOT NULL DEFAULT '' COMMENT '审核备注：审核不通过需要备注',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '-1草稿 0未激活 1已激活 2激活失败',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='公司信息表'

JSON Sample
-------------------------------------
{    "create_time": "2190-02-12T11:22:08.85384017+08:00",    "contact_name": "TwSQvqKdDWQAOXALSRnEtstua",    "license_pic": "RuuQnUzfJAuQrGJMhsLmAdXPn",    "status": 88,    "address": "aOvSdiYtwYnIqdLdxNqPvTFHw",    "address_pos": "hFOBxOJyDLQYuKlutPhwANtSp",    "verify_remark": "cFzvNjOidwrFGcnhXGMalrAkV",    "update_time": "2094-06-12T07:41:08.301103935+08:00",    "id": 24,    "name": "FcOGWcmtJcBSOctrGuBHbrsuJ",    "contact_mobile": "ZcvIfsKXCDjDroCerLRXBwFMJ"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WkCompany struct is a row record of the wk_company table in the we-work database
type WkCompany struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `json:"id"`
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `json:"name"` // 公司名称
	//[ 2] contact_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactName string `json:"contact_name"` // 公司联系人
	//[ 3] contact_mobile                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactMobile string `json:"contact_mobile"` // 公司联系电话
	//[ 4] address                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Address string `json:"address"` // 公司地址
	//[ 5] address_pos                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AddressPos string `json:"address_pos"` // 公司坐标
	//[ 6] license_pic                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	LicensePic string `json:"license_pic"` // 营业执照图片
	//[ 7] verify_remark                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	VerifyRemark string `json:"verify_remark"` // 审核备注：审核不通过需要备注
	//[ 8] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status int32 `json:"status"` // -1草稿 0未激活 1已激活 2激活失败
	//[ 9] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time `json:"create_time"`
	//[10] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `json:"update_time"`
}

type wkcompanyPages struct {
	Rows        []WkCompany `json:"rows"`
	Count       int         `json:"count"`
	CurrentPage int         `json:"current_page"`
	MaxPage     int         `json:"max_page"`
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
			Name:               "verify_remark",
			Comment:            `审核备注：审核不通过需要备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "VerifyRemark",
			GoFieldType:        "string",
			JSONFieldName:      "verify_remark",
			ProtobufFieldName:  "verify_remark",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
			Comment:            `-1草稿 0未激活 1已激活 2激活失败`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkCompany) TableName() string {
	return "wk_company"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*func (w *WkCompany) BeforeSave() error {
	return nil
}

func (w *WkCompany) BeforeCreate(tx *gorm.DB) (err error) {
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

// Get Page
func (w *WkCompany) Pages(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (pages wkcompanyPages, has bool) {
	var ret []WkCompany
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
func (w *WkCompany) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkCompany{}).Omit("CreateTime", "UpdateTime").Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkCompany) Insert(data WkCompany) uint32 {
	err := DB.Omit("CreateTime", "UpdateTime").Create(&data).Error
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
