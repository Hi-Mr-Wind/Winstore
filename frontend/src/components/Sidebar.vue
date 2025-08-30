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

const currentRoute = computed(() => route.name as string)

const menuItems = [
  { index: 'home', icon: HomeFilled, label: t('common.home'), path: '/' },
  { index: 'popular', icon: MagicStick, label: t('common.popular'), path: '/popular' },
  { index: 'category', icon: Grid, label: t('common.category'), path: '/category' },
  { index: 'installed', icon: Download, label: t('common.installed'), path: '/installed' },
  { index: 'favorites', icon: Star, label: t('common.favorites'), path: '/favorites' },
  { index: 'tools', icon: Box, label: t('common.tools'), path: '/tools' }
]

const subMenuItems = [
  { index: '7-1', label: '生产力', path: '/category/1' },
  { index: '7-2', label: '社交', path: '/category/2' },
  { index: '7-3', label: '游戏', path: '/category/3' },
  { index: '7-4', label: '娱乐', path: '/category/4' },
  { index: '7-5', label: '教育', path: '/category/5' },
  { index: '7-6', label: '工具', path: '/category/6' }
]

const bottomMenuItems = [
  { index: 'settings', icon: Setting, label: t('common.settings'), path: '/settings' },
  { index: 'help', icon: QuestionFilled, label: t('common.help'), path: '/help' }
]

const handleMenuSelect = (index: string) => {
  const menuItem = menuItems.find(item => item.index === index)
  const subMenuItem = subMenuItems.find(item => item.index === index)
  const bottomMenuItem = bottomMenuItems.find(item => item.index === index)

  if (menuItem) {
    router.push(menuItem.path)
  } else if (subMenuItem) {
    router.push(subMenuItem.path)
  } else if (bottomMenuItem) {
    router.push(bottomMenuItem.path)
  }
}
</script>

<template>
  <div class="sidebar">
    <!-- 侧边栏头部 -->
    <div class="sidebar-header">
      <div class="store-logo">
        <el-icon class="lock-icon"><Box /></el-icon>
        <span>Win Store</span>
      </div>
    </div>

    <!-- 主菜单 -->
    <el-menu
      :default-active="currentRoute"
      class="sidebar-menu"
      @select="handleMenuSelect"
    >
      <el-menu-item
        v-for="item in menuItems"
        :key="item.index"
        :index="item.index"
      >
        <el-icon><component :is="item.icon" /></el-icon>
        <span>{{ item.label }}</span>
      </el-menu-item>

      <!-- 子菜单 -->
      <el-sub-menu index="7">
        <template #title>
          <span>{{ t('common.category') }}</span>
        </template>
        <el-menu-item
          v-for="item in subMenuItems"
          :key="item.index"
          :index="item.index"
        >
          {{ item.label }}
        </el-menu-item>
      </el-sub-menu>
    </el-menu>

    <!-- 底部菜单 -->
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
.sidebar {
  width: 240px;
  height: 100vh;
  background-color: var(--bg-primary);
  border-right: 1px solid var(--border-primary);
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-primary);
}

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

.sidebar-menu {
  flex: 1;
  border: none;
  background-color: transparent;
}

.sidebar-menu :deep(.el-menu-item) {
  color: var(--text-primary);
  border-radius: 0;
  margin: 4px 8px;
  border-radius: 6px;
}

.sidebar-menu :deep(.el-menu-item:hover) {
  background-color: var(--bg-tertiary);
  color: var(--accent-primary);
}

.sidebar-menu :deep(.el-menu-item.is-active) {
  background-color: var(--accent-primary);
  color: white;
}

.sidebar-menu :deep(.el-sub-menu__title) {
  color: var(--text-primary);
  border-radius: 0;
  margin: 4px 8px;
  border-radius: 6px;
}

.sidebar-menu :deep(.el-sub-menu__title:hover) {
  background-color: var(--bg-tertiary);
  color: var(--accent-primary);
}

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
