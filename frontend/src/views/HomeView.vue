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

// 搜索功能
const searchQuery = ref('')

// 初始化数据
onMounted(() => {
  // 模拟从API获取数据
  const mockApps: AppItem[] = [
    {
      id: 1,
      name: 'Visual Studio Code',
      developer: 'Microsoft Corporation',
      icon: 'https://code.visualstudio.com/assets/images/code-stable.png',
      rating: 4.7,
      installed: false,
      description: '轻量级但功能强大的源代码编辑器',
      category: '开发工具'
    },
    {
      id: 2,
      name: 'Microsoft Edge',
      developer: 'Microsoft Corporation',
      icon: 'https://edgestatic.azureedge.net/welcome/static/favicon.ico',
      rating: 4.2,
      installed: true,
      description: '快速、安全、现代化的浏览器',
      category: '网络工具'
    },
    {
      id: 3,
      name: 'Spotify',
      developer: 'Spotify AB',
      icon: 'https://open.scdn.co/cdn/images/favicon.32c72288.ico',
      rating: 4.8,
      installed: false,
      description: '音乐流媒体服务',
      category: '娱乐'
    },
    {
      id: 4,
      name: 'Discord',
      developer: 'Discord Inc.',
      icon: 'https://discord.com/assets/847541504914fd33810e70a0ea73177e.ico',
      rating: 4.5,
      installed: false,
      description: '游戏聊天和社区平台',
      category: '社交'
    },
    {
      id: 5,
      name: 'Steam',
      developer: 'Valve Corporation',
      icon: 'https://store.steampowered.com/favicon.ico',
      rating: 4.6,
      installed: false,
      description: '数字游戏发行平台',
      category: '游戏'
    }
  ]

  const mockCategories: CategoryItem[] = [
    { id: 1, name: '生产力', icon: '💼', count: 156 },
    { id: 2, name: '社交', icon: '💬', count: 89 },
    { id: 3, name: '游戏', icon: '🎮', count: 234 },
    { id: 4, name: '娱乐', icon: '🎵', count: 123 },
    { id: 5, name: '教育', icon: '📚', count: 67 },
    { id: 6, name: '工具', icon: '🔧', count: 98 }
  ]

  appStore.setApps(mockApps)
  appStore.setCategories(mockCategories)
})

// 处理搜索
const handleSearch = () => {
  if (searchQuery.value.trim()) {
    appStore.setSearchQuery(searchQuery.value)
    router.push('/search')
  }
}

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

const handleCategoryClick = (category: CategoryItem) => {
  router.push(`/category/${category.id}`)
}

const handleExploreClick = () => {
  router.push('/popular')
}
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 搜索框 -->
      <div class="search-container">
        <el-input
          v-model="searchQuery"
          :placeholder="t('search.placeholder')"
          class="search-box"
          :prefix-icon="Search"
          size="large"
          @keyup.enter="handleSearch"
        />
      </div>

      <!-- 横幅区域 -->
      <div class="banner">
        <div class="banner-content">
          <h2 class="banner-title">{{ t('home.title') }}</h2>
          <p class="banner-desc">
            {{ t('home.subtitle') }}
          </p>
          <el-button 
            type="primary" 
            size="large" 
            class="explore-btn"
            @click="handleExploreClick"
          >
            {{ t('common.explore') }} →
          </el-button>
        </div>
      </div>

      <!-- 热门软件区域 -->
      <div class="apps-section">
        <div class="section-header">
          <h3 class="section-title">{{ t('home.popularApps') }}</h3>
          <el-link 
            type="primary" 
            class="view-all-link"
            @click="router.push('/popular')"
          >
            {{ t('common.viewAll') }} >
          </el-link>
        </div>
        
        <div class="apps-grid">
          <AppCard
            v-for="app in appStore.popularApps.slice(0, 6)"
            :key="app.id"
            :app="app"
            @install="handleInstall"
            @uninstall="handleUninstall"
            @favorite="handleFavorite"
            @click="handleAppClick"
          />
        </div>
      </div>

      <!-- 按分类浏览区域 -->
      <div class="categories-section">
        <div class="section-header">
          <h3 class="section-title">{{ t('home.browseByCategory') }}</h3>
          <el-link 
            type="primary" 
            class="view-all-link"
            @click="router.push('/category')"
          >
            {{ t('common.viewAll') }} >
          </el-link>
        </div>
        
        <div class="categories-grid">
          <div 
            v-for="category in appStore.categories" 
            :key="category.id" 
            class="category-card"
            @click="handleCategoryClick(category)"
          >
            <div class="category-icon">{{ category.icon }}</div>
            <div class="category-info">
              <h4 class="category-name">{{ category.name }}</h4>
              <p class="category-count">{{ category.count }} {{ t('home.apps') }}</p>
            </div>
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

/* 主内容区域样式 */
.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

/* 搜索框样式 */
.search-container {
  margin-bottom: 24px;
}

.search-box {
  max-width: 600px;
}

/* 横幅样式 */
.banner {
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  border-radius: 12px;
  padding: 40px;
  margin-bottom: 32px;
  color: white;
}

.banner-content {
  max-width: 600px;
}

.banner-title {
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 16px;
}

.banner-desc {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 24px;
  opacity: 0.9;
}

.explore-btn {
  background-color: white;
  color: var(--accent-primary);
  border: none;
  font-weight: 600;
}

.explore-btn:hover {
  background-color: #f0f0f0;
}

/* 区域样式 */
.apps-section,
.categories-section {
  margin-bottom: 32px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.section-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-primary);
}

.view-all-link {
  font-weight: 500;
}

/* 应用卡片网格 */
.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

/* 分类卡片网格 */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.category-card {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  text-align: center;
  box-shadow: var(--shadow-light);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  border: 1px solid var(--border-primary);
}

.category-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

.category-icon {
  font-size: 32px;
  margin-bottom: 12px;
}

.category-name {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 4px;
  color: var(--text-primary);
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
  
  .banner {
    padding: 24px;
  }
  
  .banner-title {
    font-size: 24px;
  }
  
  .apps-grid {
    grid-template-columns: 1fr;
  }
  
  .categories-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  }
}
</style>
