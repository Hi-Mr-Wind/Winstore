<!--
  事件通知窗口组件
  功能：
  - 显示后端事件通知
  - 支持不同类型的通知（成功、错误、信息、警告）
  - 可拖拽移动
  - 可最小化/关闭
  - 支持自动关闭
-->
<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElIcon } from 'element-plus'
import { Close, Minus, SuccessFilled, WarningFilled, InfoFilled, CircleCloseFilled } from '@element-plus/icons-vue'

interface EventNotification {
  id: string
  type: 'success' | 'error' | 'info' | 'warning'
  title: string
  message: string
  timestamp: Date
  softwareName?: string
  progress?: number
}

const { t } = useI18n()

// 通知列表
const notifications = ref<EventNotification[]>([])
const isMinimized = ref(false)
const isVisible = ref(false)

// 拖拽相关
const isDragging = ref(false)
const dragOffset = ref({ x: 0, y: 0 })
const windowPosition = ref({ x: 20, y: 20 })

// 计算属性
const hasNotifications = computed(() => notifications.value.length > 0)
const unreadCount = computed(() => notifications.value.length)

// 通知类型图标映射
const typeIcons = {
  success: SuccessFilled,
  error: CircleCloseFilled,
  warning: WarningFilled,
  info: InfoFilled
}

// 通知类型颜色映射
const typeColors = {
  success: '#67c23a',
  error: '#f56c6c',
  warning: '#e6a23c',
  info: '#409eff'
}

// 添加通知
const addNotification = (notification: Omit<EventNotification, 'id' | 'timestamp'>) => {
  const newNotification: EventNotification = {
    ...notification,
    id: Date.now().toString(),
    timestamp: new Date()
  }
  
  notifications.value.unshift(newNotification)
  isVisible.value = true
  
  // 自动关闭成功和错误通知
  if (notification.type === 'success' || notification.type === 'error') {
    setTimeout(() => {
      removeNotification(newNotification.id)
    }, 5000)
  }
}

// 移除通知
const removeNotification = (id: string) => {
  const index = notifications.value.findIndex(n => n.id === id)
  if (index > -1) {
    notifications.value.splice(index, 1)
  }
  
  if (notifications.value.length === 0) {
    isVisible.value = false
  }
}

// 清空所有通知
const clearAllNotifications = () => {
  notifications.value = []
  isVisible.value = false
}

// 切换最小化状态
const toggleMinimize = () => {
  isMinimized.value = !isMinimized.value
}

// 关闭窗口
const closeWindow = () => {
  isVisible.value = false
}

// 拖拽开始
const startDrag = (event: MouseEvent) => {
  isDragging.value = true
  const rect = (event.target as HTMLElement).getBoundingClientRect()
  dragOffset.value = {
    x: event.clientX - rect.left,
    y: event.clientY - rect.top
  }
  document.addEventListener('mousemove', onDrag)
  document.addEventListener('mouseup', stopDrag)
}

// 拖拽中
const onDrag = (event: MouseEvent) => {
  if (isDragging.value) {
    windowPosition.value = {
      x: event.clientX - dragOffset.value.x,
      y: event.clientY - dragOffset.value.y
    }
  }
}

// 拖拽结束
const stopDrag = () => {
  isDragging.value = false
  document.removeEventListener('mousemove', onDrag)
  document.removeEventListener('mouseup', stopDrag)
}

// 格式化时间
const formatTime = (date: Date) => {
  return date.toLocaleTimeString('zh-CN', { 
    hour: '2-digit', 
    minute: '2-digit',
    second: '2-digit'
  })
}

// 暴露方法给父组件
defineExpose({
  addNotification,
  removeNotification,
  clearAllNotifications
})

// 组件卸载时清理事件监听器
onUnmounted(() => {
  document.removeEventListener('mousemove', onDrag)
  document.removeEventListener('mouseup', stopDrag)
})
</script>

<template>
  <div 
    v-if="isVisible"
    class="event-notification-window"
    :class="{ minimized: isMinimized }"
    :style="{
      left: windowPosition.x + 'px',
      top: windowPosition.y + 'px'
    }"
  >
    <!-- 窗口标题栏 -->
    <div class="window-header" @mousedown="startDrag">
      <div class="window-title">
        <el-icon class="notification-icon"><InfoFilled /></el-icon>
        <span>事件通知</span>
        <span v-if="unreadCount > 0" class="notification-count">({{ unreadCount }})</span>
      </div>
      <div class="window-controls">
        <el-button 
          type="text" 
          size="small" 
          @click="toggleMinimize"
          class="control-btn"
        >
          <el-icon><Minus /></el-icon>
        </el-button>
        <el-button 
          type="text" 
          size="small" 
          @click="closeWindow"
          class="control-btn"
        >
          <el-icon><Close /></el-icon>
        </el-button>
      </div>
    </div>

    <!-- 窗口内容 -->
    <div v-if="!isMinimized" class="window-content">
      <!-- 工具栏 -->
      <div class="toolbar">
        <el-button 
          type="text" 
          size="small" 
          @click="clearAllNotifications"
          :disabled="!hasNotifications"
        >
          清空全部
        </el-button>
      </div>

      <!-- 通知列表 -->
      <div class="notifications-list">
        <div 
          v-for="notification in notifications" 
          :key="notification.id"
          class="notification-item"
          :class="notification.type"
        >
          <div class="notification-header">
            <el-icon 
              :color="typeColors[notification.type]"
              class="type-icon"
            >
              <component :is="typeIcons[notification.type]" />
            </el-icon>
            <span class="notification-title">{{ notification.title }}</span>
            <span class="notification-time">{{ formatTime(notification.timestamp) }}</span>
            <el-button 
              type="text" 
              size="small" 
              @click="removeNotification(notification.id)"
              class="close-btn"
            >
              <el-icon><Close /></el-icon>
            </el-button>
          </div>
          
          <div class="notification-message">
            {{ notification.message }}
          </div>
          
          <!-- 进度条（如果有） -->
          <div v-if="notification.progress !== undefined" class="notification-progress">
            <el-progress 
              :percentage="notification.progress" 
              :color="typeColors[notification.type]"
              :stroke-width="4"
            />
          </div>
          
          <!-- 软件名称（如果有） -->
          <div v-if="notification.softwareName" class="notification-software">
            软件: {{ notification.softwareName }}
          </div>
        </div>
        
        <!-- 空状态 -->
        <div v-if="!hasNotifications" class="empty-state">
          <el-icon class="empty-icon"><InfoFilled /></el-icon>
          <p>暂无事件通知</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.event-notification-window {
  position: fixed;
  width: 400px;
  max-height: 500px;
  background-color: var(--bg-primary);
  border: 1px solid var(--border-primary);
  border-radius: 8px;
  box-shadow: var(--shadow-heavy);
  z-index: 10001;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  user-select: none;
}

.event-notification-window.minimized {
  height: 40px;
}

.window-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  background-color: var(--bg-secondary);
  border-bottom: 1px solid var(--border-primary);
  cursor: move;
  user-select: none;
}

.window-title {
  display: flex;
  align-items: center;
  gap: 6px;
  font-weight: 500;
  color: var(--text-primary);
}

.notification-icon {
  font-size: 16px;
  color: var(--accent-primary);
}

.notification-count {
  color: var(--accent-primary);
  font-weight: 600;
}

.window-controls {
  display: flex;
  gap: 4px;
}

.control-btn {
  padding: 2px;
  color: var(--text-secondary);
}

.control-btn:hover {
  color: var(--text-primary);
  background-color: var(--bg-hover);
}

.window-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.toolbar {
  padding: 8px 12px;
  border-bottom: 1px solid var(--border-primary);
  background-color: var(--bg-secondary);
}

.notifications-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.notification-item {
  margin-bottom: 12px;
  padding: 12px;
  border-radius: 6px;
  border: 1px solid var(--border-primary);
  background-color: var(--bg-secondary);
}

.notification-item.success {
  border-left: 4px solid #67c23a;
}

.notification-item.error {
  border-left: 4px solid #f56c6c;
}

.notification-item.warning {
  border-left: 4px solid #e6a23c;
}

.notification-item.info {
  border-left: 4px solid #409eff;
}

.notification-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.type-icon {
  font-size: 16px;
}

.notification-title {
  flex: 1;
  font-weight: 500;
  color: var(--text-primary);
}

.notification-time {
  font-size: 12px;
  color: var(--text-secondary);
}

.close-btn {
  padding: 2px;
  color: var(--text-secondary);
}

.close-btn:hover {
  color: var(--text-primary);
}

.notification-message {
  color: var(--text-primary);
  line-height: 1.5;
  margin-bottom: 8px;
}

.notification-progress {
  margin-bottom: 8px;
}

.notification-software {
  font-size: 12px;
  color: var(--text-secondary);
  font-style: italic;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  color: var(--text-secondary);
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 16px;
  opacity: 0.5;
}

/* 滚动条样式 */
.notifications-list::-webkit-scrollbar {
  width: 6px;
}

.notifications-list::-webkit-scrollbar-track {
  background: var(--bg-secondary);
}

.notifications-list::-webkit-scrollbar-thumb {
  background: var(--border-primary);
  border-radius: 3px;
}

.notifications-list::-webkit-scrollbar-thumb:hover {
  background: var(--text-secondary);
}
</style>

