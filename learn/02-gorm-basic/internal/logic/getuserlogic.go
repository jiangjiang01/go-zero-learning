// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"gorm-basic/internal/svc"
	"gorm-basic/internal/types"
	"gorm-basic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.UserResp, err error) {
	var user model.User

	// 使用 GORM 查询用户
	err = l.svcCtx.DB.First(&user, req.ID).Error
	if err != nil {
		l.Errorf("查询用户挫败：%v", err)
		return nil, err
	}

	// 返回响应
	resp = &types.UserResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}
