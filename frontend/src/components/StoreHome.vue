<template>
  <div class="app-container">
    <!-- 顶部导航栏 -->
    <el-menu mode="horizontal" class="header-nav" :ellipsis="false">
      <div class="store-logo">Win Store</div>
      <div class="flex-grow">
<!--      <el-button @click="testRout">-->
<!--       测试路由-->
<!--      </el-button>-->
<!--        <router-link to="/store">测试路由</router-link>-->
      </div>
    </el-menu>

    <div class="main-container">
      <!-- 侧边栏 -->
      <div class="sidebar">
        <el-menu
            default-active="1"
            class="sidebar-menu"
            @select="handleMenuSelect"
        >
          <el-menu-item index="1"> <el-icon><HomeFilled /></el-icon>首页</el-menu-item>
          <el-menu-item index="2"> <el-icon><MagicStick /></el-icon>热门软件</el-menu-item>
          <el-menu-item index="3"><el-icon><Grid /></el-icon>分类</el-menu-item>
          <el-menu-item index="4"><el-icon><Download /></el-icon>已安装软件</el-menu-item>
          <el-menu-item index="5"><el-icon><Star /></el-icon>收藏室</el-menu-item>
          <el-menu-item index="6"><el-icon><Box /></el-icon>工具箱</el-menu-item>
          <el-sub-menu index="7">
            <template #title>
              <span>应用分类</span>
            </template>
            <el-menu-item index="7-1">生产力</el-menu-item>
            <el-menu-item index="7-2">社交</el-menu-item>
            <el-menu-item index="7-3">游戏</el-menu-item>
            <el-menu-item index="7-4">娱乐</el-menu-item>
            <el-menu-item index="7-5">教育</el-menu-item>
            <el-menu-item index="7-6">工具</el-menu-item>
          </el-sub-menu>
          <el-menu-item index="8">
            <el-icon><Setting /></el-icon>
            <span>设置</span>
          </el-menu-item>
          <el-menu-item index="9">
            <el-icon><QuestionFilled /></el-icon>
            <span>帮助与反馈</span>
          </el-menu-item>
        </el-menu>
      </div>

      <!-- 主内容区域 -->
      <div class="content">
        <!-- 搜索框 -->
        <div class="search-container">
          <el-input
              v-model="searchQuery"
              placeholder="搜索应用、游戏、程序和内容..."
              class="search-box"
              :prefix-icon="Search"
          />
        </div>

        <!-- 横幅区域 -->
        <div class="banner">
          <h2 class="banner-title">发现精彩应用</h2>
          <p class="banner-desc">
            探索为 Windows 11 打造的精炼应用和游戏，提升您的生产力和娱乐体验
          </p>
          <el-button type="primary" size="large">立即探索</el-button>
        </div>

        <!-- 热门应用区域 -->

        <div class="apps-section">
          <h3 class="section-title">热门软件</h3>
          <el-empty v-show="apps.length === 0" description="暂时没有数据哦"/>
          <el-scrollbar height="450px" width="100%">
          <el-row :gutter="20">
            <el-col :span="8" v-for="app in apps" :key="app.id">
              <el-card class="app-card" shadow="hover">
                <div class="app-content">
                  <div class="app-icon">
                    <el-avatar :size="60" :src="app.icon" />
                  </div>
                  <div class="app-info">
                    <h4 class="app-name">{{ app.name }}</h4>
                    <p class="app-developer">{{ app.developer }}</p>
                    <div class="app-rating">
                      <el-rate
                          v-model="app.rating"
                          disabled
                          show-score
                          text-color="#ff9900"
                          score-template="{value}"
                      />
<!--                      <span class="rating-value">{{ app.rating }}</span>-->
                    </div>
                  </div>
                  <div class="app-action">
                    <el-button
                        v-if="app.installed"
                        type="success"
                        size="small"
                        disabled
                    >
                      已安装
                    </el-button>
                    <el-button
                        v-else
                        type="primary"
                        size="small"
                    >
                      获取
                    </el-button>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
          <div class="view-all">
            <el-button type="text" @click="more">查看全部</el-button>
          </div>
          </el-scrollbar>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {defineComponent, ref} from 'vue'
import { useRouter } from 'vue-router'
import {
  Box,
  Download,
  Grid,
  HomeFilled,
  MagicStick,
  QuestionFilled,
  Search,
  Setting,
  Star
} from '@element-plus/icons-vue'
import {ElMessage} from "element-plus";
// 搜索查询
const searchQuery = ref('')

// 处理菜单选择
const handleMenuSelect = (index:string, indexPath:any) => {
  console.log('菜单选择:', index, indexPath)
}

// 展开更多
const more = ()=>{
  console.log('more')
  ElMessage({
    message: '展开更多',
    type: 'success',
  })
}

// 应用数据
const apps = ref(
    // []
    [
  {
    id: 1,
    name: 'Visual Studio Code',
    developer: 'Microsoft Corporation',
    rating: 4.0,
    installed: false,
    icon: 'https://code.visualstudio.com/favicon.ico'
  },
  {
    id: 2,
    name: 'Microsoft Edge',
    developer: 'Microsoft Corporation',
    rating: 4.2,
    installed: true,
    icon: 'https://www.microsoft.com/favicon.ico'
  },
  {
    id: 3,
    name: 'Spotify',
    developer: 'Spotify AB',
    rating: 4.8,
    installed: false,
    icon: 'https://www.spotify.com/favicon.ico'
  },
  {
    id: 4,
    name: 'Adobe Photoshop',
    developer: 'Adobe Inc.',
    rating: 4.6,
    installed: false,
    icon: 'https://www.adobe.com/favicon.ico'
  },
  {
    id: 5,
    name: 'Zoom',
    developer: 'Zoom Video Communications',
    rating: 4.3,
    installed: false,
    icon: 'https://zoom.us/favicon.ico'
  },
  {
    id: 6,
    name: 'Notion',
    developer: 'Notion Labs Inc.',
    rating: 4.9,
    installed: false,
    icon: 'https://www.notion.so/favicon.ico'
  }
]
)
const testRout = ()=>{
  console.log('testRout')
  const router = useRouter()
  // this.$router.replace({path: '/store'})
  router.replace({path: '/store'})
}

</script>

<style scoped>
.app-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

/* 顶部导航栏样式 */
.header-nav {
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 0 20px;
  display: flex;
  align-items: center;
}

.store-logo {
  font-weight: bold;
  font-size: 20px;
  color: #0078d7;
  margin-right: 30px;
  padding: 0 15px;
}

.flex-grow {
  flex-grow: 1;
}

.header-nav .el-menu-item {
  font-size: 15px;
  height: 60px;
  line-height: 60px;
}

.active-nav {
  color: #0078d7 !important;
  border-bottom: 2px solid #0078d7 !important;
}

/* 主内容区布局 */
.main-container {
  display: flex;
  margin-top: 20px;
  max-width: 1400px;
  margin-left: auto;
  margin-right: auto;
}

/* 侧边栏样式 */
.sidebar {
  width: 250px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-right: 20px;
  padding: 20px 0;
  height: fit-content;
}

.sidebar-menu {
  border-right: none;
}

.sidebar-menu .el-menu-item {
  height: 45px;
  line-height: 45px;
  font-size: 14px;
}

.sidebar-menu .el-sub-menu__title {
  height: 45px;
  line-height: 45px;
  font-size: 14px;
}

/* 内容区域样式 */
.content {
  flex: 1;
  padding: 0 20px 20px 0;
}

/* 搜索框样式 */
.search-container {
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

.search-box {
  width: 100%;
}

/* 横幅区域样式 */
.banner {
  background: linear-gradient(135deg, #0078d7, #6b69d6);
  border-radius: 8px;
  padding: 40px;
  color: white;
  margin-bottom: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.banner-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
}

.banner-desc {
  font-size: 16px;
  margin-bottom: 20px;
  opacity: 0.9;
}

/* 热门应用区域样式 */
.apps-section {
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid #5e5e5e;
  color: #5e5e5e;
}

/* 应用卡片样式 */
.app-card {
  margin-bottom: 20px;
  border: none;
  border-radius: 8px;
}

.app-card:hover {
  transform: translateY(-2px);
  transition: transform 0.2s;
}

.app-content {
  display: flex;
  align-items: center;
}

.app-icon {
  margin-right: 15px;
}

.app-info {
  flex: 1;
}

.app-name {
  font-weight: bold;
  margin-bottom: 5px;
  font-size: 16px;
}

.app-developer {
  color: #666;
  font-size: 14px;
  margin-bottom: 8px;
}

.app-rating {
  display: flex;
  align-items: center;
}

.rating-value {
  margin-left: 8px;
  color: #ff9900;
  font-weight: bold;
}

.app-action {
  margin-left: 10px;
}

.view-all {
  text-align: center;
  margin-top: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-container {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
    margin-right: 0;
    margin-bottom: 20px;
  }

  .header-nav .el-menu-item {
    font-size: 13px;
    padding: 0 10px;
  }

  .store-logo {
    font-size: 18px;
    margin-right: 15px;
  }
}
</style>