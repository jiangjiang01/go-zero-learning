// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
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
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 2. 构建查询模型
	query := l.svcCtx.DB.Model(&model.Product{})
	if req.Keyword != "" {
		likeStr := "%" + req.Keyword + "%"
		query = query.Where("name LIKE ? AND description LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	var total int64
	err = query.Count(&total).Error
	if err != nil {
		l.Errorf("查询商品总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 空结果提前处理
	if total == 0 {
		return &types.GetProductListResp{
			Products: []types.ProductInfoResp{},
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	// 4. 分页查询数据
	offset := (req.Page - 1) * req.PageSize
	var products []model.Product
	err = query.Order("created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&products).Error
	if err != nil {
		l.Errorf("查询商品列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 构建响应结果
	resp = &types.GetProductListResp{
		Products: convertToProductInfoRespList(products),
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return resp, nil
}

func convertToProductInfoRespList(products []model.Product) []types.ProductInfoResp {
	productList := make([]types.ProductInfoResp, 0, len(products))

	for _, p := range products {
		productList = append(productList, types.ProductInfoResp{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt.Unix(),
			UpdatedAt:   p.UpdatedAt.Unix(),
		})
	}

	return productList
}
