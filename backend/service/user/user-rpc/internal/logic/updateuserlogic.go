package logic

import (
	"context"
	"errors"
	"strings"

	"go-zero-learning/common/validator"
	"go-zero-learning/model"
	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户（当前用户）
func (l *UpdateUserLogic) UpdateUser(in *userrpc.UpdateUserReq) (*userrpc.UpdateUserResp, error) {
	// 1. 参数校验
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID必须大于0")
	}
	if in.Email == "" && in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "至少需要提供一个更新字段")
	}

	// 2. 查询用户是否存在
	var user model.User
	if err := l.svcCtx.DB.First(&user, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	updateFields := make(map[string]interface{})

	// 3. 处理邮箱更新
	if in.Email != "" {
		email := strings.TrimSpace(in.Email)
		// 检查邮箱格式
		if err := validator.ValidateEmail(email); err != nil {
			return nil, status.Error(codes.InvalidArgument, "邮箱格式不正确")
		}

		// 检查邮箱是否已被其他用户使用
		var existingUser model.User
		err := l.svcCtx.DB.WithContext(l.ctx).
			Where("email = ? AND id != ?", email, in.Id).
			First(&existingUser).Error
		if err == nil {
			return nil, status.Error(codes.AlreadyExists, "邮箱已存在")
		}
		// ErrRecordNotFound 表示没有找到记录，可以使用，但是其他错误需要处理
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询邮箱失败：%v", err)
			return nil, status.Error(codes.Internal, "内部错误")
		}

		// 如果邮箱有变化，则更新邮箱
		if user.Email != email {
			updateFields["email"] = email
		}
	}

	// 4. 处理密码更新
	if in.Password != "" {
		password := strings.TrimSpace(in.Password)
		// 用户密码强度验证
		if err := validator.ValidateUserPassword(password); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			l.Errorf("密码加密失败：%v", err)
			return nil, status.Error(codes.Internal, "内部错误")
		}
		// 如果密码有变化，则更新密码
		newPassword := string(hashedPassword)
		if user.Password != newPassword {
			updateFields["password"] = newPassword
		}
	}

	// 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有需要更新的字段")
	}

	// 5. 执行更新
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&user).Updates(updateFields).Error; err != nil {
		l.Errorf("更新用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 7. 重新查询最新的数据
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&user, in.Id).Error; err != nil {
		l.Error("重新查询用户失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	return &userrpc.UpdateUserResp{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
