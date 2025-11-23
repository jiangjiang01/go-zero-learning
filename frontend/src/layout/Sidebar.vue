<template>
  <div
    class="sidebar-container"
    :class="{ collapsed: appStore.sidebarCollapsed }"
  >
    <el-scrollbar>
      <el-menu
        :default-active="activeMenu"
        :collapse="appStore.sidebarCollapsed"
        :collapse-transition="false"
        router
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
      >
        <sidebar-item
          v-for="route in routes"
          :key="route.path"
          :item="route"
          :base-path="route.path"
        />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import SidebarItem from './SidebarItem.vue'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

// 获取路由列表（从 router 中获取，过滤掉隐藏的路由）
const routes = computed(() => {
  return router.getRoutes().filter((r) => {
    return r.meta && !r.meta.hidden && r.children && r.children.length > 0
  })
})

// 当前激活的菜单
const activeMenu = computed(() => {
  const { path } = route
  return path
})
</script>

<style scoped>
.sidebar-container {
  width: 210px;
  height: 100%;
  background-color: #304156;
  transition: width 0.28s;
}

.sidebar-container.collapsed {
  width: 64px;
}

:deep(.el-menu) {
  border-right: none;
}

:deep(.el-menu--collapse) {
  width: 64px;
}
</style>

