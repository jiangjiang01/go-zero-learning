package model

import "time"

// Role 角色模型
// 用于定义系统中的角色，用于权限管理
type Role struct {
	// ID 角色唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// CreatedAt 角色创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 角色更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`

	// Name 角色名称，唯一索引，不能为空，最大长度50字符
	// 用于显示和识别角色，如"管理员"、"普通用户"等
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`

	// Code 角色代码，唯一索引，不能为空，最大长度50字符
	// 用于程序内部识别角色，通常使用英文，如"admin"、"user"等
	Code string `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`

	// Desc 角色描述，不能为空，最大长度255字符
	// 用于详细说明角色的职责和权限范围
	Desc string `gorm:"type:varchar(255);not null" json:"desc"`

	// Users 角色拥有的用户列表（多对多关联）
	// Users []User `gorm:"many2many:user_roles" json:"users,omitempty"`

	// Permissions 角色拥有的权限列表（多对多关联）
	// Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions,omitempty"`
}

func (Role) TableName() string {
	return "roles"
}
