<template>
  <div class="dashboard-container p-4">
    <h1 class="text-2xl font-bold mb-6">仪表盘</h1>

    <!-- 欢迎卡片 -->
    <el-card class="mb-6 welcome-card">
      <div class="welcome-content">
        <h2 class="text-2xl font-bold mb-3">欢迎使用 Go-Zero-Learning 管理系统</h2>
        <p class="text-gray-600 mb-4">基于 go-zero 框架构建的微服务管理系统</p>
        <el-row :gutter="16">
          <el-col :span="8">
            <div class="feature-item">
              <User :size="24" class="feature-icon" />
              <div>
                <div class="feature-title">用户管理</div>
                <div class="feature-desc">完整的用户注册、登录、信息管理功能</div>
              </div>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="feature-item">
              <Lock :size="24" class="feature-icon" />
              <div>
                <div class="feature-title">安全认证</div>
                <div class="feature-desc">JWT Token 认证，保障系统安全</div>
              </div>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="feature-item">
              <Setting :size="24" class="feature-icon" />
              <div>
                <div class="feature-title">系统管理</div>
                <div class="feature-desc">统一的错误处理和响应格式</div>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-card>

    <!-- 快速操作 -->
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="8">
        <el-card shadow="hover" class="quick-card">
          <template #header>
            <div class="card-header">
              <User :size="18" />
              <span>用户管理</span>
            </div>
          </template>
          <div class="quick-content">
            <p class="text-gray-600 mb-4">管理系统的用户信息，包括用户列表、用户详情、用户编辑等功能。</p>
            <el-button type="primary" @click="goToUserManagement">前往用户管理</el-button>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="8">
        <el-card shadow="hover" class="quick-card">
          <template #header>
            <div class="card-header">
              <Setting :size="18" />
              <span>系统设置</span>
            </div>
          </template>
          <div class="quick-content">
            <p class="text-gray-600 mb-4">配置系统参数，包括 JWT 配置、日志配置、数据库配置等。</p>
            <el-button type="info" disabled>功能开发中</el-button>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :md="8">
        <el-card shadow="hover" class="quick-card">
          <template #header>
            <div class="card-header">
              <Document :size="18" />
              <span>系统信息</span>
            </div>
          </template>
          <div class="quick-content">
            <p class="text-gray-600 mb-4">查看系统运行状态、版本信息、API 文档等系统相关信息。</p>
            <el-button type="info" disabled>功能开发中</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import {
  User,
  Document,
  Lock,
  Setting
} from '@element-plus/icons-vue'

const router = useRouter()

// 前往用户管理
const goToUserManagement = () => {
  router.push('/system/user')
}
</script>

<style scoped>
.dashboard-container {
  min-height: calc(100vh - 84px);
}

.welcome-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.welcome-card :deep(.el-card__body) {
  color: white;
}

.welcome-content h2 {
  color: white;
}

.welcome-content .text-gray-600 {
  color: rgba(255, 255, 255, 0.9);
}

.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.feature-icon {
  color: white;
  margin-top: 4px;
}

.feature-title {
  font-weight: bold;
  color: white;
  margin-bottom: 4px;
}

.feature-desc {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.8);
  line-height: 1.5;
}

.quick-card {
  height: 100%;
  transition: all 0.3s;
}

.quick-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: bold;
}

.quick-content {
  min-height: 120px;
}
</style>