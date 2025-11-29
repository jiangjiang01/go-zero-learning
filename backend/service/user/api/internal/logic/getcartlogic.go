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

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO
// 性能问题（N+1 查询）：第90行在循环中查询商品，购物车项多时性能差
// 商品状态未检查：未检查商品是否上架（status = 1）

func (l *GetCartLogic) GetCart(req *types.GetCartReq) (resp *types.GetCartResp, err error) {
	// 1. 获取当前用户
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok || userID == 0 {
		return nil, errorx.ErrUnauthorized
	}

	// 2. 查询购物车
	var cart model.Cart
	err = l.svcCtx.DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 购物车不存在，返回空购物车
			return &types.GetCartResp{
				ID:          0,
				UserID:      userID,
				Items:       []types.CartItemResp{},
				TotalAmount: 0,
				ItemCount:   0,
			}, nil
		}
		l.Errorf("查询购物车失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 查询购物车项
	var cartItems []model.CartItem
	err = l.svcCtx.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error
	if err != nil {
		l.Errorf("查询购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 构建响应
	cartItemList, totalAmount, err := l.getCartItemListAndTotalAmount(cartItems)
	if err != nil {
		return nil, err
	}
	resp = &types.GetCartResp{
		ID:          cart.ID,
		UserID:      userID,
		Items:       cartItemList,
		TotalAmount: totalAmount,
		ItemCount:   len(cartItemList),
	}

	return resp, nil
}

func (l *GetCartLogic) getCartItemListAndTotalAmount(cartItems []model.CartItem) ([]types.CartItemResp, int64, error) {
	var cartItemList []types.CartItemResp
	var totalAmount int64

	for _, cartItem := range cartItems {
		// 查询商品信息
		var product model.Product
		err := l.svcCtx.DB.Where("id = ?", cartItem.ProductID).First(&product).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 商品可能已删除，跳过
				continue
			}
			l.Errorf("查询商品信息失败：%v", err)
			return nil, 0, errorx.ErrInternalError
		}

		amount := int64(cartItem.Quantity) * product.Price
		totalAmount += amount

		cartItemList = append(cartItemList, types.CartItemResp{
			ID:          cartItem.ID,
			ProductID:   cartItem.ProductID,
			ProductName: product.Name,
			Price:       product.Price,
			Quantity:    cartItem.Quantity,
			Amount:      amount,
		})
	}

	return cartItemList, totalAmount, nil
}
