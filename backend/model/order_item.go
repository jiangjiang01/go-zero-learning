package model

import "time"

// 订单项表
// 存储订单中每个商品的详细信息，一个订单可以有多个订单项
type OrderItem struct {
	// 主键
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// 关联订单（外键）
	OrderID int64 `gorm:"not null" json:"order_id"`

	// 关联商品（外键）
	ProductID int64 `gorm:"not null" json:"product_id"`

	// 商品快照信息（下单时的商品信息，避免商品信息变更影响历史订单）
	// 商品名称快照
	ProductName string `gorm:"type:varchar(100);not null" json:"product_name"`

	// 商品快照描述（可选）
	ProductDesc string `gorm:"type:text" json:"product_desc"`

	// 价格和数量信息
	Price int64 `gorm:"not null" json:"price"`

	// 购买数量
	Quantity int `gorm:"not null" json:"quantity"`

	// 小计金额（分），等于 Price * Quantity
	Amount int64 `gorm:"not null" json:"amount"`

	// 时间字段
	CreatedAt time.Time `json:"created_at"`

	// 关联关系
	Order   Order   `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
