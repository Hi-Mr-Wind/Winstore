<!--
  设置页面组件
  功能：
  - 提供应用商店的各种设置选项
  - 支持主题切换（浅色/深色/自动）
  - 支持语言切换（中文/英文）
  - 自动更新设置
  - 通知设置
  - 下载路径设置
  - 响应式设计，适配不同屏幕尺寸
-->
<script setup lang="ts">
import { useAppStore } from '@/stores/appStore'
import { useTheme } from '@/composables/useTheme'
import { setLanguage, getCurrentLanguage } from '@/i18n'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'

const appStore = useAppStore()
const { theme, setTheme } = useTheme()
const { t, locale } = useI18n()

// 处理主题切换
const handleThemeChange = (newTheme: string) => {
  setTheme(newTheme as 'light' | 'dark' | 'auto')
  appStore.updateUserPreferences({ theme: newTheme as 'light' | 'dark' | 'auto' })
}

// 处理语言切换
const handleLanguageChange = (language: string) => {
  setLanguage(language)
  appStore.updateUserPreferences({ language })
}

// 处理自动更新设置
const handleAutoUpdateChange = (value: boolean) => {
  appStore.updateUserPreferences({ autoUpdate: value })
}

// 处理通知设置
const handleNotificationsChange = (value: boolean) => {
  appStore.updateUserPreferences({ notifications: value })
}

// 处理下载路径设置
const handleDownloadPathChange = (path: string) => {
  appStore.updateUserPreferences({ downloadPath: path })
}
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏导航 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 页面标题 -->
      <div class="page-header">
        <h1 class="page-title">{{ t('settings.title') }}</h1>
        <p class="page-description">{{ t('settings.subtitle') }}</p>
      </div>

      <!-- 设置内容区域 -->
      <div class="settings-content">
        <!-- 外观设置区域 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.appearance') }}</h2>
          
          <!-- 主题设置 -->
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.theme') }}</h3>
              <p>{{ t('settings.themeDesc') }}</p>
            </div>
            <el-select
              :model-value="theme"
              @change="handleThemeChange"
              style="width: 120px"
            >
              <el-option :label="t('settings.light')" value="light" />
              <el-option :label="t('settings.dark')" value="dark" />
              <el-option :label="t('settings.auto')" value="auto" />
            </el-select>
          </div>
        </div>

        <!-- 语言设置区域 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.language') }}</h2>
          
          <!-- 语言选择 -->
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.languageDesc') }}</h3>
              <p>{{ t('settings.languageDesc') }}</p>
            </div>
            <el-select
              :model-value="locale"
              @change="handleLanguageChange"
              style="width: 120px"
            >
              <el-option :label="t('settings.chinese')" value="zh-CN" />
              <el-option :label="t('settings.english')" value="en-US" />
            </el-select>
          </div>
        </div>

        <!-- 更新设置区域 -->
        <!-- <div class="settings-section">
          <h2 class="section-title">{{ t('settings.update') }}</h2>
          
          自动更新设置 -->
          <!-- <div class="setting-item"> -->
            <!-- <div class="setting-info"> -->
              <!-- <h3>{{ t('settings.autoUpdate') }}</h3> -->
              <!-- <p>{{ t('settings.autoUpdateDesc') }}</p> -->
            <!-- </div> -->
            <!-- <el-switch -->
              <!-- :model-value="appStore.userPreferences.autoUpdate" -->
              <!-- @change="handleAutoUpdateChange" -->
            <!-- /> -->
          <!-- </div> -->
        <!-- </div> -->

        <!-- 通知设置区域 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.notifications') }}</h2>
          
          <!-- 推送通知设置 -->
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.pushNotifications') }}</h3>
              <p>{{ t('settings.pushNotificationsDesc') }}</p>
            </div>
            <el-switch
              :model-value="appStore.userPreferences.notifications"
              @change="handleNotificationsChange"
            />
          </div>
        </div>

        <!-- 下载设置区域 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.download') }}</h2>
          
          <!-- 下载路径设置 -->
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.downloadPath') }}</h3>
              <p>{{ t('settings.downloadPathDesc') }}</p>
            </div>
            <div class="path-input">
              <el-input
                :model-value="appStore.userPreferences.downloadPath"
                @change="handleDownloadPathChange"
                style="width: 300px"
              />
              <el-button>{{ t('settings.browse') }}</el-button>
            </div>
          </div>
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

/* 页面标题区域 */
.page-header {
  margin-bottom: 32px;
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

/* 设置内容区域 */
.settings-content {
  max-width: 800px;
}

/* 设置区域样式 */
.settings-section {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-primary);
}

/* 设置项样式 */
.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px 0;
  border-bottom: 1px solid var(--border-secondary);
}

.setting-item:last-child {
  border-bottom: none;
}

.setting-info {
  flex: 1;
  margin-right: 24px;
}

.setting-info h3 {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.setting-info p {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
}

/* 路径输入框样式 */
.path-input {
  display: flex;
  gap: 8px;
  align-items: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .setting-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .setting-info {
    margin-right: 0;
  }
  
  .path-input {
    width: 100%;
  }
  
  .path-input .el-input {
    width: 100% !important;
  }
}
</style>
