<template>
  <div class="header-container flex-between h-16 px-4 bg-white shadow-sm">
    <div class="flex-center">
      <el-icon
        class="cursor-pointer text-xl mr-4"
        @click="toggleSidebar"
      >
        <Fold v-if="!appStore.sidebarCollapsed" />
        <Expand v-else />
      </el-icon>
      <h1 class="text-xl font-bold text-gray-800">Admin-Gin-Vue</h1>
    </div>
    <div class="flex-center">
      <el-dropdown @command="handleCommand">
        <div class="flex-center cursor-pointer">
          <el-avatar
            :size="32"
            :src="userStore.avatar"
            class="mr-2"
          >
            {{ userStore.nickname.charAt(0) }}
          </el-avatar>
          <span class="mr-2">{{ userStore.nickname || userStore.username }}</span>
          <el-icon><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">个人中心</el-dropdown-item>
            <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useAppStore } from '@/stores/app'
import { useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { Fold, Expand, ArrowDown } from '@element-plus/icons-vue'

const userStore = useUserStore()
const appStore = useAppStore()
const router = useRouter()

const toggleSidebar = () => {
  appStore.toggleSidebar()
}

const handleCommand = async (command: string) => {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        await userStore.logout()
        router.push('/login')
      })
      .catch(() => {})
  } else if (command === 'profile') {
    // TODO: 跳转到个人中心
    console.log('个人中心')
  }
}
</script>

<style scoped>
.header-container {
  border-bottom: 1px solid #e5e7eb;
}
</style>

