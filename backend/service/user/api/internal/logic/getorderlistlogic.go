// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListReq) (resp *types.GetOrderListResp, err error) {
	// 1. 获取当前用户ID
	userId, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userId == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 参数校验和默认值设置
	if req.Page < 1 {
		req.Page = 1
	}

	if req.PageSize < 1 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 3. 构建查询模型
	query := l.svcCtx.DB.Model(&model.Order{}).Where("user_id = ?", userId)

	// 按状态筛选
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}

	// 搜索关键词
	if req.Keyword != "" {
		likeStr := "%" + req.Keyword + "%"
		query = query.Where("order_no LIKE ?", likeStr)
	}

	// 4. 查询总数
	var total int64
	err = query.Count(&total).Error
	if err != nil {
		l.Errorf("查询订单总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 空结果提前处理
	if total == 0 {
		return &types.GetOrderListResp{
			Orders:   []types.OrderInfoResp{},
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	// 5. 分页查询
	offset := (req.Page - 1) * req.PageSize
	var orders []model.Order
	err = query.Order("created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&orders).Error
	if err != nil {
		l.Errorf("查询订单列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 构建响应结果

	resp = &types.GetOrderListResp{
		Orders:   convertToOrderInfoRespList(orders),
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return
}

func convertToOrderInfoRespList(orders []model.Order) []types.OrderInfoResp {
	orderList := make([]types.OrderInfoResp, 0, len(orders))

	for _, order := range orders {
		orderList = append(orderList, types.OrderInfoResp{
			ID:          order.ID,
			OrderNo:     order.OrderNo,
			UserID:      order.UserID,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			StatusText:  getOrderStatusText(order.Status),
			Remark:      order.Remark,
			CreatedAt:   order.CreatedAt.Unix(),
			UpdatedAt:   order.UpdatedAt.Unix(),
		})
	}

	return orderList
}

// 获取订单状态文本
func getOrderStatusText(status int) string {
	switch status {
	case model.OrderStatusPending:
		return "待支付"
	case model.OrderStatusPaid:
		return "已支付"
	case model.OrderStatusShipped:
		return "已发货"
	case model.OrderStatusCompleted:
		return "已完成"
	case model.OrderStatusCancelled:
		return "已取消"
	default:
		return "未知状态"
	}
}
