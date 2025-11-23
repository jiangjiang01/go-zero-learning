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

type GetRoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleDetailLogic {
	return &GetRoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleDetailLogic) GetRoleDetail(req *types.GetRoleDetailReq) (resp *types.RoleInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 查询角色信息
	var role model.Role
	err = l.svcCtx.DB.First(&role, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		l.Errorf("查询角色失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
	resp = &types.RoleInfoResp{
		ID:        role.ID,
		Name:      role.Name,
		Code:      role.Code,
		Desc:      role.Desc,
		CreatedAt: role.CreatedAt.Unix(),
		UpdatedAt: role.UpdatedAt.Unix(),
	}

	// 4. 返回响应
	return resp, nil
}
