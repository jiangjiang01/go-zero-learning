package model

import "time"

type Category struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Desc string `gorm:"type:varchar(255)" json:"desc"`

	ParentID int64 `gorm:"default:0;index" json:"parent_id"` // 0表示顶级分类

	Sort   int `gorm:"default:0" json:"sort"`                // 排序
	Status int `gorm:"type:tinyint;default:1" json:"status"` // 状态 1:启用 0:禁用

	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

func (Category) TableName() string {
	return "categories"
}
