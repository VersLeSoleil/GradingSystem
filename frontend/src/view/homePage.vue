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
      
      <el-dropdown>
      <el-button type="primary" @click="handleLoginClick">
        未登录<el-icon class="el-icon--right"><arrow-down /></el-icon>
      </el-button>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>编辑资料</el-dropdown-item>
          <el-dropdown-item>退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
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
            v-for="post in filteredPosts"
            :key="post.post_id"
            shadow="hover"
            class="animated-card"
            @click="goToPostDetail(post)"
            style="cursor: pointer; width: 300px; height: 200px; display: flex; align-items: center; justify-content: center;"
          >
            <div>
              <div class="card-title">{{ post.title }}</div>
              <div class="card-desc">{{ post.introduction }}</div>
              <el-tag size="small" type="success">{{ post.type_name }}</el-tag>
              <div class="card-foot" style="display: flex;flex-direction: row;">
                <div class="likes-container" style="margin-top: 10px;">
                  <button @click="toggleLike(post)"  class="icon-button">
                    <el-icon><Star /></el-icon>
                  </button>
                  <span class="likes-count">{{ post.likes }}</span>
                </div>
                <div class="card-author" style="margin-left:130px;">{{ post.user_name }}</div>
              </div>
              
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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'
import { Star } from '@element-plus/icons-vue';
import logoImg from '@/assets/logo.png'
import UserInfo from './components/userInfo.vue'

const router = useRouter()
const activeMenu = ref('/home')
const selectedCategory = ref('全部')
let posts = ref([])
const loading = ref(true)
const error = ref(null)
const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线分级' },
]

const categories = ['全部', '甲状腺分级', '乳腺癌分级', '肺结节分级', '肝脏分级', '脑肿瘤分级']

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:8888/getPosts', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include' // 如果你涉及到 cookie 认证
    })

    if (!response.ok) {
      throw new Error('无法获取模型列表')
    }
    const data = await response.json()
    posts.value = data
    console.log('获取到的模型列表：', posts.value)
  } catch (err) {
    console.error('获取模型失败：', err)
    error.value = err.message
  } finally {
    loading.value = false
  }
})

const filteredPosts = computed(() => {
  if (selectedCategory.value === '全部') return posts.value
  return posts.value.filter((m) => m.type_name === selectedCategory.value)
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}

function handleLoginClick() {
  
}
async function toggleLike(post) {
  try{
    console.log('Toggling like for post:', post.post_id)
    const endpoint = `http://localhost:8888/updatePost`;
    const response = await fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        post_id: post.post_id,
        likes: post.isLiked ? post.likes - 1 : post.likes + 1,
        isLiked: !post.isLiked
      }),
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('更新点赞状态失败');
    }
  }catch (error) {
    console.error('Error toggling like:', error)
    alert('操作失败，请稍后再试')
    return
  }
  post.isLiked = !post.isLiked
  post.likes += post.isLiked ? 1 : -1
  console.log('Post liked:', post)
}
async function goToPostDetail(post) {
  try {
    console.log('id:', post.post_id)
    const response = await fetch(`http://localhost:8888/getPostByID?post_id=${parseInt(post.post_id)}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include' // 如果你需要发送 cookie
    })

    if (!response.ok) {
      throw new Error('获取帖子详情失败')
    }

    const data = await response.json()
    console.log('获取到的模型信息：', data)

    // 假设你需要把 data 传给下一个页面，可以通过 query 或 store
    router.push({
      path: `/model/${post.post_id}`,
      query: {
        postid: data.post_id,
        type_name: data.type_name,
        content: data.content,
        title: data.title,
        introduction: data.introduction,
        user_name: data.user_name,
        likes: data.likes,
      }
    })
    
  } catch (error) {
    console.error('请求失败:', error)
    alert('获取模型详情失败，请稍后再试')
  }
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
  font-size: 18px;
  margin-bottom: 6px;
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

.likes-container {
  display: flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
}
.likes-count {
  font-size: 14px;
  color: #666;
}
.icon-button {
  border: none;          /* 移除边框 */
  background: none;      /* 移除背景 */
  padding: 0;            /* 移除内边距 */
  cursor: pointer;       /* 保持手型指针 */
  outline: none;         /* 移除聚焦时的轮廓线 */
}

/* 可选：悬停效果 */
.icon-button:hover {
  opacity: 0.8;
}
</style>
