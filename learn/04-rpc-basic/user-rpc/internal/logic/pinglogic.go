package logic

import (
	"context"
	"fmt"

	"user-rpc/internal/svc"
	"user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *userrpc.Request) (*userrpc.Response, error) {
	// 记录请求日志
	l.Infof("收到 Ping 请求：%s", in.Ping)

	// 使用输入参数，返回带时间戳的响应
	pongMsg := fmt.Sprintf("pong from server, received: %s", in.Ping)

	// 记录响应日志
	l.Infof("返回 Ping 响应：%s", pongMsg)

	return &userrpc.Response{
		Pong: pongMsg,
	}, nil
}
