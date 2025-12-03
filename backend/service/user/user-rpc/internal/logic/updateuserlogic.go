package logic

import (
	"context"

	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户（当前用户）
func (l *UpdateUserLogic) UpdateUser(in *userrpc.UpdateUserReq) (*userrpc.UpdateUserResp, error) {
	// todo: add your logic here and delete this line

	return &userrpc.UpdateUserResp{}, nil
}
