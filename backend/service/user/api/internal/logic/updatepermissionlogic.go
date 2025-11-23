// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"strings"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermissionLogic) UpdatePermission(req *types.UpdatePermissionReq) (resp *types.PermissionInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 检查是否有更新的字段
	if req.Name == "" && req.Code == "" && req.Desc == "" {
		return nil, errorx.ErrPermissionNoUpdateFields
	}

	// 3. 查询权限是否存在
	var permission model.Permission
	err = l.svcCtx.DB.First(&permission, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrPermissionNotFound
		}
		l.Errorf("查询权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 标记是否有实际更新（兼容权限名称或代码与请求中的相同的情况）
	hasUpdate := false

	// 4. 如果提供了权限名称，则查询权限名称是否已存在（非自己）
	if req.Name != "" {
		name := strings.TrimSpace(req.Name)
		var existingPermission model.Permission
		err = l.svcCtx.DB.Where("name = ? AND id != ?", name, req.ID).First(&existingPermission).Error
		if err == nil {
			return nil, errorx.ErrPermissionNameExists
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询权限名称失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if permission.Name != name {
			hasUpdate = true
			permission.Name = name
		}
	}

	// 5. 如果提供了权限代码，则查询权限代码是否已存在（非自己）
	if req.Code != "" {
		code := strings.TrimSpace(req.Code)
		var existingPermission model.Permission
		err = l.svcCtx.DB.Where("code = ? AND id != ?", code, req.ID).First(&existingPermission).Error
		if err == nil {
			return nil, errorx.ErrPermissionCodeExists
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询权限代码失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if permission.Code != code {
			hasUpdate = true
			permission.Code = code
		}
	}

	// 修改描述
	if req.Desc != "" {
		desc := strings.TrimSpace(req.Desc)
		// 描述允许调整为空
		if permission.Desc != desc {
			hasUpdate = true
			permission.Desc = desc
		}
	}

	// 6. 更新权限信息
	if !hasUpdate {
		return nil, errorx.ErrPermissionNoUpdateFields
	}

	// 7. 保存
	err = l.svcCtx.DB.Save(&permission).Error
	if err != nil {
		l.Errorf("更新权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 构建响应结果
	resp = &types.PermissionInfoResp{
		ID:        permission.ID,
		Name:      permission.Name,
		Code:      permission.Code,
		Desc:      permission.Desc,
		CreatedAt: permission.CreatedAt.Unix(),
		UpdatedAt: permission.UpdatedAt.Unix(),
	}

	// 9. 返回响应
	return resp, nil
}
