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

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuReq) (resp *types.MenuInfoResp, err error) {
	// 1. 校验参数
	name := strings.TrimSpace(req.Name)
	code := strings.TrimSpace(req.Code)
	if name == "" || code == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单名称和代码不能为空")
	}
	if req.Type != 1 && req.Type != 2 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "菜单类型只能是1（菜单）或2（按钮）")
	}

	parentID := req.ParentID
	if req.ParentID < 0 {
		parentID = 0
	}

	// 2. 检查菜单代码是否已存在
	var existingMenu model.Menu
	err = l.svcCtx.DB.Where("code = ?", code).First(&existingMenu).Error
	if err == nil {
		return nil, errorx.ErrMenuCodeExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询菜单代码失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 如果提供了 ParentID，验证是否存在
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
	}

	// 4. 检查同级菜单是否存在同名菜单
	var sameLevelCount int64
	err = l.svcCtx.DB.Model(&model.Menu{}).Where("parent_id = ? AND name = ?", parentID, name).
		Count(&sameLevelCount).Error
	if err != nil {
		l.Errorf("查询同级菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}
	if sameLevelCount > 0 {
		return nil, errorx.ErrMenuAlreadyExists
	}

	// 5. 创建菜单
	menu := &model.Menu{
		Name:     name,
		Code:     code,
		Desc:     req.Desc,
		ParentID: parentID,
		Path:     req.Path,
		Icon:     req.Icon,
		Type:     req.Type,
		Sort:     req.Sort,
		Status:   1, // 默认启用
	}
	err = l.svcCtx.DB.Create(&menu).Error
	if err != nil {
		l.Errorf("创建菜单失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 构建响应
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

	// 7. 返回响应
	return resp, nil
}
