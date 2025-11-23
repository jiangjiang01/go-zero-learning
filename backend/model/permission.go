package model

import "time"

type Permission struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"` // 权限名称，唯一
	Code      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"` // 权限代码，唯一
	Desc      string    `gorm:"type:varchar(255);not null" json:"desc"`            // 权限描述
}

func (Permission) TableName() string {
	return "permissions"
}
