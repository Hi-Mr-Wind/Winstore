import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router'
import { i18n, initLanguage } from './i18n'
import App from './App.vue'
import { initTheme } from './composables/useTheme'
import './styles/theme.css'
import { useWailsEnvironment } from './composables/useWailsEnvironment'

// 创建应用实例
const app = createApp(App)

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 使用插件
app.use(createPinia())
app.use(ElementPlus)
app.use(router)
app.use(i18n)

// 初始化语言和主题
initLanguage()
initTheme()

// 初始化 Wails 环境优化
const { initializeWailsEnvironment } = useWailsEnvironment()
initializeWailsEnvironment()

// 挂载应用
app.mount('#app')
