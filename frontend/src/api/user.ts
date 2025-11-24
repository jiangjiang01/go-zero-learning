import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'
import type { UserInfo, LoginResponse } from './auth'
import type { RoleInfo } from './role'

// 用户列表响应（对接后端接口）
export interface UserListResponse {
  users: UserInfo[]
  total: number
  page: number
  page_size: number
}

// 创建用户请求（使用注册接口）
export interface CreateUserRequest {
  username: string
  email: string
  password: string
}

// 更新用户请求（对接后端接口）
export interface UpdateUserRequest {
  email?: string
  password?: string
}

/**
 * 获取用户列表
 * @param params - 查询参数
 * @returns Promise
 */
export function getUserList(params?: {
  page?: number
  page_size?: number
  keyword?: string
}): Promise<ResponseData<UserListResponse>> {
  return request({
    url: '/api/users',
    method: 'get',
    params
  })
}

/**
 * 获取用户详情
 * @param id - 用户ID
 * @returns Promise
 */
export function getUser(id: number): Promise<ResponseData<UserInfo>> {
  return request({
    url: `/api/users/${id}`,
    method: 'get'
  })
}

/**
 * 创建用户（使用注册接口）
 * @param data - 用户数据
 * @returns Promise
 */
export function createUser(data: CreateUserRequest): Promise<ResponseData<LoginResponse>> {
  return request({
    url: '/api/users',
    method: 'post',
    data
  })
}

/**
 * 更新当前用户信息
 * @param data - 用户数据
 * @returns Promise
 */
export function updateUser(data: UpdateUserRequest): Promise<ResponseData<UserInfo>> {
  return request({
    url: '/api/users/me',
    method: 'put',
    data
  })
}

/**
 * 更新指定用户信息
 * @param id - 用户ID
 * @param data - 用户数据
 * @returns Promise
 */
export function updateUserDetail(id: number, data: UpdateUserRequest): Promise<ResponseData<UserInfo>> {
  return request({
    url: `/api/users/${id}`,
    method: 'put',
    data
  })
}

/**
 * 删除用户
 * @param id - 用户ID
 * @returns Promise
 */
export function deleteUser(id: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/users/${id}`,
    method: 'delete'
  })
}

// ========== 用户角色管理 API ==========

// 用户角色列表响应
export interface UserRolesResponse {
  roles: RoleInfo[]
}

/**
 * 给用户分配角色
 * @param userId - 用户ID
 * @param roleId - 角色ID
 * @returns Promise
 */
export function assignUserRole(userId: number, roleId: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/users/${userId}/roles`,
    method: 'post',
    data: { role_id: roleId }
  })
}

/**
 * 获取用户角色列表
 * @param userId - 用户ID
 * @returns Promise
 */
export function getUserRoles(userId: number): Promise<ResponseData<UserRolesResponse>> {
  return request({
    url: `/api/users/${userId}/roles`,
    method: 'get'
  })
}

/**
 * 移除用户角色
 * @param userId - 用户ID
 * @param roleId - 角色ID
 * @returns Promise
 */
export function removeUserRole(userId: number, roleId: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/users/${userId}/roles/${roleId}`,
    method: 'delete'
  })
}

