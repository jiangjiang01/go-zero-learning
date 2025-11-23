import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 权限信息
export interface PermissionInfo {
  id: number
  name: string
  code: string
  desc: string
  created_at: number
  updated_at: number
}

// 权限列表响应
export interface PermissionListResponse {
  permissions: PermissionInfo[]
  total: number
  page: number
  page_size: number
}

// 创建权限请求
export interface CreatePermissionRequest {
  name: string
  code: string
  desc?: string
}

// 更新权限请求
export interface UpdatePermissionRequest {
  name?: string
  code?: string
  desc?: string
}

/**
 * 获取权限列表
 * @param params - 查询参数
 * @returns Promise
 */
export function getPermissionList(params?: {
  page?: number
  page_size?: number
  keyword?: string
}): Promise<ResponseData<PermissionListResponse>> {
  return request({
    url: '/api/permissions',
    method: 'get',
    params
  })
}

/**
 * 获取权限详情
 * @param id - 权限ID
 * @returns Promise
 */
export function getPermission(id: number): Promise<ResponseData<PermissionInfo>> {
  return request({
    url: `/api/permissions/${id}`,
    method: 'get'
  })
}

/**
 * 创建权限
 * @param data - 权限数据
 * @returns Promise
 */
export function createPermission(data: CreatePermissionRequest): Promise<ResponseData<PermissionInfo>> {
  return request({
    url: '/api/permissions',
    method: 'post',
    data
  })
}

/**
 * 更新权限
 * @param id - 权限ID
 * @param data - 权限数据
 * @returns Promise
 */
export function updatePermission(id: number, data: UpdatePermissionRequest): Promise<ResponseData<PermissionInfo>> {
  return request({
    url: `/api/permissions/${id}`,
    method: 'put',
    data
  })
}

/**
 * 删除权限
 * @param id - 权限ID
 * @returns Promise
 */
export function deletePermission(id: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/permissions/${id}`,
    method: 'delete'
  })
}

