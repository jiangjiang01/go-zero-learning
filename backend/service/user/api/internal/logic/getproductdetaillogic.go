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

type GetProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductDetailLogic) GetProductDetail(req *types.GetProductDetailReq) (resp *types.ProductInfoResp, err error) {
	// 1. 校验参数
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品ID不能小于等于0")
	}

	// 2. 查询商品是否存在
	var product model.Product
	err = l.svcCtx.DB.First(&product, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrProductNotFound
		}
		l.Errorf("查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
	resp = convertToProductInfoResp(product)

	return resp, nil
}

func convertToProductInfoResp(product model.Product) *types.ProductInfoResp {
	return &types.ProductInfoResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}
}
