// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"strings"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

	// 2. 查询商品是否存在
	var product model.Product
	err = l.svcCtx.DB.First(&product, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrProductNotFound
		}
		l.Errorf("查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建更新字段
	updateFields := make(map[string]interface{})

	// 处理名称更新
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "商品名称不能为空")
		}
		if name != product.Name {
			// 检查名称是否重复
			var existingProduct model.Product
			err = l.svcCtx.DB.Where("name = ? AND id != ?", name, product.ID).First(&existingProduct).Error
			if err == nil {
				return nil, errorx.ErrProductNameExists
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				l.Errorf("查询商品名称失败：%v", err)
				return nil, errorx.ErrInternalError
			}
			updateFields["name"] = name
		}
	}

	// 处理描述更新（允许设置为空）
	if req.Description != nil {
		description := strings.TrimSpace(*req.Description)
		if description != product.Description {
			updateFields["description"] = description
		}
	}

	// 处理价格更新
	if req.Price != nil {
		price := *req.Price
		if price <= 0 {
			return nil, errorx.ErrProductPriceTooLow
		}
		if price > 99999900 {
			return nil, errorx.ErrProductPriceTooHigh
		}
		if price != product.Price {
			updateFields["price"] = price
		}
	}

	// 处理状态更新
	if req.Status != nil {
		status := *req.Status
		if status != 0 && status != 1 {
			return nil, errorx.ErrProductStatusInvalid
		}
		if status != product.Status {
			updateFields["status"] = status
		}
	}

	// 处理库存更新
	if req.Stock != nil {
		stock := *req.Stock
		// 验证库存不能为负数
		if stock < 0 {
			return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "库存不能为负数")
		}
		if stock != product.Stock {
			updateFields["stock"] = stock
		}
	}

	// 4. 检查是否有字段需要更新
	if len(updateFields) == 0 {
		return nil, errorx.ErrProductNoUpdateFields
	}

	// 5. 执行更新
	err = l.svcCtx.DB.Model(&product).Updates(updateFields).Error
	if err != nil {
		l.Errorf("更新商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 6. 重新查询最新数据
	err = l.svcCtx.DB.First(&product, req.ID).Error
	if err != nil {
		l.Errorf("重新查询商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 7. 构建响应结果
	resp = convertToProductInfoResp(product)

	return resp, nil
}
