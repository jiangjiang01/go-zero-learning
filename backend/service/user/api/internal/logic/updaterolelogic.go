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

	// 2. 查询角色是否存在
	var role model.Role
	err = l.svcCtx.DB.First(&role, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		l.Errorf("查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建更新字段
	updateFields := make(map[string]interface{})
	// 处理名称更新
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "角色名称不能为空")
		}
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
			updateFields["name"] = name
		}
	}

	// 处理代码更新
	if req.Code != nil {
		code := strings.TrimSpace(*req.Code)
		if code == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "角色代码不能为空")
		}
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
			updateFields["code"] = code
		}
	}

	// 修改描述
	if req.Desc != nil {
		desc := *req.Desc
		// 描述允许调整为空
		if role.Desc != desc {
			updateFields["desc"] = desc
		}
	}

	// 4. 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, errorx.ErrRoleNoUpdateFields
	}

	// 5. 执行更新
	err = l.svcCtx.DB.Model(&role).Updates(updateFields).Error
	if err != nil {
		l.Errorf("更新角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 重新查询最新数据，确保 role 变量包含数据库最新的字段值（如 UpdatedAt 自动更新时间等）
	err = l.svcCtx.DB.First(&role, req.ID).Error
	if err != nil {
		l.Errorf("重新查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 构建响应结果
	resp = convertToRoleInfoResp(role)

	// 8. 返回响应
	return resp, nil
}

func convertToRoleInfoResp(role model.Role) *types.RoleInfoResp {
	return &types.RoleInfoResp{
		ID:        role.ID,
		Name:      role.Name,
		Code:      role.Code,
		Desc:      role.Desc,
		CreatedAt: role.CreatedAt.Unix(),
		UpdatedAt: role.UpdatedAt.Unix(),
	}
}
