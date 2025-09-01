<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue'

const helpTopics = [
  {
    id: 1,
    title: '如何安装应用？',
    content: '在应用详情页面点击"获取"按钮，系统会自动下载并安装应用。'
  },
  {
    id: 2,
    title: '如何卸载应用？',
    content: '在已安装软件页面找到要卸载的应用，点击"已安装"按钮即可卸载。'
  },
  {
    id: 3,
    title: '如何收藏应用？',
    content: '在应用卡片上点击星形图标即可收藏应用，收藏的应用会出现在收藏室中。'
  },
  {
    id: 4,
    title: '如何搜索应用？',
    content: '在首页或任何页面的搜索框中输入应用名称、开发者或关键词即可搜索。'
  }
]

const contactInfo = {
  email: 'support@winstore.com',
  website: 'https://winstore.com/support',
  phone: 'https://gitee.com/MR-wind/win-store'
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
        <h1 class="page-title">帮助与反馈</h1>
        <p class="page-description">获取帮助或向我们反馈问题</p>
      </div>

      <!-- 帮助内容 -->
      <div class="help-content">
        <!-- 常见问题 -->
        <div class="help-section">
          <h2 class="section-title">常见问题</h2>
          <el-collapse>
            <el-collapse-item 
              v-for="topic in helpTopics" 
              :key="topic.id"
              :title="topic.title"
            >
              <p>{{ topic.content }}</p>
            </el-collapse-item>
          </el-collapse>
        </div>

        <!-- 联系我们 -->
        <div class="help-section">
          <h2 class="section-title">联系我们</h2>
          <div class="contact-info">
            <div class="contact-item">
              <h3>邮箱支持</h3>
              <p>{{ contactInfo.email }}</p>
            </div>
            <div class="contact-item">
              <h3>在线支持</h3>
              <p>{{ contactInfo.website }}</p>
            </div>
            <div class="contact-item">
              <h3>仓库地址</h3>
              <p>{{ contactInfo.phone }}</p>
            </div>
          </div>
        </div>

        <!-- 反馈表单 -->
        <div class="help-section">
          <h2 class="section-title">问题反馈</h2>
          <el-form label-width="100px">
            <el-form-item label="问题类型">
              <el-select placeholder="请选择问题类型" style="width: 100%">
                <el-option label="功能问题" value="bug" />
                <el-option label="界面问题" value="ui" />
                <el-option label="性能问题" value="performance" />
                <el-option label="其他" value="other" />
              </el-select>
            </el-form-item>
            <el-form-item label="问题描述">
              <el-input 
                type="textarea" 
                :rows="4"
                placeholder="请详细描述您遇到的问题..."
              />
            </el-form-item>
            <el-form-item label="联系方式">
              <el-input placeholder="邮箱或手机号（可选）" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary">提交反馈</el-button>
            </el-form-item>
          </el-form>
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

.help-content {
  max-width: 800px;
}

.help-section {
  background-color: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #333;
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e0e0e0;
}

.contact-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.contact-item {
  padding: 16px;
  background-color: #f8f9fa;
  border-radius: 8px;
  text-align: center;
}

.contact-item h3 {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.contact-item p {
  font-size: 14px;
  color: #666;
  margin: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .contact-info {
    grid-template-columns: 1fr;
  }
}
</style>
