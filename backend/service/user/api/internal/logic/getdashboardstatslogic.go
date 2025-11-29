// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"time"

	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDashboardStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDashboardStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDashboardStatsLogic {
	return &GetDashboardStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取统计数据
func (l *GetDashboardStatsLogic) GetDashboardStats() (resp *types.DashboardStatsResp, err error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 1. 获取订单统计数据
	orderStats, err := l.getOrderStats(todayStart)
	if err != nil {
		l.Errorf("获取订单统计数据失败: %v", err)
		return nil, errorx.ErrInternalError
	}

	// 2. 获取商品统计数据
	productStats, err := l.getProductStats()
	if err != nil {
		l.Errorf("获取商品统计数据失败: %v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 获取用户统计数据
	userStats, err := l.getUserStats(todayStart)
	if err != nil {
		l.Errorf("获取用户统计数据失败: %v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 构建响应
	resp = &types.DashboardStatsResp{
		OrderStats:   *orderStats,
		ProductStats: *productStats,
		UserStats:    *userStats,
	}

	return resp, nil
}

// 获取订单统计数据
func (l *GetDashboardStatsLogic) getOrderStats(todayStart time.Time) (resp *types.OrderStats, err error) {
	var stats types.OrderStats

	// 1. 订单总数
	err = l.svcCtx.DB.Model(&model.Order{}).Count(&stats.TotalOrders).Error
	if err != nil {
		return nil, err
	}

	// 2. 今日订单数
	err = l.svcCtx.DB.Model(&model.Order{}).Where("created_at >= ?", todayStart).Count(&stats.TodayOrders).Error
	if err != nil {
		return nil, err
	}

	// 订单总金额
	var totalAmountResult struct {
		Total int64
	}
	err = l.svcCtx.DB.Model(&model.Order{}).Select("COALESCE(SUM(total_amount), 0) as total").
		Scan(&totalAmountResult).Error
	if err != nil {
		return nil, err
	}
	stats.TotalAmount = totalAmountResult.Total

	// 今日订单金额
	var todayAmountResult struct {
		Total int64
	}
	err = l.svcCtx.DB.Model(&model.Order{}).
		Where("created_at >= ?", todayStart).
		Select("COALESCE(SUM(total_amount), 0) as total").
		Scan(&todayAmountResult).Error
	if err != nil {
		return nil, err
	}
	stats.TodayAmount = todayAmountResult.Total

	// 订单状态分布
	statsCount, err := l.getOrderStatusCount()
	if err != nil {
		l.Errorf("获取订单状态分布失败: %v", err)
		return nil, err
	}

	stats.StatusCount = *statsCount

	return &stats, nil
}

func (l *GetDashboardStatsLogic) getOrderStatusCount() (resp *types.OrderStatusCount, err error) {
	var statsCount types.OrderStatusCount

	// 待支付（状态1）
	err = l.svcCtx.DB.Model(&model.Order{}).Where("status = ?", 1).Count(&statsCount.Pending).Error
	if err != nil {
		return nil, err
	}

	// 已支付（状态2）
	err = l.svcCtx.DB.Model(&model.Order{}).Where("status = ?", 2).Count(&statsCount.Paid).Error
	if err != nil {
		return nil, err
	}

	// 已发货（状态3）
	err = l.svcCtx.DB.Model(&model.Order{}).Where("status = ?", 3).Count(&statsCount.Shipped).Error
	if err != nil {
		return nil, err
	}

	// 已完成（状态4）
	err = l.svcCtx.DB.Model(&model.Order{}).Where("status = ?", 4).Count(&statsCount.Completed).Error
	if err != nil {
		return nil, err
	}

	// 已取消（状态5）
	err = l.svcCtx.DB.Model(&model.Order{}).Where("status = ?", 5).Count(&statsCount.Canceled).Error
	if err != nil {
		return nil, err
	}

	return &statsCount, nil
}

// 获取商品统计数据
func (l *GetDashboardStatsLogic) getProductStats() (resp *types.ProductStats, err error) {
	var stats types.ProductStats

	// 商品总数
	err = l.svcCtx.DB.Model(&model.Product{}).Count(&stats.TotalProducts).Error
	if err != nil {
		return nil, err
	}

	// 上架商品数
	err = l.svcCtx.DB.Model(&model.Product{}).Where("status = ?", 1).Count(&stats.OnsaleProducts).Error
	if err != nil {
		return nil, err
	}

	// 下架商品数
	err = l.svcCtx.DB.Model(&model.Product{}).Where("status = ?", 0).Count(&stats.OffsaleProducts).Error
	if err != nil {
		return nil, err
	}

	// 库存不足商品数 < 10
	err = l.svcCtx.DB.Model(&model.Product{}).Where("stock < ?", 10).Count(&stats.LowStockProducts).Error
	if err != nil {
		return nil, err
	}

	// 总库存量
	var totalStockResult struct {
		Total int64
	}
	err = l.svcCtx.DB.Model(&model.Product{}).
		Select("COALESCE(SUM(stock), 0) as total").
		Scan(&totalStockResult).Error
	if err != nil {
		return nil, err
	}
	stats.TotalStock = totalStockResult.Total

	return &stats, nil
}

// 获取用户统计数据
func (l *GetDashboardStatsLogic) getUserStats(todayStart time.Time) (resp *types.UserStats, err error) {
	var stats types.UserStats

	// 总用户数
	err = l.svcCtx.DB.Model(&model.User{}).Count(&stats.TotalUsers).Error
	if err != nil {
		return nil, err
	}

	// 今日新增用户数
	err = l.svcCtx.DB.Model(&model.User{}).Where("created_at >= ?", todayStart).Count(&stats.TodayUsers).Error
	if err != nil {
		return nil, err
	}

	// 活跃用户数（有订单）
	err = l.svcCtx.DB.Model(&model.User{}).
		Joins("INNER JOIN orders ON users.id = orders.user_id").
		Distinct("users.id").
		Count(&stats.ActiveUsers).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
