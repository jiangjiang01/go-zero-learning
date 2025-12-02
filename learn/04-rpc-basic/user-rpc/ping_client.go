package main

import (
	"context"
	"fmt"
	"time"

	"user-rpc/userrpc"
	"user-rpc/userrpcclient"

	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	// 1. 创建直连客户端（不走 etcd）
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:8080"},
	})
	// 注意：zrpc.Client 不需要显式关闭，连接会自动管理

	// 2. 基于 client 创建 UserRpc 客户端
	userRpc := userrpcclient.NewUserRpc(client)

	// 3. 构造请求并调用 Ping
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := userRpc.Ping(ctx, &userrpc.Request{
		Ping: "hello rpc",
	})
	if err != nil {
		fmt.Println("调用 Ping 出错:", err)
		return
	}

	fmt.Println("Ping 响应:", resp.Pong)
}
