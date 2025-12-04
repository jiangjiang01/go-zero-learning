package logic

import (
	"context"

	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建商品
func (l *CreateProductLogic) CreateProduct(in *productrpc.CreateProductReq) (*productrpc.CreateProductResp, error) {
	// todo: add your logic here and delete this line

	return &productrpc.CreateProductResp{}, nil
}
