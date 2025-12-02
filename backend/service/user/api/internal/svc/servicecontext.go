// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fmt"
	"go-zero-learning/common/cron"
	"go-zero-learning/common/db"
	"go-zero-learning/common/jwt"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/config"

	rediscache "go-zero-learning/common/redis"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	JWT    *jwt.JWTManager   // JWT 管理器
	Redis  *redis.Redis      // Redis 客户端
	Cron   *cron.CronManager // 定时任务管理器
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
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
		&model.Category{},
		&model.Cart{},
		&model.CartItem{},
	)
	if err != nil {
		panic(err)
	}

	// 初始化 JWT 管理器
	jwtManager := jwt.NewJWTManager(c.JWT.Secret, c.JWT.ExpireDays)

	// 初始化 Redis 连接
	err = rediscache.InitRedis(c.Redis)
	if err != nil {
		panic(err)
	}

	// 初始化定时任务管理器
	cronManager := cron.NewCronManager(db.GetDB())

	// 添加订单取消任务（每5分钟执行一次）
	_, err = cronManager.AddJob("0 */5 * * * *", func() {
		job := cron.NewOrderCancelJob(db.GetDB())
		job.Run()
	})
	if err != nil {
		panic(fmt.Errorf("添加订单取消任务失败：%w", err))
	}

	// 启动定时任务
	cronManager.Start()

	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
		JWT:    jwtManager,
		Redis:  rediscache.GetRedis(),
		Cron:   cronManager,
	}
}
