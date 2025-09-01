<!--
  热门应用页面组件
  功能：
  - 展示最受欢迎的应用列表
  - 提供应用安装、卸载、收藏功能
  - 支持应用详情页跳转
  - 展示应用的基本信息和操作按钮
  - 支持分页加载
-->
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'
import { SelectPage } from '../../wailsjs/go/apps/App'

const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

// 分页状态
const currentPage = ref(1)
const pageSize = ref(12)
const total = ref(0)
const isLoading = ref(false)
const errorMessage = ref('')
const softwareList = ref<AppItem[]>([])

// 加载软件列表
const loadSoftwareList = async (page: number = 1) => {
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    console.log(`[Popular] 加载第${page}页软件，每页${pageSize.value}条，分类ID: ""`)
    const response = await SelectPage(page, pageSize.value, '')
    
    console.log('[Popular] SelectPage 返回:', response)
    
    // 解析返回数据
    let data: any = response
    if (response?.data) {
      data = response.data
    }
    
    if (data && data.list && Array.isArray(data.list)) {
      // 转换为AppItem格式，使用正确的Go后端字段名称
      softwareList.value = data.list.map((item: any, index: number) => ({
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
      
      // 使用后端返回的总数
      total.value = data.count || 0
      currentPage.value = page
      
      console.log(`[Popular] 成功加载${softwareList.value.length}条软件，总数: ${total.value}`)
    } else {
      console.warn('[Popular] 返回数据格式不正确:', data)
      softwareList.value = []
      total.value = 0
    }
  } catch (error: any) {
    console.error('[Popular] 加载软件列表失败:', error)
    errorMessage.value = error?.message || '加载软件列表失败'
    softwareList.value = []
  } finally {
    isLoading.value = false
  }
}

// 处理分页变化
const handlePageChange = (page: number) => {
  loadSoftwareList(page)
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

// 初始化数据
onMounted(() => {
  console.log('[Popular] 热门软件页面加载，开始获取数据')
  loadSoftwareList(1)
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
        <h1 class="page-title">{{ t('popular.title') }}</h1>
        <p class="page-description">{{ t('popular.subtitle') }}</p>
      </div>

      <!-- 错误提示 -->
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
        <button class="retry-btn" @click="loadSoftwareList(1)">重试</button>
      </div>

      <!-- 热门应用列表 -->
      <div class="apps-section">
        <div v-if="isLoading && softwareList.length === 0" class="loading-state">
          <div class="loading-spinner"></div>
          <p>正在加载热门软件...</p>
        </div>
        
        <div v-else-if="softwareList.length > 0" class="apps-grid">
          <AppCard
            v-for="app in softwareList"
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
        
        <!-- 空状态显示 -->
        <div v-else class="empty-state">
          <el-empty :description="t('popular.noData')" />
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

/* 页面标题区域样式 */
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

/* 应用网格布局 */
.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}

/* 空状态样式 */
.empty-state {
  text-align: center;
  padding: 60px 20px;
}

/* 分页样式 */
.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 32px;
}

/* 响应式设计 */
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
