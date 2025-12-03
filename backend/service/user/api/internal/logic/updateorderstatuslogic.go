// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"fmt"

	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateOrderStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(req *types.UpdateOrderStatusReq) (resp *types.OrderInfoResp, err error) {
	// 1. 获取当前用户ID
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "订单ID格式不正确")
	}
	if req.Status <= 0 || req.Status > 5 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "订单状态格式不正确")
	}

	// 3. 查询订单是否存在
	var order model.Order
	err = l.svcCtx.DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrOrderNotFound
		}
		l.Errorf("查询订单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 验证状态流转规则
	if !isValidStatusTransition(order.Status, req.Status) {
		return nil, errorx.NewBusinessError(errorx.CodeOrderStatusInvalid,
			getStatusTransitionError(order.Status, req.Status))
	}
	// 5. 更新订单状态
	order.Status = req.Status
	err = l.svcCtx.DB.Save(&order).Error
	if err != nil {
		l.Errorf("更新订单状态失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 重新查询最新订单
	err = l.svcCtx.DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&order).Error
	if err != nil {
		l.Errorf("重新查询订单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 查询订单子项
	var orderItems []model.OrderItem
	err = l.svcCtx.DB.Order("created_at DESC").Where("order_id = ?", req.ID).Find(&orderItems).Error
	if err != nil {
		l.Errorf("查询订单子项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 构建响应结果
	resp = convertToOrderInfoResp(order, orderItems)

	return resp, nil
}

// isValidStatusTransition 验证订单状态流转是否合法
func isValidStatusTransition(currentStatus, newStatus int) bool {
	// 相同状态，允许（幂等性）
	if currentStatus == newStatus {
		return true
	}
	switch currentStatus {
	case model.OrderStatusPending: // 待支付
		// 可以 支付，取消
		return newStatus == model.OrderStatusPaid || newStatus == model.OrderStatusCancelled
	case model.OrderStatusPaid: // 已支付
		// 可以：发货(3)、取消(5) - 注意：实际业务中已支付可能不能取消，这里先允许
		return newStatus == model.OrderStatusShipped || newStatus == model.OrderStatusCancelled
	case model.OrderStatusShipped: // 已发货
		// 可以 完成
		return newStatus == model.OrderStatusCompleted
	case model.OrderStatusCompleted: // 已完成
		// 不能改变
		return false
	case model.OrderStatusCancelled: // 已取消
		// 不能改变
		return false
	default:
		return false
	}
}

// getStatusTransitionError 获取状态流转错误信息
func getStatusTransitionError(currentStatus, newStatus int) string {
	currentText := getOrderStatusText(currentStatus)
	newText := getOrderStatusText(newStatus)
	return fmt.Sprintf("订单状态不能从 %s 直接流转到 %s", currentText, newText)
}
