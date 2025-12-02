package cron

import (
	"context"
	"go-zero-learning/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 订单取消任务
type OrderCancelJob struct {
	db     *gorm.DB
	logger logx.Logger
}

// 创建订单取消任务
func NewOrderCancelJob(db *gorm.DB) *OrderCancelJob {
	return &OrderCancelJob{
		db:     db,
		logger: logx.WithContext(context.Background()),
	}
}

// Run 执行任务：取消超时未支付的订单（30分钟）
func (j *OrderCancelJob) Run() {
	j.logger.Infof("开始执行订单取消任务...")

	// 计算30分钟前的时间
	timeoutTime := time.Now().Add(-30 * time.Minute)

	// 查询待支付且超过30分钟的订单
	var orders []model.Order
	err := j.db.Model(&model.Order{}).Where("status = ? AND created_at < ?", 1, timeoutTime).Find(&orders).Error
	if err != nil {
		j.logger.Errorf("查询超时订单失败：%v", err)
		return
	}

	// 批量取消订单
	cancelCount := 0
	for _, order := range orders {
		// 更新订单状态为已取消（5）
		err := j.db.Model(&order).Update("status", 5).Error
		if err != nil {
			j.logger.Errorf("订单取消失败：订单ID：%d，错误：%v", order.ID, err)
			continue
		}

		// TODO: 恢复库存（后续实现）
		cancelCount++
		j.logger.Infof("订单已自动取消：订单号=%s， 订单ID=%d", order.OrderNo, order.ID)
	}

	j.logger.Infof("订单取消任务完成，共取消 %d 个订单", cancelCount)
}
