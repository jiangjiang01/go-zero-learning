// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.GetOrderDetailReq) (resp *types.OrderInfoResp, err error) {
	// 1. 获取当前用户ID
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "订单ID格式不正确")
	}

	// 3. 查询当前用户的订单详情
	var order model.Order
	err = l.svcCtx.DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrOrderNotFound
		}
		l.Errorf("查询订单详情失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 查询订单项
	var orderItems []model.OrderItem
	err = l.svcCtx.DB.Order("created_at ASC").Where("order_id = ?", req.ID).Find(&orderItems).Error
	if err != nil {
		l.Errorf("查询订单项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	resp = &types.OrderInfoResp{
		ID:          order.ID,
		OrderNo:     order.OrderNo,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		StatusText:  getOrderStatusText(order.Status),
		Remark:      order.Remark,
		Items:       convertToOrderItemRespList(orderItems),
		CreatedAt:   order.CreatedAt.Unix(),
		UpdatedAt:   order.UpdatedAt.Unix(),
	}

	return resp, nil
}

func convertToOrderItemRespList(orderItems []model.OrderItem) []types.OrderItemResp {
	var orderItemList []types.OrderItemResp

	for _, item := range orderItems {
		orderItemList = append(orderItemList, types.OrderItemResp{
			ID:          item.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			ProductDesc: item.ProductDesc,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Amount:      item.Amount,
		})
	}

	return orderItemList
}
