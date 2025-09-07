<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { GetSysInfo } from '../../wailsjs/go/apps/App.js'
import { ClipboardSetText } from '../../wailsjs/runtime/runtime.js'
import type { sysUtils } from '../../wailsjs/go/models'
import Sidebar from '@/components/Sidebar.vue'
import ImageConversionDrawer from '@/components/ImageConversionDrawer.vue'
import { 
  Monitor, 
  Cpu, 
  Memo, 
  Clock, 
  Setting,
  Refresh,
  Tools
} from '@element-plus/icons-vue'

const { t } = useI18n()

// 系统信息数据
const sysInfo = ref<sysUtils.WinSysInfo | null>(null)
const loading = ref(false)

// 图片格式转换抽屉状态
const imageConversionDrawerVisible = ref(false)

// 工具列表
const tools = [
  {
    id: 1,
    name: '图片格式转换',
    description: '支持 JPG、PNG、GIF、BMP、TIF、WEBP 等格式转换',
    icon: '🖼️',
    status: 'available',
    action: 'image-conversion'
  },
  {
    id: 2,
    name: '系统清理',
    description: '清理系统垃圾文件，释放磁盘空间',
    icon: '🧹',
    status: 'coming-soon'
  },
  {
    id: 3,
    name: '软件更新',
    description: '检查并更新已安装的软件',
    icon: '🔄',
    status: 'coming-soon'
  },
  {
    id: 4,
    name: '系统优化',
    description: '优化系统性能设置',
    icon: '⚡',
    status: 'coming-soon'
  },
  {
    id: 5,
    name: '精美壁纸',
    description: '免费的精品壁纸',
    icon: '🎨',
    status: 'coming-soon'
  }
]

// 获取系统信息
const fetchSysInfo = async () => {
  try {
    loading.value = true
    
    // 检查是否在Wails环境中
    if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
      const info = await GetSysInfo()
      sysInfo.value = info
      console.log('系统信息获取成功:', info)
      console.log('内存信息:', info.TotalMemory, 'bytes')
    } else {
      console.warn('当前不在Wails环境中，无法获取系统信息')
      ElMessage.warning('当前环境不支持此功能，请在Wails应用中运行')
    }
  } catch (error) {
    console.error('获取系统信息失败:', error)
    ElMessage.error('获取系统信息失败，请重试')
  } finally {
    loading.value = false
  }
}

// 刷新系统信息
const refreshSysInfo = async () => {
  await fetchSysInfo()
  ElMessage.success('系统信息已刷新')
}

// 格式化内存大小
const formatMemory = (bytes: number) => {
  if (!bytes || bytes <= 0) {
    return '未知'
  }
  console.log('bytes', bytes)
  const gb = bytes 
  if (gb >= 1) {
    return `${gb} GB`
  } else {
    const mb = bytes * 1024 * 1024
    return `${mb.toFixed(1)} MB`
  }
}

// 格式化运行时间
const formatUptime = (uptime: string) => {
  // 假设uptime是类似 "3 days 2 hours 15 minutes" 的格式
  return uptime || '未知'
}

// 处理长文本显示
const truncateText = (text: string, maxLength: number = 20) => {
  if (!text) return '未知'
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

// 复制文本到剪贴板
const copyToClipboard = async (text: string, label: string) => {
  try {
    if (typeof window !== 'undefined' && window.go && window.go.apps && window.go.apps.App) {
      const success = await ClipboardSetText(text)
      if (success) {
        ElMessage.success(`${label} 已复制到剪贴板`)
      } else {
        ElMessage.error('复制失败，请重试')
      }
    } else {
      // 备用方案：使用浏览器原生API
      await navigator.clipboard.writeText(text)
      ElMessage.success(`${label} 已复制到剪贴板`)
    }
  } catch (error) {
    console.error('复制失败:', error)
    ElMessage.error('复制失败，请重试')
  }
}

// 处理工具点击
const handleToolClick = (tool: any) => {
  if (tool.status === 'available') {
    if (tool.action === 'image-conversion') {
      imageConversionDrawerVisible.value = true
    } else {
      console.log(`启动工具: ${tool.name}`)
    }
  } else {
    ElMessage.info('该功能即将推出，敬请期待！')
  }
}

// 组件挂载时获取系统信息
onMounted(() => {
  fetchSysInfo()
})
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 页面标题 -->
      <div class="page-header">
        <div class="header-content">
          <div class="title-section">
            <h1 class="page-title">
              <el-icon class="title-icon"><Tools /></el-icon>
              工具箱
            </h1>
            <p class="page-description">系统维护和优化工具</p>
          </div>
          <div class="header-actions">
            <el-button
              type="primary"
              :icon="Refresh"
              :loading="loading"
              @click="refreshSysInfo"
              class="refresh-btn"
            >
              {{ loading ? '刷新中...' : '刷新' }}
            </el-button>
          </div>
        </div>
      </div>

      <!-- 系统信息展示区域 -->
      <div class="sys-info-section">
        <div class="section-header">
          <h2 class="section-title">
            <el-icon><Monitor /></el-icon>
            系统信息
          </h2>
          <p class="section-description">实时监控您的系统状态和性能指标</p>
        </div>

        <!-- 加载状态 -->
        <div v-if="loading" class="loading-state">
          <div class="loading-container">
            <el-icon class="loading-icon" :size="48">
              <Refresh />
            </el-icon>
            <h3 class="loading-title">正在获取系统信息...</h3>
            <el-skeleton :rows="4" animated class="loading-skeleton" />
          </div>
        </div>

        <!-- 系统信息卡片 -->
        <div v-else-if="sysInfo" class="sys-info-grid">
          <!-- 系统基本信息 -->
          <div class="info-card system-card">
            <div class="card-header">
              <el-icon class="card-icon"><Tools /></el-icon>
              <h3 class="card-title">系统信息</h3>
            </div>
            <div class="card-content">
              <div class="info-item">
                <span class="info-label">系统类型:</span>
                <el-tooltip 
                  :content="`点击复制: ${sysInfo.SysType}`" 
                  placement="top" 
                  :show-after="300"
                  :hide-after="0"
                >
                  <span 
                    class="info-value long-text clickable"
                    @click="copyToClipboard(sysInfo.SysType, '系统类型')"
                  >
                    {{ truncateText(sysInfo.SysType, 18) }}
                  </span>
                </el-tooltip>
              </div>
              <div class="info-item">
                <span class="info-label">系统版本:</span>
                <el-tooltip 
                  :content="`点击复制: ${sysInfo.SysVersion}`" 
                  placement="top" 
                  :show-after="300"
                  :hide-after="0"
                >
                  <span 
                    class="info-value long-text clickable"
                    @click="copyToClipboard(sysInfo.SysVersion, '系统版本')"
                  >
                    {{ truncateText(sysInfo.SysVersion, 18) }}
                  </span>
                </el-tooltip>
              </div>
              <div class="info-item">
                <span class="info-label">系统架构:</span>
                <span class="info-value">{{ sysInfo.SysArch }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">计算机名:</span>
                <span class="info-value">{{ sysInfo.ComputerName }}</span>
              </div>
            </div>
          </div>

          <!-- CPU信息 -->
          <div class="info-card cpu-card">
            <div class="card-header">
              <el-icon class="card-icon"><Cpu /></el-icon>
              <h3 class="card-title">处理器</h3>
            </div>
            <div class="card-content">
              <div class="info-item">
                <span class="info-label">CPU类型:</span>
                <el-tooltip 
                  :content="`点击复制: ${sysInfo.CpuType}`" 
                  placement="top" 
                  :show-after="300"
                  :hide-after="0"
                >
                  <span 
                    class="info-value long-text clickable"
                    @click="copyToClipboard(sysInfo.CpuType, 'CPU类型')"
                  >
                    {{ truncateText(sysInfo.CpuType, 18) }}
                  </span>
                </el-tooltip>
              </div>
              <div class="info-item">
                <span class="info-label">物理核心:</span>
                <span class="info-value">{{ sysInfo.CpuCoreNum }} 核</span>
              </div>
              <div class="info-item">
                <span class="info-label">逻辑核心:</span>
                <span class="info-value">{{ sysInfo.CpuLogicNum }} 线程</span>
              </div>
            </div>
          </div>

          <!-- 内存信息 -->
          <div class="info-card memory-card">
            <div class="card-header">
              <el-icon class="card-icon"><Monitor /></el-icon>
              <h3 class="card-title">内存</h3>
            </div>
            <div class="card-content">
              <div class="info-item">
                <span class="info-label">总内存:</span>
                <span class="info-value memory-value">
                  {{ sysInfo.TotalMemory.toLocaleString()}}GB
                  <span class="memory-debug" v-if="sysInfo.TotalMemory">
                    ({{ sysInfo.TotalMemory.toLocaleString() }} GB)
                  </span>
                </span>
              </div>
              <div class="memory-bar">
                <div class="memory-progress">
                  <div class="memory-fill" style="width: 75%"></div>
                </div>
                <!-- <span class="memory-text">75% 已使用</span> -->
              </div>
            </div>
          </div>

          <!-- 运行时间信息 -->
          <div class="info-card uptime-card">
            <div class="card-header">
              <el-icon class="card-icon"><Clock /></el-icon>
              <h3 class="card-title">运行状态</h3>
            </div>
            <div class="card-content">
              <div class="info-item">
                <span class="info-label">运行时间:</span>
                <span class="info-value">{{ formatUptime(sysInfo.Uptime) }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">最后启动:</span>
                <span class="info-value">{{ sysInfo.LastBootUpTime }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 无数据状态 -->
        <div v-else class="empty-state">
          <el-empty description="无法获取系统信息">
            <el-button type="primary" @click="fetchSysInfo">
              重试
            </el-button>
          </el-empty>
        </div>
      </div>

      <!-- 工具网格区域 -->
      <div class="tools-section">
        <div class="section-header">
          <h2 class="section-title">
            <el-icon><Setting /></el-icon>
            系统工具
          </h2>
          <p class="section-description">专业的系统维护和优化工具集合</p>
        </div>

        <div class="tools-grid">
          <div 
            v-for="tool in tools" 
            :key="tool.id" 
            class="tool-card"
            :class="{ 
              'coming-soon': tool.status === 'coming-soon',
              'available': tool.status === 'available'
            }"
            @click="handleToolClick(tool)"
          >
            <div class="tool-icon">{{ tool.icon }}</div>
            <div class="tool-info">
              <h3 class="tool-name">{{ tool.name }}</h3>
              <p class="tool-description">{{ tool.description }}</p>
              <el-tag 
                v-if="tool.status === 'coming-soon'" 
                type="info" 
                size="small"
              >
                即将推出
              </el-tag>
              <el-tag 
                v-else-if="tool.status === 'available'" 
                type="success" 
                size="small"
              >
                可用
              </el-tag>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 图片格式转换抽屉 -->
    <ImageConversionDrawer v-model="imageConversionDrawerVisible" />
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

/* 页面标题区域 */
.page-header {
  margin-bottom: 24px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
}

.title-section {
  flex: 1;
}

.page-title {
  font-size: 28px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0 0 6px 0;
  display: flex;
  align-items: center;
  gap: 10px;
}

.title-icon {
  font-size: 24px;
  color: var(--accent-primary);
}

.page-description {
  font-size: 14px;
  color: var(--text-secondary);
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.refresh-btn {
  min-width: 90px;
  height: 36px;
  border-radius: 8px;
  font-weight: 500;
}

/* 区域标题样式 */
.section-header {
  margin-bottom: 20px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0 0 6px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.section-description {
  font-size: 13px;
  color: var(--text-secondary);
  margin: 0;
}

/* 系统信息区域 */
.sys-info-section {
  margin-bottom: 32px;
}

/* 加载状态样式 */
.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 20px;
  min-height: 200px;
}

.loading-container {
  text-align: center;
  max-width: 350px;
}

.loading-icon {
  color: var(--accent-primary);
  margin-bottom: 12px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.loading-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.loading-skeleton {
  margin-top: 16px;
}

/* 系统信息网格 - 紧凑布局 */
.sys-info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 16px;
}

/* 信息卡片样式 - 紧凑设计 */
.info-card {
  background: linear-gradient(135deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
  border-radius: 12px;
  padding: 18px;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  border: 1px solid var(--border-primary);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.info-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, var(--accent-primary), var(--accent-secondary));
}

.info-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
}

/* 特殊卡片样式 */
.system-card::before {
  background: linear-gradient(90deg, #667eea, #764ba2);
}

.cpu-card::before {
  background: linear-gradient(90deg, #f093fb, #f5576c);
}

.memory-card::before {
  background: linear-gradient(90deg, #4facfe, #00f2fe);
}

.uptime-card::before {
  background: linear-gradient(90deg, #43e97b, #38f9d7);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 16px;
}

.card-icon {
  font-size: 20px;
  color: var(--accent-primary);
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 0;
  border-bottom: 1px solid var(--border-secondary);
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

.info-value {
  font-size: 13px;
  color: var(--text-primary);
  font-weight: 600;
  text-align: right;
  max-width: 140px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* 长文本样式 */
.info-value.long-text {
  cursor: help;
  position: relative;
}

.info-value.long-text:hover {
  color: var(--accent-primary);
}

/* 可点击文本样式 */
.info-value.clickable {
  cursor: pointer;
  padding: 2px 4px;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.info-value.clickable:hover {
  background-color: var(--accent-primary);
  color: white;
  transform: scale(1.05);
}

/* 内存值样式 */
.memory-value {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 3px;
}

.memory-debug {
  font-size: 9px;
  color: var(--text-tertiary);
  font-weight: 400;
  font-family: monospace;
}

/* 内存进度条 */
.memory-bar {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.memory-progress {
  width: 100%;
  height: 6px;
  background-color: var(--border-secondary);
  border-radius: 3px;
  overflow: hidden;
}

.memory-fill {
  height: 100%;
  background: linear-gradient(90deg, #4facfe, #00f2fe);
  border-radius: 3px;
  transition: width 0.3s ease;
}

.memory-text {
  font-size: 11px;
  color: var(--text-secondary);
  text-align: center;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 20px;
  min-height: 200px;
}

/* 工具区域 */
.tools-section {
  margin-bottom: 24px;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.tool-card {
  background: linear-gradient(135deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  border: 1px solid var(--border-primary);
  transition: all 0.3s ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 14px;
  position: relative;
  overflow: hidden;
}

.tool-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, var(--accent-primary), var(--accent-secondary));
  opacity: 0;
  transition: opacity 0.3s ease;
}

.tool-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
}

.tool-card:hover::before {
  opacity: 1;
}

.tool-card.coming-soon {
  opacity: 0.6;
  cursor: not-allowed;
}

.tool-card.coming-soon:hover {
  transform: none;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
}

.tool-card.coming-soon::before {
  display: none;
}

.tool-card.available {
  background: linear-gradient(135deg, var(--accent-primary) 0%, var(--accent-secondary) 100%);
  color: white;
  border: 2px solid var(--accent-primary);
}

.tool-card.available::before {
  background: linear-gradient(90deg, rgba(255,255,255,0.3), rgba(255,255,255,0.1));
  opacity: 1;
}

.tool-card.available:hover {
  transform: translateY(-5px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.2);
}

.tool-card.available .tool-name {
  color: white;
}

.tool-card.available .tool-description {
  color: rgba(255, 255, 255, 0.9);
}

.tool-icon {
  font-size: 28px;
  flex-shrink: 0;
}

.tool-info {
  flex: 1;
}

.tool-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 6px;
}

.tool-description {
  font-size: 13px;
  color: var(--text-secondary);
  margin-bottom: 6px;
  line-height: 1.4;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .header-content {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }
  
  .header-actions {
    justify-content: flex-start;
  }
  
  .refresh-btn {
    min-width: 100px;
  }
  
  .page-title {
    font-size: 22px;
  }
  
  .section-title {
    font-size: 18px;
  }
  
  .sys-info-grid {
    grid-template-columns: 1fr;
  }
  
  .tools-grid {
    grid-template-columns: 1fr;
  }
  
  .info-card {
    padding: 16px;
  }
  
  .tool-card {
    padding: 16px;
  }
}
</style>
