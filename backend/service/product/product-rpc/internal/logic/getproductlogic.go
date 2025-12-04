package logic

import (
	"context"

	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &productrpc.GetProductResp{}, nil
}
