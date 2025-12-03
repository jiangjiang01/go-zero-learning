package logic

import (
	"context"

	"go-zero-learning/model"
	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

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

// 根据 ID 查询用户基本信息（不返回密码）
func (l *GetUserLogic) GetUser(in *userrpc.GetUserReq) (*userrpc.GetUserResp, error) {
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID必须大于0")
	}

	var user model.User
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&user, in.Id).Error; err != nil {
		l.Errorf("查询用户失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	return &userrpc.GetUserResp{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
