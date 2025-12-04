package logic

import (
	"context"
	"encoding/json"

	"go-zero-learning/model"
	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 ID 查询商品详情
func (l *GetProductLogic) GetProduct(in *productrpc.GetProductReq) (*productrpc.GetProductResp, error) {
	// 1. 参数校验
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "商品ID必须大于0")
	}

	// 2. 查询商品
	var product model.Product
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&product, in.Id).Error; err != nil {
		l.Errorf("查询商品失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.NotFound, "商品不存在")
	}

	// 3. 解析图片列表
	var images []string
	if product.Images != "" {
		if err := json.Unmarshal([]byte(product.Images), &images); err != nil {
			l.Error("解析商品图片失败：id=%d, err=%v", in.Id, err)
			// 解析失败时返回空列表
			images = []string{}
		}
	}
	if images == nil {
		images = []string{}
	}

	// 4. 构建响应
	resp := &productrpc.GetProductResp{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      int32(product.Status),
		CategoryId:  product.CategoryID,
		Stock:       product.Stock,
		Images:      images,
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}

	return resp, nil
}
