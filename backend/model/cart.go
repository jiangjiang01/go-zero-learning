package model

import "time"

// Cart 购物车模型
// 每个用户只有一个购物车
type Cart struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	UserID int64 `gorm:"not null;uniqueIndex" json:"user_id"` // 用户ID，唯一索引

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	Items []CartItem `gorm:"foreignKey:CartID" json:"items,omitempty"`
}

func (Cart) TableName() string {
	return "carts"
}

// CartItem 购物车项模型
// 购物车的商品项
type CartItem struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	CartID int64 `gorm:"not null;index" json:"cart_id"` // 购物车ID

	ProductID int64 `gorm:"not null;" json:"product_id"` // 商品ID
	Quantity  int   `gorm:"not null" json:"quantity"`    // 购买数量

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	Cart    Cart    `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (CartItem) TableName() string {
	return "cart_items"
}
