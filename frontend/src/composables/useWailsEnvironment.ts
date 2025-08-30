import { ref, onMounted } from 'vue'

// 扩展 Window 接口以包含 Wails 相关属性
declare global {
  interface Window {
    go?: {
      apps?: {
        App?: any
      }
    }
    runtime?: any
  }
}

/**
 * Wails 环境检测和优化工具
 * 专门用于处理 Wails WebView 环境中的特殊问题
 */
export function useWailsEnvironment() {
  const isWailsEnvironment = ref(false)
  const isWebView = ref(false)
  const webViewType = ref<string>('')

  // 检测 Wails 环境
  const detectWailsEnvironment = () => {
    try {
      // 检查是否存在 window.go 对象
      const hasWailsGo = typeof window !== 'undefined' && 
        window.go !== undefined && 
        window.go.apps !== undefined &&
        window.go.apps.App !== undefined

      // 检查是否在 WebView 环境中
      const hasWebView = typeof window !== 'undefined' && 
        (window.navigator.userAgent.includes('Wails') || 
         window.navigator.userAgent.includes('WebView') ||
         window.navigator.userAgent.includes('Electron'))

      // 检查 Wails 运行时
      const hasWailsRuntime = typeof window !== 'undefined' && 
        (window as any).runtime !== undefined

      isWailsEnvironment.value = hasWailsGo || hasWailsRuntime
      isWebView.value = hasWebView

      // 确定 WebView 类型
      if (hasWebView) {
        if (window.navigator.userAgent.includes('Wails')) {
          webViewType.value = 'Wails'
        } else if (window.navigator.userAgent.includes('Electron')) {
          webViewType.value = 'Electron'
        } else {
          webViewType.value = 'Generic WebView'
        }
      }

      console.log('Wails 环境检测结果:', {
        isWailsEnvironment: isWailsEnvironment.value,
        isWebView: isWebView.value,
        webViewType: webViewType.value,
        userAgent: window.navigator.userAgent
      })

      return isWailsEnvironment.value
    } catch (error) {
      console.error('Wails 环境检测失败:', error)
      isWailsEnvironment.value = false
      return false
    }
  }

  // 应用 Wails WebView 特定的样式优化
  const applyWebViewOptimizations = () => {
    if (!isWebView.value) return

    console.log('应用 WebView 特定优化...')

    // 添加 WebView 特定的 CSS 类
    document.documentElement.classList.add('wails-webview')
    document.body.classList.add('wails-webview-body')

    // 优化滚动性能
    if (document.body.style) {
      const bodyStyle = document.body.style as any
      bodyStyle.webkitOverflowScrolling = 'touch'
      bodyStyle.overflowScrolling = 'touch'

      // 优化字体渲染
      bodyStyle.webkitFontSmoothing = 'antialiased'
      bodyStyle.mozOsxFontSmoothing = 'grayscale'
      bodyStyle.fontSmoothing = 'antialiased'
    }

    // 优化动画性能
    const style = document.createElement('style')
    style.textContent = `
      .wails-webview * {
        -webkit-transform: translateZ(0);
        transform: translateZ(0);
        -webkit-backface-visibility: hidden;
        backface-visibility: hidden;
      }
      
      .wails-webview .custom-dialog-overlay {
        -webkit-transform: translateZ(0);
        transform: translateZ(0);
      }
      
      .wails-webview .custom-dialog {
        -webkit-transform: translateZ(0);
        transform: translateZ(0);
      }
    `
    document.head.appendChild(style)
  }

  // 修复 WebView 中的弹窗问题
  const fixWebViewDialogIssues = () => {
    if (!isWebView.value) return

    console.log('修复 WebView 弹窗问题...')

    // 强制设置弹窗样式
    const style = document.createElement('style')
    style.textContent = `
      /* WebView 弹窗强制修复 */
      .custom-dialog-overlay {
        position: fixed !important;
        top: 0 !important;
        left: 0 !important;
        right: 0 !important;
        bottom: 0 !important;
        z-index: 999999 !important;
        display: flex !important;
        align-items: center !important;
        justify-content: center !important;
        background-color: rgba(0, 0, 0, 0.75) !important;
        backdrop-filter: blur(8px) !important;
        -webkit-backdrop-filter: blur(8px) !important;
        -webkit-transform: translateZ(0) !important;
        transform: translateZ(0) !important;
      }

      .custom-dialog {
        position: relative !important;
        top: auto !important;
        left: auto !important;
        right: auto !important;
        bottom: auto !important;
        transform: none !important;
        margin: 0 !important;
        z-index: 1000000 !important;
        max-width: 90vw !important;
        max-height: 90vh !important;
        overflow: auto !important;
        box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5) !important;
        border-radius: 16px !important;
        border: 1px solid var(--border-primary) !important;
        background: var(--bg-primary) !important;
        -webkit-transform: translateZ(0) !important;
        transform: translateZ(0) !important;
      }

      /* 确保弹窗在所有 WebView 环境中都正确显示 */
      body .custom-dialog-overlay {
        position: fixed !important;
        top: 0 !important;
        left: 0 !important;
        right: 0 !important;
        bottom: 0 !important;
        z-index: 999999 !important;
        display: flex !important;
        align-items: center !important;
        justify-content: center !important;
        background-color: rgba(0, 0, 0, 0.75) !important;
        backdrop-filter: blur(8px) !important;
        -webkit-backdrop-filter: blur(8px) !important;
      }

      body .custom-dialog {
        position: relative !important;
        top: auto !important;
        left: auto !important;
        right: auto !important;
        bottom: auto !important;
        transform: none !important;
        margin: 0 !important;
        z-index: 1000000 !important;
        max-width: 90vw !important;
        max-height: 90vh !important;
        overflow: auto !important;
        box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5) !important;
        border-radius: 16px !important;
        border: 1px solid var(--border-primary) !important;
        background: var(--bg-primary) !important;
      }
    `
    document.head.appendChild(style)
  }

  // 初始化 Wails 环境
  const initializeWailsEnvironment = () => {
    detectWailsEnvironment()
    
    if (isWebView.value) {
      applyWebViewOptimizations()
      fixWebViewDialogIssues()
    }
  }

  // 获取环境信息
  const getEnvironmentInfo = () => {
    return {
      isWailsEnvironment: isWailsEnvironment.value,
      isWebView: isWebView.value,
      webViewType: webViewType.value,
      userAgent: typeof window !== 'undefined' ? window.navigator.userAgent : '',
      hasWailsGo: typeof window !== 'undefined' && window.go !== undefined,
      hasWailsRuntime: typeof window !== 'undefined' && (window as any).runtime !== undefined
    }
  }

  // 组件挂载时自动检测
  onMounted(() => {
    initializeWailsEnvironment()
  })

  return {
    isWailsEnvironment,
    isWebView,
    webViewType,
    detectWailsEnvironment,
    applyWebViewOptimizations,
    fixWebViewDialogIssues,
    initializeWailsEnvironment,
    getEnvironmentInfo
  }
}
