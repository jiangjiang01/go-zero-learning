// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/jwt"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户注册逻辑
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.LoginResp, err error) {
	// 1. 检查用户名是否已存在
	var existingUser model.User
	err = l.svcCtx.DB.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if err != gorm.ErrRecordNotFound {
		// 说明这是数据库查询错误（不是未找到记录的错误）
		l.Errorf("查询用户失败：%v", err)
		return nil, errors.New("注册失败")
	}

	// 2. 检查邮箱是否已存在
	err = l.svcCtx.DB.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("邮箱已存在")
	}
	if err != gorm.ErrRecordNotFound {
		l.Errorf("查询邮箱失败：%v", err)
		return nil, errors.New("注册失败")
	}

	// 3. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("密码加密失败：%v", err)
		return nil, errors.New("注册失败")
	}

	// 4. 创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = l.svcCtx.DB.Create(user).Error
	if err != nil {
		l.Errorf("创建用户失败：%v", err)
		return nil, errors.New("注册失败")
	}

	// 5. 生成 Token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		l.Errorf("生成 Token 失败：%v", err)
		return nil, errors.New("注册失败")
	}

	// 6. 返回响应
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
