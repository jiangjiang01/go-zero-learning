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

type AddCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO
// 1. 没有验证商品是否上架
// 2. 没有验证数量是否大于0
// 3. 没有验证数量是否合理（可能过大）

func (l *AddCartItemLogic) AddCartItem(req *types.AddCartItemReq) (resp *types.CartItemResp, err error) {
	// 1. 获取当前用户
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 验证商品是否存在
	var product model.Product
	err = l.svcCtx.DB.Where("id = ?", req.ProductID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrProductNotFound
		}
		l.Errorf("查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 查询或创建购物车
	var cart model.Cart
	err = l.svcCtx.DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新购物车
			cart = model.Cart{UserID: userID}
			err = l.svcCtx.DB.Create(&cart).Error
			if err != nil {
				l.Errorf("创建购物车失败：%v", err)
				return nil, errorx.ErrInternalError
			}
		} else {
			l.Errorf("查询购物车失败：%v", err)
			return nil, errorx.ErrInternalError
		}
	}

	// 4. 查找购物车项
	var cartItem model.CartItem
	err = l.svcCtx.DB.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建购物车项
			cartItem = model.CartItem{
				CartID:    cart.ID,
				ProductID: req.ProductID,
				Quantity:  req.Quantity,
			}
			err = l.svcCtx.DB.Create(&cartItem).Error
			if err != nil {
				l.Errorf("创建购物车项失败：%v", err)
				return nil, errorx.ErrInternalError
			}
		} else {
			l.Errorf("查询购物车项失败：%v", err)
			return nil, errorx.ErrInternalError
		}
	} else {
		// 更新数量
		cartItem.Quantity += req.Quantity
		err = l.svcCtx.DB.Save(&cartItem).Error
		if err != nil {
			l.Errorf("更新购物车项失败：%v", err)
			return nil, errorx.ErrInternalError
		}
	}

	// 5. 构建响应
	resp = &types.CartItemResp{
		ID:          cartItem.ID,
		ProductID:   cartItem.ProductID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    cartItem.Quantity,
		Amount:      int64(cartItem.Quantity) * product.Price,
	}

	return resp, nil
}
