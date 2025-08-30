import { createI18n } from 'vue-i18n'

// 中文语言包
const zhCN = {
  common: {
    home: '首页',
    popular: '热门软件',
    category: '分类',
    installed: '已安装软件',
    favorites: '收藏室',
    tools: '工具箱',
    settings: '设置',
    help: '帮助与反馈',
    search: '搜索',
    install: '获取',
    uninstall: '卸载',
    uninstalling: '卸载中',
    favorite: '收藏',
    unfavorite: '取消收藏',
    viewAll: '查看全部',
    explore: '立即探索',
    back: '返回',
    clear: '清除',
    submit: '提交',
    cancel: '取消',
    confirm: '确认',
    loading: '加载中...',
    refresh: '刷新',
    refreshing: '刷新中...',
    noData: '暂无数据',
    notFound: '未找到',
    error: '错误',
    success: '成功',
    warning: '警告',
    info: '信息',
    actions: '操作'
  },
  home: {
    title: '发现精彩应用',
    subtitle: '探索为 Windows 11 打造的精选应用和游戏，提升您的生产力和娱乐体验',
    popularApps: '热门软件',
    browseByCategory: '按分类浏览',
    apps: '个应用'
  },
  popular: {
    title: '热门软件',
    subtitle: '发现最受欢迎的应用和游戏',
    noData: '暂无热门软件'
  },
  category: {
    apps: '个应用',
    noData: '该分类下暂无应用'
  },
  installed: {
    title: '已安装软件',
    subtitle: '管理您已安装的应用',
    noData: '您还没有安装任何应用',
    goDiscover: '去发现应用',
    loadingDescription: '正在获取已安装的软件列表，请稍候...'
  },
  favorites: {
    title: '收藏室',
    subtitle: '您收藏的应用',
    noData: '您还没有收藏任何应用',
    goDiscover: '去发现应用'
  },
  tools: {
    title: '工具箱',
    subtitle: '系统维护和优化工具',
    systemClean: '系统清理',
    systemCleanDesc: '清理系统垃圾文件，释放磁盘空间',
    softwareUpdate: '软件更新',
    softwareUpdateDesc: '检查并更新已安装的软件',
    systemOptimize: '系统优化',
    systemOptimizeDesc: '优化系统性能设置',
    backupRestore: '备份还原',
    backupRestoreDesc: '备份和还原系统设置',
    comingSoon: '即将推出'
  },
  settings: {
    title: '设置',
    subtitle: '个性化您的应用商店体验',
    appearance: '外观',
    theme: '主题',
    themeDesc: '选择您喜欢的主题',
    light: '浅色',
    dark: '深色',
    auto: '自动',
    language: '语言',
    languageDesc: '选择界面显示语言',
    chinese: '简体中文',
    english: 'English',
    update: '更新',
    autoUpdate: '自动更新',
    autoUpdateDesc: '自动检查并安装应用更新',
    notifications: '通知',
    pushNotifications: '推送通知',
    pushNotificationsDesc: '接收应用更新和推荐通知',
    download: '下载',
    downloadPath: '下载路径',
    downloadPathDesc: '设置应用下载保存位置',
    browse: '浏览'
  },
  help: {
    title: '帮助与反馈',
    subtitle: '获取帮助或向我们反馈问题',
    faq: '常见问题',
    howToInstall: '如何安装应用？',
    howToInstallDesc: '在应用详情页面点击"获取"按钮，系统会自动下载并安装应用。',
    howToUninstall: '如何卸载应用？',
    howToUninstallDesc: '在已安装软件页面找到要卸载的应用，点击"已安装"按钮即可卸载。',
    howToFavorite: '如何收藏应用？',
    howToFavoriteDesc: '在应用卡片上点击星形图标即可收藏应用，收藏的应用会出现在收藏室中。',
    howToSearch: '如何搜索应用？',
    howToSearchDesc: '在首页或任何页面的搜索框中输入应用名称、开发者或关键词即可搜索。',
    contact: '联系我们',
    emailSupport: '邮箱支持',
    onlineSupport: '在线支持',
    phoneSupport: '电话支持',
    feedback: '问题反馈',
    issueType: '问题类型',
    issueTypeDesc: '请选择问题类型',
    bug: '功能问题',
    ui: '界面问题',
    performance: '性能问题',
    other: '其他',
    issueDescription: '问题描述',
    issueDescriptionDesc: '请详细描述您遇到的问题...',
    contactInfo: '联系方式',
    contactInfoDesc: '邮箱或手机号（可选）',
    submitFeedback: '提交反馈'
  },
  search: {
    placeholder: '搜索应用、游戏、程序和内容...',
    title: '搜索结果',
    results: '找到 {count} 个结果',
    about: '关于 "{query}"',
    noResults: '没有找到相关应用',
    startSearch: '开始搜索',
    startSearchDesc: '在搜索框中输入关键词来查找应用'
  },
  app: {
    name: '应用名称',
    developer: '开发者',
    version: '版本',
    size: '大小',
    category: '分类',
    releaseDate: '发布日期',
    lastUpdated: '最后更新',
    description: '应用描述',
    screenshots: '应用截图',
    noScreenshots: '暂无截图',
    notFound: '应用不存在',
    backToHome: '返回首页',
    publisher: '发布者'
  }
}

// 英文语言包
const enUS = {
  common: {
    home: 'Home',
    popular: 'Popular',
    category: 'Category',
    installed: 'Installed',
    favorites: 'Favorites',
    tools: 'Tools',
    settings: 'Settings',
    help: 'Help & Feedback',
    search: 'Search',
    install: 'Get',
    uninstall: 'Uninstall',
    uninstalling: 'Uninstalling',
    favorite: 'Favorite',
    unfavorite: 'Unfavorite',
    viewAll: 'View All',
    explore: 'Explore Now',
    back: 'Back',
    clear: 'Clear',
    submit: 'Submit',
    cancel: 'Cancel',
    confirm: 'Confirm',
    loading: 'Loading...',
    refresh: 'Refresh',
    refreshing: 'Refreshing...',
    noData: 'No Data',
    notFound: 'Not Found',
    error: 'Error',
    success: 'Success',
    warning: 'Warning',
    info: 'Info',
    actions: 'Actions'
  },
  home: {
    title: 'Discover Amazing Apps',
    subtitle: 'Explore selected applications and games built for Windows 11 to enhance your productivity and entertainment experience',
    popularApps: 'Popular Apps',
    browseByCategory: 'Browse by Category',
    apps: 'apps'
  },
  popular: {
    title: 'Popular Apps',
    subtitle: 'Discover the most popular applications and games',
    noData: 'No popular apps available'
  },
  category: {
    apps: 'apps',
    noData: 'No apps in this category'
  },
  installed: {
    title: 'Installed Apps',
    subtitle: 'Manage your installed applications',
    noData: 'You haven\'t installed any apps yet',
    goDiscover: 'Discover Apps',
    loadingDescription: 'Getting installed software list, please wait...'
  },
  favorites: {
    title: 'Favorites',
    subtitle: 'Your favorite applications',
    noData: 'You haven\'t favorited any apps yet',
    goDiscover: 'Discover Apps'
  },
  tools: {
    title: 'Tools',
    subtitle: 'System maintenance and optimization tools',
    systemClean: 'System Clean',
    systemCleanDesc: 'Clean system junk files and free up disk space',
    softwareUpdate: 'Software Update',
    softwareUpdateDesc: 'Check and update installed software',
    systemOptimize: 'System Optimize',
    systemOptimizeDesc: 'Optimize system performance settings',
    backupRestore: 'Backup & Restore',
    backupRestoreDesc: 'Backup and restore system settings',
    comingSoon: 'Coming Soon'
  },
  settings: {
    title: 'Settings',
    subtitle: 'Personalize your app store experience',
    appearance: 'Appearance',
    theme: 'Theme',
    themeDesc: 'Choose your preferred theme',
    light: 'Light',
    dark: 'Dark',
    auto: 'Auto',
    language: 'Language',
    languageDesc: 'Choose interface display language',
    chinese: '简体中文',
    english: 'English',
    update: 'Update',
    autoUpdate: 'Auto Update',
    autoUpdateDesc: 'Automatically check and install app updates',
    notifications: 'Notifications',
    pushNotifications: 'Push Notifications',
    pushNotificationsDesc: 'Receive app updates and recommendation notifications',
    download: 'Download',
    downloadPath: 'Download Path',
    downloadPathDesc: 'Set app download save location',
    browse: 'Browse'
  },
  help: {
    title: 'Help & Feedback',
    subtitle: 'Get help or send us feedback',
    faq: 'Frequently Asked Questions',
    howToInstall: 'How to install apps?',
    howToInstallDesc: 'Click the "Get" button on the app detail page, and the system will automatically download and install the app.',
    howToUninstall: 'How to uninstall apps?',
    howToUninstallDesc: 'Find the app you want to uninstall on the installed apps page and click the "Installed" button to uninstall.',
    howToFavorite: 'How to favorite apps?',
    howToFavoriteDesc: 'Click the star icon on the app card to favorite the app. Favorited apps will appear in your favorites.',
    howToSearch: 'How to search apps?',
    howToSearchDesc: 'Enter app name, developer, or keywords in the search box on any page to search.',
    contact: 'Contact Us',
    emailSupport: 'Email Support',
    onlineSupport: 'Online Support',
    phoneSupport: 'Phone Support',
    feedback: 'Feedback',
    issueType: 'Issue Type',
    issueTypeDesc: 'Please select issue type',
    bug: 'Bug',
    ui: 'UI Issue',
    performance: 'Performance Issue',
    other: 'Other',
    issueDescription: 'Issue Description',
    issueDescriptionDesc: 'Please describe the issue you encountered in detail...',
    contactInfo: 'Contact Info',
    contactInfoDesc: 'Email or phone number (optional)',
    submitFeedback: 'Submit Feedback'
  },
  search: {
    placeholder: 'Search apps, games, programs and content...',
    title: 'Search Results',
    results: 'Found {count} results',
    about: 'about "{query}"',
    noResults: 'No related apps found',
    startSearch: 'Start Searching',
    startSearchDesc: 'Enter keywords in the search box to find apps'
  },
  app: {
    name: 'App Name',
    developer: 'Developer',
    version: 'Version',
    size: 'Size',
    category: 'Category',
    releaseDate: 'Release Date',
    lastUpdated: 'Last Updated',
    description: 'Description',
    screenshots: 'Screenshots',
    noScreenshots: 'No screenshots available',
    notFound: 'App not found',
    backToHome: 'Back to Home',
    publisher: 'Publisher'
  }
}

// 创建i18n实例
export const i18n = createI18n({
  legacy: false, // 使用Composition API
  locale: 'zh-CN', // 默认语言
  fallbackLocale: 'en-US', // 回退语言
  messages: {
    'zh-CN': zhCN,
    'en-US': enUS
  }
})

// 语言切换函数
export function setLanguage(locale: string) {
  i18n.global.locale.value = locale as 'zh-CN' | 'en-US'
  // 保存到localStorage
  localStorage.setItem('language', locale)
}

// 获取当前语言
export function getCurrentLanguage(): string {
  return i18n.global.locale.value
}

// 初始化语言设置
export function initLanguage() {
  const savedLanguage = localStorage.getItem('language')
  if (savedLanguage && ['zh-CN', 'en-US'].includes(savedLanguage)) {
    i18n.global.locale.value = savedLanguage as 'zh-CN' | 'en-US'
  }
}
