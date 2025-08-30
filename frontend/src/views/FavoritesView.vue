<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()

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
        <h1 class="page-title">收藏室</h1>
        <p class="page-description">
          您收藏的应用 ({{ appStore.favorites.length }})
        </p>
      </div>

      <!-- 应用网格 -->
      <div class="apps-grid">
        <AppCard
          v-for="app in appStore.favorites"
          :key="app.id"
          :app="app"
          @install="handleInstall"
          @uninstall="handleUninstall"
          @favorite="handleFavorite"
          @click="handleAppClick"
        />
      </div>

      <!-- 空状态 -->
      <el-empty 
        v-if="appStore.favorites.length === 0" 
        description="您还没有收藏任何应用"
      >
        <el-button type="primary" @click="router.push('/')">
          去发现应用
        </el-button>
      </el-empty>
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

.page-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.page-description {
  font-size: 16px;
  color: #666;
}

.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
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
