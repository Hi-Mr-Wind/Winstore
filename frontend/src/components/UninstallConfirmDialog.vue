<!--
  卸载确认弹窗组件 - Wails WebView 优化版本
  功能：
  - 完全自定义的弹窗实现，不依赖 Element Plus 弹窗
  - 针对 Wails WebView 环境优化
  - 确保在单 WebView 中正确显示
-->
<script setup lang="ts">
import { ref, watch, computed, onMounted, onUnmounted } from 'vue'
import { ElButton, ElIcon } from 'element-plus'
import { Warning } from '@element-plus/icons-vue'
import { useWailsEnvironment } from '@/composables/useWailsEnvironment'

interface Props {
  visible: boolean
  softwareName: string
  version: string
  publisher: string
}

interface Emits {
  (e: 'confirm'): void
  (e: 'cancel'): void
  (e: 'update:visible', value: boolean): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

// 使用 Wails 环境优化工具
const { isWebView, getEnvironmentInfo } = useWailsEnvironment()

// 计算属性确保数据正确传递
const dialogSoftwareName = computed(() => props.softwareName || '未知软件')
const dialogVersion = computed(() => props.version || '未知版本')
const dialogPublisher = computed(() => props.publisher || '未知供应商')

// 处理确认
const handleConfirm = () => {
  emit('confirm')
  emit('update:visible', false)
}

// 处理取消
const handleCancel = () => {
  emit('cancel')
  emit('update:visible', false)
}

// 处理背景点击
const handleBackdropClick = (event: Event) => {
  if (event.target === event.currentTarget) {
    handleCancel()
  }
}

// 处理 ESC 键
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    handleCancel()
  }
}

// 监听键盘事件
onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
  
  // 在 WebView 环境中应用额外的优化
  if (isWebView.value) {
    console.log('弹窗组件检测到 WebView 环境，应用特殊优化')
    console.log('环境信息:', getEnvironmentInfo())
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})

// 监听弹窗显示状态
watch(() => props.visible, (newVisible) => {
  if (newVisible) {
    // 弹窗显示时禁止背景滚动
    document.body.style.overflow = 'hidden'
    
    // 在 WebView 环境中应用额外的样式修复
    if (isWebView.value) {
      // 强制重新计算样式
      setTimeout(() => {
        const overlay = document.querySelector('.custom-dialog-overlay')
        if (overlay) {
          overlay.setAttribute('style', overlay.getAttribute('style') + '; z-index: 999999 !important;')
        }
      }, 10)
    }
  } else {
    // 弹窗隐藏时恢复背景滚动
    document.body.style.overflow = ''
  }
})
</script>

<template>
  <!-- 使用 Teleport 确保弹窗渲染到 body 根节点 -->
  <Teleport to="body">
    <Transition name="dialog-fade">
      <div 
        v-if="props.visible" 
        class="custom-dialog-overlay"
        :class="{ 'webview-optimized': isWebView }"
        @click="handleBackdropClick"
      >
        <div class="custom-dialog-container">
          <div class="custom-dialog">
            <!-- 弹窗头部 -->
            <div class="dialog-header">
              <h3 class="dialog-title">确认卸载</h3>
              <button 
                class="dialog-close-btn"
                @click="handleCancel"
                type="button"
              >
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="18" y1="6" x2="6" y2="18"></line>
                  <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
              </button>
            </div>

            <!-- 弹窗内容 -->
            <div class="dialog-content">
              <div class="warning-icon">
                <div class="warning-icon-bg">
                  <svg width="32" height="32" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M12 2L1 21h22L12 2zm0 3.17L19.83 19H4.17L12 5.17zM11 16h2v2h-2zm0-6h2v4h-2z"/>
                  </svg>
                </div>
              </div>
              
              <div class="message">
                <p class="main-message">确定要卸载 "{{ dialogSoftwareName }}" 吗？</p>
                <div class="details">
                  <div class="detail-item">
                    <span class="detail-label">版本:</span>
                    <span class="detail-value">{{ dialogVersion }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">供应商:</span>
                    <span class="detail-value">{{ dialogPublisher }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- 弹窗底部 -->
            <div class="dialog-footer">
              <ElButton 
                class="cancel-btn"
                @click="handleCancel"
              >
                取消
              </ElButton>
              <ElButton 
                type="danger" 
                class="confirm-btn"
                @click="handleConfirm"
              >
                确定卸载
              </ElButton>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
/* 弹窗遮罩层 */
.custom-dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.75);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999999;
  padding: 20px;
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
}

/* WebView 优化版本 */
.custom-dialog-overlay.webview-optimized {
  z-index: 999999 !important;
  position: fixed !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background-color: rgba(0, 0, 0, 0.75) !important;
  backdrop-filter: blur(8px) !important;
  -webkit-backdrop-filter: blur(8px) !important;
  -webkit-transform: translateZ(0) !important;
  transform: translateZ(0) !important;
}

/* 弹窗容器 */
.custom-dialog-container {
  width: 100%;
  max-width: 480px;
  max-height: 90vh;
  overflow: hidden;
}

/* 弹窗主体 */
.custom-dialog {
  background: var(--bg-primary);
  border-radius: 16px;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
  border: 1px solid var(--border-primary);
  overflow: hidden;
  animation: dialogSlideIn 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
}

/* WebView 优化版本 */
.webview-optimized .custom-dialog {
  z-index: 1000000 !important;
  position: relative !important;
  top: auto !important;
  left: auto !important;
  right: auto !important;
  bottom: auto !important;
  transform: none !important;
  margin: 0 !important;
  max-width: 90vw !important;
  max-height: 90vh !important;
  overflow: auto !important;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5) !important;
  border-radius: 16px !important;
  border: 1px solid var(--border-primary) !important;
  background: var(--bg-primary) !important;
  -webkit-transform: translateZ(0) !important;
  transform: translateZ(0) !important;
}

/* 弹窗头部 */
.dialog-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px 28px 0 28px;
  border-bottom: 1px solid var(--border-primary);
  position: relative;
}

.dialog-header::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 28px;
  right: 28px;
  height: 1px;
  background: linear-gradient(90deg, transparent, var(--border-primary), transparent);
}

.dialog-title {
  font-size: 20px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0;
  letter-spacing: -0.5px;
}

.dialog-close-btn {
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s ease;
}

.dialog-close-btn:hover {
  background: var(--bg-secondary);
  color: var(--text-primary);
}

/* 弹窗内容 */
.dialog-content {
  padding: 28px;
  text-align: center;
}

/* 警告图标 */
.warning-icon {
  margin-bottom: 28px;
  display: flex;
  justify-content: center;
}

.warning-icon-bg {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ffa726, #ff9800);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  box-shadow: 0 8px 24px rgba(255, 167, 38, 0.4);
  position: relative;
  animation: pulse 2s infinite;
}

.warning-icon-bg::before {
  content: '';
  position: absolute;
  top: -4px;
  left: -4px;
  right: -4px;
  bottom: -4px;
  background: linear-gradient(135deg, #ffa726, #ff9800);
  border-radius: 50%;
  z-index: -1;
  opacity: 0.3;
  animation: pulse 2s infinite 0.5s;
}

/* 消息内容 */
.message {
  font-size: 16px;
  line-height: 1.5;
}

.main-message {
  margin-bottom: 24px;
  font-size: 16px;
  line-height: 1.6;
  color: var(--text-primary);
  font-weight: 500;
}

/* 详情区域 */
.details {
  background: var(--bg-secondary);
  padding: 20px;
  border-radius: 12px;
  border: 1px solid var(--border-primary);
  position: relative;
  overflow: hidden;
  text-align: left;
}

.details::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #ff6b6b, #ffa726, #4ecdc4);
}

.detail-item {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  font-size: 14px;
}

.detail-item:last-child {
  margin-bottom: 0;
}

.detail-label {
  color: var(--accent-primary);
  font-weight: 600;
  min-width: 60px;
  margin-right: 8px;
}

.detail-value {
  color: var(--text-primary);
  flex: 1;
}

/* 弹窗底部 */
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 20px 28px 28px 28px;
  border-top: 1px solid var(--border-primary);
  position: relative;
}

.dialog-footer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 28px;
  right: 28px;
  height: 1px;
  background: linear-gradient(90deg, transparent, var(--border-primary), transparent);
}

/* 按钮样式 */
.cancel-btn {
  border-radius: 10px !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94) !important;
  font-size: 14px !important;
  letter-spacing: 0.3px !important;
  background: var(--bg-secondary) !important;
  border: 2px solid var(--border-primary) !important;
  color: var(--text-primary) !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.cancel-btn:hover {
  background: var(--bg-tertiary) !important;
  border-color: var(--accent-primary) !important;
  color: var(--accent-primary) !important;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15) !important;
  transform: translateY(-2px) !important;
}

.confirm-btn {
  border-radius: 10px !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94) !important;
  font-size: 14px !important;
  letter-spacing: 0.3px !important;
  background: linear-gradient(135deg, #ff6b6b, #ff5252) !important;
  border: none !important;
  color: white !important;
  box-shadow: 0 4px 16px rgba(255, 107, 107, 0.4) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2) !important;
}

.confirm-btn:hover {
  background: linear-gradient(135deg, #ff5252, #ff4444) !important;
  box-shadow: 0 6px 20px rgba(255, 107, 107, 0.5) !important;
  transform: translateY(-2px) !important;
}

/* 动画效果 */
@keyframes dialogSlideIn {
  0% {
    opacity: 0;
    transform: scale(0.8) translateY(-40px);
  }
  100% {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.8;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

/* 过渡动画 */
.dialog-fade-enter-active,
.dialog-fade-leave-active {
  transition: opacity 0.3s ease;
}

.dialog-fade-enter-from,
.dialog-fade-leave-to {
  opacity: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .custom-dialog-overlay {
    padding: 16px;
  }
  
  .custom-dialog {
    border-radius: 12px;
  }
  
  .dialog-header,
  .dialog-content,
  .dialog-footer {
    padding-left: 20px;
    padding-right: 20px;
  }
  
  .dialog-title {
    font-size: 18px;
  }
  
  .warning-icon-bg {
    width: 64px;
    height: 64px;
  }
  
  .main-message {
    font-size: 15px;
  }
}

/* 确保在深色主题下也有良好的对比度 */
.custom-dialog {
  background: var(--bg-primary);
  color: var(--text-primary);
}

/* 强制样式优先级 */
.custom-dialog-overlay {
  z-index: 999999 !important;
}

.custom-dialog {
  z-index: 1000000 !important;
}

/* WebView 特定优化 */
.webview-optimized {
  -webkit-transform: translateZ(0) !important;
  transform: translateZ(0) !important;
  -webkit-backface-visibility: hidden !important;
  backface-visibility: hidden !important;
}

.webview-optimized .custom-dialog {
  -webkit-transform: translateZ(0) !important;
  transform: translateZ(0) !important;
  -webkit-backface-visibility: hidden !important;
  backface-visibility: hidden !important;
}
</style>
