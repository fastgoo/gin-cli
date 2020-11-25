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


CREATE TABLE `wk_sms_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1注册账号 2解除绑定 3绑定新账号 4忘记密码  5安全校验',
  `ip` char(20) NOT NULL DEFAULT '',
  `mobile` char(20) NOT NULL COMMENT '手机号码',
  `code` char(10) NOT NULL COMMENT '验证码',
  `expire_time` int(11) unsigned NOT NULL COMMENT '有效时间',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0未使用  1已使用',
  `send_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '短信发送状态',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `mobile+type` (`mobile`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='短信验证码表'

JSON Sample
-------------------------------------
{    "mobile": "SHMCCKytdIcPbAYJlqIKnYmDd",    "status": 22,    "send_status": 57,    "ip": "XHkJjiJDsinJsjUTheAJIvmBI",    "type": 99,    "code": "KVdbORZNvFjnKdRcfqRJhOGEG",    "expire_time": 59,    "created_time": "2246-03-19T14:43:08.124146755+08:00",    "updated_time": "2033-04-13T07:02:12.080775949+08:00",    "id": 67}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 5] column is set for unsigned
[ 7] column is set for unsigned



*/

// WkSmsCode struct is a row record of the wk_sms_code table in the we-work database
type WkSmsCode struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32
	//[ 1] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Type int32 // 1注册账号 2解除绑定 3绑定新账号 4忘记密码  5安全校验
	//[ 2] ip                                             char(20)             null: false  primary: false  isArray: false  auto: false  col: char            len: 20      default: []
	IP string
	//[ 3] mobile                                         char(20)             null: false  primary: false  isArray: false  auto: false  col: char            len: 20      default: []
	Mobile string // 手机号码
	//[ 4] code                                           char(10)             null: false  primary: false  isArray: false  auto: false  col: char            len: 10      default: []
	Code string // 验证码
	//[ 5] expire_time                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ExpireTime uint32 // 有效时间
	//[ 6] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status int32 // 0未使用  1已使用
	//[ 7] send_status                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	SendStatus uint32 // 短信发送状态
	//[ 8] created_time                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedTime time.Time // 创建时间
	//[ 9] updated_time                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedTime time.Time // 更新时间

}

var wk_sms_codeTableInfo = &TableInfo{
	Name: "wk_sms_code",
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
			Comment:            `1注册账号 2解除绑定 3绑定新账号 4忘记密码  5安全校验`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       20,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "mobile",
			Comment:            `手机号码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       20,
			GoFieldName:        "Mobile",
			GoFieldType:        "string",
			JSONFieldName:      "mobile",
			ProtobufFieldName:  "mobile",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "code",
			Comment:            `验证码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       10,
			GoFieldName:        "Code",
			GoFieldType:        "string",
			JSONFieldName:      "code",
			ProtobufFieldName:  "code",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "expire_time",
			Comment:            `有效时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpireTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "expire_time",
			ProtobufFieldName:  "expire_time",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `0未使用  1已使用`,
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
			Name:               "send_status",
			Comment:            `短信发送状态`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "SendStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "send_status",
			ProtobufFieldName:  "send_status",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "created_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkSmsCode) TableName() string {
	return "wk_sms_code"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WkSmsCode) BeforeSave() error {
	return nil
}

/*func (w *WkSmsCode) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkSmsCode) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkSmsCode) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkSmsCode) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkSmsCode) TableInfo() *TableInfo {
	return wk_sms_codeTableInfo
}

// Get One
func (w *WkSmsCode) Get(fields string, query interface{}, args ...interface{}) (ret WkSmsCode, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkSmsCode) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkSmsCode, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Update
func (w *WkSmsCode) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkSmsCode{}).Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkSmsCode) Insert(data WkSmsCode) uint32 {
	err := DB.Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkSmsCode) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkSmsCode) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkSmsCode{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkSmsCode) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
