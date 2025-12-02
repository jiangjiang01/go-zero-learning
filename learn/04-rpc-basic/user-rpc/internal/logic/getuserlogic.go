package logic

import (
	"context"
	"fmt"

	"user-rpc/internal/svc"
	"user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 参数验证
	if in.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户 ID 必须大于 0")
	}

	// 从内存中查找用户
	l.svcCtx.UserMutex.RLock()
	user, exists := l.svcCtx.UserStore[in.Id]
	l.svcCtx.UserMutex.RUnlock()

	if !exists {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("用户不存在：ID=%d", in.Id))
	}

	l.Infof("查询用户成功：ID=%d, Username=%s, Email=%s", in.Id, user.Username, user.Email)

	return &userrpc.GetUserResp{
		User: user,
	}, nil
}
