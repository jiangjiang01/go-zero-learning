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

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	// 1. 获取当前用户ID
	userID, ok := ctxdata.GetUserID(l.ctx)
	if userID == 0 || !ok {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 验证订单项不能为空
	if len(req.Items) == 0 {
		return nil, errorx.ErrOrderItemsEmpty
	}

	// 3. 使用事务创建订单创建
	var order model.Order
	var totalAmount int64
	var orderItems []model.OrderItem

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 3.1 验证每个订单项计算总金额
		totalAmount = 0
		orderItems = make([]model.OrderItem, 0, len(req.Items))

		for _, item := range req.Items {
			// 验证数量
			if item.Quantity <= 0 {
				return errorx.ErrOrderQuantityInvalid
			}

			// 验证商品是否存在
			var product model.Product
			err = tx.Where("id = ? AND status = 1", item.ProductID).
				First(&product).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errorx.ErrProductNotFound
				}
				l.Errorf("查询商品失败：%v", err)
				return errorx.ErrInternalError
			}

			// 计算小计金额
			amount := product.Price * int64(item.Quantity)
			totalAmount += amount

			// 创建订单项(先不保存，等订单创建后再保存)
			orderItem := model.OrderItem{
				ProductID:   item.ProductID,
				ProductName: product.Name,
				ProductDesc: product.Description,
				Price:       product.Price,
				Quantity:    item.Quantity,
				Amount:      amount,
			}
			orderItems = append(orderItems, orderItem)
		}

		// 3.2 创建订单
		order = model.Order{
			OrderNo:     model.GenerateOrderNo(),
			UserID:      userID,
			TotalAmount: totalAmount,
			Status:      1, // 待支付
			Remark:      req.Remark,
			// 【注意】不设置 OrderItems 字段，避免 GORM 尝试级联创建导致 product_id 为 0 的错误
		}

		err = tx.Create(&order).Error
		if err != nil {
			l.Errorf("创建订单失败：%v", err)
			return errorx.ErrInternalError
		}

		// 3.3 保存订单项（逐个创建以确保 ID 被正确填充）
		for i := range orderItems {
			orderItems[i].OrderID = order.ID
			// 添加 ProductID 验证 - 防止 product_id 为 0 导致数据库错误
			if orderItems[i].ProductID == 0 {
				l.Errorf("商品ID为0,item:%v", orderItems[i])
				return errorx.NewBusinessError(errorx.CodeInvalidParam, "商品ID无效")
			}

			// ==================== 修复方式对比 ====================
			// 问题原因：GORM 的 Create() 方法默认会跳过结构体中的零值字段，
			// 导致 order_id 和 product_id 这些整数类型的字段没有被插入到 SQL 语句中，
			// 从而引发数据库错误："Field 'product_id' doesn't have a default value"

			// ---------- 修复方式1：使用 Select 显式指定字段 ----------
			// 优点：明确指定要插入的字段，语义清晰
			// 缺点：如果字段很多，需要一个个列出来
			// err = tx.Select("order_id", "product_id", "product_name", "product_desc", "price", "quantity", "amount").
			//	Create(&orderItems[i]).Error

			// ---------- 修复方式2：使用 Omit 排除自增主键 ----------
			// 优点：代码更简洁，只需排除 ID 字段
			// 缺点：如果后续字段变化，不需要手动修改代码
			err = tx.Omit("id").Create(&orderItems[i]).Error

			// ---------- 原始问题代码（会失败）----------
			// err = tx.Create(&orderItems[i]).Error
			// 这样会导致 GORM 跳过零值字段，最终 product_id 没有被插入

			// =====================================================

			if err != nil {
				l.Errorf("创建订单项失败：%v, orderItem:%+v", err, orderItems[i])
				return errorx.ErrInternalError
			}
		}

		return nil
	})

	if err != nil {
		// 如果是业务错误，直接返回
		if _, ok := err.(*errorx.BusinessError); ok {
			return nil, err
		}
		// 其他错误统一返回内部错误
		return nil, errorx.ErrInternalError
	}

	// 4. 构建响应
	var itemsResp []types.OrderItemResp
	for _, item := range orderItems {
		itemsResp = append(itemsResp, types.OrderItemResp{
			ID:          item.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			ProductDesc: item.ProductDesc,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Amount:      item.Amount,
		})
	}

	resp = &types.CreateOrderResp{
		ID:          order.ID,
		OrderNo:     order.OrderNo,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		StatusText:  getOrderStatus(order.Status),
		Remark:      order.Remark,
		Items:       itemsResp,
		CreatedAt:   order.CreatedAt.Unix(),
	}

	return resp, nil
}

// 获取订单状态文本
func getOrderStatus(status int) string {
	switch status {
	case 1:
		return "待支付"
	case 2:
		return "已支付"
	case 3:
		return "已发货"
	case 4:
		return "已完成"
	case 5:
		return "已取消"
	default:
		return "未知状态"
	}
}
