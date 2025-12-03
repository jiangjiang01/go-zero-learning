package logic

import (
	"context"
	"errors"

	"go-zero-learning/model"
	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建用户
func (l *CreateUserLogic) CreateUser(in *userrpc.CreateUserReq) (*userrpc.CreateUserResp, error) {
	// 1. 参数校验
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "用户名不能为空")
	}
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "邮箱不能为空")
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "密码不能为空")
	}

	// 2. 检查用户名是否已存在
	var existingUser model.User
	err := l.svcCtx.DB.WithContext(l.ctx).Where("username = ?", in.Username).First(&existingUser).Error
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "用户名已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询用户名失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 3. 检查邮箱是否已存在
	err = l.svcCtx.DB.WithContext(l.ctx).Where("email = ?", in.Email).First(&existingUser).Error
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "邮箱已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询邮箱失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 4. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("密码加密失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 5. 创建用户
	user := &model.User{
		Username: in.Username,
		Email:    in.Email,
		Password: string(hashedPassword),
	}

	err = l.svcCtx.DB.WithContext(l.ctx).Create(user).Error
	if err != nil {
		l.Errorf("创建用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	l.Infof("创建用户成功：ID=%d, Username=%s, Email=%s", user.ID, user.Username, user.Email)
	return &userrpc.CreateUserResp{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
