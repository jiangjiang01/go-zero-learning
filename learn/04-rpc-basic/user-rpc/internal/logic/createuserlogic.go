package logic

import (
	"context"

	"user-rpc/internal/svc"
	"user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *userrpc.CreateUserReq) (*userrpc.CreateUserResp, error) {
	// todo: add your logic here and delete this line

	return &userrpc.CreateUserResp{}, nil
}
