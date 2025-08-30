import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/views/HomeView.vue'),
    meta: {
      title: '首页 - Win Store'
    }
  },
  {
    path: '/popular',
    name: 'popular',
    component: () => import('@/views/PopularView.vue'),
    meta: {
      title: '热门软件 - Win Store'
    }
  },
  {
    path: '/category',
    name: 'category-list',
    component: () => import('@/views/CategoryListView.vue'),
    meta: {
      title: '分类 - Win Store'
    }
  },
  {
    path: '/category/:id',
    name: 'category',
    component: () => import('@/views/CategoryView.vue'),
    meta: {
      title: '分类 - Win Store'
    }
  },
  {
    path: '/installed',
    name: 'installed',
    component: () => import('@/views/InstalledView.vue'),
    meta: {
      title: '已安装软件 - Win Store'
    }
  },
  {
    path: '/favorites',
    name: 'favorites',
    component: () => import('@/views/FavoritesView.vue'),
    meta: {
      title: '收藏室 - Win Store'
    }
  },
  {
    path: '/tools',
    name: 'tools',
    component: () => import('@/views/ToolsView.vue'),
    meta: {
      title: '工具箱 - Win Store'
    }
  },
  {
    path: '/settings',
    name: 'settings',
    component: () => import('@/views/SettingsView.vue'),
    meta: {
      title: '设置 - Win Store'
    }
  },
  {
    path: '/help',
    name: 'help',
    component: () => import('@/views/HelpView.vue'),
    meta: {
      title: '帮助与反馈 - Win Store'
    }
  },
  {
    path: '/app/:id',
    name: 'app-detail',
    component: () => import('@/views/AppDetailView.vue'),
    meta: {
      title: '应用详情 - Win Store'
    }
  },
  {
    path: '/search',
    name: 'search',
    component: () => import('@/views/SearchView.vue'),
    meta: {
      title: '搜索结果 - Win Store'
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// 路由守卫 - 设置页面标题
router.beforeEach((to, from, next) => {
  if (to.meta?.title) {
    document.title = to.meta.title as string
  }
  next()
})

export default router
