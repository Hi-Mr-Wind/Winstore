<script setup lang="ts">
import { useAppStore } from '@/stores/appStore'
import { useTheme } from '@/composables/useTheme'
import { setLanguage, getCurrentLanguage } from '@/i18n'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'

const appStore = useAppStore()
const { theme, setTheme } = useTheme()
const { t, locale } = useI18n()

const handleThemeChange = (newTheme: string) => {
  setTheme(newTheme as 'light' | 'dark' | 'auto')
  appStore.updateUserPreferences({ theme: newTheme as 'light' | 'dark' | 'auto' })
}

const handleLanguageChange = (language: string) => {
  setLanguage(language)
  appStore.updateUserPreferences({ language })
}

const handleAutoUpdateChange = (value: boolean) => {
  appStore.updateUserPreferences({ autoUpdate: value })
}

const handleNotificationsChange = (value: boolean) => {
  appStore.updateUserPreferences({ notifications: value })
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
        <h1 class="page-title">{{ t('settings.title') }}</h1>
        <p class="page-description">{{ t('settings.subtitle') }}</p>
      </div>

      <!-- 设置内容 -->
      <div class="settings-content">
        <!-- 外观设置 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.appearance') }}</h2>
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

        <!-- 语言设置 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.language') }}</h2>
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

        <!-- 更新设置 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.update') }}</h2>
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.autoUpdate') }}</h3>
              <p>{{ t('settings.autoUpdateDesc') }}</p>
            </div>
            <el-switch 
              :model-value="appStore.userPreferences.autoUpdate"
              @change="handleAutoUpdateChange"
            />
          </div>
        </div>

        <!-- 通知设置 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.notifications') }}</h2>
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

        <!-- 下载设置 -->
        <div class="settings-section">
          <h2 class="section-title">{{ t('settings.download') }}</h2>
          <div class="setting-item">
            <div class="setting-info">
              <h3>{{ t('settings.downloadPath') }}</h3>
              <p>{{ t('settings.downloadPathDesc') }}</p>
            </div>
            <el-input 
              :model-value="appStore.userPreferences.downloadPath"
              readonly
              style="width: 300px"
            >
              <template #append>
                <el-button>{{ t('settings.browse') }}</el-button>
              </template>
            </el-input>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background-color: var(--bg-secondary);
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
  color: var(--text-primary);
  margin-bottom: 8px;
}

.page-description {
  font-size: 16px;
  color: var(--text-secondary);
}

.settings-content {
  max-width: 800px;
}

.settings-section {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 20px;
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

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 0;
  border-bottom: 1px solid var(--border-secondary);
}

.setting-item:last-child {
  border-bottom: none;
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
  margin: 0;
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
    gap: 12px;
  }
  
  .setting-item .el-input {
    width: 100% !important;
  }
}
</style>
