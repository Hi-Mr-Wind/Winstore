<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue'

const tools = [
  {
    id: 1,
    name: '系统清理',
    description: '清理系统垃圾文件，释放磁盘空间',
    icon: '🧹',
    status: 'available'
  },
  {
    id: 2,
    name: '软件更新',
    description: '检查并更新已安装的软件',
    icon: '🔄',
    status: 'available'
  },
  {
    id: 3,
    name: '系统优化',
    description: '优化系统性能设置',
    icon: '⚡',
    status: 'coming-soon'
  },
  {
    id: 4,
    name: '备份还原',
    description: '备份和还原系统设置',
    icon: '💾',
    status: 'coming-soon'
  }
]

const handleToolClick = (tool: any) => {
  if (tool.status === 'available') {
    console.log(`启动工具: ${tool.name}`)
  } else {
    console.log('该功能即将推出')
  }
}
</script>

<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <Sidebar />

    <!-- 主内容区域 -->
    <div class="main-content">
      <!-- 页面标题 -->
      <div class="page-header">
        <h1 class="page-title">工具箱</h1>
        <p class="page-description">系统维护和优化工具</p>
      </div>

      <!-- 工具网格 -->
      <div class="tools-grid">
        <div 
          v-for="tool in tools" 
          :key="tool.id" 
          class="tool-card"
          :class="{ 'coming-soon': tool.status === 'coming-soon' }"
          @click="handleToolClick(tool)"
        >
          <div class="tool-icon">{{ tool.icon }}</div>
          <div class="tool-info">
            <h3 class="tool-name">{{ tool.name }}</h3>
            <p class="tool-description">{{ tool.description }}</p>
            <el-tag 
              v-if="tool.status === 'coming-soon'" 
              type="info" 
              size="small"
            >
              即将推出
            </el-tag>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background-color: #f5f5f5;
}

.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

.page-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.page-description {
  font-size: 16px;
  color: #666;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.tool-card {
  background-color: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 16px;
}

.tool-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.tool-card.coming-soon {
  opacity: 0.6;
  cursor: not-allowed;
}

.tool-card.coming-soon:hover {
  transform: none;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  font-size: 32px;
  flex-shrink: 0;
}

.tool-info {
  flex: 1;
}

.tool-name {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.tool-description {
  font-size: 14px;
  color: #666;
  margin-bottom: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .tools-grid {
    grid-template-columns: 1fr;
  }
}
</style>
