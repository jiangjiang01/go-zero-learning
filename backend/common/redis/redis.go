package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var RedisClient *redis.Redis

// 初始化 Redis 连接
// 注意：go-zero 的 Redis 配置方式不同，直接传入配置对象
func InitRedis(redisConf redis.RedisConf) error {
	RedisClient = redis.MustNewRedis(redisConf)
	ctx := context.Background()

	err := RedisClient.SetCtx(ctx, "test:connection", "ok")
	if err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}

	value, err := RedisClient.GetCtx(ctx, "test:connection")
	if err != nil {
		return fmt.Errorf("failed to get key: %w", err)
	}

	log.Println("value: ", value)
	// 清理测试键
	RedisClient.DelCtx(ctx, "test:connection")

	log.Println("Redis 连接成功！")
	return nil
}

// 获取 Redis 客户端
func GetRedis() *redis.Redis {
	return RedisClient
}
