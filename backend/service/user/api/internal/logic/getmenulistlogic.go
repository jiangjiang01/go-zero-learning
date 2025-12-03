// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"strings"

	"go-zero-learning/common/consts"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.GetMenuListReq) (resp *types.GetMenuListResp, err error) {
	// 1. 参数校验和默认值设置
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = consts.DefaultPageSize
	}
	// 限制每页最大数量，防止过大查询
	if req.PageSize > consts.MaxPageSize {
		req.PageSize = consts.MaxPageSize
	}

	// 2. 构建查询条件
	db := l.svcCtx.DB.Model(&model.Menu{})

	// 搜索关键词（菜单名称或代码）
	if req.Keyword != "" {
		keyword := strings.TrimSpace(req.Keyword)
		likeStr := "%" + keyword + "%"
		db = db.Where("name LIKE ? OR code LIKE ?", likeStr, likeStr)
	}

	// 3. 如果 all=true，获取全部菜单（不分页）
	if req.All {
		var menus []model.Menu
		err = db.Order("sort ASC, id ASC").Find(&menus).Error
		if err != nil {
			l.Errorf("查询菜单列表失败：%v", err)
			return nil, errorx.ErrInternalError
		}

		// 构建响应结果
		menuList := l.convertToMenuInfoResp(menus)
		total := int64(len(menus))
		resp = &types.GetMenuListResp{
			Menus:    menuList,
			Total:    total,
			Page:     1,
			PageSize: total,
		}

		return resp, nil
	}

	// 4. 分页模式：查询总数
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		l.Errorf("查询菜单总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 分页查询
	offset := (req.Page - 1) * req.PageSize
	var menus []model.Menu
	err = db.Order("sort ASC, id ASC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&menus).Error
	if err != nil {
		l.Errorf("查询菜单列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 构建响应结果
	menuList := l.convertToMenuInfoResp(menus)
	resp = &types.GetMenuListResp{
		Menus:    menuList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 7. 返回响应
	return resp, nil
}

// 转换菜单模型为响应格式
func (l *GetMenuListLogic) convertToMenuInfoResp(menus []model.Menu) []types.MenuInfoResp {
	menuList := make([]types.MenuInfoResp, 0, len(menus))
	for _, menu := range menus {
		menuList = append(menuList, types.MenuInfoResp{
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
		})
	}
	return menuList
}
