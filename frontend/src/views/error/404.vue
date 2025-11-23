<template>
  <div class="error-container flex-center h-screen">
    <div class="text-center">
      <h1 class="text-6xl font-bold mb-4">404</h1>
      <p class="text-gray-500 mb-4">页面不存在</p>
      <el-button type="primary" @click="$router.push('/')">返回首页</el-button>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>

