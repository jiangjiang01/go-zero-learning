// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuDetailLogic {
	return &GetMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuDetailLogic) GetMenuDetail(req *types.GetMenuDetailReq) (resp *types.MenuInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
