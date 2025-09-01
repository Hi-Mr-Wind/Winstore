/**
 * 应用商店状态管理
 * 使用Pinia进行状态管理，包含应用数据、分类、用户偏好等核心状态
 */
import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { AppItem, CategoryItem, DownloadProgress, UserPreferences, GoClassify } from '@/types/app'
import { SelectClassify } from '../../wailsjs/go/apps/App'

export const useAppStore = defineStore('app', () => {
  // ==================== 状态定义 ====================
  
  // 应用相关状态
  const apps = ref<AppItem[]>([])                    // 所有应用列表
  const categories = ref<CategoryItem[]>([])         // 所有分类列表
  const installedApps = ref<AppItem[]>([])           // 已安装应用列表
  const favorites = ref<AppItem[]>([])               // 收藏应用列表
  
  // 搜索和筛选状态
  const searchQuery = ref('')                        // 搜索关键词
  const selectedCategory = ref<string | null>(null)  // 当前选中的分类
  
  // 下载相关状态
  const downloadQueue = ref<DownloadProgress[]>([])  // 下载队列
  
  // 用户偏好设置
  const userPreferences = reactive<UserPreferences>({
    theme: 'light',                                  // 默认浅色主题
    language: 'zh-CN',                               // 默认中文
    autoUpdate: true,                                // 默认开启自动更新
    downloadPath: 'C:\\Downloads',                   // 默认下载路径
    notifications: true                              // 默认开启通知
  })

  // ==================== 计算属性 ====================
  
  // 根据搜索和分类筛选的应用列表
  const filteredApps = computed(() => {
    let filtered = apps.value

    // 按分类过滤
    if (selectedCategory.value) {
      filtered = filtered.filter(app => app.category === selectedCategory.value)
    }

    // 按搜索关键词过滤
    if (searchQuery.value.trim()) {
      const query = searchQuery.value.toLowerCase()
      filtered = filtered.filter(app => 
        app.name.toLowerCase().includes(query) ||
        app.developer.toLowerCase().includes(query) ||
        app.description.toLowerCase().includes(query)
      )
    }

    return filtered
  })

  // 热门应用列表（暂时返回前10个应用）
  const popularApps = computed(() => {
    // 暂时返回前10个应用，后续可以根据下载量、安装量等指标排序
    return apps.value.slice(0, 10)
  })

  // 判断应用是否已安装
  const isAppInstalled = computed(() => {
    return (appId: number) => installedApps.value.some(app => app.id === appId)
  })

  // 判断应用是否已收藏
  const isAppFavorited = computed(() => {
    return (appId: number) => favorites.value.some(app => app.id === appId)
  })

  // ==================== 动作方法 ====================
  
  // 设置应用列表
  const setApps = (newApps: AppItem[]) => {
    apps.value = newApps
  }

  // 设置分类列表
  const setCategories = (newCategories: CategoryItem[]) => {
    categories.value = newCategories
  }

  // 根据分类名称获取图标（作为备选方案）
  const getCategoryIcon = (categoryName: string): string => {
    const iconMap: Record<string, string> = {
      // '生产力': '💼',
      // '社交': '💬',
      // '游戏': '🎮',
      // '娱乐': '🎵',
      // '教育': '📚',
      // '工具': '🔧',
      // '开发': '💻',
      // '设计': '🎨',
      // '多媒体': '🎬',
      // '安全': '🔒',
      // '网络': '🌐',
      // '系统': '⚙️'
    }
    
    // 尝试精确匹配
    if (iconMap[categoryName]) {
      return iconMap[categoryName]
    }
    
    // 尝试模糊匹配
    for (const [key, icon] of Object.entries(iconMap)) {
      if (categoryName.includes(key) || key.includes(categoryName)) {
        return icon
      }
    }
    
    // 默认图标
    return '📱'
  }

  // 根据分类名称获取主题色
  const getCategoryColor = (categoryName: string): string => {
    const colorMap: Record<string, string> = {
      // '生产力': '#4CAF50',
      // '社交': '#2196F3',
      // '游戏': '#FF9800',
      // '娱乐': '#E91E63',
      // '教育': '#9C27B0',
      // '工具': '#607D8B',
      // '开发': '#795548',
      // '设计': '#FF5722',
      // '多媒体': '#00BCD4',
      // '安全': '#F44336',
      // '网络': '#3F51B5',
      // '系统': '#757575'
    }
    
    // 尝试精确匹配
    if (colorMap[categoryName]) {
      return colorMap[categoryName]
    }
    
    // 尝试模糊匹配
    for (const [key, color] of Object.entries(colorMap)) {
      if (categoryName.includes(key) || key.includes(categoryName)) {
        return color
      }
    }
    
    // 默认颜色
    return '#9E9E9E'
  }

  // 从Go后端获取分类数据
  const fetchCategoriesFromWails = async (): Promise<CategoryItem[]> => {
    try {
      console.log('正在从Go后端获取分类数据...')
      console.log('检查SelectClassify方法是否存在:', typeof SelectClassify)
      
      // 检查Wails环境
      if (typeof window !== 'undefined') {
        console.log('window对象存在')
        if (window.go) {
          console.log('window.go存在')
          if (window.go.apps) {
            console.log('window.go.apps存在')
            if (window.go.apps.App) {
              console.log('window.go.apps.App存在')
              console.log('SelectClassify方法:', window.go.apps.App.SelectClassify)
            } else {
              console.log('window.go.apps.App不存在')
            }
          } else {
            console.log('window.go.apps不存在')
          }
        } else {
          console.log('window.go不存在')
        }
      } else {
        console.log('window对象不存在')
      }
      
      const response = await SelectClassify()
      console.log('获取到的分类响应:', response)
      console.log('响应类型:', typeof response)
      console.log('响应结构:', Object.keys(response))
      
      // 检查响应状态
      if (response.code !== 0) {
        console.error('获取分类数据失败:', response.message)
        return []
      }
      
      // 转换Go后端的分类数据为前端格式
      const goCategories = response.data as GoClassify[]
      console.log('解析后的分类数据:', goCategories)
      console.log('数据类型:', typeof goCategories)
      console.log('是否为数组:', Array.isArray(goCategories))
      
      if (!Array.isArray(goCategories)) {
        console.error('分类数据格式错误:', response.data)
        return []
      }
      
      // 转换为CategoryItem格式，优先使用后端返回的icon和count
      const convertedCategories: CategoryItem[] = goCategories.map((goCategory, index) => {
        console.log(`[Store] 处理分类 ${goCategory.classify_name}:`, {
          id: goCategory.classify_id,
          name: goCategory.classify_name,
          icon: goCategory.icon,
          count: goCategory.count,
          hasCount: 'count' in goCategory
        })
        
        return {
          id: parseInt(goCategory.classify_id) || index + 1, // 使用后端返回的分类ID
          name: goCategory.classify_name,
          icon: goCategory.icon || getCategoryIcon(goCategory.classify_name), // 优先使用后端返回的图标
          count: goCategory.count || 0, // 使用后端返回的应用数量
          description: `${goCategory.classify_name}分类`,
          color: getCategoryColor(goCategory.classify_name) // 根据分类名称获取主题色
        }
      })
      
      console.log('转换后的分类数据:', convertedCategories)
      return convertedCategories
    } catch (error: any) {
      console.error('获取分类数据失败:', error)
      console.error('错误详情:', {
        name: error?.name,
        message: error?.message,
        stack: error?.stack
      })
      return []
    }
  }

  // 安装应用
  const installApp = async (app: AppItem) => {
    try {
      // 这里应该调用后端API进行实际安装
      console.log(`正在安装 ${app.name}...`)
      
      // 模拟安装过程
      await new Promise(resolve => setTimeout(resolve, 2000))
      
      // 添加到已安装列表
      if (!installedApps.value.find(installed => installed.id === app.id)) {
        installedApps.value.push(app)
      }
      
      // 更新应用状态
      const appIndex = apps.value.findIndex(a => a.id === app.id)
      if (appIndex !== -1) {
        apps.value[appIndex].installed = true
      }
      
      return { success: true, message: `${app.name} 安装成功` }
    } catch (error) {
      console.error('安装失败:', error)
      return { success: false, message: '安装失败，请重试' }
    }
  }

  // 卸载应用
  const uninstallApp = async (app: AppItem) => {
    try {
      console.log(`正在卸载 ${app.name}...`)
      
      // 模拟卸载过程
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      // 从已安装列表移除
      installedApps.value = installedApps.value.filter(installed => installed.id !== app.id)
      
      // 更新应用状态
      const appIndex = apps.value.findIndex(a => a.id === app.id)
      if (appIndex !== -1) {
        apps.value[appIndex].installed = false
      }
      
      return { success: true, message: `${app.name} 卸载成功` }
    } catch (error) {
      console.error('卸载失败:', error)
      return { success: false, message: '卸载失败，请重试' }
    }
  }

  // 切换应用收藏状态
  const toggleFavorite = (app: AppItem) => {
    const favoriteIndex = favorites.value.findIndex(fav => fav.id === app.id)
    
    if (favoriteIndex !== -1) {
      // 如果已收藏，则取消收藏
      favorites.value.splice(favoriteIndex, 1)
    } else {
      // 如果未收藏，则添加收藏
      favorites.value.push(app)
    }
  }

  // 设置搜索关键词
  const setSearchQuery = (query: string) => {
    searchQuery.value = query
  }

  // 设置选中的分类
  const setSelectedCategory = (category: string | null) => {
    selectedCategory.value = category
  }

  // 添加应用到下载队列
  const addToDownloadQueue = (app: AppItem) => {
    const download: DownloadProgress = {
      percentage: 0,
      downloaded: '0 MB',
      total: '0 MB',
      speed: '0 MB/s',
      status: 'pending'
    }
    
    downloadQueue.value.push(download)
    
    // 模拟下载过程
    simulateDownload(download, app)
  }

  // 模拟下载过程
  const simulateDownload = (download: DownloadProgress, app: AppItem) => {
    download.status = 'downloading'
    
    const interval = setInterval(() => {
      if (download.percentage < 100) {
        download.percentage += Math.random() * 10
        download.downloaded = `${Math.floor(download.percentage * 50)} MB`
        download.total = '500 MB'
        download.speed = `${(Math.random() * 5 + 2).toFixed(1)} MB/s`
      } else {
        download.status = 'completed'
        clearInterval(interval)
        
        // 下载完成后自动安装
        installApp(app)
      }
    }, 500)
  }

  // 更新用户偏好设置
  const updateUserPreferences = (preferences: Partial<UserPreferences>) => {
    Object.assign(userPreferences, preferences)
  }

  // 清除搜索条件
  const clearSearch = () => {
    searchQuery.value = ''
    selectedCategory.value = null
  }

  // ==================== 返回状态和方法 ====================
  return {
    // 状态
    apps,
    categories,
    installedApps,
    favorites,
    searchQuery,
    selectedCategory,
    downloadQueue,
    userPreferences,
    
    // 计算属性
    filteredApps,
    popularApps,
    isAppInstalled,
    isAppFavorited,
    
    // 动作方法
    setApps,
    setCategories,
    fetchCategoriesFromWails,
    installApp,
    uninstallApp,
    toggleFavorite,
    setSearchQuery,
    setSelectedCategory,
    addToDownloadQueue,
    updateUserPreferences,
    clearSearch
  }
})
