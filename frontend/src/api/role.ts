import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 角色信息
export interface RoleInfo {
  id: number
  name: string
  code: string
  desc: string
  created_at: number
  updated_at: number
}

// 角色列表响应
export interface RoleListResponse {
  roles: RoleInfo[]
  total: number
  page: number
  page_size: number
}

// 创建角色请求
export interface CreateRoleRequest {
  name: string
  code: string
  desc?: string
}

// 更新角色请求
export interface UpdateRoleRequest {
  name?: string
  code?: string
  desc?: string
}

/**
 * 获取角色列表
 * @param params - 查询参数
 * @returns Promise
 */
export function getRoleList(params?: {
  page?: number
  page_size?: number
  keyword?: string
}): Promise<ResponseData<RoleListResponse>> {
  return request({
    url: '/api/roles',
    method: 'get',
    params
  })
}

/**
 * 获取角色详情
 * @param id - 角色ID
 * @returns Promise
 */
export function getRole(id: number): Promise<ResponseData<RoleInfo>> {
  return request({
    url: `/api/roles/${id}`,
    method: 'get'
  })
}

/**
 * 创建角色
 * @param data - 角色数据
 * @returns Promise
 */
export function createRole(data: CreateRoleRequest): Promise<ResponseData<RoleInfo>> {
  return request({
    url: '/api/roles',
    method: 'post',
    data
  })
}

/**
 * 更新角色
 * @param id - 角色ID
 * @param data - 角色数据
 * @returns Promise
 */
export function updateRole(id: number, data: UpdateRoleRequest): Promise<ResponseData<RoleInfo>> {
  return request({
    url: `/api/roles/${id}`,
    method: 'put',
    data
  })
}

/**
 * 删除角色
 * @param id - 角色ID
 * @returns Promise
 */
export function deleteRole(id: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/roles/${id}`,
    method: 'delete'
  })
}

