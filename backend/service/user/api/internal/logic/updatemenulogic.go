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

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (resp *types.MenuInfoResp, err error) {
	name := strings.TrimSpace(req.Name)
	code := strings.TrimSpace(req.Code)

	// 1. 校验参数
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单ID不能为空")
	}
	if name == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单名称不能为空")
	}
	if code == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单代码不能为空")
	}
	if req.Type != 1 && req.Type != 2 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单类型只能是1（菜单）或2（按钮）")
	}
	if req.Status != 0 && req.Status != 1 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "状态只能是0（禁用）或1（启用）")
	}
	if req.ParentID < 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "父菜单ID不能为负数")
	}

	// 2. 检查菜单是否存在
	var menu model.Menu
	err = l.svcCtx.DB.First(&menu, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMenuNotFound
		}
		l.Errorf("查询菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 如果提供了 ParentID，验证父菜单是否存在（到这儿不可能是负数了）
	parentID := req.ParentID
	if parentID > 0 {
		var parentMenu model.Menu
		err = l.svcCtx.DB.First(&parentMenu, parentID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "父菜单不存在")
			}
			l.Errorf("查询父菜单失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		// 防止循环引用：不能将父菜单设置为自己的子菜单
		if parentID == req.ID {
			return nil, errorx.ErrMenuCircularRef
		}
	}

	// 4. 检查菜单代码是否已存在（非自己）
	var existingMenu model.Menu
	err = l.svcCtx.DB.Where("code = ? AND id != ?", code, req.ID).First(&existingMenu).Error
	if err == nil {
		return nil, errorx.ErrMenuCodeExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询菜单代码失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 检查同级菜单是否存在同名菜单（非自己）
	var sameLevelCount int64
	err = l.svcCtx.DB.Model(&model.Menu{}).Where("parent_id = ? AND name = ? AND id != ?", parentID, name, req.ID).
		Count(&sameLevelCount).Error
	if err != nil {
		l.Errorf("查询同级菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}
	if sameLevelCount > 0 {
		return nil, errorx.ErrMenuAlreadyExists
	}

	// 6. 如果设置为禁用，检查是否有子菜单
	if req.Status == 0 && menu.Status != 0 {
		var childCount int64
		err = l.svcCtx.DB.Model(&model.Menu{}).Where("parent_id = ?", req.ID).Count(&childCount).Error
		if err != nil {
			l.Errorf("查询子菜单失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if childCount > 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单下有子菜单，不能禁用")
		}
	}

	// 7. 直接更新所有字段（使用 Save 方法）
	menu.Name = name
	menu.Code = code
	menu.Desc = req.Desc
	menu.ParentID = parentID
	menu.Path = req.Path
	menu.Icon = req.Icon
	menu.Type = req.Type
	menu.Sort = req.Sort
	menu.Status = req.Status

	err = l.svcCtx.DB.Save(&menu).Error
	if err != nil {
		l.Errorf("更新菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 构建响应（不需要重新查询，menu 对象已经是最新的）
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

	// 9. 返回响应
	return resp, nil
}
