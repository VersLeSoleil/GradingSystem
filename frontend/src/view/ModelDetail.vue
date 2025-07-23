<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput, ElTree, ElButton } from 'element-plus'
import logoImg from '@/assets/logo.png'
const router = useRouter()
const route = useRoute()

// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线分级' },
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


function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}

function goBack() {
  router.back()
}

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
    <el-main style="padding: 32px 0 0 0; background: #f6f8fa;">
      <!-- 1. 模型简介 -->
      <el-card>
        
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