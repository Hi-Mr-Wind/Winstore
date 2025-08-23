<template>
  <main>
  <div class="app-container">
    <!-- 侧边栏 -->
    <div class="sidebar">
      <div class="store-logo">
        <el-icon size="24"><Platform /></el-icon>
        <span>win store</span>
      </div>

      <div class="nav-menu">
        <div class="nav-item" :class="{ active: activeNav === 'home' }" @click="activeNav = 'home'">
          <el-icon><HomeFilled /></el-icon>
          <span>首页</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'popular' }" @click="activeNav = 'popular'">
          <el-icon><MagicStick /></el-icon>
          <span>热门软件</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'categories' }" @click="activeNav = 'categories'">
          <el-icon><Grid /></el-icon>
          <span>分类</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'installed' }" @click="activeNav = 'installed'">
          <el-icon><Download /></el-icon>
          <span>已安装软件</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'favorites' }" @click="activeNav = 'favorites'">
          <el-icon><Star /></el-icon>
          <span>收藏室</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'tools' }" @click="activeNav = 'tools'">
          <el-icon><Box /></el-icon>
          <span>工具箱</span>
        </div>

        <div class="category-title">应用分类</div>

        <div class="nav-item" v-for="category in appCategories" :key="category.id"
             :class="{ active: activeNav === category.id }" @click="activeNav = category.id">
          <el-icon><component :is="category.icon" /></el-icon>
          <span>{{ category.name }}</span>
        </div>

        <div class="category-title">其他</div>

        <div class="nav-item" :class="{ active: activeNav === 'settings' }" @click="activeNav = 'settings'">
          <el-icon><Setting /></el-icon>
          <span>设置</span>
        </div>
        <div class="nav-item" :class="{ active: activeNav === 'help' }" @click="activeNav = 'help'">
          <el-icon><QuestionFilled /></el-icon>
          <span>帮助与反馈</span>
        </div>
      </div>
    </div>

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 搜索框 -->
      <div class="search-container">
        <el-input
            v-model="searchQuery"
            placeholder="搜索应用、游戏、程序和内容..."
            size="large"
            :prefix-icon="Search"
            clearable
        />
      </div>

      <!-- 宣传横幅 -->
      <div class="hero-section">
        <h2 class="hero-title">发现精彩应用</h2>
        <p class="hero-description">为大众提供一个官方的，简单的正版软件下载渠道</p>
        <el-button type="primary" size="large" @click="exploreApps">立即探索</el-button>
      </div>

      <!-- 热门软件 -->
      <h2 class="section-title">热门软件</h2>
      <div class="apps-grid">
        <div class="app-card" v-for="app in filteredApps" :key="app.id">
          <div class="app-image" :style="{ backgroundColor: app.color }">
            {{ app.name }}
          </div>
          <div class="app-details">
            <h3 class="app-name">{{ app.name }}</h3>
            <p class="app-developer">{{ app.developer }}</p>
            <div class="app-rating">
              <span class="stars">★★★★☆</span>
              <span>{{ app.rating }}</span>
            </div>
            <div class="app-actions">
              <el-button v-if="!app.installed" type="primary" size="small" @click="installApp(app)">查看详情</el-button>
              <el-button v-else type="info" size="small" disabled>已安装</el-button>
              <span v-if="app.installed" class="installed-badge">已安装</span>
            </div>
          </div>
        </div>
      </div>

      <div class="view-all">
        <el-button type="primary" link @click="showAllApps">查看全部</el-button>
      </div>
    </div>
  </div>
  </main>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  Search,
  HomeFilled,
  Grid,
  Download,
  Star,
  Box,
  Setting,
  QuestionFilled,
  Platform, Fries, MagicStick
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

// 响应式数据
const activeNav = ref('popular')
const searchQuery = ref('')
const apps = ref([])

// 应用分类
const appCategories = ref([
  { id: 'productivity', name: '生产力', icon: 'Briefcase' },
  { id: 'social', name: '社交', icon: 'ChatRound' },
  { id: 'games', name: '游戏', icon: 'VideoPlay' },
  { id: 'entertainment', name: '娱乐', icon: 'Film' },
  { id: 'education', name: '教育', icon: 'Notebook' },
  { id: 'tools', name: '工具', icon: 'Tools' }
])

// 过滤应用列表
const filteredApps = computed(() => {
  if (!searchQuery.value) {
    return apps.value
  }
  return apps.value.filter(app =>
      app.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      app.developer.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// 初始化应用数据
const initApps = () => {
  apps.value = [
    { id: 1, name: 'Visual Studio Code', developer: 'Microsoft Corporation', rating: '4.7', installed: false, color: '#0078d4' },
    { id: 2, name: 'Microsoft Edge', developer: 'Microsoft Corporation', rating: '4.2', installed: true, color: '#107c10' },
    { id: 3, name: 'Spotify', developer: 'Spotify AB', rating: '4.5', installed: false, color: '#1ed760' },
    { id: 4, name: 'Adobe Photoshop', developer: 'Adobe Inc.', rating: '4.8', installed: false, color: '#001e36' },
    { id: 5, name: 'Zoom', developer: 'Zoom Video Communications', rating: '4.3', installed: true, color: '#2d8cff' },
    { id: 6, name: 'Netflix', developer: 'Netflix, Inc.', rating: '4.6', installed: false, color: '#e50914' }
  ]
}

// 安装应用
const installApp = (app) => {
  app.installed = true
  ElMessage.success(`已成功安装 ${app.name}`)
}

// 探索应用
const exploreApps = () => {
  ElMessage({
    message: 'Congrats, this is a success message.',
    type: 'success',
  })
  console.log('Explore apps')
}

// 查看全部应用
const showAllApps = () => {
  ElMessage({
    message: '查看全部应用',
    type: 'success',
  })
}

// 组件挂载时初始化数据
onMounted(() => {
  initApps()
})
</script>

<style scoped>
.app-container {
  display: flex;
  min-height: 100vh;
  background-color: #f5f5f5;
}

/* 侧边栏样式 */
.sidebar {
  width: 280px;
  background: linear-gradient(135deg, #2996ea 0%, #106ebe 100%);
  color: white;
  padding: 20px 0;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.1);
}

.store-logo {
  padding: 0 24px 20px;
  font-size: 24px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 10px;
}

.nav-menu {
  padding: 10px 0;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 12px 24px;
  cursor: pointer;
  transition: all 0.3s;
  gap: 12px;
}

.nav-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.nav-item.active {
  background-color: rgba(255, 255, 255, 0.2);
  border-left: 4px solid #fff;
}

.nav-item .el-icon {
  font-size: 20px;
}

.category-title {
  padding: 20px 24px 10px;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
}

/* 主内容区域样式 */
.main-content {
  flex: 1;
  padding: 20px 30px;
  overflow-y: auto;
}

.search-container {
  margin-bottom: 30px;
}

.hero-section {
  background: linear-gradient(to right, #4b6cb7, #182848);
  border-radius: 12px;
  padding: 40px;
  color: white;
  margin-bottom: 30px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.hero-title {
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 15px;
}

.hero-description {
  font-size: 16px;
  margin-bottom: 20px;
  opacity: 0.9;
  max-width: 600px;
}

.section-title {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 20px;
  color: #1f1f1f;
}

.apps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.app-card {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  transition: transform 0.3s, box-shadow 0.3s;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.app-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
}

.app-image {
  height: 160px;
  background-color: #0078d4;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 18px;
  font-weight: 500;
}

.app-details {
  padding: 16px;
}

.app-name {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 6px;
}

.app-developer {
  font-size: 14px;
  color: #666;
  margin-bottom: 10px;
}

.app-rating {
  display: flex;
  align-items: center;
  gap: 5px;
  margin-bottom: 15px;
  color: #666;
}

.stars {
  color: #ffb900;
}

.app-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.view-all {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}

.installed-badge {
  background-color: #e1f5fe;
  color: #0288d1;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 14px;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 992px) {
  .sidebar {
    width: 80px;
  }

  .store-logo span,
  .nav-item span,
  .category-title {
    display: none;
  }

  .store-logo {
    justify-content: center;
    padding: 0 0 20px 0;
  }

  .nav-item {
    justify-content: center;
    padding: 12px 0;
  }
}

@media (max-width: 768px) {
  .app-container {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
    padding: 10px 0;
  }

  .store-logo {
    padding: 0 0 10px 0;
  }

  .nav-menu {
    display: flex;
    overflow-x: auto;
    padding: 5px 0;
  }

  .nav-item {
    flex-direction: column;
    padding: 8px 12px;
    gap: 5px;
    font-size: 12px;
  }

  .nav-item .el-icon {
    font-size: 18px;
  }

  .category-title {
    display: none;
  }

  .main-content {
    padding: 15px;
  }

  .hero-section {
    padding: 20px;
  }

  .hero-title {
    font-size: 22px;
  }

  .apps-grid {
    grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  }
}
</style>