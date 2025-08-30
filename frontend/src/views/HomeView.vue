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
import { Search, Grid } from '@element-plus/icons-vue'
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
    console.log('开始尝试从Go后端获取分类数据...')
    
    // 检查Wails环境
    if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
      console.log('Wails环境检测成功，window.go对象存在')
      
      const wailsCategories = await appStore.fetchCategoriesFromWails()
      if (wailsCategories.length > 0) {
        console.log('成功从Go后端获取分类数据:', wailsCategories)
        appStore.setCategories(wailsCategories)
      } else {
        console.log('从Go后端获取分类数据失败，分类列表为空')
        appStore.setCategories([]) // 设置为空数组，显示空状态
      }
    } else {
      console.log('Wails环境检测失败，分类列表为空')
      appStore.setCategories([]) // 设置为空数组，显示空状态
    }
  } catch (error) {
    console.error('获取分类数据时出错，分类列表为空:', error)
    appStore.setCategories([]) // 设置为空数组，显示空状态
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
          <p class="banner-description">{{ t('home.subtitle') }}</p>
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
            :description="t('home.noCategories')" 
            :image-size="120"
          >
            <template #image>
              <el-icon class="empty-icon"><Grid /></el-icon>
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 16px;
  padding: 48px;
  margin-bottom: 40px;
  color: white;
  text-align: center;
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
}

.banner-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse"><path d="M 10 0 L 0 0 0 10" fill="none" stroke="rgba(255,255,255,0.1)" stroke-width="0.5"/></pattern></defs><rect width="100" height="100" fill="url(%23grid)"/></svg>');
  opacity: 0.3;
}

.banner-section::after {
  content: '';
  position: absolute;
  top: -50%;
  right: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 70%);
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(180deg); }
}

.banner-content {
  position: relative;
  z-index: 1;
}

.banner-title {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 16px;
  text-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.banner-description {
  font-size: 18px;
  margin-bottom: 32px;
  opacity: 0.9;
  line-height: 1.6;
}

.banner-section .el-button {
  font-size: 16px;
  padding: 12px 32px;
  border-radius: 8px;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
  transition: all 0.3s ease;
}

.banner-section .el-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0,0,0,0.2);
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
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
  background-color: var(--bg-primary);
  border-radius: 12px;
  border: 2px dashed var(--border-primary);
}

.empty-icon {
  font-size: 80px;
  color: var(--text-tertiary);
  margin-bottom: 16px;
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
