// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRolePermissionLogic {
	return &RemoveRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveRolePermissionLogic) RemoveRolePermission(req *types.RemoveRolePermissionReq) (resp *types.RemoveRolePermissionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
