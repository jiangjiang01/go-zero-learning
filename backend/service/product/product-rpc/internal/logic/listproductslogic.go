package logic

import (
	"context"

	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品列表（分页，支持 keyword 搜索）
func (l *ListProductsLogic) ListProducts(in *productrpc.ListProductReq) (*productrpc.ListProductResp, error) {
	// todo: add your logic here and delete this line

	return &productrpc.ListProductResp{}, nil
}
