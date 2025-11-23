<template>
  <div v-if="!item.meta?.hidden">
    <el-menu-item
      v-if="!hasChildren || (hasChildren && onlyOneChild)"
      :index="resolvePath"
    >
      <el-icon v-if="item.meta?.icon">
        <component :is="getIconComponent(item.meta.icon)" />
      </el-icon>
      <template #title>{{ item.meta?.title }}</template>
    </el-menu-item>

    <el-sub-menu
      v-else
      :index="resolvePath"
    >
      <template #title>
        <el-icon v-if="item.meta?.icon">
          <component :is="getIconComponent(item.meta.icon)" />
        </el-icon>
        <span>{{ item.meta?.title }}</span>
      </template>
      <sidebar-item
        v-for="child in item.children"
        :key="child.path"
        :item="child"
        :base-path="resolvePath"
      />
    </el-sub-menu>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { RouteRecordRaw } from 'vue-router'
import { isExternal } from '@/utils/route'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

interface Props {
  item: RouteRecordRaw
  basePath: string
}

const props = defineProps<Props>()

// 获取图标组件
const getIconComponent = (iconName?: string) => {
  if (!iconName) return null
  return (ElementPlusIconsVue as any)[iconName] || null
}

const hasChildren = computed(() => {
  return props.item.children && props.item.children.length > 0
})

const onlyOneChild = computed(() => {
  // 如果父级路由有 redirect 属性，说明是父级菜单，应该显示为子菜单
  if (props.item.redirect) return false
  if (!props.item.children) return false
  const showingChildren = props.item.children.filter((child) => {
    return !child.meta?.hidden
  })
  return showingChildren.length === 1
})

const resolvePath = computed(() => {
  if (isExternal(props.item.path)) {
    return props.item.path
  }
  // 处理路径拼接
  if (props.item.path.startsWith('/')) {
    // 如果路径以 / 开头，直接使用
    return props.item.path
  }
  // 如果 basePath 是根路径，直接拼接
  if (props.basePath === '/') {
    return `/${props.item.path}`
  }
  // 其他情况，在 basePath 和 item.path 之间加 /
  return `${props.basePath}/${props.item.path}`
})
</script>

