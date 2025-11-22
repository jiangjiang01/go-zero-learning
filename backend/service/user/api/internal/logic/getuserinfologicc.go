// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.UserInfoResp, err error) {
	// 1. 从上下文中获取用户 ID（由中间件设置）
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok {
		return nil, errorx.ErrNoUserInfo
	}

	// 2. 查询用户信息
	var user model.User
	err = l.svcCtx.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户信息失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 返回用户信息
	resp = &types.UserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}
