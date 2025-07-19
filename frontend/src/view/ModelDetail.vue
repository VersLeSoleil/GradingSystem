<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput, ElTree, ElButton } from 'element-plus'

const router = useRouter()
const route = useRoute()

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

// 模型数据（id改为英文唯一名）
const models = [
  { id: 'GoogleNet', name: 'GoogleNet', desc: '高准确率甲状腺分级AI模型', category: '甲状腺分级', author: 'Google' },
  { id: 'breast_cancer', name: '乳腺癌辅助诊断', desc: '乳腺癌影像分级智能模型', category: '乳腺癌分级', author: 'MedAI' },
  { id: 'lung_nodule', name: '肺结节检测', desc: '肺结节分级与检测一体化模型', category: '肺结节分级', author: 'AI Health' },
  { id: 'liver_tumor', name: '肝脏肿瘤分级', desc: '肝脏肿瘤分级AI模型', category: '肝脏分级', author: 'AI Lab' },
  { id: 'brain_tumor', name: '脑肿瘤分级', desc: '脑肿瘤影像分级模型', category: '脑肿瘤分级', author: 'MedAI'},
]

const modelId = computed(() => route.params.id)
const model = computed(() => models.find(m => m.id === modelId.value))

const filteredModels = computed(() => {
  if (selectedCategory.value === '全部') return models
  return models.filter(m => m.category === selectedCategory.value)
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}

function goBack() {
  router.back()
}

// 模拟github文件树
const fileTree = ref([
  {
    label: 'model/',
    children: [
      { label: 'README.md' },
      { label: 'main.py' },
      { label: 'requirements.txt' },
      {
        label: 'src/',
        children: [
          { label: 'model.py' },
          { label: 'utils.py' },
        ]
      },
      { label: 'data/' },
    ]
  }
])
const treeProps = { children: 'children', label: 'label' }

// 讨论区模拟
const comments = ref([
  { user: 'Alice', time: '2025-06-28 10:00', content: '这个模型效果很棒！' },
  { user: 'Bob', time: '2025-06-28 11:20', content: '希望能开源更多数据集。' },
])
const newComment = ref('')
function addComment() {
  if (newComment.value.trim()) {
    comments.value.push({
      user: '你',
      time: new Date().toLocaleString(),
      content: newComment.value
    })
    newComment.value = ''
  }
}

const activeIndex = ref('readme')
const readmeContent = ref(`# GoogleNet\n\n本模型基于GoogleNet架构，适用于甲状腺分级任务。\n\n- 高准确率\n- 支持多种医学影像格式\n- 详细文档与示例\n\n## 快速开始\n\n1. 安装依赖\n2. 运行main.py\n3. 查看结果\n`)
function handleSelect(key) {
  activeIndex.value = key
}
</script>

<template>
  <el-container style="min-height: 100vh; background: #f6f8fa; flex-direction: column;">
    <!-- 顶部导航栏 -->
    <el-header height="64px" style="background: #409EFF; color: #fff; display: flex; align-items: center; padding: 0 32px;">
      <div style="font-size: 20px; font-weight: bold; margin-right: 32px;">疾病影像智能分级系统</div>
      <el-input placeholder="搜索..." clearable style="width: 300px; margin-right: 32px;" prefix-icon="el-icon-search" />
      <el-menu :default-active="activeMenu" mode="horizontal" background-color="#409EFF" text-color="#fff" active-text-color="#ffd04b" @select="handleMenuSelect" style="flex: 1; min-width: 500px; border-bottom: none;">
        <el-menu-item v-for="item in navMenus" :key="item.index" :index="item.index">
          <el-icon v-if="item.icon"><component :is="item.icon" /></el-icon>
          {{ item.label }}
        </el-menu-item>
      </el-menu>
      <el-avatar size="40" src="https://element-plus.org/images/element-plus-logo.svg" style="margin-left: 32px; background: #fff;" />
    </el-header>
    <el-main style="padding: 32px 0 0 0; background: #f6f8fa;">
      <!-- 1. 模型简介 -->
      <el-card>
        <div v-if="model" class="model-intro-card">
          <div style="display: flex; align-items: center;">
            <div style="flex: 1;">
              <div style="font-size: 50px; font-weight: bold;">{{ model.name }}</div>
              <div style="color: #888; margin: 4px 0;">{{ model.desc }}</div>
              <el-tag size="small" type="info">{{ model.category }}</el-tag>
            </div>
            <div style="margin-left: 16px; color: #409EFF;">{{ model.author }}</div>
          </div>
        </div>
        <div v-else class="model-intro-card" style="text-align:center;">未找到该模型</div>
        <el-divider />
        <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          @select="handleSelect"
          style="margin-bottom: 24px; border-radius: 8px; background: #fff; box-shadow: 0 2px 8px #f0f1f2;"
        >
          <el-menu-item index="readme">README</el-menu-item>
          <el-menu-item index="tree">File</el-menu-item>
        </el-menu>
        <div class="model-tree-card">
          <div v-if="activeIndex === 'readme'" style="font-size: 16px; font-family: monospace; white-space: pre-wrap; background: #f8f8f8; padding: 24px; border-radius: 8px; min-height: 200px;">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 12px;">README.md</div>
            <div>{{ readmeContent }}</div>
          </div>
          <div v-else>
            
            <el-tree
              :data="fileTree"
              :props="treeProps"
              accordion
              highlight-current
              default-expand-all
              style="background: #f6f8fa;"
            />
          </div>
        </div>
      </el-card>
      <!-- 3. 讨论区 -->
      <el-card class="model-discuss-card">
        <div style="font-size: 18px; font-weight: bold; margin-bottom: 12px;">讨论区</div>
        <div v-for="(item, idx) in comments" :key="idx" style="margin-bottom: 16px;">
          <div style="font-weight: bold;">{{ item.user }}</div>
          <div style="color: #888; font-size: 13px; margin-bottom: 4px;">{{ item.time }}</div>
          <div>{{ item.content }}</div>
        </div>
        <el-input v-model="newComment" type="textarea" rows="2" placeholder="发表你的看法..." style="margin-bottom: 8px;" />
        <el-button type="primary" @click="addComment">发表评论</el-button>
      </el-card>
    </el-main>
  </el-container>
</template>

<style scoped>
.el-header {
  box-shadow: 0 2px 8px #f0f1f2;
}
.model-intro-card {
  margin: 0 auto 24px auto;
  max-width: 900px;
  border-radius: 8px;
}
.model-tree-card {
  margin: 0 auto 24px auto;
  max-width: 900px;
  border-radius: 8px;
  background: #f6f8fa;
}
.model-discuss-card {
  margin: 0 auto 32px auto;
  max-width: 900px;
  border-radius: 8px;
}
.readme-content {
  margin: 0 auto 24px auto;
  max-width: 900px;
  padding: 16px;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}
</style>