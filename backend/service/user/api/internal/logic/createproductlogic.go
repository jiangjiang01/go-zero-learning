// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"strings"

	"go-zero-learning/common/consts"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/product/product-rpc/productrpc"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (resp *types.ProductInfoResp, err error) {
	// 1. 参数校验
	name := strings.TrimSpace(req.Name)
	description := strings.TrimSpace(req.Description)

	if name == "" {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品名称不能为空")
	}

	if len(name) > 100 {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品名称长度不能超过100个字符")
	}

	if req.Price < consts.MinProductPrice {
		return nil, errorx.ErrProductPriceTooLow
	}

	// 2. 调用 ProductRpc 创建商品
	rpcResp, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &productrpc.CreateProductReq{
		Name:        name,
		Description: description,
		Price:       req.Price,
		Status: func() int32 {
			if req.Status != nil {
				return int32(*req.Status)
			}
			return -1 // -1 表示使用默认值
		}(),
		Stock: func() int64 {
			if req.Stock != nil {
				return *req.Stock
			}
			return 0
		}(),
		Images: req.Images,
	})
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.CreateProduct", errorx.RpcErrorMapper{
			AlreadyExistsErr: errorx.ErrProductNameExists,
		})
		if rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 3. 构建响应
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

	// 4. 清除商品列表缓存（新增商品后，列表需要更新）
	l.clearProductListCache()

	return resp, nil
}

// clearProductListCache 清除所有商品列表缓存
func (l *CreateProductLogic) clearProductListCache() {
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
