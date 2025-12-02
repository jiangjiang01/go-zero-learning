package logic

import (
	"context"
	"fmt"

	"user-rpc/internal/svc"
	"user-rpc/userrpc"

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

func (l *GetUserLogic) GetUser(in *userrpc.GetUserReq) (*userrpc.GetUserResp, error) {
	// 从内存中查找用户
	l.svcCtx.UserMutex.RLock()
	user, exists := l.svcCtx.UserStore[in.Id]
	l.svcCtx.UserMutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("用户不存在：ID=%d", in.Id)
	}

	l.Infof("查询用户成功：ID=%d, Username=%s, Email=%s", in.Id, user.Username, user.Email)

	return &userrpc.GetUserResp{
		User: user,
	}, nil
}
