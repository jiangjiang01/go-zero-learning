package logic

import (
	"context"
	"errors"
	"go-zero-learning/common/errorx"
	"go-zero-learning/common/validator"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UpdateUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserByIdLogic {
	return &UpdateUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserByIdLogic) UpdateUserById(req *types.UpdateUserByIdReq) (resp *types.UserInfoResp, err error) {
	// 1. 检查是否有需要更新的字段
	if req.Email == "" && req.Password == "" {
		return nil, errorx.ErrNoUpdateFields
	}

	// 2. 路径参数中获取用户ID
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 3. 查询用户是否存在
	var user model.User
	if err = l.svcCtx.DB.First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	hasUpdate := false // 标记是否有实际更新（兼容用户更新邮箱或密码时，当前邮箱或密码与请求中的相同的情况）

	// 4. 提供了邮箱
	if req.Email != "" {
		email := strings.TrimSpace(req.Email)
		// 检查邮箱格式
		if err = validator.ValidateEmail(email); err != nil {
			return nil, err
		}

		// 检查邮箱是否已被其他用户使用
		var existingUser model.User
		err = l.svcCtx.DB.Where("email = ? AND id != ?", email, req.ID).
			First(&existingUser).Error
		if err == nil {
			return nil, errorx.ErrEmailExists
		}
		// ErrRecordNotFound 表示没有找到记录，可以使用，但是其他错误需要处理
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询邮箱失败：%v", err)
			return nil, errorx.ErrInternalError
		}

		// 如果邮箱有变化，则更新邮箱
		if user.Email != email {
			hasUpdate = true
			user.Email = email
		}
	}

	// 5. 提供了密码
	if req.Password != "" {
		password := strings.TrimSpace(req.Password)
		// 用户密码强度验证
		if err = validator.ValidateUserPassword(password); err != nil {
			return nil, err
		}
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			l.Errorf("密码加密失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		// 如果密码有变化，则更新密码
		newPassword := string(hashedPassword)
		if user.Password != newPassword {
			user.Password = newPassword
			hasUpdate = true
		}
	}

	// 6. 检查是否有实际更新
	if !hasUpdate {
		return nil, errorx.ErrNoUpdateFields
	}

	// 7. 保存更新
	if err := l.svcCtx.DB.Save(&user).Error; err != nil {
		l.Errorf("更新用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 返回更新后的用户信息
	resp = &types.UserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}
