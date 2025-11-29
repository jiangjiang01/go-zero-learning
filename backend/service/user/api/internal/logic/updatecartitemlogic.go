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

type UpdateCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 实现更新购物车项数量逻辑
// TODO
// 1. 没有验证数量是否大于 0
// 2. 没有验证数量是否合理（可能过大）
// 3. 没有验证商品是否上架

func (l *UpdateCartItemLogic) UpdateCartItem(req *types.UpdateCartItemReq) (resp *types.CartItemResp, err error) {
	// 1. 获取当前用户
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 查询购物车项
	var cartItem model.CartItem
	err = l.svcCtx.DB.Where("id = ?", req.ItemID).First(&cartItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.NewBusinessError(errorx.CodeNotFound, "购物车项不存在")
		}
		l.Errorf("查询购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 验证购物车是否属于当前用户
	var cart model.Cart
	err = l.svcCtx.DB.Where("id = ? AND user_id = ?", cartItem.CartID, userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.NewBusinessError(errorx.CodeNotFound, "购物车不存在")
		}
		l.Errorf("查询购物车失败：%v", err)
		return nil, err
	}

	// 4. 更新数量
	cartItem.Quantity = req.Quantity
	err = l.svcCtx.DB.Save(&cartItem).Error
	if err != nil {
		l.Errorf("更新购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 查询商品
	var product model.Product
	err = l.svcCtx.DB.Where("id = ?", cartItem.ProductID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrProductNotFound
		}
		l.Errorf("查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应
	amount := int64(cartItem.Quantity) * product.Price
	resp = &types.CartItemResp{
		ID:          cartItem.ID,
		ProductID:   cartItem.ProductID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    cartItem.Quantity,
		Amount:      amount,
	}

	return resp, nil
}
