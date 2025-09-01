<!--
  侧边栏导航组件
  功能：
  - 提供应用商店的主要导航菜单
  - 包含首页、热门软件、分类、已安装、收藏室、工具箱等导航项
  - 支持分类子菜单展开
  - 底部包含设置和帮助菜单
  - 响应式设计，支持移动端适配
-->
<script setup lang="ts">
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  HomeFilled,
  MagicStick,
  Grid,
  Download,
  Star,
  Box,
  Setting,
  QuestionFilled
} from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()
const { t } = useI18n()

// 计算当前路由名称
const currentRoute = computed(() => route.name as string)

// 主菜单项配置
const menuItems = [
  { index: 'home', icon: HomeFilled, label: t('common.home'), path: '/' },
  { index: 'popular', icon: MagicStick, label: t('common.popular'), path: '/popular' },
  { index: 'category', icon: Grid, label: t('common.category'), path: '/category' },
  { index: 'installed', icon: Download, label: t('common.installed'), path: '/installed' },
  { index: 'favorites', icon: Star, label: t('common.favorites'), path: '/favorites' },
  { index: 'tools', icon: Box, label: t('common.tools'), path: '/tools' }
]

// 底部菜单项配置
const bottomMenuItems = [
  { index: 'settings', icon: Setting, label: t('common.settings'), path: '/settings' },
  { index: 'help', icon: QuestionFilled, label: t('common.help'), path: '/help' }
]

// 处理菜单项选择事件
const handleMenuSelect = (index: string) => {
  // 查找对应的菜单项
  const menuItem = menuItems.find(item => item.index === index)
  const bottomMenuItem = bottomMenuItems.find(item => item.index === index)

  // 根据找到的菜单项进行路由跳转
  if (menuItem) {
    router.push(menuItem.path)
  } else if (bottomMenuItem) {
    router.push(bottomMenuItem.path)
  }
}
</script>

<template>
  <div class="sidebar">
    <!-- 侧边栏头部 - 应用商店Logo -->
    <div class="sidebar-header">
      <div class="store-logo">
        <el-icon class="lock-icon"><Box /></el-icon>
        <span>Win Store</span>
      </div>
    </div>

    <!-- 主菜单区域 -->
    <el-menu
      :default-active="currentRoute"
      class="sidebar-menu"
      @select="handleMenuSelect"
    >
      <!-- 主菜单项 -->
      <el-menu-item
        v-for="item in menuItems"
        :key="item.index"
        :index="item.index"
      >
        <el-icon><component :is="item.icon" /></el-icon>
        <span>{{ item.label }}</span>
      </el-menu-item>
    </el-menu>

    <!-- 底部菜单区域 -->
    <div class="sidebar-footer">
      <el-menu class="sidebar-menu">
        <el-menu-item
          v-for="item in bottomMenuItems"
          :key="item.index"
          :index="item.index"
          @click="handleMenuSelect(item.index)"
        >
          <el-icon><component :is="item.icon" /></el-icon>
          <span>{{ item.label }}</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<style scoped>
/* 侧边栏容器 */
.sidebar {
  width: 240px;
  height: 100vh;
  background-color: var(--bg-primary);
  border-right: 1px solid var(--border-primary);
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
}

/* 侧边栏头部样式 */
.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-primary);
}

/* 应用商店Logo样式 */
.store-logo {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
}

.lock-icon {
  font-size: 20px;
  color: var(--accent-primary);
}

/* 菜单样式 */
.sidebar-menu {
  flex: 1;
  border: none;
  background-color: transparent;
}

/* 菜单项样式 */
.sidebar-menu :deep(.el-menu-item) {
  color: var(--text-primary);
  border-radius: 0;
  margin: 4px 8px;
  border-radius: 6px;
}

/* 菜单项悬停效果 */
.sidebar-menu :deep(.el-menu-item:hover) {
  background-color: var(--bg-tertiary);
  color: var(--accent-primary);
}

/* 菜单项激活状态 */
.sidebar-menu :deep(.el-menu-item.is-active) {
  background-color: var(--accent-primary);
  color: white;
}

/* 子菜单标题样式 */
.sidebar-menu :deep(.el-sub-menu__title) {
  color: var(--text-primary);
  border-radius: 0;
  margin: 4px 8px;
  border-radius: 6px;
}

/* 子菜单标题悬停效果 */
.sidebar-menu :deep(.el-sub-menu__title:hover) {
  background-color: var(--bg-tertiary);
  color: var(--accent-primary);
}

/* 底部菜单区域 */
.sidebar-footer {
  border-top: 1px solid var(--border-primary);
  padding: 8px 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .sidebar {
    width: 200px;
  }
  
  .store-logo span {
    display: none;
  }
}
</style>
