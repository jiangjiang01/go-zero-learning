package logic

import (
	"context"
	"encoding/json"
	"strings"

	"go-zero-learning/common/consts"
	"go-zero-learning/model"
	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. 参数校验和默认值设置
	page := in.Page
	if page < 1 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize < 1 {
		pageSize = consts.DefaultPageSize
	}
	if pageSize > consts.MaxPageSize {
		pageSize = consts.MaxPageSize
	}

	// 2. 构建查询条件
	query := l.svcCtx.DB.WithContext(l.ctx).Model(&model.Product{})
	if in.Keyword != "" {
		keyword := strings.TrimSpace(in.Keyword)
		likeStr := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		l.Errorf("查询商品总数失败：%v", err)
		return nil, status.Error(codes.Internal, "查询商品总数失败")
	}

	// 4. 分页查询
	offset := (page - 1) * pageSize
	var products []model.Product
	if err := query.Order("created_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&products).Error; err != nil {
		l.Errorf("查询商品列表失败：%v", err)
		return nil, status.Error(codes.Internal, "查询商品列表失败")
	}

	// 5. 转换为响应格式
	productList := make([]*productrpc.ProductItem, 0, len(products))
	for _, p := range products {
		// 解析图片列表
		var images []string
		if p.Images != "" {
			if err := json.Unmarshal([]byte(p.Images), &images); err != nil {
				l.Errorf("解析商品图片失败：id=%d, err=%v", p.ID, err)
				images = []string{}
			}
		}
		if images == nil {
			images = []string{}
		}

		productList = append(productList, &productrpc.ProductItem{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Status:      int32(p.Status),
			CategoryId:  p.CategoryID,
			Stock:       p.Stock,
			Images:      images,
			CreatedAt:   p.CreatedAt.Unix(),
			UpdatedAt:   p.UpdatedAt.Unix(),
		})
	}

	return &productrpc.ListProductResp{
		Products: productList,
		Total:    total,
	}, nil
}
