package logic

import (
	"context"

	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 ID 查询用户基本信息（不返回密码）
func (l *GetUserLogic) GetUser(in *userrpc.GetUserReq) (*userrpc.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &userrpc.GetUserResp{}, nil
}
