package model

import "time"

// Menu 菜单模型
// 用于定义系统中的菜单结构，支持树形结构，用于前端菜单展示和权限控制
type Menu struct {
	// ID 菜单唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// CreatedAt 菜单创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 菜单更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`

	// Name 菜单名称，不能为空，最大长度50字符
	// 不唯一，允许同名但不同层级，用于前端显示
	Name string `gorm:"type:varchar(50);not null" json:"name"`

	// Code 菜单代码，唯一索引，不能为空，最大长度50字符
	// 用于权限控制和程序内部识别，如"system:user"、"system:role"等
	Code string `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`

	// Desc 菜单描述，可选，最大长度255字符
	// 用于详细说明菜单的功能和用途
	Desc string `gorm:"type:varchar(255)" json:"desc"`

	// ParentID 父菜单ID，默认值为0，建立索引
	// 0表示顶级菜单，非0表示子菜单，用于构建树形结构
	ParentID int64 `gorm:"default:0;index" json:"parent_id"`

	// Path 菜单路径，可选，最大长度255字符
	// 前端路由路径，如"/system/user"、"/system/role"等
	Path string `gorm:"type:varchar(255)" json:"path"`

	// Icon 菜单图标，可选，最大长度50字符
	// 图标名称或图标类名，用于前端显示，如"User"、"Setting"等
	Icon string `gorm:"type:varchar(50)" json:"icon"`

	// Type 菜单类型，默认值为1
	// 1:菜单（可点击跳转），2:按钮（用于权限控制，不显示在菜单中）
	Type int `gorm:"type:tinyint;default:1" json:"type"`

	// Sort 排序值，默认值为0
	// 数字越小越靠前，用于控制菜单的显示顺序
	Sort int `gorm:"default:0" json:"sort"`

	// Status 菜单状态，默认值为1
	// 1:启用（正常显示），0:禁用（不显示）
	Status int `gorm:"type:tinyint;default:1" json:"status"`
}

func (Menu) TableName() string {
	return "menus"
}
