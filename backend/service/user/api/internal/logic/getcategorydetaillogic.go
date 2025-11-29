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

type GetCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDetailLogic {
	return &GetCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryDetailLogic) GetCategoryDetail(req *types.GetCategoryDetailReq) (resp *types.CategoryInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "分类ID无效")
	}

	// 2. 查询分类
	var category model.Category
	err = l.svcCtx.DB.Where("id = ?", req.ID).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCategoryNotFound
		}
		l.Errorf("查询分类失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
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
