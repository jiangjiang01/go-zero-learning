import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// Dashboard 统计数据响应类型
export interface DashboardStats {
  user_total: number          // 用户总数
  user_today: number          // 今日新增用户数
  user_enabled: number        // 启用用户数
  user_disabled: number       // 禁用用户数
  role_total: number          // 角色总数
  menu_total: number          // 菜单总数
  operation_log_today: number // 今日操作日志数
  login_log_today: number     // 今日登录日志数
}

/**
 * 获取 Dashboard 统计数据
 * @returns Promise<ResponseData<DashboardStats>>
 */
export function getDashboardStats(): Promise<ResponseData<DashboardStats>> {
  return request({
    url: '/dashboard/status',
    method: 'get'
  })
}