<!--
  应用详情页面组件
  功能：
  - 展示应用的详细信息（图标、名称、开发者、描述等）
  - 提供官方网站和下载按钮
  - 显示应用版本、大小、分类、发布日期等信息
  - 展示应用截图（目前为占位符）
  - 支持响应式设计
  - 移除评分系统，专注于应用信息展示
-->
<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import Sidebar from '@/components/Sidebar.vue'
import type { AppItem } from '@/types/app'
import { SelectByID } from '../../wailsjs/go/apps/App'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime.js'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

// 详情数据
const currentApp = ref<AppItem | null>(null)
const screenshots = ref<string[]>([])
const isLoading = ref(false)
const errorMessage = ref('')

const mapToAppItem = (item: any): AppItem => ({
  id: Number(item.app_id || item.AppID || item.id || 0),
  name: item.app_name || item.AppName || item.name || '未知软件',
  developer: item.company || item.Company || item.developer || '未知开发者',
  icon: item.app_icon || item.AppIcon || item.icon || '',
  installed: false,
  description: item.brief_introduction || item.BriefIntroduction || item.description || '',
  category: item.classify || item.Classify || item.category || '未分类',
  version: item.app_version || item.AppVersion || item.version || '',
  size: '',
  releaseDate: '',
  lastUpdated: item.update_time || item.UpdateTime || '',
  officialWebsite: item.official_website || item.OfficialWebsite || '',
  downloadUrl: item.download_url || item.DownloadURL || ''
})

const fetchDetail = async () => {
  const idParam = route.params.id as string
  if (!idParam) return
  isLoading.value = true
  errorMessage.value = ''
  currentApp.value = null
  screenshots.value = []
  try {
    const resp = await SelectByID(idParam)
    const code = (resp as any)?.code
    if (code !== 0) {
      errorMessage.value = (resp as any)?.message || '获取应用详情失败'
      return
    }
    // 兼容 data/Data
    const data = (resp as any)?.data ?? (resp as any)?.Data
    if (!data) {
      errorMessage.value = '返回数据为空'
      return
    }
    const appList = (data as any).appList ?? (data as any).AppList
    const imgList = (data as any).img ?? (data as any).Img
    if (appList) {
      currentApp.value = mapToAppItem(appList)
    }
    if (Array.isArray(imgList)) {
      // 兼容字段名 url/Url/img/imgUrl 等
      screenshots.value = imgList.map((it: any) => it.url || it.Url || it.img || it.imgUrl || it.ImgUrl).filter(Boolean)
    }
  } catch (e: any) {
    errorMessage.value = e?.message || '获取应用详情异常'
  } finally {
    isLoading.value = false
  }
}

// 处理官方网站按钮点击 - 在新窗口打开官方网站
const handleOfficialWebsite = async () => {
  if (!currentApp.value?.officialWebsite) return
  try {
    await BrowserOpenURL(currentApp.value.officialWebsite)
  } catch (e) {
    window.open(currentApp.value.officialWebsite, '_blank')
  }
}

// 处理下载按钮点击 - 在新窗口打开下载页面
const handleDownload = async () => {
  if (!currentApp.value?.downloadUrl) return
  try {
    await BrowserOpenURL(currentApp.value.downloadUrl)
  } catch (e) {
    window.open(currentApp.value.downloadUrl, '_blank')
  }
}

// 处理图片加载错误
const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement
  target.src = 'https://via.placeholder.com/120' // 替换为默认图片
}

// 初始化与路由变化时拉取详情
onMounted(fetchDetail)
watch(() => route.params.id, fetchDetail)
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏导航 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 应用详情内容 -->
      <div v-if="currentApp" class="app-detail">
        <!-- 应用头部信息区域 -->
        <div class="app-header">
          <!-- 应用图标 -->
          <div class="app-icon">
            <img 
              v-if="currentApp.icon && currentApp.icon !== ''" 
              :src="currentApp.icon" 
              :alt="currentApp.name"
              @error="handleImageError"
            />
          </div>
          
          <!-- 应用基本信息 -->
          <div class="app-info">
            <h1 class="app-name">{{ currentApp.name }}</h1>
            <p class="app-developer">{{ currentApp.developer }}</p>
            
            <!-- 操作按钮区域 -->
            <div class="app-actions">
              <!-- 官方网站按钮 -->
              <el-button
                v-if="currentApp.officialWebsite"
                type="default"
                size="large"
                @click="handleOfficialWebsite"
              >
                官方网站
              </el-button>
              
              <!-- 下载按钮 -->
              <el-button
                v-if="currentApp.downloadUrl"
                type="primary"
                size="large"
                @click="handleDownload"
              >
                立即下载
              </el-button>
            </div>
          </div>
        </div>

        <!-- 应用详细信息区域 -->
        <div class="app-details">
          <!-- 应用信息区域 -->
          <div class="detail-section">
            <h2>应用信息</h2>
            <div class="detail-grid">
              <div class="detail-item">
                <span class="label">版本：</span>
                <span>{{ currentApp.version || '未知' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">分类：</span>
                <span>{{ currentApp.category || '未分类' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">最后更新：</span>
                <span>{{ currentApp.lastUpdated || '未知' }}</span>
              </div>
            </div>
          </div>

          <!-- 应用描述区域 -->
          <div class="detail-section">
            <h2>应用描述</h2>
            <p class="app-description">{{ currentApp.description }}</p>
          </div>

          <!-- 应用截图区域 -->
          <div class="detail-section">
            <h2>应用截图</h2>
            <div class="screenshots">
              <template v-if="screenshots.length > 0">
                <el-image
                  v-for="(img, idx) in screenshots"
                  :key="idx"
                  :src="img"
                  fit="cover"
                  :preview-src-list="screenshots"
                  :initial-index="idx"
                  preview-teleported
                  lazy
                  class="screenshot"
                />
              </template>
              <el-empty v-else :description="t('app.noScreenshots')" />
            </div>
          </div>
        </div>
      </div>

      <!-- 应用不存在时的空状态 -->
      <el-empty 
        v-else 
        :description="errorMessage || '应用不存在'"
      >
        <el-button type="primary" @click="router.push('/')">
          返回首页
        </el-button>
      </el-empty>
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

/* 应用详情容器 */
.app-detail {
  max-width: 1000px;
}

/* 应用头部信息区域 */
.app-header {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 32px;
  margin-bottom: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
  display: flex;
  gap: 24px;
}

/* 应用图标区域 */
.app-icon {
  flex-shrink: 0;
}

/* 应用信息区域 */
.app-info {
  flex: 1;
}

/* 应用名称样式 */
.app-name {
  font-size: 32px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

/* 开发者名称样式 */
.app-developer {
  font-size: 16px;
  color: var(--text-secondary);
  margin-bottom: 24px;
}

/* 操作按钮区域 */
.app-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* 应用详细信息区域 */
.app-details {
  display: grid;
  gap: 24px;
}

/* 详情区域样式 */
.detail-section {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
}

/* 详情区域标题 */
.detail-section h2 {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-primary);
}

/* 详情信息网格布局 */
.detail-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  column-gap: 24px;
  row-gap: 12px;
}

/* 详情项样式 */
.detail-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 0;
}

/* 详情项标签样式 */
.detail-item .label {
  font-weight: 600;
  color: var(--text-secondary);
  min-width: 72px;
  text-align: right;
}

/* 应用描述样式 */
.app-description {
  font-size: 16px;
  line-height: 1.6;
  color: var(--text-primary);
}

/* 截图区域样式 */
.screenshots {
  min-height: 200px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 12px;
}

.screenshot {
  width: 100%;
  border-radius: 8px;
  box-shadow: var(--shadow-light);
  border: 1px solid var(--border-primary);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .app-header {
    flex-direction: column;
    text-align: center;
    padding: 24px;
  }
  
  .app-name {
    font-size: 24px;
  }
  
  .app-actions {
    justify-content: center;
  }
  
  .detail-grid {
    grid-template-columns: 1fr;
  }
}
</style>
