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

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.RoleInfoResp, err error) {
	// 1. 参数校验
	if req.Name == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "角色名称不能为空")
	}
	if req.Code == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "角色代码不能为空")
	}

	// 2. 检查重复-角色名称
	var existingRole model.Role
	err = l.svcCtx.DB.Where("name = ?", req.Name).First(&existingRole).Error
	if err == nil {
		return nil, errorx.ErrRoleNameExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询角色名称失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 检查重复-角色代码
	err = l.svcCtx.DB.Where("code = ?", req.Code).First(&existingRole).Error
	if err == nil {
		return nil, errorx.ErrRoleCodeExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询角色代码失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 创建角色
	role := &model.Role{
		Name: req.Name,
		Code: req.Code,
		Desc: req.Desc,
	}

	err = l.svcCtx.DB.Create(&role).Error
	if err != nil {
		l.Errorf("创建角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 返回响应
	resp = &types.RoleInfoResp{
		ID:        role.ID,
		Name:      role.Name,
		Code:      role.Code,
		Desc:      role.Desc,
		CreatedAt: role.CreatedAt.Unix(),
		UpdatedAt: role.UpdatedAt.Unix(),
	}

	return resp, nil
}
