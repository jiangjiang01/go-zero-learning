package logic

import (
	"context"
	"errors"

	"go-zero-learning/model"
	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteUserLogic) DeleteUser(in *userrpc.DeleteUserReq) (*userrpc.DeleteUserResp, error) {
	// 1. 参数校验
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID必须大于0")
	}

	// 2. 查询用户是否存在
	var user model.User
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&user, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 3. 删除用户
	if err := l.svcCtx.DB.WithContext(l.ctx).Delete(&user).Error; err != nil {
		l.Errorf("删除用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	l.Infof("删除用户成功：ID=%d, Username=%s", user.ID, user.Username)
	return &userrpc.DeleteUserResp{
		Message: "用户删除成功",
	}, nil
}
