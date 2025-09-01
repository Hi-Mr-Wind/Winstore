<!--
  首页视图组件
  功能：
  - 展示应用商店的主页
  - 包含搜索框、轮播横幅、热门应用展示
  - 按分类展示应用
  - 提供应用发现和浏览功能
-->
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import { useAppStore } from '@/stores/appStore'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem, CategoryItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

// 搜索关键词
const searchQuery = ref('')

// 处理搜索
const handleSearch = () => {
  if (searchQuery.value.trim()) {
    router.push(`/search?q=${encodeURIComponent(searchQuery.value.trim())}`)
  }
}

// 处理应用卡片点击 - 跳转到应用详情页
const handleAppClick = (app: AppItem) => {
  router.push(`/app/${app.id}`)
}

// 处理应用安装
const handleAppInstall = (app: AppItem) => {
  appStore.installApp(app)
}

// 处理应用卸载
const handleAppUninstall = (app: AppItem) => {
  appStore.uninstallApp(app)
}

// 处理应用收藏
const handleAppFavorite = (app: AppItem) => {
  appStore.toggleFavorite(app)
}

// 处理分类点击 - 跳转到分类页面
const handleCategoryClick = (category: CategoryItem) => {
  router.push(`/category/${category.id}`)
  
}

// 处理查看全部热门应用 - 跳转到热门页面
const handleViewAllPopular = () => {
  router.push('/popular')
}

// 初始化数据
onMounted(async () => {
  // 如果应用数据为空，则初始化模拟数据
  if (appStore.apps.length === 0) {
    const mockApps: AppItem[] = [
      {
        id: 1,
        name: 'Visual Studio Code',
        developer: 'Microsoft Corporation',
        icon: 'https://code.visualstudio.com/assets/images/code-stable.png',
        installed: false,
        description: '轻量级但功能强大的源代码编辑器，支持多种编程语言和丰富的扩展插件。',
        category: '开发工具',
        version: '1.85.0',
        size: '85.2 MB',
        releaseDate: '2023-12-01',
        lastUpdated: '2023-12-15',
        officialWebsite: 'https://code.visualstudio.com/',
        downloadUrl: 'https://code.visualstudio.com/download'
      },
      {
        id: 2,
        name: 'Chrome',
        developer: 'Google LLC',
        icon: 'https://www.google.com/chrome/static/images/chrome-logo.svg',
        installed: true,
        description: '快速、安全、免费的网页浏览器',
        category: '网络工具',
        version: '120.0.6099.109',
        size: '150.5 MB',
        releaseDate: '2023-11-15',
        lastUpdated: '2023-12-10',
        officialWebsite: 'https://www.google.com/chrome/',
        downloadUrl: 'https://www.google.com/chrome/download/'
      },
      {
        id: 3,
        name: 'Notion',
        developer: 'Notion Labs Inc',
        icon: 'https://www.notion.so/images/logo-ios.png',
        installed: false,
        description: '一体化工作空间，用于笔记、文档、项目和协作',
        category: '生产力',
        version: '2.1.0',
        size: '120.8 MB',
        releaseDate: '2023-11-20',
        lastUpdated: '2023-12-05',
        officialWebsite: 'https://www.notion.so/',
        downloadUrl: 'https://www.notion.so/desktop'
      }
    ]
    appStore.setApps(mockApps)
  }

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
        {
          id: 1,
          name: '生产力',
          icon: '💼',
          count: 12,
          description: '提高工作效率的工具',
          color: '#4CAF50'
        },
        {
          id: 2,
          name: '开发工具',
          icon: '💻',
          count: 8,
          description: '程序员必备的开发环境',
          color: '#795548'
        },
        {
          id: 3,
          name: '网络工具',
          icon: '🌐',
          count: 15,
          description: '网络浏览和通信工具',
          color: '#3F51B5'
        },
        {
          id: 4,
          name: '娱乐',
          icon: '🎵',
          count: 20,
          description: '音乐、视频、游戏等娱乐应用',
          color: '#E91E63'
        }
      ]
      appStore.setCategories(defaultCategories)
    }
  } catch (error) {
    console.error('获取分类数据失败，使用默认分类:', error)
    // 出错时使用默认分类
    const fallbackCategories: CategoryItem[] = [
      {
        id: 1,
        name: '生产力',
        icon: '💼',
        count: 12,
        description: '提高工作效率的工具',
        color: '#4CAF50'
      },
      {
        id: 2,
        name: '开发工具',
        icon: '💻',
        count: 8,
        description: '程序员必备的开发环境',
        color: '#795548'
      }
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
      <!-- 搜索区域 -->
      <div class="search-section">
        <el-input
          v-model="searchQuery"
          :placeholder="t('search.placeholder')"
          size="large"
          @keyup.enter="handleSearch"
        >
          <template #append>
            <el-button @click="handleSearch">
              <el-icon><Search /></el-icon>
            </el-button>
          </template>
        </el-input>
      </div>

      <!-- 横幅区域 -->
      <div class="banner-section">
        <div class="banner-content">
          <h1 class="banner-title">{{ t('home.title') }}</h1>
          <p v-if="t('home.subtitle')" class="banner-description">{{ t('home.subtitle') }}</p>
          <el-button type="primary" size="large" @click="handleViewAllPopular">
            {{ t('common.explore') }}
          </el-button>
        </div>
      </div>

      <!-- 热门应用区域 -->
      <div class="section">
        <div class="section-header">
          <h2 class="section-title">{{ t('home.popularApps') }}</h2>
          <el-button type="text" @click="handleViewAllPopular">
            {{ t('common.viewAll') }}
          </el-button>
        </div>
        <div class="apps-grid">
          <AppCard
            v-for="app in appStore.popularApps.slice(0, 6)"
            :key="app.id"
            :app="app"
            @click="handleAppClick"
            @install="handleAppInstall"
            @uninstall="handleAppUninstall"
            @favorite="handleAppFavorite"
          />
        </div>
      </div>

      <!-- 分类浏览区域 -->
      <div class="section">
        <h2 class="section-title">{{ t('home.browseByCategory') }}</h2>
        <div v-if="appStore.categories.length > 0" class="categories-grid">
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
        <div v-else class="empty-categories">
          <el-empty 
            :image-size="120"
            description="暂无分类数据"
          >
            <template #description>
              <p>正在加载分类数据...</p>
            </template>
          </el-empty>
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

/* 搜索区域样式 */
.search-section {
  margin-bottom: 32px;
}

/* 横幅区域样式 */
.banner-section {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
  background-size: 200% 200%;
  border-radius: 16px;
  padding: 48px;
  margin-bottom: 40px;
  color: white;
  text-align: center;
  position: relative;
  overflow: hidden;
  animation: gradientShift 8s ease-in-out infinite;
  box-shadow: 0 20px 40px rgba(59, 130, 246, 0.3), 0 8px 16px rgba(139, 92, 246, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.banner-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

.banner-section::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  right: -50%;
  bottom: -50%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: rotate 20s linear infinite;
  pointer-events: none;
}

.banner-content {
  position: relative;
  z-index: 2;
}

.banner-title {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 16px;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
  animation: titleGlow 4s ease-in-out infinite;
}

.banner-description {
  font-size: 18px;
  margin-bottom: 24px;
  opacity: 0.9;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

/* 横幅按钮样式 */
.banner-section .el-button {
  position: relative;
  z-index: 3;
  background: rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  backdrop-filter: blur(10px) !important;
  -webkit-backdrop-filter: blur(10px) !important;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94) !important;
  animation: buttonPulse 2s ease-in-out infinite;
}

.banner-section .el-button:hover {
  background: rgba(255, 255, 255, 0.3) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 8px 25px rgba(255, 255, 255, 0.3) !important;
}

/* 横幅动画效果 */
@keyframes gradientShift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}



@keyframes shimmer {
  0%, 100% {
    transform: translateX(-100%);
  }
  50% {
    transform: translateX(100%);
  }
}

@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

@keyframes titleGlow {
  0%, 100% {
    text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
  }
  50% {
    text-shadow: 0 4px 20px rgba(255, 255, 255, 0.4), 0 0 30px rgba(255, 255, 255, 0.2);
  }
}

@keyframes buttonPulse {
  0%, 100% {
    box-shadow: 0 4px 15px rgba(255, 255, 255, 0.2);
  }
  50% {
    box-shadow: 0 4px 25px rgba(255, 255, 255, 0.4);
  }
}

/* 内容区域通用样式 */
.section {
  margin-bottom: 40px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.section-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-primary);
}

/* 应用网格布局 */
.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

/* 分类网格布局 */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

/* 分类卡片样式 */
.category-card {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 20px;
  box-shadow: var(--shadow-light);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  border: 1px solid var(--border-primary);
  display: flex;
  align-items: center;
  gap: 12px;
}

.category-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

.category-icon {
  font-size: 32px;
  flex-shrink: 0;
}

.category-info {
  flex: 1;
}

.category-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.category-count {
  font-size: 14px;
  color: var(--text-secondary);
}

/* 空状态样式 */
.empty-categories {
  text-align: center;
  padding: 40px 20px;
  background-color: var(--bg-primary);
  border-radius: 12px;
  border: 1px solid var(--border-primary);
}

.empty-categories .el-empty__description {
  color: var(--text-secondary);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .banner-section {
    padding: 32px 24px;
  }
  
  .banner-title {
    font-size: 28px;
  }
  
  .banner-description {
    font-size: 16px;
  }
  
  .apps-grid {
    grid-template-columns: 1fr;
  }
  
  .categories-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
