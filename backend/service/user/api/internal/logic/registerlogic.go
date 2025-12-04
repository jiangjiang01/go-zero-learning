// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/common/errorx"
	"go-zero-learning/common/validator"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
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
	// 1. 参数校验 - 邮箱格式
	if err = validator.ValidateEmail(req.Email); err != nil {
		return nil, err
	}

	// 2. 参数校验 - 用户密码强度
	if err = validator.ValidateUserPassword(req.Password); err != nil {
		return nil, err
	}

	// 3. 调用 user-rpc 创建用户
	rpcResp, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &userrpc.CreateUserReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		// 使用统一的错误映射函数
		if rpcErr := errorx.MapRpcError(err, l.Logger, "UserRpc.CreateUser", errorx.RpcErrorMapper{
			AlreadyExistsErr: errorx.ErrUserAlreadExists,
		}); rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 4. 生成 Token (API 层职责，RPC 不处理)
	token, err := l.svcCtx.JWT.GenerateToken(rpcResp.Id, rpcResp.Username)
	if err != nil {
		l.Errorf("生成 Token 失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 返回响应
	resp = &types.LoginResp{
		Token: token,
		UserInfo: types.UserInfoResp{
			ID:       rpcResp.Id,
			Username: rpcResp.Username,
			Email:    rpcResp.Email,
		},
	}

	return resp, nil
}

// 旧的注册逻辑，用于对比新旧逻辑
/*
func (l *RegisterLogic) RegisterOld(req *types.RegisterReq) (resp *types.LoginResp, err error) {
	// 1. 参数校验 - 邮箱格式
	if err = validator.ValidateEmail(req.Email); err != nil {
		return nil, err
	}

	// 2. 参数校验 - 用户密码强度
	if err = validator.ValidateUserPassword(req.Password); err != nil {
		return nil, err
	}

	// 3. 检查用户名是否已存在
	var existingUser model.User
	err = l.svcCtx.DB.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		return nil, errorx.ErrUsernameExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 说明这是数据库查询错误（不是未找到记录的错误）
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 检查邮箱是否已存在
	err = l.svcCtx.DB.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		return nil, errorx.ErrEmailExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询邮箱失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("密码加密失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = l.svcCtx.DB.Create(user).Error
	if err != nil {
		l.Errorf("创建用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 生成 Token
	token, err := l.svcCtx.JWT.GenerateToken(user.ID, user.Username)
	if err != nil {
		l.Errorf("生成 Token 失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 返回响应
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
*/
