package cron

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// CronManager 定时任务管理器
type CronManager struct {
	cron   *cron.Cron
	db     *gorm.DB
	logger logx.Logger
}

// NewCronManager 创建定时任务管理器
func NewCronManager(db *gorm.DB) *CronManager {
	return &CronManager{
		cron:   cron.New(cron.WithSeconds()), // 支持秒级精度
		db:     db,
		logger: logx.WithContext(context.Background()),
	}
}

// 启动定时任务
func (m *CronManager) Start() {
	m.cron.Start()
	m.logger.Infof("定时任务管理器已启动")
}

// 停止定时任务
func (m *CronManager) Stop() {
	ctx := m.cron.Stop()
	<-ctx.Done() // 等待所有正在执行的任务完成
	m.logger.Infof("定时任务管理器已停止")
}

// 添加定时任务
func (m *CronManager) AddJob(spec string, job func()) (cron.EntryID, error) {
	return m.cron.AddFunc(spec, job)
}
