// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/tmp/internal/svc"
	"go-zero-learning/service/user/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDetailLogic {
	return &GetCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryDetailLogic) GetCategoryDetail(req *types.GetCategoryDetailReq) (resp *types.CategoryInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
