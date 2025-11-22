// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jwt-basic/common"
	"jwt-basic/internal/svc"
	"jwt-basic/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.VerifyReq) (resp *types.VerifyResp, err error) {
	// 解析 Token
	claims, err := common.ParseToken(req.Token)
	if err != nil {
		resp = &types.VerifyResp{
			Valid:   false,
			Message: "无效的 Token: " + err.Error(),
		}
		return resp, nil
	}

	// Token 有效
	resp = &types.VerifyResp{
		Valid:   true,
		UserID:  claims.UserID,
		Message: "Token 有效，用户: " + claims.Username,
	}
	return resp, nil
}
