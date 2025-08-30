<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import { useAppStore } from '@/stores/appStore'
import Sidebar from '@/components/Sidebar.vue'
import AppCard from '@/components/AppCard.vue'
import type { AppItem } from '@/types/app'

const router = useRouter()
const appStore = useAppStore()

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
      },
      {
        id: 6,
        name: 'Notion',
        developer: 'Notion Labs Inc.',
        icon: 'https://www.notion.so/images/favicon.ico',
        rating: 4.4,
        installed: false,
        description: '一体化工作空间',
        category: '生产力'
      },
      {
        id: 7,
        name: 'Slack',
        developer: 'Slack Technologies',
        icon: 'https://a.slack-edge.com/bv1-10/slack_logo-ebd02d1.svg',
        rating: 4.3,
        installed: false,
        description: '团队协作平台',
        category: '社交'
      },
      {
        id: 8,
        name: 'Zoom',
        developer: 'Zoom Video Communications',
        icon: 'https://st1.zoom.us/zoom.ico',
        rating: 4.1,
        installed: false,
        description: '视频会议软件',
        category: '社交'
      }
    ]
    appStore.setApps(mockApps)
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
        <h1 class="page-title">热门软件</h1>
        <p class="page-description">发现最受欢迎的应用和游戏</p>
      </div>

      <!-- 应用网格 -->
      <div class="apps-grid">
        <AppCard
          v-for="app in appStore.popularApps"
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
        v-if="appStore.popularApps.length === 0" 
        description="暂无热门软件"
      />
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
