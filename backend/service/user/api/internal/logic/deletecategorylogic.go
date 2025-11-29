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

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryReq) (resp *types.DeleteCategoryResp, err error) {
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

	// 3. 检查分类下是否存在子分类
	var childCount int64
	err = l.svcCtx.DB.Model(&model.Category{}).Where("parent_id = ?", req.ID).Count(&childCount).Error
	if err != nil {
		l.Errorf("检查子分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	if childCount > 0 {
		return nil, errorx.ErrCategoryHasChildren
	}

	// 4. 检查分类下是否存在商品
	var productCount int64
	err = l.svcCtx.DB.Model(&model.Product{}).Where("category_id = ?", req.ID).Count(&productCount).Error
	if err != nil {
		l.Errorf("查询分类对应商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	if productCount > 0 {
		return nil, errorx.ErrCategoryHasProducts
	}

	// 5. 删除分类
	err = l.svcCtx.DB.Delete(&category).Error
	if err != nil {
		l.Errorf("删除分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 构建响应结果
	resp = &types.DeleteCategoryResp{
		Message: "分类删除成功",
	}

	return resp, nil
}
