/**
 * 应用路由配置
 * 定义应用的所有页面路由，包括路由守卫和页面标题设置
 */
import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

// 路由配置数组
const routes: RouteRecordRaw[] = [
  // 首页路由
  {
    path: '/',
    name: 'home',
    component: () => import('@/views/HomeView.vue'),
    meta: {
      title: '首页 - Win Store'
    }
  },
  
  // 热门应用页面路由
  {
    path: '/popular',
    name: 'popular',
    component: () => import('@/views/PopularView.vue'),
    meta: {
      title: '热门软件 - Win Store'
    }
  },
  
  // 分类列表页面路由
  {
    path: '/category',
    name: 'category-list',
    component: () => import('@/views/CategoryListView.vue'),
    meta: {
      title: '分类 - Win Store'
    }
  },
  
  // 分类详情页面路由（带参数）
  {
    path: '/category/:id',
    name: 'category',
    component: () => import('@/views/CategoryView.vue'),
    meta: {
      title: '分类 - Win Store'
    }
  },
  
  // 已安装应用页面路由
  {
    path: '/installed',
    name: 'installed',
    component: () => import('@/views/InstalledView.vue'),
    meta: {
      title: '已安装软件 - Win Store'
    }
  },
  
  // 收藏应用页面路由
  {
    path: '/favorites',
    name: 'favorites',
    component: () => import('@/views/FavoritesView.vue'),
    meta: {
      title: '收藏室 - Win Store'
    }
  },
  
  // 工具箱页面路由
  {
    path: '/tools',
    name: 'tools',
    component: () => import('@/views/ToolsView.vue'),
    meta: {
      title: '工具箱 - Win Store'
    }
  },
  
  // 设置页面路由
  {
    path: '/settings',
    name: 'settings',
    component: () => import('@/views/SettingsView.vue'),
    meta: {
      title: '设置 - Win Store'
    }
  },
  
  // 帮助页面路由
  {
    path: '/help',
    name: 'help',
    component: () => import('@/views/HelpView.vue'),
    meta: {
      title: '帮助与反馈 - Win Store'
    }
  },
  
  // 应用详情页面路由（带参数）
  {
    path: '/app/:id',
    name: 'app-detail',
    component: () => import('@/views/AppDetailView.vue'),
    meta: {
      title: '应用详情 - Win Store'
    }
  },
  
  // 搜索结果页面路由
  {
    path: '/search',
    name: 'search',
    component: () => import('@/views/SearchView.vue'),
    meta: {
      title: '搜索结果 - Win Store'
    }
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// 全局路由守卫 - 设置页面标题
router.beforeEach((to, from, next) => {
  // 如果路由有标题信息，则设置页面标题
  if (to.meta?.title) {
    document.title = to.meta.title as string
  }
  next()
})

export default router
