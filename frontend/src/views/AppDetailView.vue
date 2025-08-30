<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import type { AppItem } from '@/types/app'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

// 获取当前应用
const currentApp = computed(() => {
  const appId = parseInt(route.params.id as string)
  return appStore.apps.find(app => app.id === appId)
})

// 处理应用操作
const handleInstall = () => {
  if (currentApp.value) {
    appStore.installApp(currentApp.value)
  }
}

const handleUninstall = () => {
  if (currentApp.value) {
    appStore.uninstallApp(currentApp.value)
  }
}

const handleFavorite = () => {
  if (currentApp.value) {
    appStore.toggleFavorite(currentApp.value)
  }
}

// 初始化数据
onMounted(() => {
  if (appStore.apps.length === 0) {
    // 如果还没有数据，重新加载
    const mockApps: AppItem[] = [
      {
        id: 1,
        name: 'Visual Studio Code',
        developer: 'Microsoft Corporation',
        icon: 'https://code.visualstudio.com/assets/images/code-stable.png',
        rating: 4.7,
        installed: false,
        description: '轻量级但功能强大的源代码编辑器',
        category: '开发工具',
        version: '1.85.0',
        size: '85.2 MB',
        releaseDate: '2023-12-01',
        lastUpdated: '2023-12-15'
      }
    ]
    appStore.setApps(mockApps)
  }
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <div v-if="currentApp" class="app-detail">
        <!-- 应用头部信息 -->
        <div class="app-header">
          <div class="app-icon">
            <el-avatar :size="120" :src="currentApp.icon" />
          </div>
          <div class="app-info">
            <h1 class="app-name">{{ currentApp.name }}</h1>
            <p class="app-developer">{{ currentApp.developer }}</p>
            <div class="app-rating">
              <el-rate
                :model-value="currentApp.rating"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value}"
              />
            </div>
            <div class="app-actions">
              <el-button
                v-if="currentApp.installed"
                type="info"
                size="large"
                @click="handleUninstall"
              >
                已安装
              </el-button>
              <el-button
                v-else
                type="primary"
                size="large"
                @click="handleInstall"
              >
                获取
              </el-button>
              <el-button
                :icon="appStore.isAppFavorited(currentApp.id) ? 'StarFilled' : 'Star'"
                circle
                size="large"
                @click="handleFavorite"
              />
            </div>
          </div>
        </div>

        <!-- 应用详细信息 -->
        <div class="app-details">
          <div class="detail-section">
            <h2>应用信息</h2>
            <div class="detail-grid">
              <div class="detail-item">
                <span class="label">版本：</span>
                <span>{{ currentApp.version || '1.0.0' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">大小：</span>
                <span>{{ currentApp.size || '未知' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">分类：</span>
                <span>{{ currentApp.category || '未分类' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">发布日期：</span>
                <span>{{ currentApp.releaseDate || '未知' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">最后更新：</span>
                <span>{{ currentApp.lastUpdated || '未知' }}</span>
              </div>
            </div>
          </div>

          <div class="detail-section">
            <h2>应用描述</h2>
            <p class="app-description">{{ currentApp.description }}</p>
          </div>

          <!-- 截图区域 -->
          <div class="detail-section">
            <h2>应用截图</h2>
            <div class="screenshots">
              <el-empty description="暂无截图" />
            </div>
          </div>
        </div>
      </div>

      <!-- 应用不存在 -->
      <el-empty 
        v-else 
        description="应用不存在"
      >
        <el-button type="primary" @click="router.push('/')">
          返回首页
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

.app-detail {
  max-width: 1000px;
}

.app-header {
  background-color: white;
  border-radius: 12px;
  padding: 32px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  gap: 24px;
}

.app-icon {
  flex-shrink: 0;
}

.app-info {
  flex: 1;
}

.app-name {
  font-size: 32px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.app-developer {
  font-size: 16px;
  color: #666;
  margin-bottom: 16px;
}

.app-rating {
  margin-bottom: 24px;
}

.app-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.app-details {
  display: grid;
  gap: 24px;
}

.detail-section {
  background-color: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.detail-section h2 {
  font-size: 20px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e0e0e0;
}

.detail-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
}

.detail-item .label {
  font-weight: 600;
  color: #666;
}

.app-description {
  font-size: 16px;
  line-height: 1.6;
  color: #333;
}

.screenshots {
  min-height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .app-header {
    flex-direction: column;
    text-align: center;
    padding: 24px;
  }
  
  .app-name {
    font-size: 24px;
  }
  
  .app-actions {
    justify-content: center;
  }
  
  .detail-grid {
    grid-template-columns: 1fr;
  }
}
</style>
