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

type GetPermissionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionDetailLogic {
	return &GetPermissionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionDetailLogic) GetPermissionDetail(req *types.GetPermissionDetailReq) (resp *types.PermissionInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 查询权限信息
	var permission model.Permission
	err = l.svcCtx.DB.First(&permission, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrPermissionNotFound
		}
		l.Errorf("查询权限失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
	resp = &types.PermissionInfoResp{
		ID:        permission.ID,
		Name:      permission.Name,
		Code:      permission.Code,
		Desc:      permission.Desc,
		CreatedAt: permission.CreatedAt.Unix(),
		UpdatedAt: permission.UpdatedAt.Unix(),
	}

	// 4. 返回响应
	return resp, nil
}
