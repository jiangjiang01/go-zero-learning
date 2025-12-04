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

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *UpdateProductLogic) UpdateProduct(in *productrpc.UpdateProductReq) (*productrpc.UpdateProductResp, error) {
	// 1. 参数校验
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "商品ID必须大于0")
	}

	// 2. 查询商品是否存在
	var product model.Product
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&product, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "商品不存在")
		}
		l.Errorf("查询商品失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.Internal, "查询商品失败")
	}

	// 3. 构建更新字段
	updates := make(map[string]interface{})

	// 更新名称（如果提供）
	if in.Name != nil {
		name := strings.TrimSpace(*in.Name)
		if name == "" {
			return nil, status.Error(codes.InvalidArgument, "商品名称不能为空")
		}
		if len(name) > consts.MaxProductNameLength {
			return nil, status.Error(codes.InvalidArgument, "商品名称长度不能超过100个字符")
		}

		// 检查名称唯一性（排除当前商品）
		var existingProduct model.Product
		err := l.svcCtx.DB.WithContext(l.ctx).Where("name = ? AND id != ?", name, in.Id).First(&existingProduct).Error
		if err == nil {
			return nil, status.Error(codes.AlreadyExists, "商品名称已存在")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorf("查询商品名称失败：%v", err)
			return nil, status.Error(codes.Internal, "查询商品名称失败")
		}
		updates["name"] = name
	}

	// 更新描述（如果提供）
	if in.Description != nil {
		updates["description"] = strings.TrimSpace(*in.Description)
	}

	// 更新价格（如果提供）
	if in.Price != nil {
		if *in.Price < consts.MinProductPrice {
			return nil, status.Error(codes.InvalidArgument, "商品价格太低")
		}
		if *in.Price > consts.MaxProductPrice {
			return nil, status.Error(codes.InvalidArgument, "商品价格太高")
		}
		updates["price"] = *in.Price
	}

	// 更新状态（如果提供）
	if in.Status != nil {
		if !model.IsValidProductStatus(int(*in.Status)) {
			return nil, status.Error(codes.InvalidArgument, "无效的商品状态，状态值必须为0或1")
		}
		updates["status"] = *in.Status
	}

	// 更新库存（如果提供）
	if in.Stock != nil {
		if *in.Stock < 0 {
			return nil, status.Error(codes.InvalidArgument, "库存不能为负数")
		}
		updates["stock"] = *in.Stock
	}

	// 更新图片（如果标记需要更新）
	if in.UpdateImages {
		imagesJSON := "[]"
		if len(in.Images) > 0 {
			imagesBytes, err := json.Marshal(in.Images)
			if err != nil {
				l.Errorf("序列化图片列表失败：%v", err)
				return nil, status.Error(codes.InvalidArgument, "序列化图片列表失败")
			}
			imagesJSON = string(imagesBytes)
		}
		updates["images"] = imagesJSON
	}

	// 4. 如果没有任何字段需要更新
	if len(updates) == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有需要更新的字段")
	}

	// 5. 执行更新
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&product).Updates(updates).Error; err != nil {
		l.Errorf("更新商品失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.Internal, "更新商品失败")
	}

	// 6. 重新查询更新后的商品
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&product, in.Id).Error; err != nil {
		l.Errorf("查询更新后的商品失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.Internal, "查询更新后的商品失败")
	}

	// 7. 解析图片列表
	var images []string
	if product.Images != "" {
		if err := json.Unmarshal([]byte(product.Images), &images); err != nil {
			l.Errorf("解析商品图片失败：id=%d, err=%v", in.Id, err)
			images = []string{}
		}
	}
	if images == nil {
		images = []string{}
	}

	// 8. 构建响应
	return &productrpc.UpdateProductResp{
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
	}, nil
}
