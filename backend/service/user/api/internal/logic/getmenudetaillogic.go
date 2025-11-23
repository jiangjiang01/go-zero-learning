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

type GetMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuDetailLogic {
	return &GetMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuDetailLogic) GetMenuDetail(req *types.GetMenuDetailReq) (resp *types.MenuInfoResp, err error) {
	// 1. 校验参数
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 查询菜单信息
	var menu model.Menu
	err = l.svcCtx.DB.First(&menu, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMenuNotFound
		}
		l.Errorf("查询菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
	resp = &types.MenuInfoResp{
		ID:        menu.ID,
		Name:      menu.Name,
		Code:      menu.Code,
		Desc:      menu.Desc,
		ParentID:  menu.ParentID,
		Path:      menu.Path,
		Icon:      menu.Icon,
		Type:      menu.Type,
		Sort:      menu.Sort,
		Status:    menu.Status,
		CreatedAt: menu.CreatedAt.Unix(),
		UpdatedAt: menu.UpdatedAt.Unix(),
	}

	// 4. 返回响应
	return resp, nil
}
