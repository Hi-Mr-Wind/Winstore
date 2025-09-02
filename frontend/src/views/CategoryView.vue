<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem, CategoryItem } from '@/types/app'
import { SelectPage } from '../../wailsjs/go/apps/App'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

// 分页状态
const currentPage = ref(1)
const pageSize = ref(12)
const total = ref(0)
const isLoading = ref(false)
const errorMessage = ref('')
const softwareList = ref<AppItem[]>([])

// 获取当前分类
const currentCategory = computed(() => {
  const categoryId = route.params.id as string
  return appStore.categories.find(cat => cat.id.toString() === categoryId)
})

// 加载分类软件列表
const loadCategorySoftware = async (page: number = 1) => {
  if (!currentCategory.value) return
  
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    const categoryId = currentCategory.value.id.toString()
    console.log(`[Category] 加载分类${categoryId}第${page}页软件，每页${pageSize.value}条`)
    
    const response = await SelectPage(page, pageSize.value, categoryId)
    
    console.log('[Category] SelectPage 返回:', response)
    
    // 解析返回数据（兼容多种后端字段）
    let data: any = response
    if (response?.data) {
      data = response.data
    } else if ((response as any)?.Data) {
      data = (response as any).Data
    }

    const toAppItem = (item: any, index: number): AppItem => ({
      id: item.app_id || item.AppID || item.id || index + 1,
      name: item.app_name || item.AppName || item.name || '未知软件',
      developer: item.company || item.Company || item.developer || '未知开发者',
      icon: item.app_icon || item.AppIcon || item.icon || 'https://via.placeholder.com/64x64?text=App',
      installed: false,
      description: item.brief_introduction || item.BriefIntroduction || item.description || '暂无描述',
      category: currentCategory.value?.name || '未分类',
      version: item.app_version || item.AppVersion || item.version || '未知版本',
      size: '未知大小',
      releaseDate: '',
      lastUpdated: item.update_time || item.UpdateTime || item.lastUpdated || '',
      officialWebsite: item.official_website || item.OfficialWebsite || item.officialWebsite || '',
      downloadUrl: item.download_url || item.DownloadURL || item.downloadUrl || ''
    })

    const possibleListKeys = ['list', 'List', 'items', 'Items', 'results', 'Results']

    let list: any[] | null = null
    if (Array.isArray(data)) {
      list = data
    } else if (data && typeof data === 'object') {
      for (const key of possibleListKeys) {
        if (Array.isArray((data as any)[key])) {
          list = (data as any)[key]
          break
        }
      }
    }

    if (Array.isArray(list)) {
      softwareList.value = list.map((item: any, index: number) => toAppItem(item, index))
      const countFromResponse = (data && (data.count ?? data.total ?? (data.Total as number))) as number | undefined
      const countFromCategory = currentCategory.value && typeof currentCategory.value.count === 'number'
        ? currentCategory.value.count
        : undefined
      // 根据后端返回或分类自带数量自动计算分页总数
      total.value = typeof countFromResponse === 'number'
        ? countFromResponse
        : (typeof countFromCategory === 'number' ? countFromCategory : ((page - 1) * pageSize.value + softwareList.value.length))
      currentPage.value = page
      console.log(`[Category] 成功加载${softwareList.value.length}条软件，总数: ${total.value}`)
    } else {
      console.warn('[Category] 未找到列表数据，数据为:', data)
      softwareList.value = []
      total.value = 0
    }
  } catch (error: any) {
    console.error('[Category] 加载分类软件失败:', error)
    errorMessage.value = error?.message || '加载分类软件失败'
    softwareList.value = []
  } finally {
    isLoading.value = false
  }
}

// 处理分页变化
const handlePageChange = (page: number) => {
  loadCategorySoftware(page)
}

// 监听分类变化，重新加载数据
watch(() => route.params.id, () => {
  currentPage.value = 1
  loadCategorySoftware(1)
}, { immediate: true })

// 初始化数据
onMounted(async () => {
  // 如果分类数据为空，尝试从Go后端获取
  if (appStore.categories.length === 0) {
    try {
      console.log('CategoryView: 分类数据为空，尝试从Go后端获取...')
      const categories = await appStore.fetchCategoriesFromWails()
      if (categories.length > 0) {
        console.log('CategoryView: 成功从Go后端获取分类数据:', categories)
        appStore.setCategories(categories)
      } else {
        console.log('CategoryView: Go后端返回空分类数据，使用默认分类')
        // 使用默认分类
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
      console.error('CategoryView: 获取分类数据失败，使用默认分类:', error)
      // 出错时使用默认分类
      const fallbackCategories: CategoryItem[] = [
      
      ]
      appStore.setCategories(fallbackCategories)
    }
  }
  
  if (currentCategory.value) {
    loadCategorySoftware(1)
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
              {{ currentCategory?.count || total }} 个应用
            </p>
          </div>
        </div>
      </div>

      <!-- 错误提示 -->
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
        <button class="retry-btn" @click="loadCategorySoftware(1)">重试</button>
      </div>

      <!-- 应用网格 -->
      <div v-if="isLoading && softwareList.length === 0" class="loading-state">
        <div class="loading-spinner"></div>
        <p>正在加载分类软件...</p>
      </div>
      
      <div v-else-if="softwareList.length > 0" class="apps-grid">
        <AppCard
          v-for="app in softwareList"
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
      <div v-else class="empty-state">
        <div class="empty-content">
          <div class="empty-icon">📱</div>
          <h3 class="empty-title">该分类下暂无应用</h3>
          <p class="empty-description">
            当前分类 "{{ currentCategory?.name }}" 下还没有应用，请稍后再来查看。
          </p>
          <el-button type="primary" @click="router.push('/')">
            返回首页
          </el-button>
        </div>
      </div>

      <!-- 分页组件 -->
      <div v-if="softwareList.length > 0" class="pagination-wrapper">
        <el-pagination
          v-model:current-page="currentPage"
          :page-size="pageSize"
          :total="total"
          layout="prev, pager, next, jumper"
          @current-change="handlePageChange"
          background
        />
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

.category-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.category-icon {
  font-size: 48px;
  width: 64px;
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-primary);
  border-radius: 16px;
  box-shadow: var(--shadow-light);
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

.retry-btn {
  background-color: #dc2626;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
}

.retry-btn:hover {
  background-color: #b91c1c;
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

.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
}

.empty-content {
  max-width: 400px;
  margin: 0 auto;
}

.empty-icon {
  font-size: 64px;
  margin-bottom: 24px;
  opacity: 0.6;
}

.empty-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.empty-description {
  font-size: 16px;
  color: var(--text-secondary);
  margin-bottom: 24px;
  line-height: 1.5;
}

/* 分页样式 */
.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 32px;
}

@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .apps-grid {
    grid-template-columns: 1fr;
  }
}
</style>
