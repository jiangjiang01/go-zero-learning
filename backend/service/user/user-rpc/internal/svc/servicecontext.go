package svc

import (
	"go-zero-learning/common/db"
	"go-zero-learning/service/user/user-rpc/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 暂时和 API 共用同一个 db.InitDB
	if err := db.InitDB(c.DataSource); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
	}
}
