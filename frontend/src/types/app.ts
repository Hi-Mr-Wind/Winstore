// 应用项类型定义
export interface AppItem {
  id: number
  name: string
  developer: string
  icon: string
  rating: number
  installed: boolean
  description: string
  version?: string
  size?: string
  category?: string
  tags?: string[]
  screenshots?: string[]
  downloadUrl?: string
  releaseDate?: string
  lastUpdated?: string
}

// 分类项类型定义
export interface CategoryItem {
  id: number
  name: string
  icon: string
  count: number
  description?: string
  color?: string
}

// 下载进度类型定义
export interface DownloadProgress {
  percentage: number
  downloaded: string
  total: string
  speed: string
  status: 'pending' | 'downloading' | 'completed' | 'error'
}

// 搜索结果类型定义
export interface SearchResult {
  apps: AppItem[]
  categories: CategoryItem[]
  total: number
}

// 应用商店状态类型定义
export interface AppStoreState {
  apps: AppItem[]
  categories: CategoryItem[]
  installedApps: AppItem[]
  favorites: AppItem[]
  searchQuery: string
  selectedCategory: string | null
  downloadQueue: DownloadProgress[]
}

// 用户偏好类型定义
export interface UserPreferences {
  theme: 'light' | 'dark' | 'auto'
  language: string
  autoUpdate: boolean
  downloadPath: string
  notifications: boolean
}
