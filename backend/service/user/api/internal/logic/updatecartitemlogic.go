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
func (l *UpdateCartItemLogic) UpdateCartItem(req *types.UpdateCartItemReq) (resp *types.CartItemResp, err error) {
	// 1. 获取当前用户
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 参数验证：数量必须大于0
	if req.Quantity <= 0 {
		return nil, errorx.ErrCartItemQuantityInvalid
	}

	// 3. 参数验证：数量不能过大（防止恶意刷单）
	maxQuantity := 999
	if req.Quantity > maxQuantity {
		return nil, errorx.ErrCartItemQuantityTooLarge
	}

	// 4. 查询购物车项
	var cartItem model.CartItem
	err = l.svcCtx.DB.Where("id = ?", req.ItemID).First(&cartItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCartItemNotFound
		}
		l.Errorf("查询购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 验证购物车是否属于当前用户
	var cart model.Cart
	err = l.svcCtx.DB.Where("id = ? AND user_id = ?", cartItem.CartID, userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrForbidden
		}
		l.Errorf("查询购物车失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 查询商品是否仍然可用
	var product model.Product
	err = l.svcCtx.DB.Where("id = ? AND status = 1", cartItem.ProductID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCartProductNotAvailable
		}
		l.Errorf("查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 更新数量
	cartItem.Quantity = req.Quantity
	err = l.svcCtx.DB.Save(&cartItem).Error
	if err != nil {
		l.Errorf("更新购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 构建响应
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
