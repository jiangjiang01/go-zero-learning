package model

import "time"

type Role struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"` // 角色名称，唯一
	Code      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"` // 角色代码，唯一
	Desc      string    `gorm:"type:varchar(255);not null" json:"desc"`            // 角色描述
}

func (Role) TableName() string {
	return "roles"
}
