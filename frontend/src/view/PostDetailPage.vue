
<script setup>
import { ref, computed,onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput, ElTree, ElButton } from 'element-plus'
import { ArrowDown } from '@element-plus/icons-vue'
import logoImg from '@/assets/logo.png'
import bgImage from '@/assets/background.png';
import MarkdownIt from 'markdown-it'
import 'github-markdown-css/github-markdown-light.css'
import { Star } from '@element-plus/icons-vue';
import UserInfo from '@/view/components/userInfo.vue'
import { useUserStore } from '@/store/user'
const userStore = useUserStore()
const username = userStore.userInfo.UserName
const userInfoRef = ref(null)
const router = useRouter()
const route = useRoute()
const md = new MarkdownIt()
const renderedContent = ref('')
// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线分级' },
]
const activeMenu = ref('/home')

const postId = route.query.postid
const type_name = route.query.type_name
const content = route.query.content
const title = route.query.title
const introduction = route.query.introduction
const userName = route.query.user_name
const likes = route.query.likes

// 点赞状态和数量
const liked = ref(false)
const likeCount = ref(Number(route.query.likes) || 0)

const handleCommand = async (command) => {
  switch(command) {
    case 'editProfile':
      showUserInfo()
      break
    case 'logout':
      await handleLogout()
      break
  }
}

function showUserInfo() {
  console.log('userInfoRef:', userInfoRef.value)
  userInfoRef.value.openDialog()
}

const handleLogout = () => {
  // 清除本地存储的 token 和用户信息
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('user')
  // 如果有 Pinia/Vuex 用户信息，也要清空
  userStore.$reset && userStore.$reset()
  // 跳转到登录页
  router.push('/login')
}

// 点赞接口
async function toggleLike() {
  // 防止重复点击
  if (liked.value) {
    likeCount.value--
  } else {
    likeCount.value++
  }
  liked.value = !liked.value
  // 可选：后端同步
  try {
    const response = await fetch('http://localhost:8888/updatePost', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        post_id: postId,
        likes: likeCount.value,
        isLiked: liked.value
      }),
      credentials: 'include',
    })
    if (!response.ok) {
      throw new Error('点赞失败')
    }
  } catch (err) {
    // 回滚本地状态
    liked.value = !liked.value
    likeCount.value += liked.value ? 1 : -1
    alert('点赞失败，请稍后再试')
  }
}

onMounted(() => {
  console.log('content:', route.query.content)
  const rawMarkdown = content || `# Hello Markdown!`
  renderedContent.value = md.render(rawMarkdown)
})

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
      
      <el-dropdown @command="handleCommand">
      <el-button type="primary" @click="handleLoginClick" class="user-info-button">
        {{ username }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
      </el-button>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item command="editProfile">编辑资料</el-dropdown-item>
          <el-dropdown-item command="logout">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

    </el-header>
    <el-main :style="{
          padding: '32px 0 0 0',
          backgroundImage: `url(${bgImage})`,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
          backgroundRepeat: 'no-repeat'
        }">
      <!-- 1. 模型简介 -->
      <el-card style="width: 80%;margin: 0 auto;">
        <div class="card-header">
          <el-page-header @back="goBack" :title="首页" class="custom-page-header">
            <template #content>
              <div class="header-content-wrapper">
                <!-- 标题行 -->
                <div class="header-top-row">
                  <span class="custom-header-title">{{ route.query.title }}</span>
                  <el-tag class="type-tag">{{ route.query.type_name }}</el-tag>
                </div>
                
                <!-- 描述行 -->
                <div class="header-middle-row">
                  <div class="custom-header-intro">{{ route.query.introduction }}</div>
                </div>
                
                <!-- 底部信息行 -->
                <div class="header-bottom-row">
                  <div class="likes-container">
                    <button @click="toggleLike"  class="icon-button">
                      <el-icon :style="liked ? { color: '#f56c6c' } : {}"><Star /></el-icon>
                    </button>
                    <span class="likes-count">{{ likeCount }}</span>
                  </div>
                  <div class="card-author">{{ route.query.user_name }}</div>
                </div>
              </div>
            </template>
          </el-page-header>
        </div>
        <div class="markdown-body" style="width: 90%; margin: 0 auto;" v-html="renderedContent"></div>
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
  <UserInfo ref="userInfoRef" />
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
  margin: 3px auto 32px auto;
  width: 80%;
  border-radius: 8px;
}
.custom-markdown {
  font-family: 'Segoe UI', sans-serif;
  line-height: 1.8;
  width: 80%;
  padding: 1em;
  background: #f9f9f9;
  color: #333;
  border-radius: 8px;
  border: 1px solid #ddd;
}

.custom-markdown h1, .custom-markdown h2, .custom-markdown h3 {
  color: #2c3e50;
  margin-top: 1em;
  margin-bottom: 0.5em;
}

.custom-markdown code {
  background-color: #eee;
  padding: 0.2em 0.4em;
  border-radius: 4px;
  font-family: monospace;
}

.custom-markdown pre code {
  background-color: #272822;
  color: #f8f8f2;
  display: block;
  padding: 1em;
  overflow-x: auto;
  border-radius: 6px;
}


/* 整体卡片头部样式 */
.card-header {
  background: white;
  border-radius: 8px;
  padding: 16px 20px;
  margin-bottom: 15px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
}

/* 页面头部容器 */
.custom-page-header {
  padding: 0;
}

/* 内容容器 */
.header-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

/* 顶部标题行 */
.header-top-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.custom-header-title {
  font-size: 30px;
  font-weight: 600;
  color: #333;
}

.type-tag {
  margin-left: 8px;
  height: 24px;
  line-height: 22px;
}

/* 描述文本 */
.custom-header-intro {
  color: #666;
  font-size: 14px;
  line-height: 1.5;
  margin-top: 4px;
}

/* 底部信息行 */
.header-bottom-row {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-top: 5px;
}

.likes-container {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #f56c6c;
}

.like-icon {
  font-size: 16px;
}
.icon-button {
  color: #f56c6c; /* 设置图标颜色 */
  font-size: 16px; /* 设置图标大小 */
  border: none;          /* 移除边框 */
  background: none;      /* 移除背景 */
  padding: 0;            /* 移除内边距 */
  cursor: pointer;       /* 保持手型指针 */
  outline: none;         /* 移除聚焦时的轮廓线 */
}
.likes-count {
  font-size: 14px;
}

.card-author {
  font-size: 14px;
  color: #888;
  display: flex;
  align-items: center;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .header-top-row {
    flex-wrap: wrap;
  }
  
  .header-bottom-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
}

.user-info-button{
  color:#cf5454;
  font-size:15px;
  border: none;          /* 移除边框 */
  background: none;      /* 移除背景 */
  padding: 0;            /* 移除内边距 */
  cursor: pointer;       /* 保持手型指针 */
  outline: none;         /* 移除聚焦时的轮廓线 */
}
</style>