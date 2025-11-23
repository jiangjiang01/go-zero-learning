<template>
  <div class="layout-container flex">
    <Sidebar />
    <div class="main-container flex flex-col flex-1">
      <Header />
      <AppMain />
    </div>
  </div>
</template>

<script setup lang="ts">
import Header from './Header.vue'
import Sidebar from './Sidebar.vue'
import AppMain from './AppMain.vue'
</script>

<style scoped>
.layout-container {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.main-container {
  overflow: hidden;
}
</style>

