<script setup>
import { ref, computed } from 'vue'
import { useUserStore } from '@/store/user'
import { ElDialog, ElCard, ElAvatar, ElSkeleton, ElSkeletonItem, ElDivider, ElInput, ElButton, ElMessage } from 'element-plus'
import useravatar from '@/assets/user-avatar.png'
const dialogVisible = ref(false)
const userStore = useUserStore()
const user = computed(() => {
  const info = userStore.userInfo || {}
  return {
    ...info,
    // 处理可能存在的 { String, Valid } 结构
    Resume: info.Resume?.Valid ? info.Resume.String : info.Resume || '',
    Phone: info.Phone?.Valid ? info.Phone.String : info.Phone || '',
    Email: info.Email?.Valid ? info.Email.String : info.Email || ''
  }
})
console.log('user:', user.value)
const loading = ref(true) 
const editing = ref(false)
const editIntro = ref('')
const editNumber = ref('')
const birthday = ref(user.value.Birthday || '')
const localSex = ref(user.value.Sex === 'F' ? '2' : '1')
async function openDialog() {
  loading.value = true
  dialogVisible.value = true
  // 模拟加载时间，也可以换成 await 接口请求
  await new Promise(resolve => setTimeout(resolve, 500))

  editIntro.value = user.value.Resume || ''
  editNumber.value = user.value.Phone || ''
  loading.value = false
}

function closeDialog() {
  dialogVisible.value = false
  editing.value = false
}

function startEdit() {
  editing.value = true
}

async function saveIntro() {
  if (!editIntro.value.trim()) {
    ElMessage.error('简介不能为空')
    return
  }

  userStore.userInfo.Resume = editIntro.value
  userStore.userInfo.Phone = editNumber.value
  userStore.userInfo.Sex = localSex.value === '1' ? 'M' : 'F'
  userStore.userInfo.Birthday = birthday.value
  
  editing.value = false
  let endpoint = 'http://localhost:8888/UpdateUserInfo';
  let method = 'POST';
  let requestBody = {
    user_name: user.value.UserName,
    resume: editIntro.value,
    phone: editNumber.value,
    sex: userStore.userInfo.Sex,
    birthday: birthday.value,
    password: user.value.Password,
    avatar: user.value.Avatar,
    role: user.value.Role,
    email: user.value.Email
  };
  const response = await fetch(endpoint, {
    method: method,
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestBody),
    credentials: 'include', 
  });
  if (response.ok) {
    ElMessage.success('个人简介已更新')
  } else {
    ElMessage.error('更新失败，请稍后再试')
  }
}

defineExpose({ openDialog})

</script>
<template>
  <el-dialog
    v-model="dialogVisible"
    width="460px"
    :before-close="closeDialog"
    :close-on-click-modal="false"
    class="pretty-profile-dialog"
    :show-close="false"
  >
    <div class="profile-container">
      <div class="dialog-title-bar">完善您的个人简介</div>

      <div class="profile-content">
        <div class="profile-header">
          <el-avatar :src="useravatar || ''" size="80px" />
          <div class="profile-meta">
            <div class="name">{{ user.UserName || '未命名用户' }}</div>
            <div class="email">{{ user.Email || '暂无邮箱' }}</div>
          </div>
        </div>

        <el-divider />

        <div class="profile-info">
          <div class="intro-section">
          <div class="intro-label" v-if="!editing">生日</div>
          <el-date-picker
              :readonly="!editing"
              v-model="birthday"
              type="date"
              placeholder="选择您的生日"
              :size="size"
            />
        </div>
        <div class="intro-section">
          <div class="intro-label">性别</div>
          <el-radio-group v-model="localSex" :disabled="!editing" size="large">
            <el-radio value="1" size="large">男</el-radio>
            <el-radio value="2" size="large">女</el-radio>
          </el-radio-group>
        </div>
          <div class="intro-section">
          <div class="intro-label">手机号</div>
          <div v-if="!editing" class="number-content">{{ user.Phone || '无' }}</div>
          <el-input v-else v-model="editNumber" type="textarea" rows="2" placeholder="电话号码" />
        </div>
          
        </div>

        <el-divider />

        <div class="intro-section">
          <div class="intro-label">个人简介</div>
          <div v-if="!editing" class="intro-content">{{ user.Resume || '暂无简介' }}</div>
          <el-input v-else v-model="editIntro" type="textarea" rows="4" placeholder="介绍一下自己..." />
        </div>

        <div class="btn-area">
          <el-button v-if="!editing" type="primary" @click="startEdit" class="edit">编辑简介</el-button>
          <el-button v-else type="success" @click="saveIntro" class="edit">保存</el-button>   
          <el-button v-if="!editing" @click="closeDialog" type="danger" class="cancel">关闭</el-button>
          <el-button v-else @click="editing = false" class="cancel">取消</el-button>
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<style scoped>
/* 去除对话框本体的 padding 和圆角问题 */
::v-deep(.el-dialog) {
  border-radius: 12px;
  overflow: hidden;
  padding: 0 !important;
  margin: 0 !important;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15);
}
.edit {
  color: #fff;
  background-color: rgb(196, 87, 87);
  border-color: rgb(196, 87, 87);
  width: 80px;
}
.edit:hover,
.edit:focus {
  background: var(--el-button-hover-color);
  border-color: var(--el-button-hover-color);
  color: var(--el-button-font-color);
}
.cancel {
  color: #fff;
  background-color: rgb(105, 105, 105);
  border-color: rgb(105, 105, 105);
  width: 80px;
}
.cancel:hover,
.cancel:focus {
  background: var(--el-button-hover-color);
  border-color: var(--el-button-hover-color);
  color: var(--el-button-font-color);
}
/* 去除内部 padding */
::v-deep(.el-dialog__body) {
  padding: 0 !important;
}

/* 主体容器填满内容 */
.profile-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: white;
}

/* 顶部标题条 */
.dialog-title-bar {
  background-color: #ad3c3c;
  color: white;
  font-size: 20px;
  font-weight: bold;
  text-align: center;
  padding: 16px;
}

/* 主体内容区域 */
.profile-content {
  padding: 20px;
}

/* 信息头部 */
.profile-header {
  display: flex;
  align-items: center;
  gap: 16px;
}

.profile-meta .name {
  font-size: 26px;
  font-weight: bold;
  color: #333;
}

.profile-meta .email {
  font-size: 13px;
  color: #888;
}

/* 基本信息区域 */
.profile-info {
  margin-top: 12px;
  font-size: 14px;
  color: #444;
  line-height: 1.6;
}

/* 简介区域 */
.intro-section {
  margin-top: 12px;
}

.intro-label {
  font-weight: 600;
  margin-bottom: 6px;
  color: #444;
}

.intro-content {
  background-color: #f8f8f8;
  border-radius: 6px;
  padding: 10px;
  min-height: 60px;
  white-space: pre-wrap;
}

.number-content {
  background-color: #f8f8f8;
  border-radius: 6px;
  padding: 10px;
  min-height: 30px;
  white-space: pre-wrap;
}

/* 按钮区域 */
.btn-area {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}
</style>
