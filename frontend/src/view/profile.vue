<script setup>
import { ref, onMounted } from 'vue'
import { ElAvatar, ElCard, ElButton, ElDivider, ElInput, ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
const router = useRouter()

const user = ref({
  name: '',
  email: '',
  avatar: '',
  intro: '',
})
const navMenus = [
  { index: '/home', label: '模型库', icon: 'i-ep-house' },
  { index: '/dataset', label: '数据集', icon: 'i-ep-user' },
  { index: '/online', label: '在线分级', icon: 'i-ep-document' },
  { index: '/history', label: '历史记录', icon: 'i-ep-setting' },
]
const activeMenu = ref('')

const editing = ref(false)
const editIntro = ref('')

async function fetchUserProfile() {
  // 模拟异步请求，后续可替换为真实API
  // const res = await axios.get('/api/user/profile')
  // Object.assign(user.value, res.data)
  await new Promise(r => setTimeout(r, 500))
  Object.assign(user.value, {
    name: '张三',
    email: 'zhangsan@example.com',
    avatar: 'https://element-plus.org/images/element-plus-logo.svg',
    intro: 'AI医疗模型开发者，专注于智能分级与辅助诊断。',
  })
}

onMounted(() => {
  fetchUserProfile()
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}
function startEdit() {
  editing.value = true
  editIntro.value = user.value.intro
}
function saveIntro() {
  user.value.intro = editIntro.value
  editing.value = false
  ElMessage.success('简介已保存（仅本地模拟）')
}
</script>

<template>
    <el-container style="height: 100vh;width: 100%;">
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
      <el-avatar size="40" src="https://element-plus.org/images/element-plus-logo.svg" style="margin-left: 32px; background: #fff; cursor: pointer;" @click="goToProfile" />
    </el-header>
    <el-main style="background: #fff; padding: 24px 0; height: calc(100vh - 64px);">
      <el-container class="profile-github-layout">
    <div class="profile-sidebar">
      <el-card shadow="never" class="profile-card">
        <el-avatar :src="user.avatar" size="240px" class="profile-avatar" />
        <div class="profile-name">{{ user.name }}</div>
        <div class="profile-email">{{ user.email }}</div>
        <el-divider />
        <div class="profile-intro-label">个人简介</div>
        <div v-if="!editing" class="profile-intro">{{ user.intro }}</div>
        <el-input v-else v-model="editIntro" type="textarea" rows="3" />
        <div class="profile-btns">
          <el-button v-if="!editing" type="primary" @click="startEdit" size="small">编辑简介</el-button>
          
          <el-button v-else type="success" @click="saveIntro" size="small">保存</el-button>
          <el-button v-else @click="editing = false" size="small">取消</el-button>
          <el-button v-if="!editing" type="primary" @click="startEdit" size="small">退出登录</el-button>
        </div>
      </el-card>
    </div>
    <div class="profile-main">
      <el-card shadow="never" style="height: 600px; ">
        <div style="font-size: 18px; font-weight: bold; margin-bottom: 12px;">个人动态</div>
        <div style="color: #888;">这里可以展示用户相关的动态、操作记录等内容。</div>
      </el-card>
    </div>
  </el-container>
    </el-main>
  </el-container>
  
</template>

<style scoped>
.profile-github-layout {
  min-height: 100vh;
  background: #f6f8fa;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: flex-start;
  padding: 32px 0;
}
.profile-sidebar {
  width: 340px;
  height:90%;
  margin-right: 32px;
  margin-left: 32px;
}
.profile-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 32px 24px 24px 24px;
  height: 600px;
  border-radius: 8px;
  background: #fff;
}
.profile-avatar {
  margin-bottom: 16px;
  border: 2px solid #e1e4e8;
}
.profile-name {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 4px;
  margin-top: 8px;
}
.profile-email {
  color: #586069;
  margin-bottom: 16px;
}
.profile-intro-label {
  font-weight: bold;
  margin-bottom: 6px;
  width: 100%;
}
.profile-intro {
  min-height: 40px;
  width: 100%;
  margin-bottom: 8px;
  color: #24292f;
}
.profile-btns {
  margin-top: 12px;
  width: 100%;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}
.profile-main {
  flex: 1;
  min-width: 0;
  height:600px;
}
</style>
