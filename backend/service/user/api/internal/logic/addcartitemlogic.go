// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AddCartItemLogic) AddCartItem(req *types.AddCartItemReq) (resp *types.CartItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
