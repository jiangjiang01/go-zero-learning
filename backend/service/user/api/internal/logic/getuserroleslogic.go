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

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.GetUserRolesReq) (resp *types.GetUserRolesResp, err error) {
	// 1. 验证用户是否存在
	var user model.User
	err = l.svcCtx.DB.First(&user, req.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 2. 查询用户所有角色
	// 没有角色返回空列表，继续收集不会影响后续逻辑
	var userRoles []model.UserRole
	err = l.svcCtx.DB.Where("user_id = ?", req.UserID).Find(&userRoles).Error
	if err != nil {
		l.Errorf("查询用户角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 收集角色 ID
	roleIDs := make([]int64, 0, len(userRoles))
	for _, role := range userRoles {
		roleIDs = append(roleIDs, role.RoleID)
	}

	// 4. 查询角色详情
	var roles []model.Role
	if len(roleIDs) > 0 {
		err = l.svcCtx.DB.Where("id IN ?", roleIDs).Find(&roles).Error
		if err != nil {
			l.Errorf("查询角色详情失败：%v", err)
			return nil, errorx.ErrInternalError
		}
	}

	// 5. 构建响应
	resp = &types.GetUserRolesResp{
		Roles: convertToRoleInfoResp(roles),
	}

	// 6. 返回响应
	return resp, nil
}
