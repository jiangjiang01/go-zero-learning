// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	// 1. 校验参数
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

	// 3. 删除商品
	err = l.svcCtx.DB.Delete(&product).Error
	if err != nil {
		l.Errorf("删除商品失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 清除商品列表缓存（删除商品后， 列表需要更新）
	l.clearProductListCache()

	// 4. 构建响应结果
	resp = &types.DeleteProductResp{
		Message: "商品删除成功",
	}

	return resp, nil
}

// clearProductListCache 清除所有商品列表缓存
func (l *DeleteProductLogic) clearProductListCache() {
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
