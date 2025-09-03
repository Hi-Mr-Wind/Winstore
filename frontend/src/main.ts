import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router'
import { i18n, initLanguage } from './i18n'
import App from './App.vue'
import { initTheme } from './composables/useTheme'
import './styles/theme.css'
import { useWailsEnvironment } from './composables/useWailsEnvironment'
import { useAppStore } from './stores/appStore'
import { GetNewDataBaseVersion, UpdateDatabase, GetDatabaseVersion } from '../wailsjs/go/apps/App.js'
import { ElMessageBox, ElMessage } from 'element-plus'

// 创建应用实例
const app = createApp(App)

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 使用插件
const pinia = createPinia()
app.use(pinia)
app.use(ElementPlus)
app.use(router)
app.use(i18n)

// 初始化语言和主题
initLanguage()
initTheme()

// 初始化 Wails 环境优化
const { initializeWailsEnvironment } = useWailsEnvironment()
initializeWailsEnvironment()

// 启动时自动检查软件数据库更新（当开启自动更新时）
const startupAutoCheck = async () => {
  try {
    const store = useAppStore()
    // 启动时拉取一次当前数据库版本（仅日志用途）
    try {
      const v = await GetDatabaseVersion()
      console.debug('当前数据库版本: ', v)
    } catch {}

    if (!store.userPreferences.autoUpdate) return
    const hasNew = await GetNewDataBaseVersion()
    if (hasNew === true) {
      try {
        await ElMessageBox.confirm('检测到新的软件数据库，是否现在更新？', '提示', {
          confirmButtonText: '立即更新',
          cancelButtonText: '稍后',
          type: 'warning'
        })
        await UpdateDatabase()
        ElMessage.success('软件库已更新到最新。')
        try {
          const v = await GetDatabaseVersion()
          console.debug('更新后数据库版本: ', v)
        } catch {}
      } catch {
        // 用户取消或更新失败会在 UpdateDatabase 捕获
      }
    } else if (hasNew === false) {
      // 静默
    }
  } catch (err: any) {
    console.error('数据库更新检查失败: ', err)
  }
}

queueMicrotask(() => { startupAutoCheck() })

// 挂载应用
app.mount('#app')
