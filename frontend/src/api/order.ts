import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 订单项信息接口
export interface OrderItemInfo {
  id: number
  product_id: number
  product_name: string
  product_desc: string
  price: number
  quantity: number
  amount: number
}

// 订单信息接口
export interface OrderInfo {
  id: number
  order_no: string
  user_id: number
  total_amount: number
  status: number
  status_text: string
  remark: string
  items?: OrderItemInfo[]
  created_at: number
  updated_at: number
}

// 订单列表响应
export interface OrderListResponse {
  orders: OrderInfo[]
  total: number
  page: number
  page_size: number
}

// 创建订单项请求
export interface CreateOrderItem {
  product_id: number
  quantity: number
}

// 创建订单请求
export interface CreateOrderRequest {
  items: CreateOrderItem[]
  remark?: string
}

// 更新订单状态请求
export interface UpdateOrderStatusRequest {
  status: number
}

// 订单列表查询参数
export interface OrderListParams {
  page?: number
  page_size?: number
  status?: number
  keyword?: string
}

/**
 * 获取订单列表
 * @param params - 查询参数
 * @returns Promise
 */
export const getOrderList = (params?: OrderListParams): Promise<ResponseData<OrderListResponse>> => {
  return request.get('/api/orders', { params })
}

/**
 * 获取订单详情
 * @param id - 订单ID
 * @returns Promise
 */
export const getOrderDetail = (id: number): Promise<ResponseData<OrderInfo>> => {
  return request.get(`/api/orders/${id}`)
}

/**
 * 创建订单
 * @param data - 订单数据
 * @returns Promise
 */
export const createOrder = (data: CreateOrderRequest): Promise<ResponseData<OrderInfo>> => {
  return request.post('/api/orders', data)
}

/**
 * 更新订单状态
 * @param id - 订单ID
 * @param data - 状态数据
 * @returns Promise
 */
export const updateOrderStatus = (id: number, data: UpdateOrderStatusRequest): Promise<ResponseData<OrderInfo>> => {
  return request.put(`/api/orders/${id}/status`, data)
}

/**
 * 格式化价格显示（分转元）
 * @param price - 价格（分）
 * @returns 格式化后的价格字符串
 */
export const formatPrice = (price: number): string => {
  return (price / 100).toFixed(2)
}

/**
 * 获取订单状态文本
 * @param status - 状态值
 * @returns 状态文本
 */
export const getOrderStatusText = (status: number): string => {
  const statusMap: Record<number, string> = {
    1: '待支付',
    2: '已支付',
    3: '已发货',
    4: '已完成',
    5: '已取消'
  }
  return statusMap[status] || '未知状态'
}

/**
 * 获取订单状态标签类型
 * @param status - 状态值
 * @returns Element Plus 标签类型
 */
export const getOrderStatusType = (status: number): 'success' | 'warning' | 'danger' | 'info' => {
  const typeMap: Record<number, 'success' | 'warning' | 'danger' | 'info'> = {
    1: 'warning',  // 待支付
    2: 'info',     // 已支付
    3: 'info',     // 已发货
    4: 'success',  // 已完成
    5: 'danger'    // 已取消
  }
  return typeMap[status] || 'info'
}

