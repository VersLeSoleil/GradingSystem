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

    <!-- 搜索框区域 -->
    <div class="search-section">
      <transition name="fade-slide">
        <div class="search-content">
          <h2 class="title">模型广场 - 为您推荐智能分级解决方案</h2>
          <p class="subtitle">发现和分享优质大模型，让诊断更高效</p>
          <el-input v-model="searchText" placeholder="搜索分级方案..." prefix-icon="el-icon-search" clearable class="search-input" />
        </div>
      </transition>
    </div>

    <!-- 主体区域 -->
    <el-main style="padding: 32px; background: #fff">
      <div class="main-content">
        <!-- 左侧筛选栏 -->
        <div class="sidebar">
          <div class="sidebar-title">筛选条件</div>
          <el-menu :default-active="selectedCategory.value" @select="handleCategorySelect" class="el-menu-vertical-demo">
            <el-menu-item v-for="cat in categories" :key="cat" :index="cat">
              <el-icon><i class="el-icon-folder" /></el-icon>
              {{ cat }}
            </el-menu-item>
          </el-menu>
        </div>

        <!-- 卡片区域骨架屏 -->
        <el-skeleton :rows="7" animated v-if="changing" style="margin-bottom: 24px;" />
        <!-- 卡片区域 -->
        <transition-group name="list-fade" tag="div" class="card-grid" v-else>
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
              <el-tag size="small" type="warning" style="margin-left:5px">{{ post.model_types }}</el-tag>
              <div class="card-foot" style="display: flex;flex-direction: row;">
                <div class="likes-container" style="margin-top: 10px;">               
                    <el-icon><Star /></el-icon>            
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
        <el-input v-model="postForm.introduction" placeholder="描述（15字以内）" maxlength="15" show-word-limit />
      </el-form-item>
      <el-form-item label="模型分类">
        <el-input v-model="postForm.model_types" placeholder="模型分类（机器学习or深度学习）" />
      </el-form-item>
      <el-form-item label="疾病分类">
      <el-select
        v-model="postForm.type_name"      
        placeholder="请选择标签"
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
function handleCategorySelect(cat) {
  selectedCategory.value = cat;
}
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'
import { Star,ArrowDown } from '@element-plus/icons-vue';
import logoImg from '@/assets/logo.png'
import UserInfo from './components/userInfo.vue'
import { useUserStore } from '@/store/user'

// 在组件中
const userStore = useUserStore()
const username = userStore.userInfo.UserName
console.log('当前用户名:', username)
const token = userStore.accessToken;
const router = useRouter()
const activeMenu = ref('/home')
const selectedCategory = ref('全部')
let posts = ref([])
const loading = ref(true)
let changing = ref(false)
const error = ref(null)
const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线分级' },
]

const categories = ['全部', '甲状腺分级', '乳腺癌分级', '肺结节分级', '肝脏分级', '脑肿瘤分级']

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
function showUserInfo() {
  console.log('userInfoRef:', userInfoRef.value)
  userInfoRef.value.openDialog()
}

async function getAllposts() {
  try {
    const response = await fetch('http://localhost:8888/getPosts', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token 
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
}

onMounted(async () => {
  await getAllposts()
})


function handleMenuSelect(index) {
  changing.value = true;
  activeMenu.value = index;
  router.push(index);
  setTimeout(() => {
    changing.value = false;
  }, 250);
}


async function goToPostDetail(post) {
  try {
    console.log('id:', post.post_id)
    const response = await fetch(`http://localhost:8888/getPostByID?post_id=${parseInt(post.post_id)}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token 
        
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
        model_types: data.model_types,
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


// 发帖弹窗逻辑
const dialogVisible = ref(false)
const tagOptions = [
  { label: '甲状腺分级', value: '甲状腺分级' },
  { label: '乳腺癌分级', value: '乳腺癌分级' },
  { label: '肺结节分级', value: '肺结节分级' },
  { label: '肝脏分级', value: '肝脏分级' },
  { label: '脑肿瘤分级', value: '脑肿瘤分级' }
]
const postForm = ref({
  title: '',
  introduction: '',
  content: '',
  user_name: username || '未登录用户', // 确保后端能接受这个默认值
  type_name: '全部', // 提供默认分类
  model_types: '', 
  is_public: true // 明确设置
})
function handleClose() {
  dialogVisible.value = false
}

async function submitPost() {
  console.log('用户发帖内容：', postForm.value)
  if (!postForm.value.title || !postForm.value.content) {
    ElMessage.error('请填写标题和内容')
    return
  }
  const postData = {
  type_name: postForm.value.type_name,
  user_name: postForm.value.user_name, // 确保这个字段有值
  title: postForm.value.title,
  introduction: postForm.value.introduction,
  content: postForm.value.content,
  model_types: postForm.value.model_types, // 确保这个字段有值
  created_date: new Date().toISOString(), // 改为created_date
  updated_date: new Date().toISOString(), // 添加updated_date
  is_public:  true // 确保有默认值
}
  let endpoint = 'http://localhost:8888/createPost';
  let method = 'POST';
  const response = await fetch(endpoint, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(postData),
      credentials: 'include', 
    });

  if (!response.ok) {
    ElMessage.error('发布失败，请重试')
    console.error('发布失败:', response.statusText)
    return
  }

  const data = await response.json()
  ElMessage.success('发布成功！')
  console.log('发布成功:', data)
  // 刷新帖子列表
  await getAllposts()

  // 清空表单
  postForm.value = {
    title: '',
    introduction: '',
    type_name: '',
    content: ''
  }
  dialogVisible.value = false
}
const searchText = ref('')

watch(selectedCategory, () => {
  changing.value = true;
  setTimeout(() => {
    changing.value = false;
  }, 150);
});

watch(searchText, () => {
  changing.value = true;
  setTimeout(() => {
    changing.value = false;
  }, 150);
});

const filteredPosts = computed(() => {
  let result = posts.value;
  if (selectedCategory.value !== '全部') {
    result = result.filter(m => m.type_name === selectedCategory.value);
  }
  if (searchText.value.trim()) {
    const keyword = searchText.value.trim().toLowerCase();
    result = result.filter(
      m =>
        m.title.toLowerCase().includes(keyword) ||
        m.introduction.toLowerCase().includes(keyword) ||
        (m.model_types && m.model_types.toLowerCase().includes(keyword))
    );
  }
  return result;
});
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
  position: relative;
  max-height:400px;
}
.sidebar {
  width: 220px;
  position: sticky;
  align-self: flex-start;
  z-index: 10;
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
  /* 保证卡片区域可滚动 */
  overflow-y: auto;
  max-height: calc(100vh - 160px); /* 视实际页面高度调整 */
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

.list-fade-enter-active,
.list-fade-leave-active {
  transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
}
.list-fade-enter-from,
.list-fade-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
.list-fade-leave-from,
.list-fade-enter-to {
  opacity: 1;
  transform: translateY(0);
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
.user-info-button{
  color:#cf5454;
  font-size:15px;
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
