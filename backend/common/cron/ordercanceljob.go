package cron

import (
	"context"
	"go-zero-learning/common/consts"
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
	timeoutTime := time.Now().Add(-consts.OrderCancelTimeout)

	// 查询待支付且超过30分钟的订单
	var orders []model.Order
	err := j.db.Where("status = ? AND created_at < ?", model.OrderStatusPending, timeoutTime).Find(&orders).Error
	if err != nil {
		j.logger.Errorf("查询超时订单失败：%v", err)
		return
	}

	// 批量取消订单
	cancelCount := 0
	for _, order := range orders {
		// 使用事务确保订单取消和库存恢复的原子性
		err := j.db.Transaction(func(tx *gorm.DB) error {
			// 1. 更新订单状态为已取消
			err := tx.Model(&order).Update("status", model.OrderStatusCancelled).Error
			if err != nil {
				return err
			}

			// 2. 查询订单所有的订单项
			var orderItems []model.OrderItem
			err = tx.Where("order_id = ?", order.ID).Find(&orderItems).Error
			if err != nil {
				return err
			}

			// 3. 恢复每个商品的库存
			for _, item := range orderItems {
				// 使用数据库原子操作恢复库存
				result := tx.Model(&model.Product{}).Where("id = ?", item.ProductID).
					Update("stock", gorm.Expr("stock + ?", item.Quantity))
				if result.Error != nil {
					j.logger.Errorf("恢复库存失败，商品ID：%d, 错误：%v", item.ProductID, result.Error)
					return result.Error
				}

				if result.RowsAffected == 0 {
					j.logger.Infof("商品不存在或已删除，商品ID：%d", item.ProductID)
					// 商品不存在不影响订单取消，继续处理
				}
			}

			return nil
		})
		if err != nil {
			j.logger.Errorf("订单取消失败：订单ID：%d，错误：%v", order.ID, err)
			continue
		}

		cancelCount++
		j.logger.Infof("订单已自动取消：订单号=%s， 订单ID=%d", order.OrderNo, order.ID)
	}

	j.logger.Infof("订单取消任务完成，共取消 %d 个订单", cancelCount)
}
