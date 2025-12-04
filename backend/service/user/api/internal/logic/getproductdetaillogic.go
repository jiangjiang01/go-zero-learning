// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/product/product-rpc/productrpc"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

	// 2. 调用 ProductRpc.GetProduct
	rpcResp, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &productrpc.GetProductReq{
		Id: req.ID,
	})
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.GetProduct", errorx.RpcErrorMapper{
			NotFoundErr: errorx.ErrProductNotFound,
		})
		if rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 3. 构建响应结果
	resp = &types.ProductInfoResp{
		ID:          rpcResp.Id,
		Name:        rpcResp.Name,
		Description: rpcResp.Description,
		Price:       rpcResp.Price,
		Status:      int(rpcResp.Status),
		Stock:       rpcResp.Stock,
		Images:      rpcResp.Images,
		CreatedAt:   rpcResp.CreatedAt,
		UpdatedAt:   rpcResp.UpdatedAt,
	}

	return resp, nil
}

func convertToProductInfoResp(product model.Product) *types.ProductInfoResp {
	var images []string
	if product.Images != "" {
		json.Unmarshal([]byte(product.Images), &images)
	}
	if images == nil {
		images = []string{}
	}

	return &types.ProductInfoResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      product.Status,
		Stock:       product.Stock,
		Images:      images,
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}
}
