<template>
  <el-drawer
    v-model="visible"
    title="系统清理"
    :size="520"
    direction="rtl"
    :before-close="handleClose"
    class="system-cleanup-drawer"
  >
    <div class="drawer-content">
      <div class="feature-description">
        <div class="feature-header">
          <div class="feature-icon">🧹</div>
          <div class="feature-text">
            <h3 class="feature-title">系统清理工具</h3>
            <p class="feature-subtitle">清理过期临时文件，调用系统磁盘清理释放空间</p>
          </div>
        </div>
      </div>

      <div class="section-card">
        <div class="section-header">
          <div class="section-icon">📄</div>
          <h3 class="section-title">清理过期临时文件</h3>
        </div>
        <div class="form-row">
          <span class="setting-label">保留天数</span>
          <el-input-number
            v-model="daysToKeep"
            :min="-1"
            :max="3650"
            :step="1"
            controls-position="right"
          />
        </div>
        <div class="hint">传 0 或 -1 按默认 7 天处理</div>
        <div class="action-buttons">
          <el-button type="primary" :loading="cleaningExpired" @click="runExpiredCleanup">
            <el-icon v-if="!cleaningExpired"><Switch /></el-icon>
            {{ cleaningExpired ? '清理中...' : '开始清理' }}
          </el-button>
          <el-button @click="resetExpired" :disabled="cleaningExpired">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </div>
        <div v-if="expiredResult" class="result">
          <el-alert :type="expiredResult.success ? 'success' : 'error'" :title="expiredResult.title" :description="expiredResult.message" :closable="false" show-icon />
        </div>
      </div>

      <div class="section-card">
        <div class="section-header">
          <div class="section-icon">🖥️</div>
          <h3 class="section-title">系统磁盘清理</h3>
        </div>
        <div class="form-row">
          <span class="setting-label">清理模式参数</span>
          <el-input-number
            v-model="diskCleanupMode"
            :min="0"
            :max="65535"
            :step="1"
            controls-position="right"
          />
        </div>
        <div class="hint">不同参数对应系统磁盘清理的预设配置，留空使用默认</div>
        <div class="action-buttons">
          <el-button type="success" :loading="runningDiskCleanup" @click="runDiskCleanup">
            <el-icon v-if="!runningDiskCleanup"><Promotion /></el-icon>
            {{ runningDiskCleanup ? '执行中...' : '调用磁盘清理' }}
          </el-button>
        </div>
        <div v-if="diskCleanupResult" class="result">
          <el-alert :type="diskCleanupResult.success ? 'success' : 'error'" :title="diskCleanupResult.title" :description="diskCleanupResult.message" :closable="false" show-icon />
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh, Switch, Promotion } from '@element-plus/icons-vue'
import { SysExpiredDocuments, SysGarbageCleanup } from '../../wailsjs/go/apps/App.js'

interface Props {
  modelValue: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{ 'update:modelValue': [value: boolean] }>()

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// 过期文件清理
const daysToKeep = ref<number>(7)
const cleaningExpired = ref(false)
const expiredResult = ref<{ success: boolean; title: string; message: string } | null>(null)

const runExpiredCleanup = async () => {
  if (cleaningExpired.value) return
  try {
    cleaningExpired.value = true
    expiredResult.value = null
    await SysExpiredDocuments(daysToKeep.value as unknown as number)
    expiredResult.value = { success: true, title: '清理完成', message: '过期临时文件清理已完成' }
    ElMessage.success('清理完成')
  } catch (error: any) {
    console.error('SysExpiredDocuments error:', error)
    expiredResult.value = { success: false, title: '清理失败', message: error?.message || '执行失败' }
    ElMessage.error(expiredResult.value.message)
  } finally {
    cleaningExpired.value = false
  }
}

const resetExpired = () => {
  if (cleaningExpired.value) return
  daysToKeep.value = 7
  expiredResult.value = null
}

// 磁盘清理
const diskCleanupMode = ref<number | null>(null)
const runningDiskCleanup = ref(false)
const diskCleanupResult = ref<{ success: boolean; title: string; message: string } | null>(null)

const runDiskCleanup = async () => {
  if (runningDiskCleanup.value) return
  try {
    runningDiskCleanup.value = true
    diskCleanupResult.value = null
    const mode = (diskCleanupMode.value ?? 0)
    await SysGarbageCleanup(mode)
    diskCleanupResult.value = { success: true, title: '已调用', message: '已调用系统磁盘清理，请根据提示操作' }
    ElMessage.success('已调用系统磁盘清理')
  } catch (error: any) {
    console.error('SysGarbageCleanup error:', error)
    diskCleanupResult.value = { success: false, title: '调用失败', message: error?.message || '执行失败' }
    ElMessage.error(diskCleanupResult.value.message)
  } finally {
    runningDiskCleanup.value = false
  }
}

const handleClose = (done: () => void) => {
  if (cleaningExpired.value || runningDiskCleanup.value) {
    ElMessage.warning('任务进行中，无法关闭')
    return
  }
  done()
}
</script>

<style scoped>
.system-cleanup-drawer {
  --el-drawer-bg-color: var(--bg-primary);
  --el-drawer-padding-primary: 24px;
}

.drawer-content {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.feature-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px;
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  border-radius: 16px;
  color: white;
}

.feature-icon { font-size: 28px; }
.feature-title { margin: 0; font-size: 18px; font-weight: 700; color: #fff; }
.feature-subtitle { margin: 4px 0 0; font-size: 13px; color: rgba(255,255,255,0.9); }

.section-card {
  background: var(--bg-secondary);
  border: 1px solid var(--border-primary);
  border-radius: 16px;
  padding: 20px;
}

.section-header { display: flex; align-items: center; gap: 12px; margin-bottom: 12px; }
.section-icon { font-size: 18px; }
.section-title { margin: 0; font-size: 16px; font-weight: 600; color: var(--text-primary); }

.form-row { display: flex; align-items: center; gap: 12px; }
.setting-label { font-size: 14px; color: var(--text-primary); }
.hint { margin-top: 8px; font-size: 12px; color: var(--text-secondary); }

.action-buttons { display: flex; gap: 12px; margin-top: 12px; }

.result { margin-top: 12px; }
</style>

