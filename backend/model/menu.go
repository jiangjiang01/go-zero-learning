package model

import "time"

type Menu struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`             // 菜单名称，不唯一，允许同名但不同层级
	Code      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"` // 菜单代码，唯一（用于权限控制）
	Desc      string    `gorm:"type:varchar(255)" json:"desc"`                     // 菜单描述
	ParentID  int64     `gorm:"default:0;index" json:"parent_id"`                  // 父菜单ID,0表示顶级菜单
	Path      string    `gorm:"type:varchar(255)" json:"path"`                     // 菜单路径（前端路由）
	Icon      string    `gorm:"type:varchar(50)" json:"icon"`                      // 菜单图标
	Type      int       `gorm:"type:tinyint;default:1" json:"type"`                // 菜单类型，1:菜单,2:按钮
	Sort      int       `gorm:"default:0" json:"sort"`                             // 排序，数字越小越靠前
	Status    int       `gorm:"type:tinyint;default:1" json:"status"`              // 状态，1:启用,0:禁用
}

func (Menu) TableName() string {
	return "menus"
}
