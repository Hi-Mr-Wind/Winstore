import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { AppItem, CategoryItem, DownloadProgress, UserPreferences } from '@/types/app'

export const useAppStore = defineStore('app', () => {
  // 状态
  const apps = ref<AppItem[]>([])
  const categories = ref<CategoryItem[]>([])
  const installedApps = ref<AppItem[]>([])
  const favorites = ref<AppItem[]>([])
  const searchQuery = ref('')
  const selectedCategory = ref<string | null>(null)
  const downloadQueue = ref<DownloadProgress[]>([])
  const userPreferences = reactive<UserPreferences>({
    theme: 'light',
    language: 'zh-CN',
    autoUpdate: true,
    downloadPath: 'C:\\Downloads',
    notifications: true
  })

  // 计算属性
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

  const popularApps = computed(() => {
    return apps.value
      .sort((a, b) => b.rating - a.rating)
      .slice(0, 10)
  })

  const isAppInstalled = computed(() => {
    return (appId: number) => installedApps.value.some(app => app.id === appId)
  })

  const isAppFavorited = computed(() => {
    return (appId: number) => favorites.value.some(app => app.id === appId)
  })

  // 动作
  const setApps = (newApps: AppItem[]) => {
    apps.value = newApps
  }

  const setCategories = (newCategories: CategoryItem[]) => {
    categories.value = newCategories
  }

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

  const toggleFavorite = (app: AppItem) => {
    const favoriteIndex = favorites.value.findIndex(fav => fav.id === app.id)
    
    if (favoriteIndex !== -1) {
      favorites.value.splice(favoriteIndex, 1)
    } else {
      favorites.value.push(app)
    }
  }

  const setSearchQuery = (query: string) => {
    searchQuery.value = query
  }

  const setSelectedCategory = (category: string | null) => {
    selectedCategory.value = category
  }

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

  const updateUserPreferences = (preferences: Partial<UserPreferences>) => {
    Object.assign(userPreferences, preferences)
  }

  const clearSearch = () => {
    searchQuery.value = ''
    selectedCategory.value = null
  }

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
    
    // 动作
    setApps,
    setCategories,
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
