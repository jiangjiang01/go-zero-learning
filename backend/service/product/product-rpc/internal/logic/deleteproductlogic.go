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

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *DeleteProductLogic) DeleteProduct(in *productrpc.DeleteProductReq) (*productrpc.DeleteProductResp, error) {
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

	// 3. 执行删除
	var err error
	if in.HardDelete {
		// 硬删除：永久删除记录
		err = l.svcCtx.DB.WithContext(l.ctx).Unscoped().Delete(&product).Error
	} else {
		// 软删除：将商品状态设为禁用
		err = l.svcCtx.DB.WithContext(l.ctx).Model(&product).Update("status", model.ProductStatusDisabled).Error
	}

	if err != nil {
		l.Errorf("删除商品失败：id=%d, hard_delete=%v, err=%v", in.Id, in.HardDelete, err)
		return nil, status.Error(codes.Internal, "删除商品失败")
	}

	return &productrpc.DeleteProductResp{
		Success: true,
	}, nil
}
