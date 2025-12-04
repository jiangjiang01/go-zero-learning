package logic

import (
	"context"
	"errors"

	"go-zero-learning/model"
	"go-zero-learning/service/product/product-rpc/internal/svc"
	"go-zero-learning/service/product/product-rpc/productrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UpdateProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStatusLogic {
	return &UpdateProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品状态
func (l *UpdateProductStatusLogic) UpdateProductStatus(in *productrpc.UpdateProductStatusReq) (*productrpc.UpdateProductStatusResp, error) {
	// 1. 参数校验
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "商品ID必须大于0")
	}

	// 验证状态值
	if !model.IsValidProductStatus(int(in.Status)) {
		return nil, status.Error(codes.InvalidArgument, "无效的商品状态，状态值必须为0或1")
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

	// 3. 如果状态相同则无需更新
	if product.Status == int(in.Status) {
		return &productrpc.UpdateProductStatusResp{
			Id:        product.ID,
			Status:    int32(product.Status),
			UpdatedAt: product.UpdatedAt.Unix(),
		}, nil
	}

	// 4. 更新状态
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&product).Update("status", in.Status).Error; err != nil {
		l.Errorf("更新商品状态失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.Internal, "更新商品状态失败")
	}

	// 5. 重新查询更新后的商品
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&product, in.Id).Error; err != nil {
		l.Errorf("查询更新后的商品失败：id=%d, err=%v", in.Id, err)
		return nil, status.Error(codes.Internal, "查询更新后的商品失败")
	}

	// 6. 构建响应
	return &productrpc.UpdateProductStatusResp{
		Id:        product.ID,
		Status:    int32(product.Status),
		UpdatedAt: product.UpdatedAt.Unix(),
	}, nil
}
