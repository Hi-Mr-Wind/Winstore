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
import { SelectPage, SelectByName } from '../../wailsjs/go/apps/App'

const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

// 搜索关键词
const searchQuery = ref('')

// 热门软件数据
const popularApps = ref<AppItem[]>([])

// 加载热门软件数据
const loadPopularApps = async () => {
  try {
    console.log('[Home] 加载首页热门软件数据')
    const response = await SelectPage(1, 3, '') // 只获取前3个
    
    console.log('[Home] SelectPage 返回:', response)
    
    // 解析返回数据
    let data: any = response
    if (response?.data) {
      data = response.data
    }
    
    if (data && data.list && Array.isArray(data.list)) {
      // 转换为AppItem格式，使用正确的Go后端字段名称
      popularApps.value = data.list.map((item: any, index: number) => ({
        id: item.app_id || item.AppID || index + 1,
        name: item.app_name || item.AppName || '未知软件',
        developer: item.company || item.Company || '未知开发者',
        icon: item.app_icon || item.AppIcon || 'https://via.placeholder.com/64x64?text=App',
        installed: false,
        description: item.brief_introduction || item.BriefIntroduction || '暂无描述',
        category: item.classify || item.Classify || '未分类',
        version: item.app_version || item.AppVersion || '未知版本',
        size: '未知大小', // 后端没有提供大小字段
        releaseDate: '', // 后端没有提供发布日期字段
        lastUpdated: item.update_time || item.UpdateTime || '',
        officialWebsite: item.official_website || item.OfficialWebsite || '',
        downloadUrl: item.download_url || item.DownloadURL || ''
      }))
      
      console.log(`[Home] 成功加载${popularApps.value.length}个热门软件`)
    } else {
      console.warn('[Home] 返回数据格式不正确:', data)
      popularApps.value = []
    }
  } catch (error: any) {
    console.error('[Home] 加载热门软件失败:', error)
    popularApps.value = []
  }
}

// 处理搜索
const handleSearch = async () => {
  if (!searchQuery.value.trim()) return
  
  try {
    console.log('[Home] 搜索关键词:', searchQuery.value.trim())
    const response = await SelectByName(searchQuery.value.trim())
    
    console.log('[Home] SelectByName 返回:', response)
    
    // 解析返回数据
    let data: any = response
    if (response?.data) {
      data = response.data
    }
    
    if (Array.isArray(data)) {
      // 转换为AppItem格式并设置到store，使用正确的Go后端字段名称
      const searchResults: AppItem[] = data.map((item: any, index: number) => ({
        id: item.app_id || item.AppID || index + 1,
        name: item.app_name || item.AppName || '未知软件',
        developer: item.company || item.Company || '未知开发者',
        icon: item.app_icon || item.AppIcon || 'https://via.placeholder.com/64x64?text=App',
        installed: false, // 暂时设为false，后续可以检查是否已安装
        description: item.brief_introduction || item.BriefIntroduction || '暂无描述',
        category: item.classify || item.Classify || '未分类',
        version: item.app_version || item.AppVersion || '未知版本',
        size: '未知大小', // 后端没有提供大小字段
        releaseDate: '', // 后端没有提供发布日期字段
        lastUpdated: item.update_time || item.UpdateTime || '',
        officialWebsite: item.official_website || item.OfficialWebsite || '',
        downloadUrl: item.download_url || item.DownloadURL || ''
      }))
      
      appStore.setApps(searchResults)
      console.log(`[Home] 搜索完成，找到${searchResults.length}个结果`)
      
      // 跳转到搜索结果页面
      router.push(`/search?q=${encodeURIComponent(searchQuery.value.trim())}`)
    } else {
      console.warn('[Home] 搜索结果不是数组:', data)
      // 如果没有搜索结果，可以显示提示或跳转到搜索页面
      router.push(`/search?q=${encodeURIComponent(searchQuery.value.trim())}`)
    }
  } catch (error: any) {
    console.error('[Home] 搜索失败:', error)
    // 搜索失败时也跳转到搜索页面，让用户看到错误提示
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
  // 加载热门软件数据
  await loadPopularApps()

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
          name: '游戏',
          icon: '🎮',
          count: 25,
          description: '娱乐游戏应用',
          color: '#FF9800'
        },
        {
          id: 4,
          name: '多媒体',
          icon: '🎵',
          count: 15,
          description: '音乐、视频、图片处理',
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
      },
      {
        id: 3,
        name: '游戏',
        icon: '🎮',
        count: 25,
        description: '娱乐游戏应用',
        color: '#FF9800'
      },
      {
        id: 4,
        name: '多媒体',
        icon: '🎵',
        count: 15,
        description: '音乐、视频、图片处理',
        color: '#E91E63'
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
            v-for="app in popularApps"
            :key="app.id"
            :app="app"
            @click="handleAppClick"
            @install="handleAppInstall"
            @uninstall="handleAppUninstall"
            @favorite="handleAppFavorite"
            :show-favorite="false"
            :show-install="false"
            :show-view="true"
          />
        </div>
      </div>

      <!-- 分类浏览区域 -->
      <div class="section">
        <h2 class="section-title">{{ t('home.browseByCategory') }}</h2>
        
        <!-- 调试信息 -->
        <div v-if="appStore.categories.length === 0" class="debug-info">
          <p>分类数据为空，正在加载中...</p>
        </div>
        <!-- <div v-else class="debug-info">
          <p>已加载 {{ appStore.categories.length }} 个分类</p>
          <p>分类数据: {{ JSON.stringify(appStore.categories, null, 2) }}</p>
        </div> -->
        
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
  border-radius: 16px;
  padding: 48px;
  margin-bottom: 40px;
  color: white;
  text-align: center;
}

.banner-title {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 16px;
}

.banner-description {
  font-size: 18px;
  margin-bottom: 24px;
  opacity: 0.9;
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

/* 统一首页卡片图标外观（复用 AppCard 内部样式） */
.app-card .app-icon {
  width: 72px;
  height: 72px;
  border-radius: 16px;
  background: var(--icon-bg);
}

/* 分类网格布局 */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

/* 调试信息样式 */
.debug-info {
  background-color: #f0f9ff;
  border: 1px solid #0ea5e9;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 16px;
  font-size: 14px;
  color: #0369a1;
}

.debug-info p {
  margin: 4px 0;
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
