package model

import "time"

// Permission 权限模型
// 用于定义系统中的权限，用于细粒度的权限控制
type Permission struct {
	// ID 权限唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// CreatedAt 权限创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 权限更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`

	// Name 权限名称，唯一索引，不能为空，最大长度50字符
	// 用于显示和识别权限，如"用户管理"、"角色管理"等
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`

	// Code 权限代码，唯一索引，不能为空，最大长度50字符
	// 用于程序内部识别权限，通常使用英文，如"user:create"、"role:delete"等
	Code string `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`

	// Desc 权限描述，不能为空，最大长度255字符
	// 用于详细说明权限的作用和范围
	Desc string `gorm:"type:varchar(255);not null" json:"desc"`
}

func (Permission) TableName() string {
	return "permissions"
}
