import { ref, watch, onMounted } from 'vue'

export type Theme = 'light' | 'dark' | 'auto'

const currentTheme = ref<Theme>('light')

// 检测系统主题
function getSystemTheme(): 'light' | 'dark' {
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

// 应用主题到DOM
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

// 获取实际主题（考虑auto模式）
function getActualTheme(): 'light' | 'dark' {
  if (currentTheme.value === 'auto') {
    return getSystemTheme()
  }
  return currentTheme.value
}

// 监听系统主题变化
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

// 设置主题
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

// 获取当前主题
export function getTheme(): Theme {
  return currentTheme.value
}

// 获取实际主题
export function getEffectiveTheme(): 'light' | 'dark' {
  return getActualTheme()
}

// 初始化主题
export function initTheme() {
  const savedTheme = localStorage.getItem('theme') as Theme
  if (savedTheme && ['light', 'dark', 'auto'].includes(savedTheme)) {
    setTheme(savedTheme)
  } else {
    setTheme('light')
  }
}

// 主题composable
export function useTheme() {
  const isDark = ref(false)
  
  // 监听主题变化
  watch(currentTheme, () => {
    isDark.value = getActualTheme() === 'dark'
  }, { immediate: true })
  
  // 初始化
  onMounted(() => {
    initTheme()
    const cleanup = watchSystemTheme()
    
    // 组件卸载时清理
    return cleanup
  })
  
  return {
    theme: currentTheme,
    isDark,
    setTheme,
    getTheme,
    getEffectiveTheme
  }
}
