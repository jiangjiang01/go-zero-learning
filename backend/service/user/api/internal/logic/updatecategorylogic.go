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

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) (resp *types.CategoryInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类ID无效")
	}

	// 2. 查询分类是否存在
	var category model.Category
	err = l.svcCtx.DB.Where("id = ?", req.ID).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCategoryNotFound
		}
		l.Errorf("查询分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建更新字段
	updateFields := make(map[string]interface{})

	// 处理分类名称更新
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类名称不能为空")
		}
		// 查询分类名称同级是否存在
		var sameLevelCount int64
		err = l.svcCtx.DB.Model(&model.Category{}).
			Where("name = ? AND parent_id = ? AND id != ?", name, category.ParentID, category.ID).
			Count(&sameLevelCount).Error
		if err != nil {
			l.Errorf("查询同级分类名称失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if sameLevelCount > 0 {
			return nil, errorx.ErrCategoryNameExists
		}
		if category.Name != name {
			updateFields["name"] = name
		}
	}

	// 处理分类描述更新
	if req.Desc != nil {
		desc := strings.TrimSpace(*req.Desc)
		if desc != category.Desc {
			updateFields["desc"] = desc
		}
	}

	// 处理父分类ID更新
	if req.ParentID != nil {
		parentID := *req.ParentID

		// 判断父分类是否存在
		var parentCategory model.Category
		err = l.svcCtx.DB.Where("id = ?", parentID).First(&parentCategory).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.ErrCategoryParentInvalid
			}
			l.Errorf("查询父分类失败：%v", err)
			return nil, errorx.ErrInternalError
		}

		// 循环引用检查函数(不能将父分类设置为自己的子分类)
		hasCircular, err := l.isCircularReference(parentID, req.ID)
		if err != nil {
			l.Errorf("循环引用检查失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		if hasCircular {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "不能将父分类设置为自己的子分类")
		}

		// 更新父分类ID
		if category.ParentID != parentID {
			updateFields["parent_id"] = parentID
		}
	}

	// 处理分类排序更新
	if req.Sort != nil {
		if category.Sort != *req.Sort {
			updateFields["sort"] = *req.Sort
		}
	}

	// 处理分类状态更新
	if req.Status != nil {
		status := *req.Status
		if status != 1 && status != 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类状态只能是1（启用）或0（禁用）")
		}
		if category.Status != status {
			updateFields["status"] = status
		}
	}

	// 4. 更新分类
	if len(updateFields) == 0 {
		return nil, errorx.ErrCategoryNoUpdateFields
	}

	err = l.svcCtx.DB.Model(&category).Updates(updateFields).Error
	if err != nil {
		l.Errorf("更新分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	// 重新查询
	err = l.svcCtx.DB.First(&category, req.ID).Error
	if err != nil {
		l.Errorf("重新查询分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

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

func (l *UpdateCategoryLogic) isCircularReference(parentID, categoryID int64) (bool, error) {
	// 顶级分类没有父分类，直接返回false
	if parentID == 0 {
		return false, nil
	}

	// 自己不能是自己的父类
	if parentID == categoryID {
		return true, nil
	}

	// 检查父分类的分类链(往上查找父分类链)，看是否包含当前分类
	currentID := parentID
	visited := make(map[int64]bool)
	maxDepth := 100 // 最大深度，防止无限循环

	for depth := 0; depth < maxDepth; depth++ {
		if currentID == categoryID {
			// 发现循环引用
			return true, nil
		}
		if visited[currentID] {
			// 已经访问过，避免重复查询
			break
		}

		visited[currentID] = true

		var parent model.Category
		err := l.svcCtx.DB.Where("id = ?", currentID).First(&parent).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 父分类不存在，无循环引用
				break
			}
			// 其他错误，返回错误
			return false, err
		}

		if parent.ParentID == 0 {
			// 顶级分类
			break
		}

		currentID = parent.ParentID
	}

	return false, nil
}
