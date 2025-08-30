/**
 * 应用商店类型定义文件
 * 包含应用、分类、下载进度、搜索结果等核心数据结构的类型定义
 */

// 应用项类型定义
export interface AppItem {
  id: number                    // 应用唯一标识符
  name: string                  // 应用名称
  developer: string             // 开发者名称
  icon: string                  // 应用图标URL
  installed: boolean            // 是否已安装
  description: string           // 应用描述
  version?: string              // 应用版本号
  size?: string                 // 应用大小
  category?: string             // 应用分类
  tags?: string[]               // 应用标签
  screenshots?: string[]        // 应用截图URL数组
  downloadUrl?: string          // 下载链接
  officialWebsite?: string      // 官方网站链接
  releaseDate?: string          // 发布日期
  lastUpdated?: string          // 最后更新日期
}

// 分类项类型定义
export interface CategoryItem {
  id: number                    // 分类唯一标识符
  name: string                  // 分类名称
  icon: string                  // 分类图标（emoji或URL）
  count: number                 // 该分类下的应用数量
  description?: string          // 分类描述
  color?: string                // 分类主题色
}

// 下载进度类型定义
export interface DownloadProgress {
  percentage: number            // 下载进度百分比
  downloaded: string            // 已下载大小
  total: string                 // 总大小
  speed: string                 // 下载速度
  status: 'pending' | 'downloading' | 'completed' | 'error'  // 下载状态
}

// 搜索结果类型定义
export interface SearchResult {
  apps: AppItem[]               // 匹配的应用列表
  categories: CategoryItem[]    // 匹配的分类列表
  total: number                 // 搜索结果总数
}

// 应用商店状态类型定义
export interface AppStoreState {
  apps: AppItem[]               // 所有应用列表
  categories: CategoryItem[]    // 所有分类列表
  installedApps: AppItem[]      // 已安装应用列表
  favorites: AppItem[]          // 收藏应用列表
  searchQuery: string           // 搜索关键词
  selectedCategory: string | null  // 当前选中的分类
  downloadQueue: DownloadProgress[]  // 下载队列
}

// 用户偏好设置类型定义
export interface UserPreferences {
  theme: 'light' | 'dark' | 'auto'  // 主题模式
  language: string                  // 界面语言
  autoUpdate: boolean               // 是否自动更新
  downloadPath: string              // 下载路径
  notifications: boolean            // 是否启用通知
}

// Go后端分类数据结构
export interface GoClassify {
  classify_id: string              // 分类ID
  classify_name: string            // 分类名称
  parent_level?: string            // 父级分类（可选）
  icon?: string                    // 分类图标
}
