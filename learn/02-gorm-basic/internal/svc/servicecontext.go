// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"gorm-basic/common"
	"gorm-basic/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB // 直接使用 gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	err := common.InitDB(c.DataSource)
	if err != nil {
		panic(err) // 启动时如果数据库连接失败，直接 panic
	}
	return &ServiceContext{
		Config: c,
		DB:     common.GetDB(), // 使用全局变量（下一步会优化）
	}
}
