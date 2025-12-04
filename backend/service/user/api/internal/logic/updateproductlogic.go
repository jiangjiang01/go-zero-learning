// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/common/errorx"
	"go-zero-learning/service/product/product-rpc/productrpcclient"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (resp *types.ProductInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品ID不能小于等于0")
	}

	// 2. 构建 RPC 请求
	rpcReq := &productrpcclient.UpdateProductReq{
		Id: req.ID,
	}

	// 处理可选字段
	if req.Name != nil {
		rpcReq.Name = req.Name
	}
	if req.Description != nil {
		rpcReq.Description = req.Description
	}
	if req.Price != nil {
		rpcReq.Price = req.Price
	}
	if req.Status != nil {
		status := int32(*req.Status)
		rpcReq.Status = &status
	}
	if req.Stock != nil {
		rpcReq.Stock = req.Stock
	}

	// 处理图片更新
	// 如果 Images 字段不为空或为空数组，则标记需要更新图片
	if req.Images != nil {
		rpcReq.UpdateImages = true
		rpcReq.Images = req.Images
	}

	// 3. 调用 ProductRpc 更新商品
	rpcResp, err := l.svcCtx.ProductRpc.UpdateProduct(l.ctx, rpcReq)
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.UpdateProduct", errorx.RpcErrorMapper{
			NotFoundErr:      errorx.ErrProductNotFound,
			AlreadyExistsErr: errorx.ErrProductNameExists,
		})
		return nil, rpcErr
	}

	// 4. 清除商品列表缓存（数据已更新，需要清除缓存）
	l.clearProductListCache()

	// 5. 构建响应结果
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

func (l *UpdateProductLogic) clearProductListCache() {
	pattern := "product:list:*"
	keys, err := l.svcCtx.Redis.KeysCtx(l.ctx, pattern)
	if err != nil {
		l.Infof("获取缓存键失败：%v", err)
		return
	}

	// 批量删除缓存键
	if len(keys) > 0 {
		for _, key := range keys {
			l.svcCtx.Redis.DelCtx(l.ctx, key)
		}
		l.Infof("已清除 %d 个商品列表缓存", len(keys))
	}
}
