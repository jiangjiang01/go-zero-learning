// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionDetailLogic {
	return &GetPermissionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionDetailLogic) GetPermissionDetail(req *types.GetPermissionDetailReq) (resp *types.PermissionInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
