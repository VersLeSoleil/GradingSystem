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
      
      <el-avatar size="36" src="https://element-plus.org/images/element-plus-logo.svg" @click="goToProfile" style="cursor: pointer" />
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
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'

import logoImg from '@/assets/logo.png'
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

function goToProfile() {
  router.push('/profile')
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
</style>
