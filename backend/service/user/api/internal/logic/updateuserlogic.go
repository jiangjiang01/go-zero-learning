package logic

import (
	"context"
	"errors"
	"go-zero-learning/common/ctxdata"
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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UserInfoResp, err error) {
	// 1. 检查是否有需要更新的字段
	if req.Email == nil && req.Password == nil {
		return nil, errorx.ErrNoUpdateFields
	}

	// 2. 从上下文中获取用户 ID （由中间件设置）
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok {
		return nil, errorx.ErrNoUserInfo
	}

	// 3. 查询用户是否存在
	var user model.User
	if err = l.svcCtx.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 处理字段更新
	updateFields := make(map[string]interface{})

	// 处理邮箱更新
	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)
		// 检查邮箱格式
		if err = validator.ValidateEmail(email); err != nil {
			return nil, err
		}

		// 检查邮箱是否已被其他用户使用
		var existingUser model.User
		err = l.svcCtx.DB.Where("email = ? AND id != ?", email, userID).
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
			updateFields["email"] = email
		}
	}

	// 处理密码更新
	if req.Password != nil {
		password := strings.TrimSpace(*req.Password)
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
			updateFields["password"] = newPassword
		}
	}

	// 5. 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, errorx.ErrNoUpdateFields
	}

	// 6. 执行更新
	if err := l.svcCtx.DB.Model(&user).Updates(updateFields).Error; err != nil {
		l.Errorf("更新用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 重新查询最新的数据
	if err := l.svcCtx.DB.First(&user, userID).Error; err != nil {
		l.Error("重新查询用户失败：%v", err)
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
