// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/errorx"
	"go-zero-learning/common/jwt"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

// 用户登录逻辑
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. 检查用户名是否存在
	var user model.User
	err = l.svcCtx.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrInvalidPassword
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 2. 校验密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errorx.ErrInvalidPassword
	}

	// 3. 生成 Token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		l.Errorf("生成 Token 失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 返回响应
	resp = &types.LoginResp{
		Token: token,
		UserInfo: types.UserInfoResp{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	return resp, nil
}
