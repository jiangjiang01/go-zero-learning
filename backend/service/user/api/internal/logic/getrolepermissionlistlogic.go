// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/tmp/internal/svc"
	"go-zero-learning/service/user/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionListLogic {
	return &GetRolePermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionListLogic) GetRolePermissionList(req *types.GetRolePermissionListReq) (resp *types.GetRolePermissionListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
