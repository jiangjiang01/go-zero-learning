<template>
  <div class="app-main-container">
    <router-view v-slot="{ Component }">
      <transition
        name="fade-transform"
        mode="out-in"
      >
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.app-main-container {
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: #f0f2f5;
}

/* 页面切换动画 */
.fade-transform-leave-active,
.fade-transform-enter-active {
  transition: all 0.3s;
}

.fade-transform-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>

