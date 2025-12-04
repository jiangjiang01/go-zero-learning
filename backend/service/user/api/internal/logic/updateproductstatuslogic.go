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

type UpdateProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStatusLogic {
	return &UpdateProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductStatusLogic) UpdateProductStatus(req *types.UpdateProductStatusReq) (resp *types.UpdateProductStatusResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品ID不能小于等于0")
	}

	// 2. 调用 ProductRpc 更新商品状态
	_, err = l.svcCtx.ProductRpc.UpdateProductStatus(l.ctx, &productrpcclient.UpdateProductStatusReq{
		Id:     req.ID,
		Status: int32(req.Status),
	})
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.UpdateProductStatus", errorx.RpcErrorMapper{
			NotFoundErr: errorx.ErrProductNotFound,
		})
		return nil, rpcErr
	}

	// 3. 清除商品列表缓存（状态已更新，需要清除缓存）
	l.clearProductListCache()

	// 4. 构建响应结果
	resp = &types.UpdateProductStatusResp{
		Message: "商品状态更新成功",
	}

	return resp, nil
}

// clearProductListCache 清除所有商品列表缓存
func (l *UpdateProductStatusLogic) clearProductListCache() {
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
