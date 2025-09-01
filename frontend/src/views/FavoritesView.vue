<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Sidebar from '@/components/Sidebar.vue'
import CustomCollectDialog from '@/components/CustomCollectDialog.vue'
import { CollectInsert, CollectSelectAll, CollectDelete, CollectUpdate } from '../../wailsjs/go/apps/App'
import type { model } from '../../wailsjs/go/models'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime'

const router = useRouter()

// 状态
const favorites = ref<model.Collect[]>([])
const dialogVisible = ref(false)
const editDialogVisible = ref(false)
const editTarget = ref<model.Collect | null>(null)
const isLoading = ref(false)
const errorMessage = ref('')

// 工具：健壮解析后端返回为 Collect[]
function normalizeCollectList(raw: any): model.Collect[] {
  try {
    // 顶层为字符串 → 先尝试整体 parse
    if (typeof raw === 'string') {
      try {
        const parsed = JSON.parse(raw)
        return normalizeCollectList(parsed)
      } catch (e) {
        console.warn('[Favorites] 顶层为字符串且无法直接JSON.parse，尝试提取数组子串', e)
        const match = raw.match(/\[[\s\S]*\]/)
        if (match) {
          try { return normalizeCollectList(JSON.parse(match[0])) } catch (e2) { console.error('[Favorites] 数组子串解析失败', e2) }
        }
        // 尝试逐行解析
        const byLines: model.Collect[] = []
        raw.split('\n').forEach((line: string) => {
          const t = line.trim()
          if (!t) return
          try {
            const obj = JSON.parse(t)
            if (obj && typeof obj === 'object') byLines.push(obj)
          } catch {}
        })
        if (byLines.length) return byLines
        return []
      }
    }

    // 顶层为数组
    if (Array.isArray(raw)) {
      return raw as model.Collect[]
    }

    // BaseModel 包装 { code, message, data }
    if (raw && typeof raw === 'object' && 'data' in raw) {
      const data = (raw as any).data
      if (typeof data === 'string') {
        try { return normalizeCollectList(JSON.parse(data)) } catch (e) {
          console.error('[Favorites] data为字符串但解析失败', e)
          return []
        }
      }
      if (Array.isArray(data)) return data as model.Collect[]
      return normalizeCollectList(data)
    }

    // 其他未知结构
    return []
  } catch (e) {
    console.error('[Favorites] normalizeCollectList 异常:', e)
    return []
  }
}

// 拉取收藏
const fetchFavorites = async () => {
  isLoading.value = true
  errorMessage.value = ''
  try {
    console.log('[Favorites] 调用 CollectSelectAll() 获取收藏列表...')
    const res = await CollectSelectAll()
    console.log('[Favorites] CollectSelectAll 返回原始数据:', res)
    const list = normalizeCollectList(res)
    favorites.value = list
    console.log('[Favorites] 解析后的长度:', favorites.value.length, '示例:', favorites.value[0])

    // 如果为空，尝试回退 window.go 调用，观察差异
    if (favorites.value.length === 0) {
      const w: any = (globalThis as any)
      const fn = w?.go?.apps?.App?.CollectSelectAll
      if (typeof fn === 'function') {
        console.log('[Favorites] 列表为空，回退直接调用 window.go.apps.App.CollectSelectAll()')
        const raw = await fn()
        console.log('[Favorites] 回退原始返回:', raw)
        const list2 = normalizeCollectList(raw)
        favorites.value = list2
        console.log('[Favorites] 回退解析后的长度:', favorites.value.length, '示例:', favorites.value[0])
      }
    }
  } catch (err: any) {
    console.error('[Favorites] 获取收藏列表失败:', err)
    errorMessage.value = String(err?.message || err) || '获取收藏列表失败'
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  console.log('[Favorites] 页面挂载，开始拉取收藏列表')
  fetchFavorites()
})

// 删除收藏
const handleDelete = async (id: string | number) => {
  try {
    const idStr = String(id)
    console.log('[Favorites] 准备删除收藏，id=', idStr)
    await CollectDelete(idStr)
    console.log('[Favorites] 删除成功，刷新列表')
    fetchFavorites()
  } catch (err) {
    console.error('[Favorites] 删除收藏失败:', err)
  }
}

// 添加收藏
const handleAdd = async (form: { appName: string; downloadUrl: string; appUrl: string; appIcon: string }) => {
  const payload: model.Collect = {
    id: '',
    appName: form.appName,
    downloadUrl: form.downloadUrl,
    appUrl: form.appUrl,
    appIcon: form.appIcon,
    createTime: ''
  }
  try {
    console.log('[Favorites] 准备添加收藏，payload=', payload)
    const res = await CollectInsert(payload)
    console.log('[Favorites] CollectInsert 返回:', res)
    await fetchFavorites()
  } catch (err) {
    console.error('[Favorites] 添加收藏失败:', err)
  }
}

// 打开编辑
const openEdit = (item: model.Collect) => {
  editTarget.value = item
  editDialogVisible.value = true
}

// 提交编辑
const handleEditConfirm = async (form: { appName: string; downloadUrl: string; appUrl: string; appIcon: string }) => {
  if (!editTarget.value) return
  const payload: model.Collect = {
    id: String(editTarget.value.id || ''),
    appName: form.appName,
    downloadUrl: form.downloadUrl,
    appUrl: form.appUrl,
    appIcon: form.appIcon,
    createTime: editTarget.value.createTime || ''
  }
  try {
    console.log('[Favorites] 准备更新收藏，payload=', payload)
    const res = await CollectUpdate(payload)
    console.log('[Favorites] CollectUpdate 返回:', res)
    editDialogVisible.value = false
    editTarget.value = null
    await fetchFavorites()
  } catch (err) {
    console.error('[Favorites] 更新收藏失败:', err)
  }
}

// 图标加载失败回退
const onIconError = (e: Event) => {
  const el = e.target as HTMLImageElement
  el.onerror = null
  el.src = 'data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24"><rect width="100%" height="100%" fill="%23f0f0f0"/><path fill="%23999" d="M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>'
}

// 使用 Wails Runtime 打开外部链接
const openExternal = (url?: string) => {
  if (!url) return
  try {
    BrowserOpenURL(url)
  } catch (e) {
    console.error('[Favorites] 打开外部链接失败:', e)
    // 兜底使用 window.open
    // try { window.open(url, '_blank') } catch {}
  }
}
</script>

<template>
  <div class="app-container">
    <Sidebar />
    <div class="main-content">
      <div class="page-header">
        <h1 class="page-title">收藏室</h1>
        <p class="page-description">管理您本地收藏的应用</p>
        <div class="actions">
          <button class="add-btn" @click="dialogVisible = true">+ 添加收藏</button>
          <button class="refresh-btn" :disabled="isLoading" @click="fetchFavorites">{{ isLoading ? '刷新中...' : '刷新' }}</button>
        </div>
      </div>

      <div v-if="errorMessage" class="error-line">{{ errorMessage }}</div>

      <div v-if="favorites.length > 0" class="collect-list">
        <div class="collect-item" v-for="(item, idx) in favorites" :key="item.id || idx">
          <div class="collect-left">
            <img class="collect-icon" :src="item.appIcon" alt="icon" @error="onIconError" />
            <div class="collect-info">
              <div class="collect-title">{{ item.appName }}</div>
              <div class="collect-meta">
                <span class="meta-label">收藏时间：</span>
                <span class="meta-value">{{ item.createTime || '未知' }}</span>
              </div>
              <div class="collect-links">
                <a v-if="item.downloadUrl" href="javascript:;" @click.prevent="openExternal(item.downloadUrl)">下载</a>
                <span v-if="item.downloadUrl && item.appUrl">|</span>
                <a v-if="item.appUrl" href="javascript:;" @click.prevent="openExternal(item.appUrl)">官网</a>
                <span v-if="item.downloadUrl || item.appUrl" class="sep">|</span>
                <button class="link-edit" @click="openEdit(item)">编辑</button>
              </div>
            </div>
          </div>
          <div class="collect-actions">
            <button class="delete-btn" @click="handleDelete(item.id)">删除</button>
          </div>
        </div>
      </div>

      <div v-else class="empty-collect">
        <p>您还没有收藏任何应用</p>
        <button class="add-btn" @click="dialogVisible = true">+ 添加收藏</button>
      </div>

      <CustomCollectDialog :visible="dialogVisible" @confirm="handleAdd" @cancel="dialogVisible = false" @update:visible="val => dialogVisible = val" />
      <CustomCollectDialog :visible="editDialogVisible" :initial="editTarget ? { appName: editTarget.appName, downloadUrl: editTarget.downloadUrl || '', appUrl: editTarget.appUrl || '', appIcon: editTarget.appIcon || '' } : undefined" title="编辑收藏" confirmText="保存" @confirm="handleEditConfirm" @cancel="editDialogVisible = false" @update:visible="val => editDialogVisible = val" />
    </div>
  </div>
</template>

<style scoped>
/* 优化配色与卡片视觉 */
.app-container { display: flex; height: 100vh; background: var(--bg-secondary); }
.main-content { flex: 1; padding: 24px; overflow-y: auto; }
.page-header { margin-bottom: 24px; display: flex; align-items: center; gap: 16px; }
.actions { margin-left: auto; display: flex; gap: 10px; }
.page-title { font-size: 28px; font-weight: 700; color: var(--text-primary); }
.page-description { font-size: 14px; color: var(--text-secondary); }
.error-line { color: #c62828; background: #ffebee; border: 1px solid #ffcdd2; padding: 8px 10px; border-radius: 8px; margin-bottom: 12px; }

.add-btn, .refresh-btn, .edit-btn, .delete-btn { border: none; border-radius: 10px; padding: 10px 16px; font-size: 13px; font-weight: 700; cursor: pointer; transition: transform .15s ease, box-shadow .2s ease; }
.add-btn { background: linear-gradient(135deg, #3b82f6, #8b5cf6); color: #fff; box-shadow: 0 6px 16px rgba(59,130,246,.25); }
.add-btn:hover { transform: translateY(-2px); box-shadow: 0 10px 24px rgba(59,130,246,.35); }
.refresh-btn { background: var(--bg-tertiary); color: var(--text-primary); border: 1px solid var(--border-primary); }
.refresh-btn:disabled { opacity: .6; cursor: default; }
.refresh-btn:hover { transform: translateY(-2px); box-shadow: var(--shadow-light); }

.collect-list { display: flex; flex-direction: column; gap: 14px; }
.collect-item { display: flex; align-items: center; justify-content: space-between; background: var(--bg-primary); border: 1px solid var(--border-primary); border-radius: 14px; box-shadow: var(--shadow-light); padding: 16px 18px; }
.collect-item:hover { box-shadow: var(--shadow-medium); }
.collect-left { display: flex; align-items: center; gap: 14px; }
.collect-icon { width: 48px; height: 48px; border-radius: 12px; object-fit: cover; background: var(--bg-tertiary); border: 1px solid var(--border-primary); }
.collect-info { display: flex; flex-direction: column; gap: 6px; }
.collect-title { font-size: 16px; font-weight: 700; color: var(--text-primary); letter-spacing: .2px; }
.collect-links a { color: var(--accent-primary); text-decoration: none; margin-right: 6px; font-size: 13px; }
.collect-links a:hover { text-decoration: underline; }
.collect-actions { display: flex; gap: 10px; }
.edit-btn { background: #10b981; color: #fff; box-shadow: 0 4px 14px rgba(16,185,129,.25); }
.edit-btn:hover { transform: translateY(-2px); box-shadow: 0 8px 20px rgba(16,185,129,.35); }
.delete-btn { background: #ef4444; color: #fff; box-shadow: 0 4px 14px rgba(239,68,68,.25); }
.delete-btn:hover { transform: translateY(-2px); box-shadow: 0 8px 20px rgba(239,68,68,.35); }

.empty-collect { text-align: center; margin-top: 80px; color: var(--text-secondary); }
.empty-collect .add-btn { margin-top: 12px; }

/* 新增内联编辑样式 */
.collect-links { display: flex; align-items: center; gap: 6px; flex-wrap: wrap; }
.collect-links .sep { color: var(--text-tertiary); }
.link-edit { background: transparent; border: none; padding: 0; margin: 0; color: var(--accent-primary); font-weight: 700; cursor: pointer; }
.link-edit:hover { text-decoration: underline; }

/* 收藏时间样式 */
.collect-meta { font-size: 12px; color: var(--text-secondary); margin: 2px 0 6px; }
.meta-label { margin-right: 4px; }

@media (max-width: 768px) {
  .main-content { padding: 16px; }
  .collect-item { flex-direction: column; align-items: stretch; gap: 12px; }
  .collect-actions { justify-content: flex-end; }
}
</style>
