// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"hello-api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
