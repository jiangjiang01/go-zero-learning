// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RemoveUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserRoleLogic {
	return &RemoveUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserRoleLogic) RemoveUserRole(req *types.RemoveUserRoleReq) (resp *types.RemoveUserRoleResp, err error) {
	// 1. 验证用户是否存在
	var user model.User
	err = l.svcCtx.DB.First(&user, req.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 2. 验证角色是否存在
	var role model.Role
	err = l.svcCtx.DB.First(&role, req.RoleID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		l.Errorf("查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 删除关联关系
	result := l.svcCtx.DB.Where("user_id = ? AND role_id = ?", req.UserID, req.RoleID).
		Delete(&model.UserRole{})
	if result.Error != nil {
		l.Errorf("移除角色失败：%v", result.Error)
		return nil, errorx.ErrInternalError
	}

	// 4. 检查是否真的删除了（简单版本：不检查）
	if result.RowsAffected == 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "用户未分配该角色")
	}

	// 5. 构建响应
	resp = &types.RemoveUserRoleResp{
		Message: "角色移除成功",
	}

	return resp, nil
}
