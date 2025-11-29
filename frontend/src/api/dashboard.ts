import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 订单状态统计
export interface OrderStatusCount {
  pending: number
  paid: number
  shipped: number
  completed: number
  canceled: number
}

// 订单统计
export interface OrderStats {
  total_orders: number
  today_orders: number
  total_amount: number
  today_amount: number
  status_count: OrderStatusCount
}

// 商品统计
export interface ProductStats {
  total_products: number
  onsale_products: number
  offsale_products: number
  low_stock_products: number
  total_stock: number
}

// 用户统计
export interface UserStats {
  total_users: number
  today_users: number
  active_users: number
}

// Dashboard 统计响应
export interface DashboardStatsResponse {
  order_stats: OrderStats
  product_stats: ProductStats
  user_stats: UserStats
}

/**
 * 获取 Dashboard 统计数据
 * @returns Promise
 */
export const getDashboardStats = (): Promise<ResponseData<DashboardStatsResponse>> => {
  return request.get('/api/dashboard/stats')
}

/**
 * 格式化价格显示（分转元）
 * @param price - 价格（分）
 * @returns 格式化后的价格字符串
 */
export const formatPrice = (price: number): string => {
  return (price / 100).toFixed(2)
}

