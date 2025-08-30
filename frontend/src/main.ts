import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { i18n, initLanguage } from './i18n'
import { initTheme } from './composables/useTheme'
import './styles/theme.css'

// 创建应用实例
const app = createApp(App)

// 创建Pinia实例
const pinia = createPinia()

// 初始化语言和主题
initLanguage()
initTheme()

// 使用插件
app.use(pinia)
app.use(router)
app.use(i18n)

// 挂载应用
app.mount('#app')
