<template>
  <el-container style="height: 100vh; width: 100%;">
    <!-- 顶部导航栏 -->
    <el-header
      height="64px"
      style="display: flex; align-items: center; justify-content: space-between; padding: 0 40px; background: #fff; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);"
    >
      <div style="display: flex; align-items: center; flex-shrink: 0;">
        <img :src="logoImg" alt="logo" style="width: 32px; height: 32px; margin-right: 12px" />
        <div style="font-weight: 600; font-size: 18px">
          分级喵 <span style="color: #999; font-size: 14px; margin-left: 4px">Image Grading System</span>
        </div>
      </div>
      <div
    style="
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
      display: flex;
      justify-content: center;
      white-space: nowrap;
    "
  >
    <el-menu
      :default-active="activeMenu"
      mode="horizontal"
      :ellipsis="false"
      background-color="transparent"
      text-color="#333"
      active-text-color="#CD5C5C"
      @select="handleMenuSelect"
      style="border-bottom: none"
    >
      <el-menu-item v-for="item in navMenus" :key="item.index" :index="item.index">
        {{ item.label }}
      </el-menu-item>
    </el-menu>
  </div>
      
      <el-avatar size="large" src="https://element-plus.org/images/element-plus-logo.svg" @click="showUserInfo" style="cursor: pointer" />
    </el-header>

    <!-- 搜索框区域 -->
    <div class="search-section">
      <transition name="fade-slide">
        <div class="search-content">
          <h2 class="title">模型广场 - 为您推荐智能分级解决方案</h2>
          <p class="subtitle">发现和分享优质大模型，让诊断更高效</p>
          <el-input placeholder="搜索公共模型..." prefix-icon="el-icon-search" clearable class="search-input" />
        </div>
      </transition>
    </div>

    <!-- 主体区域 -->
    <el-main style="padding: 32px; background: #fff">
      <div class="main-content">
        <!-- 左侧筛选栏 -->
        <div class="sidebar">
          <div class="sidebar-title">筛选条件</div>
          <el-menu :default-active="selectedCategory" @select="selectedCategory = $event" class="el-menu-vertical-demo">
            <el-menu-item v-for="cat in categories" :key="cat" :index="cat">
              <el-icon><i class="el-icon-folder" /></el-icon>
              {{ cat }}
            </el-menu-item>
          </el-menu>
        </div>

        <!-- 卡片区域 -->
        <transition-group name="list-fade" tag="div" class="card-grid">
          <el-card
            v-for="model in filteredModels"
            :key="model.id"
            shadow="hover"
            class="animated-card"
            @click="goToModelDetail(model)"
          >
            <div>
              <div class="card-title">{{ model.name }}</div>
              <div class="card-desc">{{ model.desc }}</div>
              <el-tag size="small" type="success">{{ model.category }}</el-tag>
              <div class="card-author">{{ model.author }}</div>
            </div>
          </el-card>
        </transition-group>
      </div>
    </el-main>
  </el-container>

  <el-tooltip content="我要发帖" placement="top">
    <el-button
      class="fab"
      type="primary"
      circle
      @click="dialogVisible = true"
    >
      <span style="font-size: 20px">+</span>
    </el-button>
  </el-tooltip>

  <!-- 发帖弹窗 -->
  <el-dialog
    v-model="dialogVisible"
    title="我要分享"
    width="700px"
    height="700px"
    :before-close="handleClose"
  >
    <el-form :model="postForm" label-width="80px">
      <el-form-item label="标题">
        <el-input v-model="postForm.title" placeholder="标题" />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="postForm.describe" placeholder="描述" />
      </el-form-item>
      <el-form-item label="标签">
      <el-select
        v-model="postForm.tags"
        multiple
        placeholder="请选择标签,可多选"
        style="width: 100%"
      >
        <el-option
          v-for="tag in tagOptions"
          :key="tag.value"
          :label="tag.label"
          :value="tag.value"
        />
      </el-select>
    </el-form-item>

      <el-form-item label="内容">
        <el-input
          type="textarea"
          v-model="postForm.content"
          rows="15"
          placeholder="请输入内容"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" @click="submitPost">发布</el-button>
    </template>
  </el-dialog>
  <UserInfo ref="userInfoRef" />
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'

import logoImg from '@/assets/logo.png'
import UserInfo from './components/userInfo.vue'

const router = useRouter()
const activeMenu = ref('/home')
const selectedCategory = ref('全部')

const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '我的模型' },
]

const categories = ['全部', '甲状腺分级', '乳腺癌分级', '肺结节分级', '肝脏分级', '脑肿瘤分级']

const models = [
  { id: 'GoogleNet', name: 'GoogleNet', desc: '高准确率甲状腺分级AI模型', category: '甲状腺分级', author: 'Google' },
  { id: 'breast_cancer', name: '乳腺癌辅助诊断', desc: '乳腺癌影像分级智能模型', category: '乳腺癌分级', author: 'MedAI' },
  { id: 'lung_nodule', name: '肺结节检测', desc: '肺结节分级与检测一体化模型', category: '肺结节分级', author: 'AI Health' },
  { id: 'liver_tumor', name: '肝脏肿瘤分级', desc: '肝脏肿瘤分级AI模型', category: '肝脏分级', author: 'AI Lab' },
  { id: 'brain_tumor', name: '脑肿瘤分级', desc: '脑肿瘤影像分级模型', category: '脑肿瘤分级', author: 'MedAI' },
]

const filteredModels = computed(() => {
  if (selectedCategory.value === '全部') return models
  return models.filter((m) => m.category === selectedCategory.value)
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}

function goToModelDetail(model) {
  router.push({ path: `/model/${model.id}` })
}

const userInfoRef = ref(null)
function showUserInfo() {
  console.log('userInfoRef:', userInfoRef.value)
  userInfoRef.value.openDialog()
}

// 发帖弹窗逻辑
const dialogVisible = ref(false)
const tagOptions = [
  { label: '深度学习', value: 'deep_learning' },
  { label: '机器学习', value: 'machine_learning' },
  { label: '乳腺癌', value: 'breast_cancer' },
  { label: '肺结节', value: 'lung_nodule' },
  { label: '肝脏', value: 'liver_tumor' },
  { label: '脑肿瘤', value: 'brain_tumor' }
]
const postForm = ref({
  title: '',
  describe: '',
  tags: [],
  content: ''
})
function handleClose() {
  dialogVisible.value = false
}

function submitPost() {
  console.log('用户发帖内容：', postForm.value)
  if (!postForm.value.title || !postForm.value.content) {
    ElMessage.error('请填写标题和内容')
    return
  }
  const postData = {
    title: postForm.value.title,
    description: postForm.value.describe,
    tags: postForm.value.tags,
    content: postForm.value.content,
    // 可以添加其他需要的字段，如用户ID、时间戳等
    createdAt: new Date().toISOString()
  }
  axios.post('https://your-api-endpoint.com/posts', postData, {
    headers: {
      'Content-Type': 'application/json',
      // 如果需要认证，可以添加token
      // 'Authorization': `Bearer ${yourToken}`
    }
  })
  .then(response => {
    ElMessage.success('发布成功！')
    console.log('发布成功:', response.data)
    // 清空表单
    postForm.value = { 
      title: '', 
      describe: '', 
      tags: [], 
      content: '' 
    }
    dialogVisible.value = false
  })
  .catch(error => {
    ElMessage.error('发布失败，请重试')
    console.error('发布失败:', error)
  })
}

</script>

<style scoped>
.search-section {
  background: #f9f9f9;
  padding: 40px 0;
  text-align: center;
  animation: fadeInDown 0.8s ease;
}
.search-content .title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 12px;
}
.search-content .subtitle {
  color: #888;
  font-size: 14px;
  margin-bottom: 20px;
}
.search-input {
  width: 480px;
}
.main-content {
  display: flex;
  gap: 32px;
}
.sidebar {
  width: 220px;
}
.sidebar-title {
  font-weight: bold;
  margin-bottom: 16px;
}
.card-grid {
  flex: 1;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}
.animated-card {
  transition: all 0.3s ease;
  border-radius: 12px;
  animation: fadeInUp 0.5s ease;
}
.animated-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.08);
}
.card-title {
  font-weight: bold;
  font-size: 16px;
}
.card-desc {
  color: #888;
  font-size: 13px;
  margin: 4px 0;
}
.card-author {
  text-align: right;
  margin-top: 12px;
  color: #409eff;
  font-size: 13px;
}
.fade-slide-enter-active {
  transition: all 0.5s ease;
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.list-fade-enter-active {
  transition: all 0.6s ease;
}
.list-fade-enter-from {
  opacity: 0;
  transform: translateY(20px);
}
.header-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
  padding: 0 32px;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.header-logo {
  font-size: 20px;
  font-weight: bold;
  color: #409EFF;
  white-space: nowrap;
}

.nav-center {
  flex: 1;
  display: flex;
  justify-content: center;
  overflow: hidden;
}

.nav-menu {
  max-width: 600px;
  flex-shrink: 0;
  display: flex;
  justify-content: center;
}
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 弹窗美化 */
.fab {
  position: fixed;
  bottom: 40px;
  right: 40px;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 999;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
}
</style>
