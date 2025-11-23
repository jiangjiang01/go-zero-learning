import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { usePermissionStore } from '@/stores/permission'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

// 配置 NProgress
NProgress.configure({ showSpinner: false })

// 路由配置
const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/login.vue'),
    meta: { title: '登录', hidden: true }
  },
  {
    path: '/',
    component: () => import('@/layout/index.vue'),
    redirect: '/dashboard',
    meta: { title: '首页', icon: 'HomeFilled' },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '仪表盘', icon: 'House' }
      }
    ]
  },
  {
    path: '/system',
    component: () => import('@/layout/index.vue'),
    redirect: '/system/user',
    meta: { title: '系统管理', icon: 'Setting' },
    children: [
      {
        path: 'user',
        name: 'SystemUser',
        component: () => import('@/views/system/user/index.vue'),
        meta: { title: '用户管理', icon: 'User' }
      },
      {
        path: 'role',
        name: 'SystemRole',
        component: () => import('@/views/system/role/index.vue'),
        meta: { title: '角色管理', icon: 'UserFilled' }
      }
    ]
  },
  {
    path: '/404',
    name: '404',
    component: () => import('@/views/error/404.vue'),
    meta: { title: '404', hidden: true }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach(async (to, _from, next) => {
  NProgress.start()

  const userStore = useUserStore()
  const permissionStore = usePermissionStore()

  // 判断是否已登录
  if (userStore.token) {
    // 已登录，访问登录页则跳转到首页
    if (to.path === '/login') {
      next({ path: '/' })
      NProgress.done()
    } else {
      // 获取用户信息
      if (!userStore.userInfo) {
        try {
          await userStore.getUserInfo()
        } catch (error) {
          // 获取用户信息失败，清除 token，跳转到登录页
          await userStore.logout()
          next({ path: '/login', query: { redirect: to.fullPath } })
          NProgress.done()
          return
        }
      }

      // 获取菜单
      if (permissionStore.menus.length === 0) {
        try {
          await permissionStore.getMenus()
        } catch (error) {
          console.error('获取菜单失败:', error)
        }
      }

      next()
    }
  } else {
    // 未登录
    if (to.path === '/login') {
      next()
    } else {
      // 需要登录的页面，跳转到登录页
      next({ path: '/login', query: { redirect: to.fullPath } })
    }
    NProgress.done()
  }
})

router.afterEach(() => {
  NProgress.done()
})

export default router

