package model

import "time"

// RolePermission 角色权限关联表
// 用于建立角色和权限的多对多关系
type RolePermission struct {
	// ID 角色权限唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// RoleID 角色ID，外键，关联角色表
	RoleID int64 `gorm:"type:bigint;not null;index" json:"role_id"`

	// PermissionID 权限ID，外键，关联权限表
	PermissionID int64 `gorm:"type:bigint;not null;index" json:"permission_id"`

	// CreatedAt 用户角色创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 用户角色更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
