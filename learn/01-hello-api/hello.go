// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"hello-api/internal/config"
	"hello-api/internal/handler"
	"hello-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hello-api.yaml", "the config file")

func main() {
	// 1. 解析命令行参数，获取配置文件路径
	flag.Parse()

	// 2. 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 3. 创建HTTP服务器
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 4. 创建服务上下文（依赖注入容器）
	ctx := svc.NewServiceContext(c)

	// 5. 注册路由处理器
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 6. 启动服务器
	server.Start()
}
