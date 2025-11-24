// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"go-zero-learning/common/db"
	"go-zero-learning/common/jwt"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	JWT    *jwt.JWTManager // JWT 管理器
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	err := db.InitDB(c.DataSource)
	if err != nil {
		panic(err)
	}

	// 自动迁移（创建表）
	err = db.GetDB().AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.Menu{},
		&model.UserRole{},
		&model.RolePermission{},
	)
	if err != nil {
		panic(err)
	}

	// 初始化 JWT 管理器
	jwtManager := jwt.NewJWTManager(c.JWT.Secret, c.JWT.ExpireDays)

	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
		JWT:    jwtManager,
	}
}
