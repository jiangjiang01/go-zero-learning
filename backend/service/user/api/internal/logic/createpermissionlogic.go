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

type CreatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePermissionLogic) CreatePermission(req *types.CreatePermissionReq) (resp *types.PermissionInfoResp, err error) {
	// 1. 参数校验
	if req.Name == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "权限名称不能为空")
	}
	if req.Code == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "权限代码不能为空")
	}

	// 2. 检查重复-权限名称
	var existingPermission model.Permission
	err = l.svcCtx.DB.Where("name = ?", req.Name).First(&existingPermission).Error
	if err == nil {
		return nil, errorx.ErrPermissionNameExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询权限名称失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 检查重复-权限代码
	err = l.svcCtx.DB.Where("code = ?", req.Code).First(&existingPermission).Error
	if err == nil {
		return nil, errorx.ErrPermissionCodeExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询权限代码失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 创建权限
	permission := &model.Permission{
		Name: req.Name,
		Code: req.Code,
		Desc: req.Desc,
	}

	err = l.svcCtx.DB.Create(&permission).Error
	if err != nil {
		l.Errorf("创建权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 返回响应
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
