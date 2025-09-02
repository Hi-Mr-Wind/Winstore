<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'
import { SelectByName } from '../../wailsjs/go/apps/App'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

// 搜索状态
const searchQuery = ref('')
const searchResults = ref<AppItem[]>([])
const isLoading = ref(false)
const errorMessage = ref('')

// 从路由参数获取搜索关键词
const getSearchQueryFromRoute = () => {
  const query = route.query.q as string
  if (query) {
    searchQuery.value = query
    appStore.setSearchQuery(query)
    performSearch(query)
  }
}

// 执行搜索
const performSearch = async (query: string) => {
  if (!query.trim()) return
  
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    console.log('[Search] 执行搜索:', query)
    const response = await SelectByName(query.trim())
    
    console.log('[Search] SelectByName 返回:', response)
    
    // 解析返回数据 - 增强的数据解析逻辑
    console.log('[Search] 原始响应:', response)
    console.log('[Search] 响应类型:', typeof response)
    console.log('[Search] 响应结构:', Object.keys(response || {}))
    
    let data: any = response
    
    // 尝试多种数据格式
    if (response?.data) {
      data = response.data
      console.log('[Search] 使用 response.data:', data)
    } else if ((response as any)?.Data) {
      data = (response as any).Data
      console.log('[Search] 使用 response.Data:', data)
    } else if ((response as any)?.list) {
      data = (response as any).list
      console.log('[Search] 使用 response.list:', data)
    } else if ((response as any)?.List) {
      data = (response as any).List
      console.log('[Search] 使用 response.List:', data)
    }
    
    console.log('[Search] 最终数据:', data)
    console.log('[Search] 数据类型:', typeof data)
    console.log('[Search] 是否为数组:', Array.isArray(data))
    
    // 处理数组格式
    if (Array.isArray(data)) {
      // 转换为AppItem格式，使用正确的Go后端字段名称
      searchResults.value = data.map((item: any, index: number) => {
        console.log(`[Search] 处理项目 ${index}:`, item)
        
        return {
          id: item.app_id || item.AppID || item.id || index + 1,
          name: item.app_name || item.AppName || item.name || '未知软件',
          developer: item.company || item.Company || item.developer || '未知开发者',
          icon: item.app_icon || item.AppIcon || item.icon || 'https://via.placeholder.com/64x64?text=App',
          installed: false, // 暂时设为false，后续可以检查是否已安装
          description: item.brief_introduction || item.BriefIntroduction || item.description || '暂无描述',
          category: item.classify || item.Classify || item.category || '未分类',
          version: item.app_version || item.AppVersion || item.version || '未知版本',
          size: '未知大小', // 后端没有提供大小字段
          releaseDate: '', // 后端没有提供发布日期字段
          lastUpdated: item.update_time || item.UpdateTime || item.lastUpdated || '',
          officialWebsite: item.official_website || item.OfficialWebsite || item.officialWebsite || '',
          downloadUrl: item.download_url || item.DownloadURL || item.downloadUrl || ''
        }
      })
      
      console.log(`[Search] 搜索完成，找到${searchResults.value.length}个结果`)
      console.log('[Search] 转换后的结果:', searchResults.value)
    } else if (data && typeof data === 'object') {
      // 尝试从对象中提取数组
      console.log('[Search] 数据是对象，尝试提取数组')
      const possibleArrays = ['list', 'List', 'data', 'Data', 'items', 'Items', 'results', 'Results']
      
      for (const key of possibleArrays) {
        if (data[key] && Array.isArray(data[key])) {
          console.log(`[Search] 在 ${key} 中找到数组:`, data[key])
          searchResults.value = data[key].map((item: any, index: number) => ({
            id: item.app_id || item.AppID || item.id || index + 1,
            name: item.app_name || item.AppName || item.name || '未知软件',
            developer: item.company || item.Company || item.developer || '未知开发者',
            icon: item.app_icon || item.AppIcon || item.icon || 'https://via.placeholder.com/64x64?text=App',
            installed: false,
            description: item.brief_introduction || item.BriefIntroduction || item.description || '暂无描述',
            category: item.classify || item.Classify || item.category || '未分类',
            version: item.app_version || item.AppVersion || item.version || '未知版本',
            size: '未知大小',
            releaseDate: '',
            lastUpdated: item.update_time || item.UpdateTime || item.lastUpdated || '',
            officialWebsite: item.official_website || item.OfficialWebsite || item.officialWebsite || '',
            downloadUrl: item.download_url || item.DownloadURL || item.downloadUrl || ''
          }))
          console.log(`[Search] 从 ${key} 提取到 ${searchResults.value.length} 个结果`)
          break
        }
      }
      
      if (searchResults.value.length === 0) {
        console.warn('[Search] 无法从对象中提取数组数据:', data)
        errorMessage.value = '搜索结果格式不正确，无法解析数据'
      }
    } else {
      console.warn('[Search] 搜索结果格式不正确:', data)
      searchResults.value = []
      errorMessage.value = '搜索结果格式不正确'
    }
  } catch (error: any) {
    console.error('[Search] 搜索失败:', error)
    console.error('[Search] 错误详情:', {
      name: error?.name,
      message: error?.message,
      stack: error?.stack,
      response: error?.response
    })
    
    // 提供更具体的错误信息
    if (error?.message?.includes('Network')) {
      errorMessage.value = '网络连接失败，请检查网络设置'
    } else if (error?.message?.includes('timeout')) {
      errorMessage.value = '搜索超时，请重试'
    } else if (error?.response) {
      errorMessage.value = `搜索失败: ${error.response.status} - ${error.response.statusText}`
    } else {
      errorMessage.value = error?.message || '搜索失败，请重试'
    }
    
    searchResults.value = []
  } finally {
    isLoading.value = false
  }
}

// 处理搜索
const handleSearch = (query: string) => {
  searchQuery.value = query
  appStore.setSearchQuery(query)
  if (query.trim()) {
    performSearch(query)
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

const clearSearch = () => {
  searchQuery.value = ''
  searchResults.value = []
  errorMessage.value = ''
  appStore.clearSearch()
  router.push('/')
}

// 监听路由变化
watch(() => route.query.q, (newQuery) => {
  if (newQuery !== searchQuery.value) {
    getSearchQueryFromRoute()
  }
})

// 初始化
onMounted(() => {
  getSearchQueryFromRoute()
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 搜索框 -->
      <div class="search-header">
        <div class="search-container">
          <el-input
            :model-value="searchQuery"
            placeholder="搜索应用、游戏、程序和内容..."
            class="search-box"
            :prefix-icon="Search"
            size="large"
            @input="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-button @click="clearSearch">清除</el-button>
        </div>
      </div>

      <!-- 搜索结果 -->
      <div class="search-results">
        <div class="results-header">
          <h1 class="page-title">搜索结果</h1>
          <p class="results-count">
            找到 {{ searchResults.length }} 个结果
            <span v-if="searchQuery">
              关于 "{{ searchQuery }}"
            </span>
          </p>
        </div>

        <!-- 加载状态 -->
        <div v-if="isLoading" class="loading-state">
          <div class="loading-spinner"></div>
          <p>正在搜索...</p>
        </div>

        <!-- 错误提示 -->
        <div v-else-if="errorMessage" class="error-message">
          {{ errorMessage }}
          <el-button @click="performSearch(searchQuery)" type="primary" size="small">
            重试
          </el-button>
        </div>

        <!-- 应用网格 -->
        <div v-else-if="searchResults.length > 0" class="apps-grid">
          <AppCard
            v-for="app in searchResults"
            :key="app.id"
            :app="app"
            @install="handleInstall"
            @uninstall="handleUninstall"
            @favorite="handleFavorite"
            @click="handleAppClick"
            :show-favorite="false"
            :show-install="false"
            :show-view="true"
          />
        </div>

        <!-- 空状态 -->
        <el-empty 
          v-if="searchResults.length === 0 && searchQuery" 
          description="没有找到相关应用"
        >
          <el-button type="primary" @click="clearSearch">
            返回首页
          </el-button>
        </el-empty>

        <!-- 初始状态 -->
        <div v-if="!searchQuery" class="initial-state">
          <h2>开始搜索</h2>
          <p>在搜索框中输入关键词来查找应用</p>
        </div>
      </div>
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

.search-header {
  margin-bottom: 24px;
}

.search-container {
  display: flex;
  gap: 12px;
  align-items: center;
  max-width: 600px;
}

.search-box {
  flex: 1;
}

.search-results {
  max-width: 1200px;
}

.results-header {
  margin-bottom: 24px;
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.results-count {
  font-size: 16px;
  color: #666;
}

/* 加载状态样式 */
.loading-state {
  text-align: center;
  padding: 60px 20px;
  color: var(--text-secondary);
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid var(--accent-primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 16px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 错误提示样式 */
.error-message {
  background-color: #fef2f2;
  border: 1px solid #fecaca;
  color: #dc2626;
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.initial-state {
  text-align: center;
  padding: 60px 20px;
  color: #666;
}

.initial-state h2 {
  font-size: 24px;
  margin-bottom: 12px;
}

.initial-state p {
  font-size: 16px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .search-container {
    flex-direction: column;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .apps-grid {
    grid-template-columns: 1fr;
  }
}
</style>
