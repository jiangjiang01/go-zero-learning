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
	// 1. 校验参数
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单ID不能为空")
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

	// 3. 处理父菜单验证
	// ParentID=0 表示这是顶级菜单（无父菜单），此时不需要验证父菜单
	// ParentID>0 表示有父菜单，需要验证父菜单是否存在
	// ParentID<0 已经在参数验证阶段被拒绝，不会执行到这里
	parentID := req.ParentID
	if parentID > 0 {
		// 验证指定的父菜单是否存在
		var parentMenu model.Menu
		err = l.svcCtx.DB.First(&parentMenu, parentID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "父菜单不存在")
			}
			l.Errorf("查询父菜单失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		// 防止直接循环引用：不能将自己的ID设置为父菜单ID
		// 注：只检查直接循环引用(A->A)，间接循环引用(A->B->A)暂不处理
		if parentID == req.ID {
			return nil, errorx.ErrMenuCircularRef
		}
	}

	// 4. 处理字段更新
	updateFields := make(map[string]interface{})

	// 处理菜单代码
	if req.Code != nil {
		code := strings.TrimSpace(*req.Code)
		if code == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单代码不能为空")
		}

		var existingMenu model.Menu
		err = l.svcCtx.DB.Where("code = ? AND id != ?", code, req.ID).First(&existingMenu).Error
		if err == nil {
			return nil, errorx.ErrMenuCodeExists
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询菜单代码失败：%v", err)
			return nil, errorx.ErrInternalError
		}

		if menu.Code != code {
			updateFields["code"] = code
		}
	}

	// 处理菜单名称更新
	if req.Name != nil {
		// 检查同级菜单是否存在同名菜单（非自己）
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单名称不能为空")
		}

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

		if menu.Name != name {
			updateFields["name"] = name
		}
	}

	// 处理菜单描述更新, 允许设置为空
	if req.Desc != nil {
		desc := *req.Desc
		if menu.Desc != desc {
			updateFields["desc"] = desc
		}
	}

	// 处理菜单路径更新, 允许设置为空
	if req.Path != nil {
		path := *req.Path
		if menu.Path != path {
			updateFields["path"] = path
		}
	}

	// 处理菜单图标更新, 允许设置为空
	if req.Icon != nil {
		icon := *req.Icon
		if menu.Icon != icon {
			updateFields["icon"] = icon
		}
	}

	// 处理菜单类型更新
	if req.Type != nil {
		typeValue := *req.Type
		if typeValue != 1 && typeValue != 2 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单类型只能是1（菜单）或2（按钮）")
		}
		if menu.Type != typeValue {
			updateFields["type"] = typeValue
		}
	}

	// 处理菜单排序更新
	if req.Sort != nil {
		sortValue := *req.Sort
		if menu.Sort != sortValue {
			updateFields["sort"] = sortValue
		}
	}

	// 处理菜单状态更新
	if req.Status != nil {
		statusValue := *req.Status
		if statusValue != 1 && statusValue != 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单状态只能是1（启用）或0（禁用）")
		}
		// 如果设置为禁用，检查是否有子菜单
		var childCount int64
		err = l.svcCtx.DB.Model(&model.Menu{}).Where("parent_id = ?", req.ID).Count(&childCount).Error
		if err != nil {
			l.Errorf("查询子菜单失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if childCount > 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单下有子菜单，不能禁用")
		}

		if menu.Status != statusValue {
			updateFields["status"] = statusValue
		}
	}

	// 6. 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, errorx.ErrMenuNoUpdateFields
	}

	// 7. 执行更新
	err = l.svcCtx.DB.Model(&menu).Updates(updateFields).Error
	if err != nil {
		l.Errorf("更新菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 重新查询最新的数据
	err = l.svcCtx.DB.First(&menu, req.ID).Error
	if err != nil {
		l.Errorf("重新查询菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 9. 构建响应
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

	return resp, nil
}
