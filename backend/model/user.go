package model

import "time"

// User 用户模型
// 用于存储系统用户的基本信息
type User struct {
	// ID 用户唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// Username 用户名，唯一索引，不能为空，最大长度50字符
	// 用于用户登录和显示
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`

	// Email 用户邮箱，唯一索引，不能为空，最大长度100字符
	// 用于用户登录和联系
	Email string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`

	// Password 用户密码，经过bcrypt加密存储，不能为空，最大长度255字符
	// json:"_"表示不序列化到JSON响应中，避免密码泄露
	Password string `gorm:"type:varchar(255);not null" json:"_"`

	// Roles 用户拥有的角色列表（多对多关系）
	// 注意：这里暂时不添加，等关联表创建后再添加
	// Roles []Role `gorm:"many2many:user_roles;" json:"roles,omitempty"`

	// CreatedAt 用户创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 用户更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
