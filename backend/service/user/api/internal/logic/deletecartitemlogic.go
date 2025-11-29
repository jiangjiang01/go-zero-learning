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

type DeleteCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartItemLogic {
	return &DeleteCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 删除购物车项逻辑
func (l *DeleteCartItemLogic) DeleteCartItem(req *types.DeleteCartItemReq) (resp *types.DeleteCartItemResp, err error) {
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
		return nil, errorx.ErrInternalError
	}

	// 4. 删除购物车项
	err = l.svcCtx.DB.Delete(&cartItem).Error
	if err != nil {
		l.Errorf("删除购物车项失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应
	resp = &types.DeleteCartItemResp{
		Message: "购物车项删除成功",
	}

	return resp, nil
}
