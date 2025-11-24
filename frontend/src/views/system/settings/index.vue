<template>
  <div class="settings-container p-4">
    <el-card>
      <template #header>
        <div class="card-header">
          <el-icon><Setting /></el-icon>
          <span class="ml-2">系统设置</span>
        </div>
      </template>

      <div class="settings-content">
        <!-- 外观设置 -->
        <div class="settings-section">
          <h3 class="section-title">外观设置</h3>
          <div class="settings-item">
            <div class="settings-item-label">
              <span>深色模式</span>
              <span class="settings-item-desc">切换深色/浅色主题</span>
            </div>
            <el-switch
              v-model="isDarkMode"
              :active-icon="Moon"
              :inactive-icon="Sunny"
            />
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore } from '@/stores/app'
import { Setting, Moon, Sunny } from '@element-plus/icons-vue'

const appStore = useAppStore()

// 深色模式状态
const isDarkMode = computed({
  get: () => appStore.isDarkMode,
  set: (value: boolean) => {
    appStore.setDarkMode(value)
  }
})
</script>

<style scoped lang="scss">
.settings-container {
  min-height: calc(100vh - 120px);
}

.card-header {
  display: flex;
  align-items: center;
  font-size: 16px;
  font-weight: 600;
}

.settings-content {
  padding: 20px 0;
}

.settings-section {
  margin-bottom: 32px;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e5e7eb;
}

.settings-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 0;
  border-bottom: 1px solid #f3f4f6;

  &:last-child {
    border-bottom: none;
  }
}

.settings-item-label {
  display: flex;
  flex-direction: column;
  gap: 4px;

  span:first-child {
    font-size: 14px;
    font-weight: 500;
    color: #374151;
  }
}

.settings-item-desc {
  font-size: 12px;
  color: #6b7280;
}

/* 深色模式样式 */
:global(.dark) {
  .section-title {
    border-bottom-color: #374151;
    color: #f9fafb;
  }

  .settings-item {
    border-bottom-color: #374151;
  }

  .settings-item-label {
    span:first-child {
      color: #f9fafb;
    }
  }

  .settings-item-desc {
    color: #9ca3af;
  }
}
</style>

