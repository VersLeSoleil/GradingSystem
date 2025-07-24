<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElAvatar, ElMenu, ElMenuItem, ElInput, ElButton, ElIcon } from 'element-plus';
import { Position } from '@element-plus/icons-vue';
import logoImg from '@/assets/logo.png';
import { ArrowDown } from '@element-plus/icons-vue';
import { callDeepSeekAPI } from '@/view/tools/aiChatGen';
import MarkdownIt from 'markdown-it';
import { useUserStore } from '@/store/user'
import UserInfo from './components/userInfo.vue'
const userStore = useUserStore()
const router = useRouter();
const md = new MarkdownIt();
const username = userStore.userInfo.UserName
// 状态管理
const activeMenu = ref('/aichat');
const userInput = ref('');
const chatHistory = ref([]);
const isLoading = ref(false);
const isTyping = ref(false);

const presetQuestions = [
  '为我推荐甲状腺超声影像分级方案',
  '当前最热门的模型搭配方案是什么？',
  '哪些模型适合肺部B超图像分级？'
];

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
const userInfoRef = ref(null)
function showUserInfo() {
  console.log('userInfoRef:', userInfoRef.value)
  userInfoRef.value.openDialog()
}
// 初始化欢迎消息
onMounted(() => {
  chatHistory.value.push({
    from: 'ai',
    text: '您好！我是分级喵AI助手，欢迎询问我医学影像分级相关的问题！',
    loading: false
  });
});

const handleSend = async () => {
  const text = userInput.value.trim();
  if (!text || isLoading.value) return;

  // 添加用户消息
  chatHistory.value.push({ from: 'user', text });
  userInput.value = '';
  
  // 调用API服务
  await callDeepSeek(text);
};

const callDeepSeek = async (prompt) => {
  isLoading.value = true;
  const loadingMsg = { from: 'ai', text: '', loading: true };
  chatHistory.value.push(loadingMsg);

  try {
    const { success, content } = await callDeepSeekAPI(
      prompt,
      chatHistory.value.filter(msg => !msg.loading)
    );

    loadingMsg.text = content;
    loadingMsg.html = md.render(content);
    loadingMsg.loading = false;
  } finally {
    isLoading.value = false;
  }
};

const sendPreset = (q) => {
  userInput.value = q;
  handleSend();
};

function handleMenuSelect(index) {
  activeMenu.value = index;
  router.push(index);
}

const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线分级' },
]
</script>

<template>
  <el-container class="app-container">
    <!-- 顶部导航栏 -->
    <el-header class="app-header">
      <div class="header-left">
        <img :src="logoImg" alt="logo" class="logo" />
        <div class="title">
          分级喵 <span class="subtitle">Image Grading System</span>
        </div>
      </div>
      <div class="header-center">
        <el-menu
          :default-active="activeMenu"
          mode="horizontal"
          :ellipsis="false"
          active-text-color="#CD5C5C"
          @select="handleMenuSelect"
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

    <!-- 主体区域 -->
    <el-main class="chat-container">
      <!-- 聊天内容区域 -->
      <div class="chat-window">
        <!-- 预设问题 -->
        <div v-if="chatHistory.length <= 1" class="preset-questions">
          <div
            v-for="(q, idx) in presetQuestions"
            :key="'preset-' + idx"
            class="preset-bubble"
            @click="sendPreset(q)"
          >
            {{ q }}
          </div>
        </div>

        <!-- 聊天记录 -->
        <div 
          v-for="(msg, index) in chatHistory" 
          :key="index" 
          class="message" 
          :class="msg.from"
        >
          <div class="message-bubble" :class="{ loading: msg.loading }">
            <span v-if="msg.loading" class="typing-indicator">
              <span class="dot"></span>
              <span class="dot"></span>
              <span class="dot"></span>
            </span>
            <template v-else>
              <div v-if="msg.html" v-html="msg.html"></div>
              <div v-else>{{ msg.text }}</div>
            </template>
          </div>
        </div>
      </div>

      <!-- 输入区域 -->
      <div class="input-area">
        <el-input
          v-model="userInput"
          placeholder="输入你的问题，例如：帮我推荐一个模型组合..."
          class="chat-input"
          :disabled="isLoading"
          @keyup.enter="handleSend"
        />
        <el-button 
          type="primary" 
          class="send-button"
          :disabled="!userInput.trim() || isLoading"
          @click="handleSend"
        >
          <el-icon><Position /></el-icon>
        </el-button>
      </div>
    </el-main>
  </el-container>
  <UserInfo ref="userInfoRef" />
</template>

<style scoped>
:root {
  --el-color-primary: #CD5C5C; /* 主色调 - 这里使用砖红色为例 */
  --el-color-primary-light-3: #e89292; /* 浅色变体 */
  --el-color-primary-light-5: #f2c4c4; /* 更浅的变体 */
  --el-color-primary-light-7: #f9e1e1; /* 最浅的变体 */
  --el-color-primary-light-8: #fcefef; /* 超浅变体 */
  --el-color-primary-light-9: #fef7f7; /* 极浅变体 */
  --el-color-primary-dark-2: #a83f3f; /* 深色变体 */
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

/* 按钮悬停和激活状态 */
.el-button--primary {
  --el-button-hover-bg-color: var(--el-color-primary-light-3);
  --el-button-active-bg-color: var(--el-color-primary-dark-2);
}
.app-container {
  height: 100vh;
  width: 100%;
  background-color: #f5f7fa;
}

.app-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 40px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  position: relative;
  z-index: 10;
}

.header-left {
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.logo {
  width: 32px;
  height: 32px;
  margin-right: 12px;
}

.title {
  font-weight: 600;
  font-size: 18px;
}

.subtitle {
  color: #999;
  font-size: 14px;
  margin-left: 4px;
}

.header-center {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.user-avatar {
  cursor: pointer;
  transition: transform 0.3s;
}

.user-avatar:hover {
  transform: scale(1.1);
}

.chat-container {
  padding: 20px;
  display: flex;
  flex-direction: column;
  height: calc(100vh - 64px);
  background-color: #f5f7fa;
}

.chat-window {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  margin-bottom: 16px;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.preset-questions {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-bottom: 16px;
}

.preset-bubble {
  background: #f0f7ff;
  padding: 8px 16px;
  border-radius: 18px;
  cursor: pointer;
  font-size: 14px;
  color: #a83f3f;
  transition: all 0.3s;
  border: 1px solid #e0e0e0;
}

.preset-bubble:hover {
  background: #d9ebff;
  transform: translateY(-2px);
}

.message {
  display: flex;
}

.message.user {
  justify-content: flex-end;
}

.message.ai {
  justify-content: flex-start;
}

.message-bubble {
  max-width: 75%;
  padding: 12px 16px;
  border-radius: 18px;
  font-size: 15px;
  line-height: 1.5;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  position: relative;
}

.message.user .message-bubble {
  background: #CD5C5C;
  color: white;
  border-bottom-right-radius: 4px;
}

.message.ai .message-bubble {
  background: #f1f3f4;
  color: #202124;
  border-color: #CD5C5C !important;
  border-bottom-left-radius: 4px;
}

.message-bubble.loading {
  background: #f1f3f4;
  min-width: 80px;
}

.typing-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  height: 20px;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #888;
  animation: dotPulse 1.4s infinite ease-in-out;
}

.dot:nth-child(1) {
  animation-delay: 0s;
}

.dot:nth-child(2) {
  animation-delay: 0.2s;
}

.dot:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes dotPulse {
  0%, 60%, 100% {
    transform: translateY(0);
    opacity: 0.6;
  }
  30% {
    transform: translateY(-5px);
    opacity: 1;
  }
}

.input-area {
  display: flex;
  gap: 12px;
  padding: 12px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
}

.chat-input {
  flex: 1;
}

.chat-input :deep(.el-input__inner) {
  border-radius: 18px;
  padding: 12px 16px;
  border: 1px solid #e0e0e0;
}

.send-button {
  width: 48px;
  background-color: #CD5C5C;
  height: 48px;
  border-radius: 50%;
  background-color: #CD5C5C !important;
  display: flex;
  align-items: center;
  justify-content: center;
}

.send-button :deep(.el-icon) {
  font-size: 20px;
}

/* 滚动条样式 */
.chat-window::-webkit-scrollbar {
  width: 6px;
}

.chat-window::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.chat-window::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.chat-window::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 消息动画 */
.message-enter-active {
  transition: all 0.3s ease;
}

.message-enter-from {
  opacity: 0;
  transform: translateY(10px);
}
</style>