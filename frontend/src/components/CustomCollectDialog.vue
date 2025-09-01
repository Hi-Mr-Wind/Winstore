<template>
  <Teleport to="body">
    <Transition name="dialog-fade">
      <div v-if="visible" class="custom-dialog-overlay" @click="onBackdropClick">
        <div class="custom-dialog-container">
          <div class="custom-dialog">
            <div class="dialog-header">
              <h3 class="dialog-title">{{ title || '添加收藏' }}</h3>
              <button class="dialog-close-btn" @click="onCancel" type="button">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="18" y1="6" x2="6" y2="18"></line>
                  <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
              </button>
            </div>
            <div class="dialog-content">
              <form @submit.prevent="onConfirm">
                <div class="form-group">
                  <label>应用名称</label>
                  <input v-model="form.appName" type="text" required placeholder="请输入应用名称" />
                </div>
                <div class="form-group">
                  <label>下载地址</label>
                  <input v-model="form.downloadUrl" type="url" required placeholder="请输入下载地址" />
                </div>
                <div class="form-group">
                  <label>官网地址</label>
                  <input v-model="form.appUrl" type="url" required placeholder="请输入官网地址" />
                </div>
                <div class="form-group">
                  <label>图标路径</label>
                  <input v-model="form.appIcon" type="url" required placeholder="请输入图标URL" />
                </div>
                <div class="dialog-footer">
                  <button type="button" class="cancel-btn" @click="onCancel">取消</button>
                  <button type="submit" class="confirm-btn">{{ confirmText || '添加' }}</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { reactive, watch, defineProps, defineEmits } from 'vue'

const props = defineProps<{ visible: boolean; initial?: { appName: string; downloadUrl: string; appUrl: string; appIcon: string }; title?: string; confirmText?: string }>()
const emit = defineEmits(['confirm', 'cancel', 'update:visible'])

const form = reactive({
  appName: '',
  downloadUrl: '',
  appUrl: '',
  appIcon: ''
})

watch(() => props.visible, (val) => {
  if (val) {
    if (props.initial) {
      form.appName = props.initial.appName || ''
      form.downloadUrl = props.initial.downloadUrl || ''
      form.appUrl = props.initial.appUrl || ''
      form.appIcon = props.initial.appIcon || ''
    } else {
      form.appName = ''
      form.downloadUrl = ''
      form.appUrl = ''
      form.appIcon = ''
    }
  }
})

const onConfirm = () => {
  if (!form.appName || !form.downloadUrl || !form.appUrl || !form.appIcon) return
  emit('confirm', { ...form })
  emit('update:visible', false)
}
const onCancel = () => {
  emit('cancel')
  emit('update:visible', false)
}
const onBackdropClick = (e: Event) => {
  if (e.target === e.currentTarget) onCancel()
}
</script>

<style scoped>
.custom-dialog-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background-color: rgba(0,0,0,0.75);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999999;
  padding: 20px;
}
.custom-dialog-container {
  width: 100%;
  max-width: 420px;
  max-height: 90vh;
  overflow: hidden;
}
.custom-dialog {
  background: var(--bg-primary);
  border-radius: 16px;
  box-shadow: 0 25px 50px -12px rgba(0,0,0,0.5);
  border: 1px solid var(--border-primary);
  overflow: hidden;
  animation: dialogSlideIn 0.3s cubic-bezier(0.25,0.46,0.45,0.94);
}
.dialog-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px 28px 0 28px;
  border-bottom: 1px solid var(--border-primary);
  position: relative;
}
.dialog-title {
  font-size: 20px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0;
}
.dialog-close-btn {
  width: 32px; height: 32px;
  border: none; background: transparent;
  border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: var(--text-secondary);
  transition: all 0.2s ease;
}
.dialog-close-btn:hover {
  background: var(--bg-secondary);
  color: var(--text-primary);
}
.dialog-content {
  padding: 28px;
}
.form-group {
  margin-bottom: 18px;
  display: flex;
  flex-direction: column;
}
.form-group label {
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 6px;
}
.form-group input {
  padding: 8px 12px;
  border: 1px solid var(--border-primary);
  border-radius: 6px;
  font-size: 15px;
  color: var(--text-primary);
  background: var(--bg-secondary);
  outline: none;
  transition: border-color 0.2s;
}
.form-group input:focus {
  border-color: var(--accent-primary);
}
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding-top: 10px;
}
.cancel-btn {
  border-radius: 10px;
  font-weight: 600;
  padding: 10px 22px;
  font-size: 14px;
  background: var(--bg-secondary);
  border: 2px solid var(--border-primary);
  color: var(--text-primary);
  transition: all 0.3s;
}
.cancel-btn:hover {
  background: var(--bg-tertiary);
  border-color: var(--accent-primary);
  color: var(--accent-primary);
}
.confirm-btn {
  border-radius: 10px;
  font-weight: 600;
  padding: 10px 22px;
  font-size: 14px;
  background: linear-gradient(135deg, #3b82f6, #8b5cf6);
  border: none;
  color: white;
  box-shadow: 0 4px 16px rgba(59,130,246,0.15);
  text-shadow: 0 1px 2px rgba(0,0,0,0.1);
  transition: all 0.3s;
}
.confirm-btn:hover {
  background: linear-gradient(135deg, #8b5cf6, #3b82f6);
  box-shadow: 0 6px 20px rgba(59,130,246,0.25);
  transform: translateY(-2px);
}
@keyframes dialogSlideIn {
  0% { opacity: 0; transform: scale(0.8) translateY(-40px); }
  100% { opacity: 1; transform: scale(1) translateY(0); }
}
.dialog-fade-enter-active,
.dialog-fade-leave-active {
  transition: opacity 0.3s ease;
}
.dialog-fade-enter-from,
.dialog-fade-leave-to {
  opacity: 0;
}
</style>
