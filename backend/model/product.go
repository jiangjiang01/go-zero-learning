package model

import "time"

// Product 商品模型
// 用于存储系统中的商品信息，包括名称、描述、价格、状态等
type Product struct {
	// 1. 主键放最前面
	// ID 商品唯一标识符，主键，自增
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// 2. 业务字段
	// Name 商品名称，唯一索引，不能为空，最大长度100字符
	// 用于商品展示和搜索，同一系统中商品名称不能重复
	Name string `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`

	// Description 商品描述，可以为空，用于详细说明商品特性
	// 使用 text 类型支持较长的描述内容
	Description string `gorm:"type:text;not null" json:"description"`

	// Price 商品价格，以分为单位存储，不能为空
	// 例如：1元 = 100分，避免浮点数精度问题
	// 前端显示时需要除以100转换为元
	Price int64 `gorm:"not null" json:"price"`

	// Status 商品状态，默认为1（启用）
	// 1: 启用（上架）, 0: 禁用（下架）
	Status int `gorm:"type:tinyint;default:1" json:"status"`

	// 3. 时间字段放最后（约定俗成）
	// CreatedAt 商品创建时间，自动记录
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 商品更新时间，自动记录
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}
