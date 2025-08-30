<script setup lang="ts">
import { computed } from 'vue'
import { Star, StarFilled } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n'
import type { AppItem } from '@/types/app'

interface Props {
  app: AppItem
  showFavorite?: boolean
  showInstall?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  showFavorite: true,
  showInstall: true
})

const emit = defineEmits<{
  install: [app: AppItem]
  uninstall: [app: AppItem]
  favorite: [app: AppItem]
  click: [app: AppItem]
}>()

const { t } = useI18n()

const isInstalled = computed(() => props.app.installed)
const isFavorited = computed(() => false) // 这里应该从store获取

const handleInstall = () => {
  emit('install', props.app)
}

const handleUninstall = () => {
  emit('uninstall', props.app)
}

const handleFavorite = () => {
  emit('favorite', props.app)
}

const handleClick = () => {
  emit('click', props.app)
}
</script>

<template>
  <div class="app-card" @click="handleClick">
    <!-- 应用图标 -->
    <div class="app-icon">
      <el-avatar :size="60" :src="app.icon" />
    </div>

    <!-- 应用信息 -->
    <div class="app-info">
      <h4 class="app-name">{{ app.name }}</h4>
      <p class="app-developer">{{ app.developer }}</p>
      <div class="app-rating">
        <el-rate
          :model-value="app.rating"
          disabled
          show-score
          text-color="#ff9900"
          score-template="{value}"
        />
      </div>
    </div>

    <!-- 操作按钮 -->
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

.app-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

.app-icon {
  display: flex;
  justify-content: center;
}

.app-info {
  flex: 1;
}

.app-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
  line-height: 1.4;
}

.app-developer {
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 8px;
  line-height: 1.3;
}

.app-rating {
  margin-bottom: 12px;
}

.app-actions {
  display: flex;
  gap: 8px;
  justify-content: center;
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
