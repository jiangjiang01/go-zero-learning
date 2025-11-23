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

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.RoleInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 检查是否有更新的字段
	if req.Name == "" && req.Code == "" && req.Desc == "" {
		return nil, errorx.ErrRoleNoUpdateFields
	}

	// 3. 查询角色是否存在
	var role model.Role
	err = l.svcCtx.DB.First(&role, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		l.Errorf("查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 标记是否有实际更新（兼容角色名称或代码与请求中的相同的情况）
	hasUpdate := false

	// 4. 如果提供了角色名称，则查询角色名称是否已存在（非自己）
	if req.Name != "" {
		name := strings.TrimSpace(req.Name)
		var existingRole model.Role
		err = l.svcCtx.DB.Where("name = ? AND id != ?", name, req.ID).First(&existingRole).Error
		if err == nil {
			return nil, errorx.ErrRoleNameExists
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询角色名称失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if role.Name != name {
			hasUpdate = true
			role.Name = name
		}
	}

	// 5. 如果提供了角色代码，则查询角色代码是否已存在（非自己）
	if req.Code != "" {
		code := strings.TrimSpace(req.Code)
		var existingRole model.Role
		err = l.svcCtx.DB.Where("code = ? AND id != ?", code, req.ID).First(&existingRole).Error
		if err == nil {
			return nil, errorx.ErrRoleCodeExists
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询角色代码失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if role.Code != code {
			hasUpdate = true
			role.Code = code
		}
	}

	// 修改描述
	if req.Desc != "" {
		desc := strings.TrimSpace(req.Desc)
		// 描述允许调整为空
		if role.Desc != desc {
			hasUpdate = true
			role.Desc = desc
		}
	}

	// 6. 更新角色信息
	if !hasUpdate {
		return nil, errorx.ErrRoleNoUpdateFields
	}

	// 7. 保存
	err = l.svcCtx.DB.Save(&role).Error
	if err != nil {
		l.Errorf("更新角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 构建响应结果
	resp = &types.RoleInfoResp{
		ID:        role.ID,
		Name:      role.Name,
		Code:      role.Code,
		Desc:      role.Desc,
		CreatedAt: role.CreatedAt.Unix(),
		UpdatedAt: role.UpdatedAt.Unix(),
	}

	// 9. 返回响应
	return resp, nil
}
