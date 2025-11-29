import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 分类信息接口
export interface CategoryInfo {
  id: number
  name: string
  desc: string
  parent_id: number
  sort: number
  status: number
  created_at: number
  updated_at: number
}

// 分类列表响应
export interface CategoryListResponse {
  categories: CategoryInfo[]
  total: number
  page: number
  page_size: number
}

// 创建分类请求
export interface CreateCategoryRequest {
  name: string
  desc?: string
  parent_id?: number
  sort?: number
  status?: number
}

// 更新分类请求
export interface UpdateCategoryRequest {
  name?: string
  desc?: string
  parent_id: number
  sort?: number
  status?: number
}

// 分类列表查询参数
export interface CategoryListParams {
  page?: number
  page_size?: number
  keyword?: string
  all?: boolean
}

/**
 * 获取分类列表
 * @param params - 查询参数
 * @returns Promise
 */
export const getCategoryList = (params?: CategoryListParams): Promise<ResponseData<CategoryListResponse>> => {
  return request.get('/api/categories', { params })
}

/**
 * 获取分类详情
 * @param id - 分类ID
 * @returns Promise
 */
export const getCategoryDetail = (id: number): Promise<ResponseData<CategoryInfo>> => {
  return request.get(`/api/categories/${id}`)
}

/**
 * 创建分类
 * @param data - 分类数据
 * @returns Promise
 */
export const createCategory = (data: CreateCategoryRequest): Promise<ResponseData<CategoryInfo>> => {
  return request.post('/api/categories', data)
}

/**
 * 更新分类
 * @param id - 分类ID
 * @param data - 分类数据
 * @returns Promise
 */
export const updateCategory = (id: number, data: UpdateCategoryRequest): Promise<ResponseData<CategoryInfo>> => {
  return request.put(`/api/categories/${id}`, data)
}

/**
 * 删除分类
 * @param id - 分类ID
 * @returns Promise
 */
export const deleteCategory = (id: number): Promise<ResponseData<{ message: string }>> => {
  return request.delete(`/api/categories/${id}`)
}

/**
 * 获取订单状态文本
 * @param status - 状态值
 * @returns 状态文本
 */
export const getCategoryStatusText = (status: number): string => {
  return status === 1 ? '启用' : '禁用'
}

/**
 * 获取订单状态标签类型
 * @param status - 状态值
 * @returns Element Plus 标签类型
 */
export const getCategoryStatusType = (status: number): 'success' | 'danger' => {
  return status === 1 ? 'success' : 'danger'
}

