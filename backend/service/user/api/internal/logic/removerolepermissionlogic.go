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

type RemoveRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRolePermissionLogic {
	return &RemoveRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveRolePermissionLogic) RemoveRolePermission(req *types.RemoveRolePermissionReq) (resp *types.RemoveRolePermissionResp, err error) {
	// 1. 验证角色是否存在
	var role model.Role
	err = l.svcCtx.DB.First(&role, req.RoleID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		l.Errorf("查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 2. 验证权限是否存在
	var permission model.Permission
	err = l.svcCtx.DB.First(&permission, req.PermissionID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrPermissionNotFound
		}
		l.Errorf("查询权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 删除关联关系
	result := l.svcCtx.DB.Where("role_id = ? AND permission_id = ?", req.RoleID, req.PermissionID).
		Delete(&model.RolePermission{})
	if result.Error != nil {
		l.Errorf("移除权限失败：%v", result.Error)
		return nil, errorx.ErrInternalError
	}

	// 检查是否真的删除了
	if result.RowsAffected == 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "角色未分配权限")
	}

	// 4. 构建响应
	resp = &types.RemoveRolePermissionResp{
		Message: "权限移除成功",
	}

	return resp, nil
}
