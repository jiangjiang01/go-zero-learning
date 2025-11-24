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

type GetRolePermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionListLogic {
	return &GetRolePermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionListLogic) GetRolePermissionList(req *types.GetRolePermissionListReq) (resp *types.GetRolePermissionListResp, err error) {
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

	// 2. 查询角色所有权限
	var rolePermissions []model.RolePermission
	err = l.svcCtx.DB.Where("role_id = ?", req.RoleID).Find(&rolePermissions).Error
	if err != nil {
		l.Errorf("查询角色权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 收集权限 ID
	permissionIDs := make([]int64, 0, len(rolePermissions))
	for _, p := range rolePermissions {
		permissionIDs = append(permissionIDs, p.PermissionID)
	}

	// 4. 查询权限详情（列表）
	var permissions []model.Permission
	err = l.svcCtx.DB.Where("id IN ?", permissionIDs).Find(&permissions).Error
	if err != nil {
		l.Errorf("查询权限详情失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应
	resp = &types.GetRolePermissionListResp{
		Permissions: convertToPermissionInfoResp(permissions),
	}

	return resp, nil
}
