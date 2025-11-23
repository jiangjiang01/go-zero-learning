// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource string    `json:"dataSource"` // 数据库连接字符串dsn
	JWT        JWTConfig `json:"jwt"`        // JWT 配置
}

type JWTConfig struct {
	Secret     string `json:"secret"`     // JWT 密钥
	ExpireDays int    `json:"expireDays"` // Token 过期天数
}
