package svc

import (
	"sync"
	"user-rpc/internal/config"
	"user-rpc/userrpc"
)

type ServiceContext struct {
	Config config.Config
	// 内存存储（后续优化）
	UserStore map[int64]*userrpc.User
	UserMutex sync.RWMutex // 读写锁
	NextID    int64        // 下一个用户 ID
	IdMutex   sync.Mutex   // 确保 nextID 的原子性
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserStore: make(map[int64]*userrpc.User),
		NextID:    1,
	}
}
