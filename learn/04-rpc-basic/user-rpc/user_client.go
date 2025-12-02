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
	// 创建 RPC 客户端
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:8080"},
	})
	userRpc := userrpcclient.NewUserRpc(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// 1. 创建用户
	fmt.Println("创建用户")
	createResp, err := userRpc.CreateUser(ctx, &userrpc.CreateUserReq{
		Username: "test",
		Email:    "test@example.com",
	})
	if err != nil {
		fmt.Println("创建用户失败：", err)
		return
	}
	fmt.Printf("创建用户成功：ID=%d, Username=%s, Email=%s\n", createResp.User.Id, createResp.User.Username, createResp.User.Email)

	// 2. 获取用户
	getResp, err := userRpc.GetUser(ctx, &userrpc.GetUserReq{
		Id: createResp.User.Id,
	})
	if err != nil {
		fmt.Println("获取用户失败：", err)
		return
	}
	fmt.Printf("获取用户成功：ID=%d, Username=%s, Email=%s\n", getResp.User.Id, getResp.User.Username, getResp.User.Email)

	// 3. 获取不存在的用户（测试错误处理）
	fmt.Println("获取不存在的用户")
	getResp, err = userRpc.GetUser(ctx, &userrpc.GetUserReq{
		Id: 9999,
	})
	if err != nil {
		fmt.Println("获取不存在的用户失败：", err)
		return
	}
	fmt.Printf("获取不存在的用户成功：ID=%d, Username=%s, Email=%s\n", getResp.User.Id, getResp.User.Username, getResp.User.Email)
}
