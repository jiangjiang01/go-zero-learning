package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// 先只加一个 DataSource, 用来连同一个 MySQL
	DataSource string `json:"dataSource"`
}
