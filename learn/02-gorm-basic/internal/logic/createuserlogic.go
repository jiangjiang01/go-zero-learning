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

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.UserResp, err error) {
	// 创建用户
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
	}

	// 使用 GORM 插入数据
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		l.Errorf("创建用户失败：%v", err)
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
