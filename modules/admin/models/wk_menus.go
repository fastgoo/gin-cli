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


CREATE TABLE `wk_menus` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单ID',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `key` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单key',
  `component` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单组件',
  `icon` varchar(20) NOT NULL DEFAULT '' COMMENT '菜单icon',
  `routers` text COMMENT '菜单路由权限',
  `sort` tinyint(4) unsigned NOT NULL DEFAULT '99' COMMENT '菜单顺序',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单路由',
  `is_action` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是操作权限',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是默认，默认不可删除，修改',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '0 正常 1不显示',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单列表'

JSON Sample
-------------------------------------
{    "updated_time": "2049-06-16T19:55:58.06482824+08:00",    "key": "NIpbfYgeBeHLUBahzAvovGAuc",    "component": "WfSHewYrorThfxHywaLWplbmI",    "path": "KqVwxCQRIqOlAdcbPdvqeVXdl",    "is_action": 17,    "status": 65,    "created_time": "2026-11-17T19:25:01.961300337+08:00",    "title": "WmhAgPPyswZbydcJqqzZbysgG",    "id": 95,    "pid": 39,    "icon": "rXyOmTOuJeaWepQFEfMvWMiYd",    "routers": "XSNbQrKpoardDZGsXbOslLFhe",    "sort": 39,    "is_default": 90}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 7] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned



*/

// WkMenus struct is a row record of the wk_menus table in the we-work database
type WkMenus struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `json:"id"`
	//[ 1] pid                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Pid uint32 `json:"pid"` // 父级菜单ID
	//[ 2] title                                          varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Title string `json:"title"` // 菜单名称
	//[ 3] key                                            varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Key string `json:"key"` // 菜单key
	//[ 4] component                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Component string `json:"component"` // 菜单组件
	//[ 5] icon                                           varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Icon string `json:"icon"` // 菜单icon
	//[ 6] routers                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Routers sql.NullString `json:"routers"` // 菜单路由权限
	//[ 7] sort                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [99]
	Sort uint32 `json:"sort"` // 菜单顺序
	//[ 8] path                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Path string `json:"path"` // 菜单路由
	//[ 9] is_action                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsAction int32 `json:"is_action"` // 是否是操作权限
	//[10] is_default                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsDefault uint32 `json:"is_default"` // 是否是默认，默认不可删除，修改
	//[11] status                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Status uint32 `json:"status"` // 0 正常 1不显示
	//[12] created_time                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedTime time.Time `json:"created_time"`
	//[13] updated_time                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedTime time.Time `json:"updated_time"`
}

type wkmenusPages struct {
	Rows        []WkMenus `json:"rows"`
	Count       int       `json:"count"`
	CurrentPage int       `json:"current_page"`
	MaxPage     int       `json:"max_page"`
}

var wk_menusTableInfo = &TableInfo{
	Name: "wk_menus",
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
			Name:               "pid",
			Comment:            `父级菜单ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Pid",
			GoFieldType:        "uint32",
			JSONFieldName:      "pid",
			ProtobufFieldName:  "pid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
			Comment:            `菜单名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "key",
			Comment:            `菜单key`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Key",
			GoFieldType:        "string",
			JSONFieldName:      "key",
			ProtobufFieldName:  "key",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "component",
			Comment:            `菜单组件`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Component",
			GoFieldType:        "string",
			JSONFieldName:      "component",
			ProtobufFieldName:  "component",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "icon",
			Comment:            `菜单icon`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Icon",
			GoFieldType:        "string",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "routers",
			Comment:            `菜单路由权限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Routers",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "routers",
			ProtobufFieldName:  "routers",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sort",
			Comment:            `菜单顺序`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Sort",
			GoFieldType:        "uint32",
			JSONFieldName:      "sort",
			ProtobufFieldName:  "sort",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "path",
			Comment:            `菜单路由`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Path",
			GoFieldType:        "string",
			JSONFieldName:      "path",
			ProtobufFieldName:  "path",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "is_action",
			Comment:            `是否是操作权限`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsAction",
			GoFieldType:        "int32",
			JSONFieldName:      "is_action",
			ProtobufFieldName:  "is_action",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "is_default",
			Comment:            `是否是默认，默认不可删除，修改`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsDefault",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_default",
			ProtobufFieldName:  "is_default",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "status",
			Comment:            `0 正常 1不显示`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "created_time",
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
			GoFieldName:        "CreatedTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "updated_time",
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
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WkMenus) TableName() string {
	return "wk_menus"
}

// BeforeSave invoked before saving, return an error if field is not populated.
/*func (w *WkMenus) BeforeSave() error {
	return nil
}

func (w *WkMenus) BeforeCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) AfterCreate(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) BeforeUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) AfterUpdate(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) BeforeDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) AfterDelete(tx *gorm.DB) (err error) {
    return
}

func (w *WkMenus) AfterFind(tx *gorm.DB) (err error) {
    return
}*/

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WkMenus) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WkMenus) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WkMenus) TableInfo() *TableInfo {
	return wk_menusTableInfo
}

// Get One
func (w *WkMenus) Get(fields string, query interface{}, args ...interface{}) (ret WkMenus, has bool) {
	err := DB.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return
	}
	has = true
	return
}

// Get List
func (w *WkMenus) List(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (ret []WkMenus, has bool) {
	err := DB.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil || len(ret) == 0 {
		return
	}
	has = true
	return
}

// Get Page
func (w *WkMenus) Pages(fields string, order string, page int, nums int, query interface{}, args ...interface{}) (pages wkmenusPages, has bool) {
	var ret []WkMenus
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
func (w *WkMenus) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	if DB.Model(WkMenus{}).Omit("CreateTime", "UpdateTime").Where(query, args...).Updates(data).RowsAffected == 0 {
		return false
	}
	return true
}

// Insert
func (w *WkMenus) Insert(data WkMenus) uint32 {
	err := DB.Omit("CreateTime", "UpdateTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// Create
func (w *WkMenus) Create() uint32 {
	err := DB.Create(&w).Error
	if err != nil {
		return 0
	}
	return w.ID
}

// Count
func (w *WkMenus) Count(query interface{}, args ...interface{}) int64 {
	var count int64
	err := DB.Model(&WkMenus{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// Sum
func (w *WkMenus) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := DB.Table(w.TableName()).Select("sum("+field+") as nums").Where(query, args...).Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}
