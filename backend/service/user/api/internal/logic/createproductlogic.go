// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

	if req.Price <= 0 {
		return nil, errorx.ErrProductPriceTooLow
	}
	// if req.Price > 99999900 {
	// 	return nil, errorx.ErrProductPriceTooHigh
	// }

	// 2. 查询商品名称是否已存在
	var existingProduct model.Product
	err = l.svcCtx.DB.Where("name = ?", name).First(&existingProduct).Error
	if err == nil {
		return nil, errorx.ErrProductNameExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf("查询商品名称失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 创建商品
	status := 1 // 默认启用
	if req.Status != nil {
		status = *req.Status
	}
	stock := int64(0) // 默认库存为0
	if req.Stock != nil {
		// 不能为负数
		if *req.Stock < 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "库存不能为负数")
		}
		stock = *req.Stock
	}

	// 处理图片: 将 []string 转换为 JSON 字符串
	imagesJSON := "[]" // 默认空数据
	if len(req.Images) > 0 {
		imagesBytes, err := json.Marshal(req.Images)
		if err != nil {
			l.Errorf("序列化图片列表失败：%v", err)
			return nil, errorx.ErrInternalError
		}
		imagesJSON = string(imagesBytes)
	}

	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       req.Price,
		Status:      status,
		Stock:       stock,
		Images:      imagesJSON, // 存储JSON字符串
	}
	err = l.svcCtx.DB.Create(&product).Error
	if err != nil {
		l.Errorf("创建商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 清除商品列表缓存（新增商品后，列表需要更新）
	l.clearProductListCache()

	// 4. 构建响应（需要将JSON字符串解析回数组）
	var images []string
	if product.Images != "" {
		err = json.Unmarshal([]byte(product.Images), &images)
		if err != nil {
			l.Errorf("解析图片列表失败：%v", err)
			images = []string{} // 解析失败时使用空数组
		}
	}
	resp = &types.ProductInfoResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      product.Status,
		Stock:       product.Stock,
		Images:      images, // 返回数组
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}

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
