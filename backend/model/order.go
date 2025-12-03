package model

import (
	"fmt"
	"time"
)

// 订单状态常量
const (
	OrderStatusPending   = 1 // 待支付
	OrderStatusPaid      = 2 // 已支付
	OrderStatusShipped   = 3 // 已发货
	OrderStatusCompleted = 4 // 已完成
	OrderStatusCancelled = 5 // 已取消
)

// 订单模型
type Order struct {
	// 主键
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// 订单编号（用户看到的友好编号）
	// 格式：ORD + 年月日 + 6位序号，如：ORD20240127000001
	OrderNo string `gorm:"type:varchar(32);uniqueIndex;not null" json:"order_no"`

	// 关联用户（外键）
	UserID int64 `gorm:"not null" json:"user_id"`

	// 订单金额信息
	// 订单总金额（分），等于所有 OrderItem 的金额之和
	TotalAmount int64 `gorm:"not null" json:"total_amount"`

	// 订单状态
	// 1-待支付，2-已支付，3-已发货，4-已完成，5-已取消
	Status int `gorm:"type:tinyint;default:1" json:"status"`

	// 订单备注(可选)
	Remark string `gorm:"type:varchar(500)" json:"remark"`

	// 时间字段
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}

// 生成订单编号
// 格式：ORD + YYYYMMDD + 6位序号，如：ORD20240127000001
func GenerateOrderNo() string {
	now := time.Now()
	dateStr := now.Format("20060102") // YYYYMMDD
	// 这里简化处理，实际项目中应该用 Redis 或数据库序号
	timeStr := now.Format("150405") // HHMMSS 作为序号
	return fmt.Sprintf("ORD%s%s", dateStr, timeStr)
}
