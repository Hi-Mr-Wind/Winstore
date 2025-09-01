<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem, CategoryItem } from '@/types/app'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

// 获取当前分类
const currentCategory = computed(() => {
  const categoryId = route.params.id as string
  return appStore.categories.find(cat => cat.id.toString() === categoryId)
})

// 获取当前分类的应用
const categoryApps = computed(() => {
  if (!currentCategory.value) return []
  return appStore.apps.filter(app => app.category === currentCategory.value?.name)
})

// 初始化数据
onMounted(() => {
  if (appStore.categories.length === 0) {
    const mockCategories: CategoryItem[] = [
    ]
    appStore.setCategories(mockCategories)
  }
})

// 处理应用操作
const handleInstall = (app: AppItem) => {
  appStore.installApp(app)
}

const handleUninstall = (app: AppItem) => {
  appStore.uninstallApp(app)
}

const handleFavorite = (app: AppItem) => {
  appStore.toggleFavorite(app)
}

const handleAppClick = (app: AppItem) => {
  router.push(`/app/${app.id}`)
}
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 页面标题 -->
      <div class="page-header">
        <div class="category-info">
          <div class="category-icon">{{ currentCategory?.icon }}</div>
          <div>
            <h1 class="page-title">{{ currentCategory?.name || '分类' }}</h1>
            <p class="page-description">
              {{ currentCategory?.count || 0 }} 个应用
            </p>
          </div>
        </div>
      </div>

      <!-- 应用网格 -->
      <div class="apps-grid">
        <AppCard
          v-for="app in categoryApps"
          :key="app.id"
          :app="app"
          @install="handleInstall"
          @uninstall="handleUninstall"
          @favorite="handleFavorite"
          @click="handleAppClick"
        />
      </div>

      <!-- 空状态 -->
      <el-empty 
        v-if="categoryApps.length === 0" 
        description="该分类下暂无应用"
      />
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background-color: #f5f5f5;
}

.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

.page-header {
  margin-bottom: 32px;
}

.category-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.category-icon {
  font-size: 48px;
  background: white;
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.page-description {
  font-size: 16px;
  color: #666;
}

.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .category-icon {
    font-size: 36px;
    padding: 12px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .apps-grid {
    grid-template-columns: 1fr;
  }
}
</style>
