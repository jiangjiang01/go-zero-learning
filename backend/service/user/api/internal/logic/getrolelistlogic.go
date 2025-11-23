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

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.GetRoleListReq) (resp *types.GetRoleListResp, err error) {
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
	db := l.svcCtx.DB.Model(&model.Role{})

	// 搜索关键词
	if req.Keyword != "" {
		keyword := strings.TrimSpace(req.Keyword)
		likeStr := "%" + keyword + "%"
		db = db.Where("name LIKE ? OR code LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		l.Errorf("查询角色总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 分页查询
	offset := (req.Page - 1) * req.PageSize
	var roles []model.Role
	err = db.Offset(int(offset)).Limit(int(req.PageSize)).Find(&roles).Error
	if err != nil {
		l.Errorf("查询角色列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	roleList := convertToRoleInfoResp(roles)
	resp = &types.GetRoleListResp{
		Roles:    roleList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 6. 返回响应
	return resp, nil
}

func convertToRoleInfoResp(roles []model.Role) []types.RoleInfoResp {
	roleList := make([]types.RoleInfoResp, 0, len(roles))

	for _, role := range roles {
		roleList = append(roleList, types.RoleInfoResp{
			ID:        role.ID,
			Name:      role.Name,
			Code:      role.Code,
			Desc:      role.Desc,
			CreatedAt: role.CreatedAt.Unix(),
			UpdatedAt: role.UpdatedAt.Unix(),
		})
	}

	return roleList
}
