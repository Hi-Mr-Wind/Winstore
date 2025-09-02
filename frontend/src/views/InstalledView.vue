<!--
  已安装软件页面组件
  功能：
  - 从Wails后端获取已安装软件列表
  - 显示软件名称、版本、供应商信息
  - 提供卸载功能，带确认框
  - 监听卸载进度事件
  - 支持响应式设计
-->
<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { GetSoftwareList, UninstallSoftware } from '../../wailsjs/go/apps/App.js'
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime.js'
import Sidebar from '@/components/Sidebar.vue'
import EventNotificationWindow from '@/components/EventNotificationWindow.vue'
import UninstallConfirmDialog from '@/components/UninstallConfirmDialog.vue'
import type { sysUtils } from '../../wailsjs/go/models'
import { Box, Search, Refresh } from '@element-plus/icons-vue'

// 扩展Window接口以包含Wails的go对象
declare global {
  interface Window {
    go?: {
      apps?: {
        App?: any
      }
    }
  }
}

const router = useRouter()
const { t } = useI18n()

// 软件列表数据
const softwareList = ref<sysUtils.SoftwareList[]>([])
const loading = ref(false)
const uninstallingSoftware = ref<string | null>(null)

// 搜索相关状态
const searchQuery = ref('')

// 事件通知窗口引用
const eventNotificationRef = ref<InstanceType<typeof EventNotificationWindow> | null>(null)

// 卸载确认弹窗状态
const uninstallDialogVisible = ref(false)
const selectedSoftware = ref<sysUtils.SoftwareList | null>(null)

// 计算过滤后的软件列表
const filteredSoftwareList = computed(() => {
  if (!searchQuery.value.trim()) {
    return softwareList.value
  }
  
  const query = searchQuery.value.toLowerCase().trim()
  return softwareList.value.filter(software => 
    software.Name.toLowerCase().includes(query) ||
    software.Version.toLowerCase().includes(query) ||
    software.Publisher.toLowerCase().includes(query)
  )
})

// 清空搜索
const clearSearch = () => {
  searchQuery.value = ''
}

// 获取已安装软件列表
const fetchSoftwareList = async () => {
  try {
    loading.value = true
    
    // 检查是否在Wails环境中
    if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
      const software = await GetSoftwareList()
      softwareList.value = software
    } else {
      // 不在Wails环境中，显示提示信息
      console.warn('当前不在Wails环境中，无法获取软件列表')
      ElMessage.warning('当前环境不支持此功能，请在Wails应用中运行')
    }
  } catch (error) {
    console.error('获取软件列表失败:', error)
    ElMessage.error('获取软件列表失败，请重试')
  } finally {
    loading.value = false
  }
}

// 刷新软件列表
const refreshSoftwareList = async () => {
  try {
    // 清空搜索
    searchQuery.value = ''
    
    // 重新获取软件列表
    await fetchSoftwareList()
    
    ElMessage.success('软件列表已刷新')
  } catch (error) {
    console.error('刷新软件列表失败:', error)
    ElMessage.error('刷新软件列表失败，请重试')
  }
}

// 处理卸载软件
const handleUninstall = (software: sysUtils.SoftwareList) => {
  selectedSoftware.value = software
  uninstallDialogVisible.value = true
}

// 确认卸载
const confirmUninstall = async () => {
  if (!selectedSoftware.value) return
  
  try {
    const software = selectedSoftware.value
    
    // 检查是否在Wails环境中
    if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
      // 开始卸载
      uninstallingSoftware.value = software.Name
      
      // 调用卸载方法
      await UninstallSoftware(software.Name)
      
      ElMessage.success(`正在卸载 ${software.Name}...`)
      
      // 添加到事件通知窗口
      eventNotificationRef.value?.addNotification({
        type: 'info',
        title: '开始卸载',
        message: `正在卸载 ${software.Name}...`,
        softwareName: software.Name
      })
    } else {
      // 不在Wails环境中，显示模拟卸载
      ElMessage.warning('当前环境不支持卸载功能，请在Wails应用中运行')
      uninstallingSoftware.value = null
    }
    
    // 关闭弹窗
    uninstallDialogVisible.value = false
    selectedSoftware.value = null
    
  } catch (error) {
    console.error('卸载失败:', error)
    ElMessage.error(`卸载 ${selectedSoftware.value?.Name} 失败，请重试`)
    uninstallingSoftware.value = null
    
    // 关闭弹窗
    uninstallDialogVisible.value = false
    selectedSoftware.value = null
  }
}

// 取消卸载
const cancelUninstall = () => {
  uninstallDialogVisible.value = false
  selectedSoftware.value = null
}

// 监听卸载进度事件
const setupUninstallListeners = () => {
  // 检查是否在Wails环境中
  if (typeof window === 'undefined' || !window.go || !window.go.apps || !window.go.apps.App) {
    return
  }
  
  // 为每个软件设置事件监听器
  softwareList.value.forEach(software => {
    const eventName = `UninstallSoftware_${software.Name}`
    
    EventsOn(eventName, (data: any) => {
      console.log(`卸载进度 - ${software.Name}:`, data)
      
      // 根据返回的数据类型处理不同的情况
      if (typeof data === 'string') {
        // 如果是字符串，可能是状态信息
        if (data.includes('完成') || data.includes('success') || data.includes('完成')) {
          ElMessage.success(`${software.Name} 卸载完成`)
          uninstallingSoftware.value = null
          // 重新获取软件列表
          fetchSoftwareList()
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'success',
            title: '卸载完成',
            message: `${software.Name} 已成功卸载`,
            softwareName: software.Name
          })
        } else if (data.includes('失败') || data.includes('error') || data.includes('失败')) {
          ElMessage.error(`${software.Name} 卸载失败`)
          uninstallingSoftware.value = null
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'error',
            title: '卸载失败',
            message: `${software.Name} 卸载失败`,
            softwareName: software.Name
          })
        } else {
          // 其他进度信息
          ElMessage.info(`${software.Name}: ${data}`)
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'info',
            title: '卸载进度',
            message: data,
            softwareName: software.Name
          })
        }
      } else if (typeof data === 'object') {
        // 如果是对象，可能包含更详细的信息
        if (data.status === 'completed' || data.status === 'success') {
          ElMessage.success(`${software.Name} 卸载完成`)
          uninstallingSoftware.value = null
          fetchSoftwareList()
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'success',
            title: '卸载完成',
            message: `${software.Name} 已成功卸载`,
            softwareName: software.Name
          })
        } else if (data.status === 'failed' || data.status === 'error') {
          ElMessage.error(`${software.Name} 卸载失败: ${data.message || '未知错误'}`)
          uninstallingSoftware.value = null
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'error',
            title: '卸载失败',
            message: data.message || '未知错误',
            softwareName: software.Name
          })
        } else if (data.progress) {
          // 显示进度信息
          ElMessage.info(`${software.Name} 卸载进度: ${data.progress}%`)
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'info',
            title: '卸载进度',
            message: `正在卸载 ${software.Name}`,
            softwareName: software.Name,
            progress: data.progress
          })
        } else {
          // 其他信息
          ElMessage.info(`${software.Name}: ${data.message || JSON.stringify(data)}`)
          
          // 添加到事件通知窗口
          eventNotificationRef.value?.addNotification({
            type: 'info',
            title: '卸载信息',
            message: data.message || JSON.stringify(data),
            softwareName: software.Name
          })
        }
      }
    })
  })
}

// 清理事件监听器
const cleanupUninstallListeners = () => {
  // 检查是否在Wails环境中
  if (typeof window === 'undefined' || !window.go || !window.go.apps || !window.go.apps.App) {
    return
  }
  
  softwareList.value.forEach(software => {
    const eventName = `UninstallSoftware_${software.Name}`
    EventsOff(eventName)
  })
}

// 组件挂载时获取数据
onMounted(() => {
  fetchSoftwareList()
})

// 监听软件列表变化，重新设置事件监听器
const setupListeners = () => {
  cleanupUninstallListeners()
  setupUninstallListeners()
}

// 监听软件列表变化
watch(softwareList, () => {
  setupListeners()
}, { deep: true })

// 组件卸载时清理事件监听器
onUnmounted(() => {
  cleanupUninstallListeners()
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏导航 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 事件通知窗口 -->
      <EventNotificationWindow ref="eventNotificationRef" />
      
      <!-- 卸载确认弹窗 -->
      <UninstallConfirmDialog
        v-model:visible="uninstallDialogVisible"
        :software-name="selectedSoftware?.Name || ''"
        :version="selectedSoftware?.Version || ''"
        :publisher="selectedSoftware?.Publisher || ''"
        @confirm="confirmUninstall"
        @cancel="cancelUninstall"
      />
      
      <!-- 页面标题区域 -->
      <div class="page-header">
        <div class="header-content">
          <div class="title-section">
            <h1 class="page-title">{{ t('installed.title') }}</h1>
            <p class="page-description">
              {{ t('installed.subtitle') }} ({{ filteredSoftwareList.length }}/{{ softwareList.length }})
            </p>
          </div>
          <div class="header-actions">
            <el-button
              type="primary"
              :icon="Refresh"
              :loading="loading"
              @click="refreshSoftwareList"
              class="refresh-btn"
            >
              {{ loading ? t('common.refreshing') : t('common.refresh') }}
            </el-button>
          </div>
        </div>
      </div>

      <!-- 搜索区域 -->
      <div class="search-section">
        <div class="search-container">
          <el-input
            v-model="searchQuery"
            :placeholder="t('search.placeholder')"
            class="search-input"
            clearable
            @clear="clearSearch"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
          <div class="search-info">
            <span class="search-count">
              {{ filteredSoftwareList.length }} / {{ softwareList.length }}
            </span>
            <el-button 
              v-if="searchQuery" 
              type="text" 
              size="small" 
              @click="clearSearch"
            >
              清空搜索
            </el-button>
          </div>
        </div>
      </div>

      <!-- 软件列表区域 -->
      <div class="software-section">
        <!-- 加载状态 -->
        <div v-if="loading" class="loading-state">
          <div class="loading-container">
            <el-icon class="loading-icon" :size="48">
              <Refresh />
            </el-icon>
            <h3 class="loading-title">{{ t('common.loading') }}</h3>
            <p class="loading-description">{{ t('installed.loadingDescription') }}</p>
            <el-skeleton :rows="8" animated class="loading-skeleton" />
          </div>
        </div>

        <!-- 软件列表 -->
        <div v-else-if="filteredSoftwareList.length > 0" class="software-table">
          <el-table 
            :data="filteredSoftwareList" 
            style="width: 100%"
            height="calc(100vh - 300px)"
            :header-cell-style="{ background: 'var(--bg-secondary)', color: 'var(--text-primary)' }"
          >
            <!-- 软件名称列 -->
            <el-table-column prop="Name" :label="t('app.name')" min-width="200">
              <template #default="{ row }">
                <div class="software-name">
                  <el-icon class="software-icon"><Box /></el-icon>
                  <span>{{ row.Name }}</span>
                </div>
              </template>
            </el-table-column>

            <!-- 版本列 -->
            <el-table-column prop="Version" :label="t('app.version')" width="120">
              <template #default="{ row }">
                <el-tag size="small" type="info">{{ row.Version }}</el-tag>
              </template>
            </el-table-column>

            <!-- 供应商列 -->
            <el-table-column prop="Publisher" :label="t('app.publisher')" min-width="200">
              <template #default="{ row }">
                <span class="publisher-name">{{ row.Publisher }}</span>
              </template>
            </el-table-column>

            <!-- 操作列 -->
            <el-table-column :label="t('common.actions')" width="120" fixed="right">
              <template #default="{ row }">
                <el-button
                  type="danger"
                  size="small"
                  :loading="uninstallingSoftware === row.Name"
                  :disabled="uninstallingSoftware !== null"
                  @click="handleUninstall(row)"
                >
                  {{ uninstallingSoftware === row.Name ? t('common.uninstalling') : t('common.uninstall') }}
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <!-- 空状态 -->
        <div v-else-if="softwareList.length === 0" class="empty-state">
          <el-empty :description="t('installed.noData')">
            <el-button type="primary" @click="router.push('/')">
              {{ t('installed.goDiscover') }}
            </el-button>
          </el-empty>
        </div>

        <!-- 搜索无结果状态 -->
        <div v-else-if="searchQuery && filteredSoftwareList.length === 0" class="empty-state">
          <el-empty :description="`没有找到包含 '${searchQuery}' 的软件`">
            <el-button type="primary" @click="clearSearch">
              清空搜索
            </el-button>
          </el-empty>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 应用容器 */
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

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 24px;
}

.title-section {
  flex: 1;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0 0 8px 0;
}

.page-description {
  font-size: 16px;
  color: var(--text-secondary);
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.refresh-btn {
  min-width: 100px;
  height: 40px;
  border-radius: 8px;
  font-weight: 500;
}

/* 搜索区域样式 */
.search-section {
  margin-bottom: 24px;
}

.search-container {
  display: flex;
  align-items: center;
  gap: 16px;
  background-color: var(--bg-primary);
  padding: 16px;
  border-radius: 12px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
}

.search-input {
  flex: 1;
  max-width: 400px;
}

.search-info {
  display: flex;
  align-items: center;
  gap: 12px;
  color: var(--text-secondary);
  font-size: 14px;
}

.search-count {
  font-weight: 500;
  color: var(--text-primary);
}

/* 软件区域样式 */
.software-section {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
  overflow: hidden;
}

/* 加载状态样式 */
.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
  min-height: 400px;
}

.loading-container {
  text-align: center;
  max-width: 400px;
}

.loading-icon {
  color: var(--accent-primary);
  margin-bottom: 16px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.loading-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.loading-description {
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 24px;
  line-height: 1.5;
}

.loading-skeleton {
  margin-top: 20px;
}

/* 表格样式优化 */
:deep(.el-table) {
  background-color: var(--bg-primary);
  color: var(--text-primary);
}

:deep(.el-table th) {
  background-color: var(--bg-secondary) !important;
  color: var(--text-primary) !important;
  font-weight: 600;
  border-bottom: 1px solid var(--border-primary);
}

:deep(.el-table td) {
  border-bottom: 1px solid var(--border-secondary);
}

:deep(.el-table--border) {
  border: 1px solid var(--border-primary);
}

:deep(.el-table__body-wrapper) {
  background-color: var(--bg-primary);
}

:deep(.el-table__header-wrapper) {
  background-color: var(--bg-secondary);
}

/* 软件名称样式 */
.software-name {
  display: flex;
  align-items: center;
  gap: 8px;
}

.software-icon {
  font-size: 16px;
  color: var(--accent-primary);
}

/* 统一软件图标容器样式（若后续在表格中加入图标列可复用） */
.uniform-icon-box {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: var(--icon-bg);
  border: 1px solid var(--border-primary);
}

.uniform-icon-box img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  padding: 6px;
}

/* 供应商名称样式 */
.publisher-name {
  color: var(--text-secondary);
  font-size: 14px;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
  min-height: 400px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .header-content {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }
  
  .header-actions {
    justify-content: flex-start;
  }
  
  .refresh-btn {
    min-width: 120px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .search-container {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }
  
  .search-input {
    max-width: none;
  }
  
  .search-info {
    justify-content: space-between;
  }
  
  .software-table {
    margin: 0 -16px;
  }
}
</style>
