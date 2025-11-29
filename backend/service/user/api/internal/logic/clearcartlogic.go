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

type ClearCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 清空购买车
func (l *ClearCartLogic) ClearCart(req *types.ClearCartReq) (resp *types.ClearCartResp, err error) {
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
			// return nil, errorx.NewBusinessError(errorx.CodeNotFound, "购物车不存在")
			// 购物车不存在，直接返回成功
			return &types.ClearCartResp{
				Message: "购物车清空成功",
			}, nil
		}
		l.Errorf("查询购物车失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 删除所有购物车项
	err = l.svcCtx.DB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{}).Error
	if err != nil {
		l.Errorf("删除购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 构建响应
	resp = &types.ClearCartResp{
		Message: "购物车清空成功",
	}

	return resp, nil
}
