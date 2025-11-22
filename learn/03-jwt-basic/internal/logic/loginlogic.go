// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"jwt-basic/common"
	"jwt-basic/internal/svc"
	"jwt-basic/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 简单的用户名密码验证（实际项目中应该查询数据库）
	// 这里为了演示，使用硬编码
	if req.Username == "admin" && req.Password == "123456" {
		// 生成 Token
		token, err := common.GenerateToken(1, req.Username)
		if err != nil {
			l.Errorf("生成 Token 失败: %v", err)
			return nil, err
		}

		resp = &types.LoginResp{
			Token: token,
		}
		return resp, nil
	}

	// 用户名或密码错误
	return nil, errors.New("用户名或密码错误")
}
