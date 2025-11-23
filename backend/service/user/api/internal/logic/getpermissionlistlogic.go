// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"strings"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionListLogic {
	return &GetPermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionListLogic) GetPermissionList(req *types.GetPermissionListReq) (resp *types.GetPermissionListResp, err error) {
	// 1. 参数校验
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 2. 构建查询条件
	db := l.svcCtx.DB.Model(&model.Permission{})

	// 搜索关键词
	if req.Keyword != "" {
		keyword := strings.TrimSpace(req.Keyword)
		likeStr := "%" + keyword + "%"
		db = db.Where("name LIKE ? OR code LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		l.Errorf("查询权限总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 分页查询
	offset := (req.Page - 1) * req.PageSize
	var permissions []model.Permission
	err = db.Offset(int(offset)).Limit(int(req.PageSize)).Find(&permissions).Error
	if err != nil {
		l.Errorf("查询权限列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	permissionList := convertToPermissionInfoResp(permissions)
	resp = &types.GetPermissionListResp{
		Permissions: permissionList,
		Total:       total,
		Page:        req.Page,
		PageSize:    req.PageSize,
	}

	// 6. 返回响应
	return resp, nil
}

func convertToPermissionInfoResp(permissions []model.Permission) []types.PermissionInfoResp {
	permissionList := make([]types.PermissionInfoResp, 0, len(permissions))

	for _, permission := range permissions {
		permissionList = append(permissionList, types.PermissionInfoResp{
			ID:        permission.ID,
			Name:      permission.Name,
			Code:      permission.Code,
			Desc:      permission.Desc,
			CreatedAt: permission.CreatedAt.Unix(),
			UpdatedAt: permission.UpdatedAt.Unix(),
		})
	}

	return permissionList
}
