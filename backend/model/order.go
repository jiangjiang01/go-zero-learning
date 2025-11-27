package model

import "time"

// 订单模型
type Order struct {
	ID     int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int64 `gorm:"not null" json:"user_id"`

	// 商品信息
	ProductID int64 `gorm:"not null" json:"product_id"`
	Quantity  int   `gorm:"not null" json:"quantity"`

	// 价格信息
	Price      int64 `gorm:"not null" json:"price"`
	TotalPrice int64 `gorm:"not null" json:"total_price"`

	// 订单状态：1-待支付，2-已支付，3-已发货，4-已完成，5-已取消
	Status int `gorm:"type:tinyint;default:1" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}
