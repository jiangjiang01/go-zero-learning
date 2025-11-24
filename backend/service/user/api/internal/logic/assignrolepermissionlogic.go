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

type AssignRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolePermissionLogic {
	return &AssignRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRolePermissionLogic) AssignRolePermission(req *types.AssignRolePermissionReq) (resp *types.AssignRolePermissionResp, err error) {
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

	// 3. 检查是否已经分配过（简单版本：不检查，直接插入，可能会有重复）
	// 问题：这里没有检查重复，可能会插入重复数据
	rolePermission := &model.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}

	err = l.svcCtx.DB.Create(&rolePermission).Error
	if err != nil {
		l.Errorf("分配权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 返回响应
	resp = &types.AssignRolePermissionResp{
		Message: "权限分配成功",
	}

	return resp, nil
}
