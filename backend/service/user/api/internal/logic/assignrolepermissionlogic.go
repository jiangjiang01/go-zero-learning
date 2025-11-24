// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/tmp/internal/svc"
	"go-zero-learning/service/user/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolePermissionLogic {
	return &AssignRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRolePermissionLogic) AssignRolePermission(req *types.AssignRolePermissionReq) (resp *types.AssignRolePermissionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
