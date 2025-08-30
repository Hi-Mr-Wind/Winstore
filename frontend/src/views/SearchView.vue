<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()

// 搜索结果
const searchResults = computed(() => appStore.filteredApps)

// 处理搜索
const handleSearch = (query: string) => {
  appStore.setSearchQuery(query)
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
  appStore.clearSearch()
  router.push('/')
}
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
            :model-value="appStore.searchQuery"
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
            <span v-if="appStore.searchQuery">
              关于 "{{ appStore.searchQuery }}"
            </span>
          </p>
        </div>

        <!-- 应用网格 -->
        <div class="apps-grid">
          <AppCard
            v-for="app in searchResults"
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
          v-if="searchResults.length === 0 && appStore.searchQuery" 
          description="没有找到相关应用"
        >
          <el-button type="primary" @click="clearSearch">
            返回首页
          </el-button>
        </el-empty>

        <!-- 初始状态 -->
        <div v-if="!appStore.searchQuery" class="initial-state">
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
