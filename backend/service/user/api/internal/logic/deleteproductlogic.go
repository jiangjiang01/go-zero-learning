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

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	// 1. 校验参数
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品ID不能小于等于0")
	}

	// 2. 调用 ProductRpc 删除商品（默认使用硬删除）
	_, err = l.svcCtx.ProductRpc.DeleteProduct(l.ctx, &productrpcclient.DeleteProductReq{
		Id:         req.ID,
		HardDelete: true, // 硬删除
	})
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.DeleteProduct", errorx.RpcErrorMapper{
			NotFoundErr: errorx.ErrProductNotFound,
		})
		return nil, rpcErr
	}

	// 3. 清除商品列表缓存（删除商品后，列表需要更新）
	l.clearProductListCache()

	// 4. 构建响应结果
	resp = &types.DeleteProductResp{
		Message: "商品删除成功",
	}

	return resp, nil
}

// clearProductListCache 清除所有商品列表缓存
func (l *DeleteProductLogic) clearProductListCache() {
	pattern := "product:list:*"
	keys, err := l.svcCtx.Redis.KeysCtx(l.ctx, pattern)
	if err != nil {
		l.Infof("获取缓存键失败：%v", err)
		return
	}

	if len(keys) > 0 {
		for _, key := range keys {
			l.svcCtx.Redis.DelCtx(l.ctx, key)
		}
		l.Infof("已清除 %d 个商品列表缓存", len(keys))
	}
}
