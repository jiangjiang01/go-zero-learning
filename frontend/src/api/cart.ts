import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 购物车项信息接口
export interface CartItemInfo {
  id: number
  product_id: number
  product_name: string
  price: number
  quantity: number
  amount: number
}

// 购物车信息接口
export interface CartInfo {
  id: number
  user_id: number
  items: CartItemInfo[]
  total_amount: number
  item_count: number
}

// 添加商品到购物车请求
export interface AddCartItemRequest {
  product_id: number
  quantity: number
}

// 更新购物车项数量请求
export interface UpdateCartItemRequest {
  quantity: number
}

/**
 * 获取购物车
 * @returns Promise
 */
export const getCart = (): Promise<ResponseData<CartInfo>> => {
  return request.get('/api/cart')
}

/**
 * 添加商品到购物车
 * @param data - 商品数据
 * @returns Promise
 */
export const addCartItem = (data: AddCartItemRequest): Promise<ResponseData<CartItemInfo>> => {
  return request.post('/api/cart/items', data)
}

/**
 * 更新购物车项数量
 * @param itemId - 购物车项ID
 * @param data - 更新数据
 * @returns Promise
 */
export const updateCartItem = (itemId: number, data: UpdateCartItemRequest): Promise<ResponseData<CartItemInfo>> => {
  return request.put(`/api/cart/items/${itemId}`, data)
}

/**
 * 删除购物车项
 * @param itemId - 购物车项ID
 * @returns Promise
 */
export const deleteCartItem = (itemId: number): Promise<ResponseData<{ message: string }>> => {
  return request.delete(`/api/cart/items/${itemId}`)
}

/**
 * 清空购物车
 * @returns Promise
 */
export const clearCart = (): Promise<ResponseData<{ message: string }>> => {
  return request.delete('/api/cart')
}

/**
 * 格式化价格显示（分转元）
 * @param price - 价格（分）
 * @returns 格式化后的价格字符串
 */
export const formatPrice = (price: number): string => {
  return (price / 100).toFixed(2)
}

