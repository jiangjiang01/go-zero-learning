package svc

import (
	"go-zero-learning/common/db"
	"go-zero-learning/service/product/product-rpc/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接

	if err := db.InitDB(c.DataSource); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
	}
}
