// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RpcPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRpcPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcPingLogic {
	return &RpcPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RpcPingLogic) RpcPing(req *types.RpcPingReq) (resp *types.RpcPingResp, err error) {
	// 调用 user-rpc 的 Ping
	rpcResp, err := l.svcCtx.UserRpc.Ping(l.ctx, &userrpc.Request{
		Ping: req.Message,
	})
	if err != nil {
		l.Errorf("调用 UserRpc.Ping 失败：%v", err)
		return
	}

	resp = &types.RpcPingResp{
		Pong: rpcResp.Pong,
	}

	return resp, nil
}
