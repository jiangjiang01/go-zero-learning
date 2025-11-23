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

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp *types.DeleteMenuResp, err error) {
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

	// 3. 检查是否有子菜单
	var childCount int64
	err = l.svcCtx.DB.Model(&model.Menu{}).Where("parent_id = ?", req.ID).Count(&childCount).Error
	if err != nil {
		l.Errorf("检查子菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}
	if childCount > 0 {
		return nil, errorx.ErrMenuHasChildren
	}

	// 4. 删除菜单
	err = l.svcCtx.DB.Delete(&menu).Error
	if err != nil {
		l.Errorf("删除菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	resp = &types.DeleteMenuResp{
		Message: "菜单删除成功",
	}

	// 6. 返回响应
	return resp, nil
}
