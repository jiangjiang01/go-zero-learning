import { defineStore } from 'pinia'

// 从 localStorage 读取深色模式设置
const getDarkMode = (): boolean => {
  const stored = localStorage.getItem('darkMode')
  if (stored !== null) {
    return stored === 'true'
  }
  // 如果没有存储的值，检查系统偏好
  return window.matchMedia('(prefers-color-scheme: dark)').matches
}

export const useAppStore = defineStore('app', {
  state: () => ({
    sidebarCollapsed: false,
    loading: false,
    isDarkMode: getDarkMode()
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
    },

    // 切换深色模式
    toggleDarkMode() {
      this.isDarkMode = !this.isDarkMode
      localStorage.setItem('darkMode', String(this.isDarkMode))
      this.applyDarkMode()
    },

    // 设置深色模式
    setDarkMode(dark: boolean) {
      this.isDarkMode = dark
      localStorage.setItem('darkMode', String(this.isDarkMode))
      this.applyDarkMode()
    },

    // 应用深色模式
    applyDarkMode() {
      if (this.isDarkMode) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    },

    // 初始化深色模式（应用保存的设置）
    initDarkMode() {
      this.applyDarkMode()
    }
  }
})

