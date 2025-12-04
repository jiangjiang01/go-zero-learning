// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"

	"go-zero-learning/common/consts"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/product/product-rpc/productrpc"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取商品列表接口
func (l *GetProductListLogic) GetProductList(req *types.GetProductListReq) (resp *types.GetProductListResp, err error) {
	// 1. 参数校验和默认值设置
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = consts.DefaultPageSize
	}
	if req.PageSize > consts.MaxPageSize {
		req.PageSize = consts.MaxPageSize
	}

	// 2. 构建缓存键
	// 缓存键规范化（防止键冲突和特殊字符问题）
	// 问题：
	// 特殊字符（空格、冒号、换行等）可能导致键无法正确匹配
	// 长关键词导致键过长，浪费内存
	// 不同编码的关键词可能生成相同的键（少见但可能）
	// 解决：使用 MD5 哈希对关键词进行处理，生成固定长度，安全的值
	var keywordHash string
	if req.Keyword != "" {
		// 使用 MD5 哈希处理关键词，生成固定长度的安全键
		hash := md5.Sum([]byte(req.Keyword))
		keywordHash = hex.EncodeToString(hash[:])[:8] // 取前8位，足够唯一
	} else {
		keywordHash = "empty"
	}
	cacheKey := fmt.Sprintf("product:list:page:%d:size:%d:keyword:%s", req.Page, req.PageSize, keywordHash)

	// 3. 尝试从缓存获取
	cacheData, err := l.svcCtx.Redis.GetCtx(l.ctx, cacheKey)
	if err == nil && cacheData != "" {
		// 缓存命中，反序列化数据
		var cacheResp types.GetProductListResp
		err = json.Unmarshal([]byte(cacheData), &cacheResp)
		if err == nil {
			l.Infof("从缓存中获取商品列表：%s", cacheKey)
			return &cacheResp, nil
		}
		// 如果反序列化失败，继续查询数据库
		l.Errorf("缓存数据反序列化失败，继续查询数据库：%v", err)
	}

	// 4. 从 ProductRpc 获取商品列表
	rpcResp, err := l.svcCtx.ProductRpc.ListProducts(l.ctx, &productrpc.ListProductReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		// 使用统一的错误映射函数
		rpcErr := errorx.MapRpcError(err, l.Logger, "ProductRpc.ListProducts", errorx.RpcErrorMapper{})
		if rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 4.3 空结果提前处理
	if rpcResp.Total == 0 && len(rpcResp.Products) == 0 {
		emptyResp := &types.GetProductListResp{
			Products: []types.ProductInfoResp{},
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}

		// 即使是空结果也缓存（防止缓存穿透）
		data, _ := json.Marshal(emptyResp)
		// 防止缓存雪崩：避免大量缓存同时过期，给过期时间加入随机偏移
		// 随机过期时间范围：60-90秒，避免同时过期
		randomExpire := consts.CacheEmptyResultTTL + rand.Intn(consts.CacheEmptyRandomRange)
		_ = l.svcCtx.Redis.SetexCtx(l.ctx, cacheKey, string(data), randomExpire)

		return emptyResp, nil
	}

	// 5. 构建响应结果
	resp = &types.GetProductListResp{
		Products: convertToProductInfoRespList(rpcResp.Products),
		Total:    rpcResp.Total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 6. 存入缓存（过期时间5分钟）
	data, err := json.Marshal(resp)
	if err == nil {
		// 基础时间 300 秒 + 随机 0-60 秒，避免同时过期
		randomExpire := consts.CacheResultTTL + rand.Intn(consts.CacheResultRandomRange)
		err = l.svcCtx.Redis.SetexCtx(l.ctx, cacheKey, string(data), randomExpire)
		if err != nil {
			l.Errorf("缓存商品列表失败：%v", err)
			// 注意：缓存失败不影响返回数据
		} else {
			l.Infof("商品列表已缓存：%s（过期时间：%d秒）", cacheKey, randomExpire)
		}
	}

	return resp, nil
}

func convertToProductInfoRespList(products []*productrpc.ProductItem) []types.ProductInfoResp {
	productList := make([]types.ProductInfoResp, 0, len(products))

	for _, product := range products {
		productList = append(productList, types.ProductInfoResp{
			ID:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Status:      int(product.Status),
			Stock:       product.Stock,
			Images:      product.Images,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return productList
}
