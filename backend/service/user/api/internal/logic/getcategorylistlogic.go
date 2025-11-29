// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListReq) (resp *types.GetCategoryListResp, err error) {

	// 1. 默认值设置
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 2. 构建查询条件
	query := l.svcCtx.DB.Model(&model.Category{})

	// 搜索关键词
	if req.Keyword != "" {
		likeStr := "%" + req.Keyword + "%"
		query = query.Where("name LIKE ?", likeStr)
	}

	// 3. 是否查询全部
	if req.All {
		var categories []model.Category
		err = query.Order("sort ASC, created_at DESC").Find(&categories).Error
		if err != nil {
			l.Errorf("查询分类列表失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		categoryList := l.convertToCategoryInfoResp(categories)
		total := int64(len(categories))
		resp = &types.GetCategoryListResp{
			Categories: categoryList,
			Total:      total,
			Page:       int64(1),
			PageSize:   int64(total),
		}
		return resp, nil
	}

	// 4. 查询总数
	var total int64
	err = query.Count(&total).Error
	if err != nil {
		l.Errorf("查询分类总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 分页查询
	offset := (req.Page - 1) * req.PageSize
	var categories []model.Category
	err = query.Order("sort ASC, created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&categories).Error
	if err != nil {
		l.Errorf("查询分类列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 构建响应
	categoryList := l.convertToCategoryInfoResp(categories)
	resp = &types.GetCategoryListResp{
		Categories: categoryList,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
	}

	return resp, nil
}

func (l *GetCategoryListLogic) convertToCategoryInfoResp(categories []model.Category) []types.CategoryInfoResp {
	categoryList := make([]types.CategoryInfoResp, 0, len(categories))

	for _, category := range categories {
		categoryList = append(categoryList, types.CategoryInfoResp{
			ID:        category.ID,
			Name:      category.Name,
			Desc:      category.Desc,
			ParentID:  category.ParentID,
			Sort:      category.Sort,
			Status:    category.Status,
			CreatedAt: category.CreatedAt.Unix(),
			UpdatedAt: category.UpdatedAt.Unix(),
		})
	}

	return categoryList
}
