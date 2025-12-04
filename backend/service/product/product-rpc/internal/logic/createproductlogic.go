package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"go-zero-learning/common/consts"
	"go-zero-learning/model"
	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	// 1. 参数校验
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "商品名称不能为空")
	}
	if len(name) > 100 {
		return nil, status.Error(codes.InvalidArgument, "商品名称长度不能超过100个字符")
	}

	if in.Price < consts.MinProductPrice {
		return nil, status.Error(codes.InvalidArgument, "商品价格太低")
	}

	if in.Price > consts.MaxProductPrice {
		return nil, status.Error(codes.InvalidArgument, "商品价格太高")
	}

	// 2. 查询商品名称是否已存在
	var existingProduct model.Product
	err := l.svcCtx.DB.WithContext(l.ctx).Where("name = ?", name).First(&existingProduct).Error
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "商品名称已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询商品名称失败：%v", err)
		return nil, status.Error(codes.Internal, "查询商品名称失败")
	}

	// 3. 设置默认值
	productStatus := model.ProductStatusEnabled // 默认启用
	// 只有明确传入 0 或 1 时才使用，其他值（包括 -1）都使用默认值
	if in.Status == int32(model.ProductStatusDisabled) || in.Status == int32(model.ProductStatusEnabled) {
		productStatus = int(in.Status)
	}
	stock := int64(0) // 默认库存为0
	if in.Stock != 0 {
		if in.Stock < 0 {
			return nil, status.Error(codes.InvalidArgument, "库存不能为负数")
		}
		stock = in.Stock
	}

	// 4. 处理图片: 将 []string 转换为 JSON 字符串
	imagesJSON := "[]" // 默认空数据
	if len(in.Images) > 0 {
		imagesBytes, err := json.Marshal(in.Images)
		if err != nil {
			l.Errorf("序列化图片列表失败：%v", err)
			return nil, status.Error(codes.InvalidArgument, "序列化图片列表失败")
		}
		imagesJSON = string(imagesBytes)
	}

	// 5. 创建商品
	product := &model.Product{
		Name:        name,
		Description: strings.TrimSpace(in.Description),
		Price:       in.Price,
		Status:      productStatus,
		CategoryID:  in.CategoryId,
		Stock:       stock,
		Images:      imagesJSON, // 存储JSON字符串
	}
	err = l.svcCtx.DB.WithContext(l.ctx).Create(&product).Error
	if err != nil {
		l.Errorf("创建商品失败：%v", err)
		return nil, status.Error(codes.Internal, "创建商品失败")
	}

	// 6. 解析图片列表用于响应
	var images []string
	if product.Images != "" {
		err = json.Unmarshal([]byte(product.Images), &images)
		if err != nil {
			l.Errorf("解析图片列表失败：%v", err)
			images = []string{} // 解析失败时使用空数组
		}
	}
	if images == nil {
		images = []string{}
	}

	// 7. 构建响应
	resp := &productrpc.CreateProductResp{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      int32(product.Status),
		CategoryId:  product.CategoryID,
		Stock:       product.Stock,
		Images:      images,
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}

	return resp, nil
}
