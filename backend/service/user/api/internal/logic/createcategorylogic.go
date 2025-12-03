// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"strings"

	"go-zero-learning/common/consts"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) (resp *types.CategoryInfoResp, err error) {
	// 1. 校验参数
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类名称不能为空")
	}
	if len(name) > consts.MaxCategoryNameLength {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类名称长度不能超过100个字符")
	}

	// 2. 检查父分类是否有效
	if req.ParentID > 0 {
		var parentCategory model.Category
		err = l.svcCtx.DB.Where("id = ?", req.ParentID).First(&parentCategory).Error // fixed: should check id
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.ErrCategoryParentInvalid
			}
			l.Errorf("查询父分类失败：%v", err)
			return nil, errorx.ErrInternalError
		}
	}

	// 3. 检查同级分类名称是否存在
	var existingCategory model.Category
	err = l.svcCtx.DB.Where("name = ? AND parent_id = ?", name, req.ParentID).First(&existingCategory).Error
	if err == nil {
		// 同级分类名称已存在
		return nil, errorx.ErrCategoryNameExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询同级分类名称失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 创建分类
	status := 1 // 默认启用
	if req.Status != nil {
		status = *req.Status
	}

	sort := 0
	if req.Sort > 0 {
		sort = req.Sort
	}

	category := &model.Category{
		Name:     name,
		Desc:     strings.TrimSpace(req.Desc),
		ParentID: req.ParentID,
		Sort:     sort,
		Status:   status,
	}

	err = l.svcCtx.DB.Create(&category).Error
	if err != nil {
		l.Errorf("创建分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应
	resp = &types.CategoryInfoResp{
		ID:        category.ID,
		Name:      category.Name,
		Desc:      category.Desc,
		ParentID:  category.ParentID,
		Sort:      category.Sort,
		Status:    category.Status,
		CreatedAt: category.CreatedAt.Unix(),
		UpdatedAt: category.UpdatedAt.Unix(),
	}

	return resp, nil
}
