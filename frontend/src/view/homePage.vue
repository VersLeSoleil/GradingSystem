<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'

const router = useRouter()

// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型库', icon: 'i-ep-house' },
  { index: '/dataset', label: '数据集', icon: 'i-ep-user' },
  { index: '/online', label: '在线分级', icon: 'i-ep-document' },
  { index: '/history', label: '历史记录', icon: 'i-ep-setting' },
]
const activeMenu = ref('/home')

// 标签分类
const categories = [
  '全部',
  '甲状腺分级',
  '乳腺癌分级',
  '肺结节分级',
  '肝脏分级',
  '脑肿瘤分级'
]
const selectedCategory = ref('全部')

// 模型数据
const models = [
  { id: 1, name: '甲状腺AI分级模型', desc: '高准确率甲状腺分级AI模型', category: '甲状腺分级', author: 'AI Lab', avatar: 'https://element-plus.org/images/element-plus-logo.svg' },
  { id: 2, name: '乳腺癌辅助诊断', desc: '乳腺癌影像分级智能模型', category: '乳腺癌分级', author: 'MedAI', avatar: 'https://element-plus.org/images/element-plus-logo.svg' },
  { id: 3, name: '肺结节检测', desc: '肺结节分级与检测一体化模型', category: '肺结节分级', author: 'AI Health', avatar: 'https://element-plus.org/images/element-plus-logo.svg' },
  { id: 4, name: '肝脏肿瘤分级', desc: '肝脏肿瘤分级AI模型', category: '肝脏分级', author: 'AI Lab', avatar: 'https://element-plus.org/images/element-plus-logo.svg' },
  { id: 5, name: '脑肿瘤分级', desc: '脑肿瘤影像分级模型', category: '脑肿瘤分级', author: 'MedAI', avatar: 'https://element-plus.org/images/element-plus-logo.svg' },
]

const filteredModels = computed(() => {
  if (selectedCategory.value === '全部') return models
  return models.filter(m => m.category === selectedCategory.value)
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}
</script>

<template>
  <el-container style="height: 100vh;width: 100%;">
    <!-- 顶部导航栏 -->
    <el-header height="64px" style="background: #409EFF; color: #fff; display: flex; align-items: center; padding: 0 32px;">
      <div style="font-size: 20px; font-weight: bold; margin-right: 32px;">疾病智能分级系统</div>
      <el-input placeholder="搜索..." clearable style="width: 300px; margin-right: 32px;" prefix-icon="el-icon-search" />
      <el-menu :default-active="activeMenu" mode="horizontal" background-color="#409EFF" text-color="#fff" active-text-color="#ffd04b" @select="handleMenuSelect" style="flex: 1; min-width: 500px; border-bottom: none;">
        <el-menu-item v-for="item in navMenus" :key="item.index" :index="item.index">
          <el-icon v-if="item.icon"><component :is="item.icon" /></el-icon>
          {{ item.label }}
        </el-menu-item>
      </el-menu>
      <el-avatar size="40" src="https://element-plus.org/images/element-plus-logo.svg" style="margin-left: 32px; background: #fff;" />
    </el-header>
    <el-main style="background: #fff; padding: 24px 0; height: calc(100vh - 64px);">
      <div style="display: flex; gap: 24px; height: 100%;">
        <!-- 左侧模型卡片区 -->
        <div style="flex: 2; display: flex; flex-direction: column; gap: 20px; overflow-y: auto; margin-left: 32px;">
          <el-card v-for="model in filteredModels" :key="model.id" shadow="hover" style="margin-bottom: 0;">
            <div style="display: flex; align-items: center;">
              <el-avatar :src="model.avatar" size="large" style="margin-right: 16px;" />
              <div style="flex: 1;">
                <div style="font-size: 18px; font-weight: bold;">{{ model.name }}</div>
                <div style="color: #888; margin: 4px 0;">{{ model.desc }}</div>
                <el-tag size="small" type="info">{{ model.category }}</el-tag>
              </div>
              <div style="margin-left: 16px; color: #409EFF;">{{ model.author }}</div>
            </div>
          </el-card>
        </div>
        <!-- 右侧标签分类区 -->
        <div style="flex: 1; min-width: 200px;">
          <div style="font-weight: bold; margin-bottom: 12px;">标签分类</div>
          <div style="display: flex; flex-wrap: wrap; gap: 10px;">
            <el-tag
              v-for="cat in categories"
              :key="cat"
              :type="selectedCategory === cat ? 'primary' : 'info'"
              @click="selectedCategory = cat"
              style="cursor: pointer; user-select: none;"
            >
              {{ cat }}
            </el-tag>
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>
.el-header {
  box-shadow: 0 2px 8px #f0f1f2;
}
</style>