// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DataSource string          `json:"dataSource"` // 数据库连接字符串dsn
	JWT        JWTConfig       `json:"jwt"`        // JWT 配置
	Upload     UploadConfig    `json:"upload"`     // 文件上传配置
	Redis      redis.RedisConf `json:"redis"`      // Redis 配置

	UserRpc zrpc.RpcClientConf `json:"userRpc"` // 用户 UserRpc 客户端配置
}

type JWTConfig struct {
	Secret     string `json:"secret"`     // JWT 密钥
	ExpireDays int    `json:"expireDays"` // Token 过期天数
}

type UploadConfig struct {
	Path         string   `json:"path"`         // 上传文件根目录
	MaxSize      int64    `json:"maxSize"`      // 最大文件大小（字节）
	AllowedTypes []string `json:"allowedTypes"` // 允许的文件类型
	BaseURL      string   `json:"baseURL"`      // 静态资源访问地址
}
