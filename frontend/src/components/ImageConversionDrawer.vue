<template>
  <el-drawer
    v-model="visible"
    title="图片格式转换"
    :size="600"
    direction="rtl"
    :before-close="handleClose"
    class="image-conversion-drawer"
  >
    <div class="drawer-content">
      <!-- 功能说明 -->
      <div class="feature-description">
        <div class="feature-header">
          <div class="feature-icon">🖼️</div>
          <div class="feature-text">
            <h3 class="feature-title">图片格式转换工具</h3>
            <p class="feature-subtitle">支持将图片转换为 JPG、PNG、GIF、BMP、TIF、WEBP 等格式</p>
          </div>
        </div>
      </div>

      <!-- 文件选择区域 -->
      <div class="file-selection">
        <div class="section-header">
          <div class="section-icon">📁</div>
          <h3 class="section-title">选择图片文件</h3>
        </div>
        <div class="file-input-group">
          <el-input
            v-model="inputFilePath"
            placeholder="请选择要转换的图片文件"
            readonly
            class="file-path-input"
            size="large"
          >
            <template #prepend>
              <el-icon><Picture /></el-icon>
            </template>
            <template #append>
              <el-button @click="selectInputFile" :icon="FolderOpened" type="primary">
                选择文件
              </el-button>
            </template>
          </el-input>
        </div>
        
        <!-- 文件信息显示 -->
        <div v-if="inputFileInfo" class="file-info">
          <div class="file-info-item">
            <span class="label">文件名:</span>
            <span class="value">{{ inputFileInfo.name }}</span>
          </div>
          <!-- <div class="file-info-item">
            <span class="label">文件大小:</span>
            <span class="value">{{ formatFileSize(inputFileInfo.size) }}</span>
          </div> -->
          <div class="file-info-item">
            <span class="label">当前格式:</span>
            <span class="value">{{ inputFileInfo.extension?.toUpperCase() || '未知' }}</span>
          </div>
        </div>
      </div>

      <!-- 输出设置区域 -->
      <div class="output-settings">
        <div class="section-header">
          <div class="section-icon">⚙️</div>
          <h3 class="section-title">输出设置</h3>
        </div>
        
        <!-- 目标格式选择 -->
        <div class="format-selection">
          <label class="setting-label">目标格式:</label>
          <el-radio-group 
            v-model="targetFormat" 
            class="format-radio-group"
          >
            <div class="format-grid">
              <el-radio :value="0" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">JPG</div>
                  <div class="format-desc">JPEG 格式，适合照片</div>
                </div>
              </el-radio>
              <el-radio :value="1" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">PNG</div>
                  <div class="format-desc">PNG 格式，支持透明背景</div>
                </div>
              </el-radio>
              <el-radio :value="2" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">GIF</div>
                  <div class="format-desc">GIF 格式，支持动画</div>
                </div>
              </el-radio>
              <el-radio :value="3" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">BMP</div>
                  <div class="format-desc">BMP 格式，无损压缩</div>
                </div>
              </el-radio>
              <el-radio :value="4" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">TIF</div>
                  <div class="format-desc">TIFF 格式，高质量</div>
                </div>
              </el-radio>
              <el-radio :value="5" class="format-radio">
                <div class="format-radio-content">
                  <div class="format-name">WEBP</div>
                  <div class="format-desc">WebP 格式，现代格式</div>
                </div>
              </el-radio>
            </div>
          </el-radio-group>
        </div>

        <!-- 输出路径设置 -->
        <div class="output-path-group">
          <label class="setting-label">输出路径:</label>
          <div class="file-input-group">
            <el-input
              v-model="outputFilePath"
              placeholder="选择输出文件夹"
              readonly
              class="file-path-input"
              size="large"
            >
              <template #prepend>
                <el-icon><FolderOpened /></el-icon>
              </template>
              <template #append>
                <el-button @click="selectOutputFile" :icon="FolderOpened" type="success">
                  选择文件夹
                </el-button>
              </template>
            </el-input>
          </div>
        </div>
      </div>

      <!-- 转换预览 -->
      <div v-if="inputFileInfo && targetFormat" class="conversion-preview">
        <div class="section-header">
          <div class="section-icon">👁️</div>
          <h3 class="section-title">转换预览</h3>
        </div>
        <div class="preview-card">
          <div class="preview-item">
            <div class="preview-icon">📄</div>
            <div class="preview-content">
              <div class="preview-label">源文件</div>
              <div class="preview-value">{{ inputFileInfo.name }}</div>
            </div>
          </div>
          <div class="preview-arrow">
            <el-icon><Right /></el-icon>
          </div>
          <div class="preview-item">
            <div class="preview-icon">🎯</div>
            <div class="preview-content">
              <div class="preview-label">目标文件</div>
              <div class="preview-value">{{ getOutputFileName() }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="action-buttons">
        <el-button
          type="primary"
          :loading="converting"
          :disabled="!canConvert"
          @click="startConversion"
          class="convert-btn"
          size="large"
        >
          <el-icon v-if="!converting"><Switch /></el-icon>
          {{ converting ? '转换中...' : '开始转换' }}
        </el-button>
        <el-button @click="resetForm" :disabled="converting" size="large">
          <el-icon><Refresh /></el-icon>
          重置
        </el-button>
      </div>

      <!-- 转换进度 -->
      <div v-if="converting" class="conversion-progress">
        <el-progress
          :percentage="conversionProgress"
          :status="conversionStatus"
          :stroke-width="8"
          class="progress-bar"
        />
        <p class="progress-text">{{ conversionMessage }}</p>
      </div>

      <!-- 转换结果 -->
      <div v-if="conversionResult" class="conversion-result">
        <el-alert
          :title="conversionResult.success ? '转换成功' : '转换失败'"
          :description="conversionResult.message"
          :type="conversionResult.success ? 'success' : 'error'"
          :closable="false"
          show-icon
        />
        
        <!-- 错误详情显示 -->
        <div v-if="!conversionResult.success && conversionResult.details" class="error-details">
          <el-collapse>
            <el-collapse-item title="查看详细错误信息" name="error-details">
              <div class="error-details-content">
                <pre>{{ conversionResult.details }}</pre>
              </div>
            </el-collapse-item>
          </el-collapse>
        </div>
        
        <div v-if="conversionResult.success" class="result-actions">
          <el-button type="primary" @click="openOutputFolder">
            <el-icon><FolderOpened /></el-icon>
            打开输出文件夹
          </el-button>
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { FolderOpened, Right, Switch, Picture, Refresh } from '@element-plus/icons-vue'
import { ImageConversionType } from '../../wailsjs/go/apps/App.js'
import { OpenImageDialog, OpenDirectoryDialog } from '../../wailsjs/go/apps/App.js'

// Props
interface Props {
  modelValue: boolean
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

// 响应式数据
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const inputFilePath = ref('')
const outputFilePath = ref('')
const targetFormat = ref<number | null>(null)
const converting = ref(false)
const conversionProgress = ref(0)
const conversionStatus = ref<'success' | 'exception' | 'warning' | undefined>()
const conversionMessage = ref('')
const conversionResult = ref<{ success: boolean; message: string; details?: string } | null>(null)

// 文件信息
const inputFileInfo = ref<{
  name: string
  size: number
  extension?: string
} | null>(null)

// 格式选项
const formatOptions = [
  { value: 0, label: 'JPG', description: 'JPEG 格式，适合照片' },
  { value: 1, label: 'PNG', description: 'PNG 格式，支持透明背景' },
  { value: 2, label: 'GIF', description: 'GIF 格式，支持动画' },
  { value: 3, label: 'BMP', description: 'BMP 格式，无损压缩' },
  { value: 4, label: 'TIF', description: 'TIFF 格式，高质量' },
  { value: 5, label: 'WEBP', description: 'WebP 格式，现代格式' }
]

// 计算属性
const canConvert = computed(() => {
  return inputFilePath.value && 
         outputFilePath.value && 
         targetFormat.value !== null && 
         !converting.value
})

// 监听输入文件路径变化
watch(inputFilePath, (newPath) => {
  if (newPath) {
    const fileName = newPath.split(/[\\/]/).pop() || ''
    const extension = fileName.split('.').pop()?.toLowerCase()
    inputFileInfo.value = {
      name: fileName,
      size: 0, // 无法直接获取文件大小，设为0
      extension
    }
    // 不自动设置输出路径，让用户自己选择
  } else {
    inputFileInfo.value = null
  }
})

// 选择输入文件
const selectInputFile = async () => {
  try {
    const result = await OpenImageDialog()
    
    if (result) {
      inputFilePath.value = result
      conversionResult.value = null
    }
  } catch (error) {
    console.error('选择图片文件失败:', error)
    ElMessage.error('选择图片文件失败，请重试')
  }
}

// 选择输出文件夹
const selectOutputFile = async () => {
  if (!inputFileInfo.value || targetFormat.value === null) {
    ElMessage.warning('请先选择输入文件和目标格式')
    return
  }

  try {
    const result = await OpenDirectoryDialog()
    
    if (result) {
      const outputFileName = getOutputFileName()
      outputFilePath.value = `${result}`
    }
  } catch (error) {
    console.error('选择输出文件夹失败:', error)
    ElMessage.error('选择输出文件夹失败，请重试')
  }
}

// 更新输出路径（仅在用户选择文件夹后）
const updateOutputPath = () => {
  // 只有当用户已经选择了输出文件夹时，才更新输出路径
  // 这里不做任何自动设置，让用户自己选择
}

// 获取输出文件名
const getOutputFileName = () => {
  if (!inputFileInfo.value || targetFormat.value === null) {
    return ''
  }
  
  const baseName = inputFileInfo.value.name.replace(/\.[^/.]+$/, '')
  const extension = getFormatExtension(targetFormat.value)
  return `${baseName}.${extension}`
}

// 获取格式扩展名
const getFormatExtension = (format: number) => {
  const extensions = ['jpg', 'png', 'gif', 'bmp', 'tif', 'webp']
  return extensions[format] || 'jpg'
}

// 获取格式显示名称
const getFormatDisplayName = (format: number) => {
  const option = formatOptions.find(opt => opt.value === format)
  return option?.label || '未知格式'
}

// 格式化文件大小
const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '未知大小'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// 开始转换
const startConversion = async () => {
  if (!canConvert.value) {
    ElMessage.warning('请完善所有必要信息')
    return
  }

  try {
    converting.value = true
    conversionProgress.value = 0
    conversionStatus.value = undefined
    conversionMessage.value = '准备转换...'
    conversionResult.value = null

    // 模拟进度更新
    const progressInterval = setInterval(() => {
      if (conversionProgress.value < 90) {
        conversionProgress.value += Math.random() * 20
        conversionMessage.value = '正在转换图片格式...'
      }
    }, 200)

    // 调用后端转换方法
    await ImageConversionType(inputFilePath.value, outputFilePath.value, targetFormat.value!)

    clearInterval(progressInterval)
    conversionProgress.value = 100
    conversionStatus.value = 'success'
    conversionMessage.value = '转换完成!'
    
    conversionResult.value = {
      success: true,
      message: `图片已成功转换为 ${getFormatDisplayName(targetFormat.value!)} 格式`
    }

    ElMessage.success('图片转换成功!')
  } catch (error) {
    console.error('图片转换失败:', error)
    conversionProgress.value = 0
    conversionStatus.value = 'exception'
    conversionMessage.value = '转换失败'
    
    // 提取详细的错误信息
    let errorMessage = '未知错误'
    let errorDetails = ''
    
    if (error instanceof Error) {
      errorMessage = error.message
      errorDetails = error.stack || ''
    } else if (typeof error === 'string') {
      errorMessage = error
    } else if (error && typeof error === 'object') {
      // 处理可能的Go后端错误对象
      const errorObj = error as any
      if (errorObj.message) {
        errorMessage = errorObj.message
      } else if (errorObj.error) {
        errorMessage = errorObj.error
      } else if (errorObj.msg) {
        errorMessage = errorObj.msg
      }
      
      // 尝试获取更多错误详情
      if (errorObj.details) {
        errorDetails = errorObj.details
      } else if (errorObj.stack) {
        errorDetails = errorObj.stack
      }
    }
    
    conversionResult.value = {
      success: false,
      message: `转换失败: ${errorMessage}`,
      details: errorDetails
    }

    // 显示详细的错误信息
    ElMessage.error(`图片转换失败: ${errorMessage}`)
    
    // 在控制台输出完整错误信息用于调试
    console.error('完整错误信息:', {
      error,
      message: errorMessage,
      details: errorDetails,
      inputPath: inputFilePath.value,
      outputPath: outputFilePath.value,
      targetFormat: targetFormat.value
    })
  } finally {
    converting.value = false
  }
}

// 打开输出文件夹
const openOutputFolder = async () => {
  if (outputFilePath.value) {
    try {
      const outputDir = outputFilePath.value.substring(0, outputFilePath.value.lastIndexOf('/') || outputFilePath.value.lastIndexOf('\\'))
      // 使用系统默认程序打开文件夹
      if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
        // 这里可以调用系统API打开文件夹，暂时显示消息
        ElMessage.info(`输出文件夹: ${outputDir}`)
      } else {
        ElMessage.info(`输出文件夹: ${outputDir}`)
      }
    } catch (error) {
      console.error('打开文件夹失败:', error)
      ElMessage.error('打开文件夹失败')
    }
  }
}

// 重置表单
const resetForm = () => {
  if (converting.value) {
    ElMessage.warning('转换进行中，无法重置')
    return
  }

  ElMessageBox.confirm(
    '确定要重置所有设置吗？',
    '确认重置',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    inputFilePath.value = ''
    outputFilePath.value = ''
    targetFormat.value = null
    inputFileInfo.value = null
    conversionResult.value = null
    conversionProgress.value = 0
    conversionMessage.value = ''
    ElMessage.success('已重置所有设置')
  }).catch(() => {
    // 用户取消
  })
}

// 关闭抽屉
const handleClose = (done: () => void) => {
  if (converting.value) {
    ElMessage.warning('转换进行中，无法关闭')
    return
  }
  done()
}
</script>

<style scoped>
.image-conversion-drawer {
  --el-drawer-bg-color: var(--bg-primary);
  --el-drawer-padding-primary: 24px;
}

.drawer-content {
  display: flex;
  flex-direction: column;
  gap: 28px;
  height: 100%;
  padding: 8px 0;
}

.feature-description {
  margin-bottom: 12px;
}

.feature-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px;
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  border-radius: 16px;
  color: white;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.feature-icon {
  font-size: 32px;
  flex-shrink: 0;
}

.feature-text {
  flex: 1;
}

.feature-title {
  font-size: 20px;
  font-weight: 700;
  margin: 0 0 6px 0;
  color: white;
}

.feature-subtitle {
  font-size: 14px;
  margin: 0;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.4;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.section-icon {
  font-size: 20px;
  flex-shrink: 0;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.file-selection,
.output-settings,
.conversion-preview {
  background: var(--bg-secondary);
  border-radius: 16px;
  padding: 24px;
  border: 1px solid var(--border-primary);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.file-selection:hover,
.output-settings:hover,
.conversion-preview:hover {
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.file-input-group {
  margin-bottom: 16px;
}

.file-path-input {
  width: 100%;
}

.file-info {
  background: var(--bg-primary);
  border-radius: 8px;
  padding: 12px;
  border: 1px solid var(--border-secondary);
}

.file-info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 0;
}

.file-info-item:not(:last-child) {
  border-bottom: 1px solid var(--border-secondary);
}

.file-info-item .label {
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

.file-info-item .value {
  font-size: 13px;
  color: var(--text-primary);
  font-weight: 600;
}

.setting-label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.format-selection,
.output-path-group {
  margin-bottom: 16px;
}

.format-select {
  width: 100%;
}

.format-radio-group {
  width: 100%;
}

.format-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-top: 12px;
}

.format-radio {
  width: 100%;
  margin: 0;
  padding: 16px 12px;
  border: 2px solid var(--border-secondary);
  border-radius: 12px;
  background: var(--bg-primary);
  transition: all 0.3s ease;
  cursor: pointer;
  min-height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.format-radio:hover {
  border-color: var(--accent-primary);
  background: var(--bg-secondary);
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.format-radio.is-checked {
  border-color: var(--accent-primary);
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
}

.format-radio.is-checked .format-name {
  color: white;
}

.format-radio.is-checked .format-desc {
  color: rgba(255, 255, 255, 0.9);
}

.format-radio-content {
  display: flex;
  flex-direction: column;
  gap: 6px;
  text-align: center;
  width: 100%;
  overflow: hidden;
}

.format-name {
  font-weight: 700;
  color: var(--text-primary);
  font-size: 16px;
  margin: 0;
}

.format-desc {
  font-size: 11px;
  color: var(--text-secondary);
  line-height: 1.3;
  margin: 0;
  word-wrap: break-word;
  overflow-wrap: break-word;
  hyphens: auto;
}

.conversion-preview {
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  color: white;
}

.preview-card {
  display: flex;
  align-items: center;
  gap: 20px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 20px;
  backdrop-filter: blur(10px);
}

.preview-item {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 12px;
  text-align: left;
}

.preview-icon {
  font-size: 24px;
  flex-shrink: 0;
}

.preview-content {
  flex: 1;
}

.preview-label {
  font-size: 12px;
  opacity: 0.8;
  margin-bottom: 4px;
  font-weight: 500;
}

.preview-value {
  font-size: 14px;
  font-weight: 600;
  word-break: break-all;
  line-height: 1.3;
}

.preview-arrow {
  font-size: 24px;
  opacity: 0.8;
  color: rgba(255, 255, 255, 0.9);
}

.action-buttons {
  display: flex;
  gap: 16px;
  justify-content: center;
  margin-top: 12px;
  padding: 20px;
  background: var(--bg-secondary);
  border-radius: 16px;
  border: 1px solid var(--border-primary);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
}

.convert-btn {
  min-width: 160px;
  height: 48px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.convert-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
}

.conversion-progress {
  text-align: center;
  padding: 20px;
  background: var(--bg-secondary);
  border-radius: 12px;
  border: 1px solid var(--border-primary);
}

.progress-bar {
  margin-bottom: 12px;
}

.progress-text {
  font-size: 14px;
  color: var(--text-secondary);
  margin: 0;
}

.conversion-result {
  text-align: center;
}

.result-actions {
  margin-top: 16px;
}

.error-details {
  margin-top: 12px;
}

.error-details-content {
  background: var(--bg-primary);
  border: 1px solid var(--border-secondary);
  border-radius: 6px;
  padding: 12px;
  max-height: 200px;
  overflow-y: auto;
}

.error-details-content pre {
  margin: 0;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  color: var(--text-primary);
  white-space: pre-wrap;
  word-break: break-word;
  line-height: 1.4;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .drawer-content {
    gap: 20px;
  }
  
  .feature-header {
    padding: 16px;
    gap: 12px;
  }
  
  .feature-icon {
    font-size: 24px;
  }
  
  .feature-title {
    font-size: 18px;
  }
  
  .feature-subtitle {
    font-size: 13px;
  }
  
  .file-selection,
  .output-settings,
  .conversion-preview {
    padding: 20px;
  }
  
  .section-header {
    gap: 10px;
    margin-bottom: 14px;
  }
  
  .section-icon {
    font-size: 18px;
  }
  
  .section-title {
    font-size: 16px;
  }
  
  .format-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
  }
  
  .format-radio {
    padding: 14px 10px;
    min-height: 70px;
  }
  
  .format-name {
    font-size: 15px;
  }
  
  .format-desc {
    font-size: 10px;
  }
  
  .preview-card {
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  }
  
  .preview-item {
    flex-direction: column;
    text-align: center;
    gap: 8px;
  }
  
  .preview-arrow {
    transform: rotate(90deg);
  }
  
  .action-buttons {
    flex-direction: column;
    padding: 16px;
  }
  
  .convert-btn {
    width: 100%;
    min-width: auto;
  }
}

@media (max-width: 480px) {
  .format-grid {
    grid-template-columns: 1fr;
  }
  
  .feature-header {
    flex-direction: column;
    text-align: center;
    gap: 8px;
  }
}
</style>
