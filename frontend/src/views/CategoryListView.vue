<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import type { CategoryItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

// 初始化数据
onMounted(() => {
  if (appStore.categories.length === 0) {
    const mockCategories: CategoryItem[] = [
      { id: 1, name: '生产力', icon: '💼', count: 156 },
      { id: 2, name: '社交', icon: '💬', count: 89 },
      { id: 3, name: '游戏', icon: '🎮', count: 234 },
      { id: 4, name: '娱乐', icon: '🎵', count: 123 },
      { id: 5, name: '教育', icon: '📚', count: 67 },
      { id: 6, name: '工具', icon: '🔧', count: 98 }
    ]
    appStore.setCategories(mockCategories)
  }
})

const handleCategoryClick = (category: CategoryItem) => {
  router.push(`/category/${category.id}`)
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
        <h1 class="page-title">{{ t('common.category') }}</h1>
        <p class="page-description">{{ t('home.browseByCategory') }}</p>
      </div>

      <!-- 分类网格 -->
      <div class="categories-grid">
        <div 
          v-for="category in appStore.categories" 
          :key="category.id" 
          class="category-card"
          @click="handleCategoryClick(category)"
        >
          <div class="category-icon">{{ category.icon }}</div>
          <div class="category-info">
            <h3 class="category-name">{{ category.name }}</h3>
            <p class="category-count">{{ category.count }} {{ t('home.apps') }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background-color: var(--bg-secondary);
}

.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

.page-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.page-description {
  font-size: 16px;
  color: var(--text-secondary);
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.category-card {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--shadow-light);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  border: 1px solid var(--border-primary);
  display: flex;
  align-items: center;
  gap: 16px;
}

.category-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

.category-icon {
  font-size: 48px;
  flex-shrink: 0;
}

.category-info {
  flex: 1;
}

.category-name {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.category-count {
  font-size: 14px;
  color: var(--text-secondary);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .categories-grid {
    grid-template-columns: 1fr;
  }
  
  .category-card {
    padding: 20px;
  }
  
  .category-icon {
    font-size: 36px;
  }
}
</style>
