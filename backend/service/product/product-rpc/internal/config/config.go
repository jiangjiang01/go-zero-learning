package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// 数据库连接配置
	DataSource string `json:"dataSource"`
}
