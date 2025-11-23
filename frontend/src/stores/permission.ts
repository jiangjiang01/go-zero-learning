import { defineStore } from 'pinia'

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    routes: [] as any[],
    menus: [] as any[],
    permissions: [] as string[]
  }),

  actions: {
    // 获取用户菜单（暂时返回空数组，等后端实现菜单接口后再补充）
    async getMenus() {
      try {
        // TODO: 等后端实现菜单接口后，调用 API 获取菜单
        this.menus = []
        return Promise.resolve({ data: [] })
      } catch (error) {
        return Promise.reject(error)
      }
    },

    // 设置权限
    setPermissions(permissions: string[]) {
      this.permissions = permissions
    }
  }
})

