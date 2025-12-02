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
	// 生成 ID
	l.svcCtx.IdMutex.Lock()
	id := l.svcCtx.NextID
	l.svcCtx.NextID++
	l.svcCtx.IdMutex.Unlock()

	// 创建用户对象
	user := &userrpc.User{
		Id:       id,
		Username: in.Username,
		Email:    in.Email,
	}

	// 存储到内存
	l.svcCtx.UserMutex.Lock()
	l.svcCtx.UserStore[id] = user
	l.svcCtx.UserMutex.Unlock()

	l.Infof("创建用户成功：ID=%d, Username=%s, Email=%s", id, in.Username, in.Email)

	return &userrpc.CreateUserResp{
		User: user,
	}, nil
}
