package logic

import (
	"context"
	"errors"
	"go-zero-learning/common/ctxdata"
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
	if req.Email == "" && req.Password == "" {
		return nil, errors.New("至少需要提供一个更新字段")
	}

	// 2. 从上下文中获取用户 ID （由中间件设置）
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok {
		return nil, errors.New("未找到用户信息")
	}

	// 3. 查询用户是否存在
	var user model.User
	if err = l.svcCtx.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errors.New("更新失败")
	}

	hasUpdate := false // 标记是否有实际更新（兼容用户更新邮箱或密码时，当前邮箱或密码与请求中的相同的情况）

	// 4. 提供了邮箱
	if req.Email != "" {
		email := strings.TrimSpace(req.Email)
		// 检查邮箱格式（简单验证）
		if !strings.Contains(email, "@") {
			return nil, errors.New("邮箱格式不正确")
		}
		// 检查邮箱是否已被其他用户使用
		var existingUser model.User
		err = l.svcCtx.DB.Where("email = ? AND id != ?", email, userID).
			First(&existingUser).Error
		if err == nil {
			return nil, errors.New("邮箱已被使用")
		}
		// ErrRecordNotFound 表示没有找到记录，可以使用，但是其他错误需要处理
		if err != gorm.ErrRecordNotFound {
			return nil, errors.New("更新失败")
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
		// 密码长度验证（至少6位）
		if len(password) < 6 {
			return nil, errors.New("密码至少需要6位")
		}
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			l.Errorf("密码加密失败：%v", err)
			return nil, errors.New("更新失败")
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
		return nil, errors.New("没有需要更新的字段")
	}

	// 7. 保存更新
	if err := l.svcCtx.DB.Save(user).Error; err != nil {
		l.Errorf("更新用户失败：%v", err)
		return nil, errors.New("更新失败")
	}

	// 8. 返回更新后的用户信息
	resp = &types.UserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}
