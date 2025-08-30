<!--
  热门应用页面组件
  功能：
  - 展示最受欢迎的应用列表
  - 提供应用安装、卸载、收藏功能
  - 支持应用详情页跳转
  - 展示应用的基本信息和操作按钮
-->
<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

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
  // 热门软件页面使用假数据，不需要 Wails 环境
  console.log('热门软件页面加载，使用假数据')
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

      <!-- 热门应用列表 -->
      <div class="apps-section">
        <div v-if="appStore.popularApps.length > 0" class="apps-grid">
          <AppCard
            v-for="app in appStore.popularApps"
            :key="app.id"
            :app="app"
            @click="handleAppClick"
            @install="handleAppInstall"
            @uninstall="handleAppUninstall"
            @favorite="handleAppFavorite"
          />
        </div>
        
        <!-- 空状态显示 -->
        <div v-else class="empty-state">
          <el-empty :description="t('popular.noData')" />
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

/* 页面标题区域样式 */
.page-header {
  margin-bottom: 32px;
}

/* Wails 状态显示样式 */
.wails-status {
  margin-top: 16px;
  display: flex;
  align-items: center;
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

/* 应用区域样式 */
.apps-section {
  margin-bottom: 32px;
}

/* 应用网格布局 */
.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
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
