package model

import "time"

// UserRole 用户角色关联表
// 用于建立用户和角色的多对多关系
type UserRole struct {
	// ID 用户角色唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// UserID 用户ID，外键，关联用户表
	UserID int64 `gorm:"type:bigint;not null;index" json:"user_id"`

	// RoleID 角色ID，外键，关联角色表
	RoleID int64 `gorm:"type:bigint;not null;index" json:"role_id"`

	// CreatedAt 用户角色创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 用户角色更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
