import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 菜单信息
export interface MenuInfo {
  id: number
  name: string
  code: string
  desc: string
  parent_id: number
  path: string
  icon: string
  type: number // 1-菜单，2-按钮
  sort: number
  status: number // 1-启用，0-禁用
  created_at: number
  updated_at: number
}

// 菜单列表响应
export interface MenuListResponse {
  menus: MenuInfo[]
  total: number
  page: number
  page_size: number
}

// 创建菜单请求
export interface CreateMenuRequest {
  name: string
  code: string
  desc?: string
  parent_id?: number
  path?: string
  icon?: string
  type: number
  sort?: number
  status?: number
}

// 更新菜单请求（所有字段必填）
export interface UpdateMenuRequest {
  name: string
  code: string
  desc: string
  parent_id: number
  path: string
  icon: string
  type: number
  sort: number
  status: number
}

/**
 * 获取菜单列表
 * @param params - 查询参数
 * @returns Promise
 */
export function getMenuList(params?: {
  page?: number
  page_size?: number
  keyword?: string
  all?: boolean
}): Promise<ResponseData<MenuListResponse>> {
  return request({
    url: '/api/menus',
    method: 'get',
    params
  })
}

/**
 * 获取菜单详情
 * @param id - 菜单ID
 * @returns Promise
 */
export function getMenu(id: number): Promise<ResponseData<MenuInfo>> {
  return request({
    url: `/api/menus/${id}`,
    method: 'get'
  })
}

/**
 * 创建菜单
 * @param data - 菜单数据
 * @returns Promise
 */
export function createMenu(data: CreateMenuRequest): Promise<ResponseData<MenuInfo>> {
  return request({
    url: '/api/menus',
    method: 'post',
    data
  })
}

/**
 * 更新菜单
 * @param id - 菜单ID
 * @param data - 菜单数据
 * @returns Promise
 */
export function updateMenu(id: number, data: UpdateMenuRequest): Promise<ResponseData<MenuInfo>> {
  return request({
    url: `/api/menus/${id}`,
    method: 'put',
    data
  })
}

/**
 * 删除菜单
 * @param id - 菜单ID
 * @returns Promise
 */
export function deleteMenu(id: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/menus/${id}`,
    method: 'delete'
  })
}

