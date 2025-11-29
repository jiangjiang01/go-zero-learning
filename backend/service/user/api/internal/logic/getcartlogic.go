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

// 获取购物车逻辑
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

// 获取购物车项中的所有商品ID
func (l *GetCartLogic) getProductIDs(cartItems []model.CartItem) []int64 {
	productIDs := make([]int64, 0, len(cartItems))

	for _, cartItem := range cartItems {
		productIDs = append(productIDs, cartItem.ProductID)
	}

	return productIDs
}

// 获取购物车项列表和总金额
func (l *GetCartLogic) getCartItemListAndTotalAmount(cartItems []model.CartItem) ([]types.CartItemResp, int64, error) {
	// 空列表提前返回
	if len(cartItems) == 0 {
		return []types.CartItemResp{}, 0, nil
	}

	// 收集所有的商品ID
	productIDs := l.getProductIDs(cartItems)

	// 一次性批量查询所有商品（解决 N+1 查询问题）
	var products []model.Product
	err := l.svcCtx.DB.Where("id IN ? AND status = 1", productIDs).Find(&products).Error
	if err != nil {
		l.Errorf("查询商品信息失败：%v", err)
		return nil, 0, errorx.ErrInternalError
	}

	// 构建商品ID到商品的映射
	productMap := make(map[int64]model.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	// 构建响应
	var cartItemList []types.CartItemResp
	var totalAmount int64

	for _, cartItem := range cartItems {
		// 查询商品信息
		product, exists := productMap[cartItem.ProductID]
		if !exists {
			// 商品可能已经下架，跨过
			continue
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
