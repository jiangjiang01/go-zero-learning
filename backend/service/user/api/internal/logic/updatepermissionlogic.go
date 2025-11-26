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

	// 2. 查询权限是否存在
	var permission model.Permission
	err = l.svcCtx.DB.First(&permission, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrPermissionNotFound
		}
		l.Errorf("查询权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 处理字段更新
	updateFields := make(map[string]interface{})

	// 处理权限名称更新
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
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
			updateFields["name"] = name
		}
	}

	// 处理权限代码更新
	if req.Code != nil {
		code := strings.TrimSpace(*req.Code)
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
			updateFields["code"] = code
		}
	}

	// 处理描述更新
	if req.Desc != nil {
		desc := strings.TrimSpace(*req.Desc)
		// 描述允许调整为空
		if permission.Desc != desc {
			updateFields["desc"] = desc
		}
	}

	// 4. 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, errorx.ErrPermissionNoUpdateFields
	}

	// 5. 执行更新
	err = l.svcCtx.DB.Model(&permission).Updates(updateFields).Error
	if err != nil {
		l.Errorf("更新权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 重新查询最新的数据
	if err := l.svcCtx.DB.First(&permission, req.ID).Error; err != nil {
		l.Error("重新查询权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 构建响应结果
	resp = &types.PermissionInfoResp{
		ID:        permission.ID,
		Name:      permission.Name,
		Code:      permission.Code,
		Desc:      permission.Desc,
		CreatedAt: permission.CreatedAt.Unix(),
		UpdatedAt: permission.UpdatedAt.Unix(),
	}

	return resp, nil
}
