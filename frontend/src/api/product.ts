import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 商品信息接口
export interface ProductInfo {
  id: number
  name: string
  description: string
  price: number
  status: number
  stock: number
  images: string[]  // 图片URL列表
  created_at: number
  updated_at: number
}

// 商品列表响应
export interface ProductListResponse {
  products: ProductInfo[]
  total: number
  page: number
  page_size: number
}

// 创建商品请求
export interface CreateProductRequest {
  name: string
  description?: string
  price: number
  status?: number
  stock?: number
  images?: string[]  // 图片URL列表
}

// 更新商品请求
export interface UpdateProductRequest {
  name?: string
  description?: string
  price?: number
  status?: number
  stock?: number
  images?: string[]  // 图片URL列表
}

// 商品列表查询参数
export interface ProductListParams {
  page?: number
  page_size?: number
  keyword?: string
}

/**
 * 获取商品列表
 * @param params - 查询参数
 * @returns Promise
 */
export const getProductList = (params?: ProductListParams): Promise<ResponseData<ProductListResponse>> => {
  return request.get('/api/products', { params })
}

/**
 * 获取商品详情
 * @param id - 商品ID
 * @returns Promise
 */
export const getProductDetail = (id: number): Promise<ResponseData<ProductInfo>> => {
  return request.get(`/api/products/${id}`)
}

/**
 * 创建商品
 * @param data - 商品数据
 * @returns Promise
 */
export const createProduct = (data: CreateProductRequest): Promise<ResponseData<ProductInfo>> => {
  return request.post('/api/products', data)
}

/**
 * 更新商品
 * @param id - 商品ID
 * @param data - 更新数据
 * @returns Promise
 */
export const updateProduct = (id: number, data: UpdateProductRequest): Promise<ResponseData<ProductInfo>> => {
  return request.put(`/api/products/${id}`, data)
}

/**
 * 删除商品
 * @param id - 商品ID
 * @returns Promise
 */
export const deleteProduct = (id: number): Promise<ResponseData<{ message: string }>> => {
  return request.delete(`/api/products/${id}`)
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
 * 价格转换为分
 * @param price - 价格（元）
 * @returns 价格（分）
 */
export const priceToFen = (price: number): number => {
  return Math.round(price * 100)
}

/**
 * 获取商品状态文本
 * @param status - 状态值
 * @returns 状态文本
 */
export const getProductStatusText = (status: number): string => {
  return status === 1 ? '上架' : '下架'
}

/**
 * 获取商品状态标签类型
 * @param status - 状态值
 * @returns Element Plus 标签类型
 */
export const getProductStatusType = (status: number): 'success' | 'danger' => {
  return status === 1 ? 'success' : 'danger'
}
