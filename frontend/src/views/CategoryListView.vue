<!--
  分类列表页面组件
  功能：
  - 展示所有应用分类的列表
  - 每个分类显示图标、名称和应用数量
  - 点击分类卡片跳转到对应的分类详情页
  - 支持响应式设计
  - 提供分类浏览功能
-->
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

// 处理分类卡片点击 - 跳转到分类详情页
const handleCategoryClick = (category: CategoryItem) => {
  router.push(`/category/${category.id}`)
}

// 初始化数据
onMounted(async () => {
  // 尝试从Go后端获取分类数据
  try {
    console.log('正在从Go后端获取分类数据...')
    const categories = await appStore.fetchCategoriesFromWails()
    if (categories.length > 0) {
      console.log('成功从Go后端获取分类数据:', categories)
      appStore.setCategories(categories)
    } else {
      console.log('Go后端返回空分类数据，使用默认分类')
      // 如果Go后端没有数据，使用默认分类
      const defaultCategories: CategoryItem[] = [
      ]
      appStore.setCategories(defaultCategories)
    }
  } catch (error) {
    console.error('获取分类数据失败，使用默认分类:', error)
    // 出错时使用默认分类
    const fallbackCategories: CategoryItem[] = [
      
    ]
    appStore.setCategories(fallbackCategories)
  }
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏导航 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 页面标题区域 -->
      <div class="page-header">
        <h1 class="page-title">{{ t('common.category') }}</h1>
        <p class="page-description">{{ t('home.browseByCategory') }}</p>
      </div>

      <!-- 分类网格区域 -->
      <div class="categories-grid">
        <div 
          v-for="category in appStore.categories" 
          :key="category.id" 
          class="category-card"
          @click="handleCategoryClick(category)"
        >
          <!-- 分类图标 -->
          <div class="category-icon">{{ category.icon }}</div>
          
          <!-- 分类信息 -->
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
/* 应用容器 - 整体布局 */
.app-container {
  display: flex;
  height: 100vh;
  background-color: var(--bg-secondary);
}

/* 主内容区域 */
.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

/* 页面标题区域 */
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

/* 分类网格布局 */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

/* 分类卡片样式 */
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

/* 分类卡片悬停效果 */
.category-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

/* 分类图标样式 */
.category-icon {
  font-size: 48px;
  flex-shrink: 0;
}

/* 分类信息区域 */
.category-info {
  flex: 1;
}

/* 分类名称样式 */
.category-name {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

/* 分类应用数量样式 */
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
