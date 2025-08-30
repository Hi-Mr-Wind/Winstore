<!--
  应用详情页面组件
  功能：
  - 展示应用的详细信息（图标、名称、开发者、描述等）
  - 提供官方网站和下载按钮
  - 显示应用版本、大小、分类、发布日期等信息
  - 展示应用截图（目前为占位符）
  - 支持响应式设计
  - 移除评分系统，专注于应用信息展示
-->
<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/appStore'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import type { AppItem } from '@/types/app'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const { t } = useI18n()

// 获取当前应用数据
const currentApp = computed(() => {
  const appId = parseInt(route.params.id as string)
  return appStore.apps.find(app => app.id === appId)
})

// 处理官方网站按钮点击 - 在新窗口打开官方网站
const handleOfficialWebsite = () => {
  if (currentApp.value?.officialWebsite) {
    window.open(currentApp.value.officialWebsite, '_blank')
  }
}

// 处理下载按钮点击 - 在新窗口打开下载页面
const handleDownload = () => {
  if (currentApp.value?.downloadUrl) {
    window.open(currentApp.value.downloadUrl, '_blank')
  }
}

// 初始化数据
onMounted(() => {
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
      }
    ]
    appStore.setApps(mockApps)
  }
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏导航 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 应用详情内容 -->
      <div v-if="currentApp" class="app-detail">
        <!-- 应用头部信息区域 -->
        <div class="app-header">
          <!-- 应用图标 -->
          <div class="app-icon">
            <el-avatar :size="120" :src="currentApp.icon" />
          </div>
          
          <!-- 应用基本信息 -->
          <div class="app-info">
            <h1 class="app-name">{{ currentApp.name }}</h1>
            <p class="app-developer">{{ currentApp.developer }}</p>
            
            <!-- 操作按钮区域 -->
            <div class="app-actions">
              <!-- 官方网站按钮 -->
              <el-button
                v-if="currentApp.officialWebsite"
                type="default"
                size="large"
                @click="handleOfficialWebsite"
              >
                {{ t('app.officialWebsite') }}
              </el-button>
              
              <!-- 下载按钮 -->
              <el-button
                v-if="currentApp.downloadUrl"
                type="primary"
                size="large"
                @click="handleDownload"
              >
                {{ t('app.download') }}
              </el-button>
            </div>
          </div>
        </div>

        <!-- 应用详细信息区域 -->
        <div class="app-details">
          <!-- 应用信息区域 -->
          <div class="detail-section">
            <h2>{{ t('app.appInfo') }}</h2>
            <div class="detail-grid">
              <div class="detail-item">
                <span class="label">{{ t('app.version') }}：</span>
                <span>{{ currentApp.version || '1.0.0' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">{{ t('app.size') }}：</span>
                <span>{{ currentApp.size || '未知' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">{{ t('app.category') }}：</span>
                <span>{{ currentApp.category || '未分类' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">{{ t('app.releaseDate') }}：</span>
                <span>{{ currentApp.releaseDate || '未知' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">{{ t('app.lastUpdated') }}：</span>
                <span>{{ currentApp.lastUpdated || '未知' }}</span>
              </div>
            </div>
          </div>

          <!-- 应用描述区域 -->
          <div class="detail-section">
            <h2>{{ t('app.description') }}</h2>
            <p class="app-description">{{ currentApp.description }}</p>
          </div>

          <!-- 应用截图区域 -->
          <div class="detail-section">
            <h2>{{ t('app.screenshots') }}</h2>
            <div class="screenshots">
              <el-empty :description="t('app.noScreenshots')" />
            </div>
          </div>
        </div>
      </div>

      <!-- 应用不存在时的空状态 -->
      <el-empty 
        v-else 
        :description="t('app.notFound')"
      >
        <el-button type="primary" @click="router.push('/')">
          {{ t('app.backToHome') }}
        </el-button>
      </el-empty>
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

/* 应用详情容器 */
.app-detail {
  max-width: 1000px;
}

/* 应用头部信息区域 */
.app-header {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 32px;
  margin-bottom: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
  display: flex;
  gap: 24px;
}

/* 应用图标区域 */
.app-icon {
  flex-shrink: 0;
}

/* 应用信息区域 */
.app-info {
  flex: 1;
}

/* 应用名称样式 */
.app-name {
  font-size: 32px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

/* 开发者名称样式 */
.app-developer {
  font-size: 16px;
  color: var(--text-secondary);
  margin-bottom: 24px;
}

/* 操作按钮区域 */
.app-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* 应用详细信息区域 */
.app-details {
  display: grid;
  gap: 24px;
}

/* 详情区域样式 */
.detail-section {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
}

/* 详情区域标题 */
.detail-section h2 {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-primary);
}

/* 详情信息网格布局 */
.detail-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

/* 详情项样式 */
.detail-item {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
}

/* 详情项标签样式 */
.detail-item .label {
  font-weight: 600;
  color: var(--text-secondary);
}

/* 应用描述样式 */
.app-description {
  font-size: 16px;
  line-height: 1.6;
  color: var(--text-primary);
}

/* 截图区域样式 */
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
