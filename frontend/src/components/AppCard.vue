<!--
  应用卡片组件
  功能：
  - 展示单个应用的基本信息（图标、名称、开发者）
  - 提供安装/卸载按钮
  - 提供收藏/取消收藏功能
  - 支持点击跳转到应用详情页
  - 响应式设计，适配不同屏幕尺寸
-->
<script setup lang="ts">
import { computed } from 'vue'
import { Star, StarFilled } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n'
import type { AppItem } from '@/types/app'

// 组件属性定义
interface Props {
  app: AppItem                    // 应用数据对象
  showFavorite?: boolean          // 是否显示收藏按钮
  showInstall?: boolean           // 是否显示安装按钮
  showView?: boolean              // 是否显示查看按钮
}

// 设置默认属性值
const props = withDefaults(defineProps<Props>(), {
  showFavorite: true,
  showInstall: true,
  showView: false
})

// 组件事件定义
const emit = defineEmits<{
  click: [app: AppItem]           // 点击应用卡片事件
  install: [app: AppItem]         // 安装应用事件
  uninstall: [app: AppItem]       // 卸载应用事件
  favorite: [app: AppItem]        // 收藏应用事件
}>()

const { t } = useI18n()

// 计算属性：判断应用是否已安装
const isInstalled = computed(() => props.app.installed)

// 计算属性：判断应用是否已收藏（这里应该从store获取，暂时返回false）
const isFavorited = computed(() => false)

// 处理卡片点击事件
const handleClick = () => {
  emit('click', props.app)
}

// 处理安装按钮点击事件
const handleInstall = () => {
  emit('install', props.app)
}

// 处理卸载按钮点击事件
const handleUninstall = () => {
  emit('uninstall', props.app)
}

// 处理收藏按钮点击事件
const handleFavorite = () => {
  emit('favorite', props.app)
}

// 处理图片加载失败
const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement;
  target.src = 'https://via.placeholder.com/60'; // 设置一个默认图片
};
</script>

<template>
  <div class="app-card" @click="handleClick">
    <!-- 应用图标区域 -->
    <div class="app-icon">
      <img 
        :src="app.icon" 
        :alt="app.name"
        class="app-icon-img"
        @error="handleImageError"
      />
    </div>
    
    <!-- 应用信息区域 -->
    <div class="app-info">
      <h4 class="app-name"><span class="label">软件名：</span>{{ app.name }}</h4>
      <p class="app-developer"><span class="label">发行商：</span>{{ app.developer }}</p>
    </div>
    
    <!-- 操作按钮区域 -->
    <div class="app-actions">
      <!-- 收藏按钮 -->
      <el-button
        v-if="showFavorite"
        :icon="isFavorited ? StarFilled : Star"
        circle
        size="small"
        :type="isFavorited ? 'warning' : 'default'"
        @click.stop="handleFavorite"
      />
      
      <!-- 查看按钮 -->
      <el-button
        v-if="showView"
        type="primary"
        size="small"
        @click.stop="handleClick"
      >
        查看
      </el-button>
      
      <!-- 安装/卸载按钮 -->
      <el-button
        v-if="showInstall"
        :type="isInstalled ? 'info' : 'primary'"
        size="small"
        @click.stop="isInstalled ? handleUninstall() : handleInstall()"
      >
        {{ isInstalled ? t('common.installed') : t('common.install') }}
      </el-button>
    </div>
  </div>
</template>

<style scoped>
/* 应用卡片容器 */
.app-card {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 20px;
  box-shadow: var(--shadow-light);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  border: 1px solid var(--border-primary);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 卡片悬停效果 */
.app-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

/* 应用图标区域 */
.app-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 72px;
  height: 72px;
  margin: 0 auto;
  background: var(--icon-bg);
  border: 1px solid var(--border-primary);
  border-radius: 16px;
  box-shadow: var(--shadow-light);
}

/* 应用图标图片样式 */
.app-icon-img {
  width: 100%;
  height: 100%;
  padding: 8px;
  border-radius: 16px;
  object-fit: contain;
  background-color: transparent;
}

/* 应用信息区域 */
.app-info {
  text-align: center;
}

/* 应用名称样式 */
.app-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
  line-height: 1.4;
}

.label {
  color: var(--text-secondary);
  font-weight: 500;
  margin-right: 6px;
}

/* 开发者名称样式 */
.app-developer {
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 0;
}

/* 操作按钮区域 */
.app-actions {
  display: flex;
  gap: 8px;
  justify-content: center;
  align-items: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .app-card {
    padding: 16px;
  }
  
  .app-name {
    font-size: 14px;
  }
  
  .app-developer {
    font-size: 12px;
  }
}
</style>
