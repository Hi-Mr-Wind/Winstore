/**
 * 主题管理组合式函数
 * 提供主题切换功能，支持浅色、深色和自动模式
 * 自动检测系统主题偏好，并持久化用户设置
 */
import { ref, watch, onMounted } from 'vue'

// 主题类型定义
export type Theme = 'light' | 'dark' | 'auto'

// 当前主题状态
const currentTheme = ref<Theme>('light')

/**
 * 检测系统主题偏好
 * @returns 'light' | 'dark' 系统当前主题
 */
function getSystemTheme(): 'light' | 'dark' {
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

/**
 * 应用主题到DOM元素
 * @param theme 要应用的主题
 */
function applyTheme(theme: 'light' | 'dark') {
  const root = document.documentElement
  
  if (theme === 'dark') {
    root.classList.add('dark')
    root.classList.remove('light')
  } else {
    root.classList.add('light')
    root.classList.remove('dark')
  }
  
  // 设置CSS变量
  root.style.setProperty('--theme-mode', theme)
}

/**
 * 获取实际主题（考虑auto模式）
 * @returns 'light' | 'dark' 实际应用的主题
 */
function getActualTheme(): 'light' | 'dark' {
  if (currentTheme.value === 'auto') {
    return getSystemTheme()
  }
  return currentTheme.value
}

/**
 * 监听系统主题变化
 * @returns 清理函数
 */
function watchSystemTheme() {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  
  const handleChange = () => {
    if (currentTheme.value === 'auto') {
      applyTheme(getSystemTheme())
    }
  }
  
  mediaQuery.addEventListener('change', handleChange)
  
  // 返回清理函数
  return () => {
    mediaQuery.removeEventListener('change', handleChange)
  }
}

/**
 * 设置主题
 * @param theme 要设置的主题
 */
export function setTheme(theme: Theme) {
  currentTheme.value = theme
  
  if (theme === 'auto') {
    applyTheme(getSystemTheme())
  } else {
    applyTheme(theme)
  }
  
  // 保存到localStorage
  localStorage.setItem('theme', theme)
}

/**
 * 获取当前主题
 * @returns 当前主题设置
 */
export function getTheme(): Theme {
  return currentTheme.value
}

/**
 * 获取实际主题
 * @returns 实际应用的主题
 */
export function getEffectiveTheme(): 'light' | 'dark' {
  return getActualTheme()
}

/**
 * 初始化主题设置
 * 从localStorage读取保存的主题设置
 */
export function initTheme() {
  const savedTheme = localStorage.getItem('theme') as Theme
  if (savedTheme && ['light', 'dark', 'auto'].includes(savedTheme)) {
    setTheme(savedTheme)
  } else {
    setTheme('light')
  }
}

/**
 * 主题管理组合式函数
 * 提供主题状态和操作方法
 */
export function useTheme() {
  const isDark = ref(false)
  
  // 监听主题变化
  watch(currentTheme, () => {
    isDark.value = getActualTheme() === 'dark'
  }, { immediate: true })
  
  // 组件挂载时初始化
  onMounted(() => {
    initTheme()
    const cleanup = watchSystemTheme()
    
    // 组件卸载时清理
    return cleanup
  })
  
  return {
    theme: currentTheme,           // 当前主题设置
    isDark,                       // 是否为深色主题
    setTheme,                     // 设置主题方法
    getTheme,                     // 获取主题方法
    getEffectiveTheme            // 获取实际主题方法
  }
}
