import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    sidebarCollapsed: false,
    loading: false
  }),

  actions: {
    // 切换侧边栏折叠状态
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed
    },

    // 设置侧边栏折叠状态
    setSidebarCollapsed(collapsed: boolean) {
      this.sidebarCollapsed = collapsed
    },

    // 设置加载状态
    setLoading(loading: boolean) {
      this.loading = loading
    }
  }
})

