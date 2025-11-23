<template>
  <div class="dashboard-container p-4">
    <h1 class="text-2xl font-bold mb-6">仪表盘</h1>

    <!-- 统计卡片 -->
    <el-row :gutter="20" class="mb-6">
      <!-- 用户统计 -->
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon user-icon">
              <el-icon :size="40">
                <User />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.user_total || 0 }}</div>
              <div class="stat-label">用户总数</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon today-icon">
              <el-icon :size="40">
                <UserFilled />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.user_today || 0 }}</div>
              <div class="stat-label">今日新增</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon enabled-icon">
              <el-icon :size="40">
                <CircleCheck />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.user_enabled || 0 }}</div>
              <div class="stat-label">启用用户</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon disabled-icon">
              <el-icon :size="40">
                <CircleClose />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.user_disabled || 0 }}</div>
              <div class="stat-label">禁用用户</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 系统统计 -->
    <el-row :gutter="20" class="mb-6">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon role-icon">
              <el-icon :size="40">
                <Avatar />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.role_total || 0 }}</div>
              <div class="stat-label">角色总数</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon menu-icon">
              <el-icon :size="40">
                <Menu />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.menu_total || 0 }}</div>
              <div class="stat-label">菜单总数</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon log-icon">
              <el-icon :size="40">
                <Document />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.operation_log_today || 0 }}</div>
              <div class="stat-label">今日操作日志</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-content">
            <div class="stat-icon login-icon">
              <el-icon :size="40">
                <Key />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.login_log_today || 0 }}</div>
              <div class="stat-label">今日登录日志</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 欢迎信息 -->
    <el-card>
      <div class="welcome-content">
        <h2 class="text-xl font-bold mb-2">欢迎使用 Admin-Gin-Vue 管理系统</h2>
        <p class="text-gray-600">系统运行正常，所有功能正常使用</p>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getDashboardStats, type DashboardStats } from '@/api/dashboard'
import { ElMessage } from 'element-plus'
import {
  User,
  UserFilled,
  CircleCheck,
  CircleClose,
  Avatar,
  Menu,
  Document,
  Key
} from '@element-plus/icons-vue'

const stats = ref<DashboardStats>({
  user_total: 0,
  user_today: 0,
  user_enabled: 0,
  user_disabled: 0,
  role_total: 0,
  menu_total: 0,
  operation_log_today: 0,
  login_log_today: 0
})

const loading = ref(false)

// 获取统计数据
const fetchStats = async () => {
  loading.value = true
  try {
    const res = await getDashboardStats()
    if (res.code === 200) {
      stats.value = res.data
    } else {
      ElMessage.error(res.message || '获取统计数据失败')
    }
  } catch (error: any) {
    console.error('获取统计数据失败:', error)
    ElMessage.error('获取统计数据失败')
  } finally {
    loading.value = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchStats()
})
</script>

<style scoped>
.dashboard-container {
  min-height: calc(100vh - 84px);
}

.stat-card {
  cursor: pointer;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.stat-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.user-icon {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.today-icon {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.enabled-icon {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.disabled-icon {
  background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
}

.role-icon {
  background: linear-gradient(135deg, #30cfd0 0%, #330867 100%);
}

.menu-icon {
  background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%);
  color: #333;
}

.log-icon {
  background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%);
}

.login-icon {
  background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%);
  color: #333;
}

.stat-info {
  flex: 1;
  text-align: right;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  line-height: 1;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.welcome-content {
  text-align: center;
  padding: 40px 0;
}
</style>