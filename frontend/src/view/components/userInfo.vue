<script setup>
import { ref, computed } from 'vue'
import { useUserStore } from '@/store/user'
import { ElDialog, ElCard, ElAvatar, ElSkeleton, ElSkeletonItem, ElDivider, ElInput, ElButton, ElMessage } from 'element-plus'
const dialogVisible = ref(false)
const userStore = useUserStore()
const user = computed(() => userStore.userInfo || {})
const loading = ref(true) 
const editing = ref(false)
const editIntro = ref('')

async function openDialog() {
  loading.value = true
  dialogVisible.value = true
  // 模拟加载时间，也可以换成 await 接口请求
  await new Promise(resolve => setTimeout(resolve, 500))

  editIntro.value = user.value.Resume || ''
  loading.value = false
}

function closeDialog() {
  dialogVisible.value = false
  editing.value = false
}

function startEdit() {
  editing.value = true
}

function saveIntro() {
  userStore.userInfo.Resume = editIntro.value
  editing.value = false
  ElMessage.success('简介已保存（仅本地模拟）')
}

defineExpose({ openDialog})

</script>

<template>
  <el-dialog
    v-model="dialogVisible"
    title="个人信息"
    width="540px"
    :before-close="closeDialog"
    :close-on-click-modal="false"
  >
    <el-card shadow="hover" class="profile-card">
      <el-skeleton :loading="loading" animated>
        <template #template>
          <el-skeleton-item variant="circle" class="profile-avatar-skeleton" />
          <el-skeleton-item variant="text" class="skeleton-line short" />
          <el-skeleton-item variant="text" class="skeleton-line medium" />
          <el-divider />
          <el-skeleton-item variant="text" class="skeleton-line full" v-for="i in 5" :key="i" />
        </template>

        <template #default>
          <el-avatar :src="user.Avatar || ''" size="120px" class="profile-avatar" />
          <div class="profile-name">{{ user.UserName || '未命名用户' }}</div>
          <div class="profile-email">{{ user.Email || '暂无邮箱' }}</div>

          <el-divider />

          <div class="profile-info">
            <div class="profile-field"><strong>性别：</strong>{{ user.Sex || '未知' }}</div>
            <div class="profile-field"><strong>生日：</strong>{{ user.Birthday ? new Date(user.Birthday).toLocaleDateString() : '无' }}</div>
            <div class="profile-field"><strong>角色：</strong>{{ user.Role || '未知' }}</div>
            <div class="profile-field"><strong>手机号：</strong>{{ user.Phone || '无' }}</div>
            <div class="profile-field"><strong>注册时间：</strong>{{ user.CreatedDate ? new Date(user.CreatedDate).toLocaleString() : '未知' }}</div>
          </div>

          <el-divider />

          <div class="profile-intro-label">个人简介</div>
          <div v-if="!editing" class="profile-intro">{{ user.Resume || '暂无简介' }}</div>
          <el-input v-else v-model="editIntro" type="textarea" rows="3" placeholder="请输入个人简介" />

          <div class="profile-btns">
            <el-button v-if="!editing" type="primary" @click="startEdit" size="small" icon="Edit">编辑简介</el-button>
            <el-button v-else type="success" @click="saveIntro" size="small" icon="Check">保存</el-button>
            <el-button v-else @click="editing = false" size="small" icon="Close">取消</el-button>
            <el-button type="danger" size="small" @click="closeDialog" icon="CloseBold">关闭</el-button>
          </div>
        </template>
      </el-skeleton>
    </el-card>
  </el-dialog>
</template>


<style scoped>
.profile-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 28px 24px;
  border-radius: 10px;
  background: #fdfdfd;
  transition: box-shadow 0.3s ease;
}
.profile-avatar {
  margin-bottom: 12px;
  border: 3px solid #e1e4e8;
}
.profile-name {
  font-size: 24px;
  font-weight: 600;
  color: #1f2d3d;
  margin-top: 4px;
}
.profile-email {
  font-size: 14px;
  color: #909399;
  margin-bottom: 12px;
}
.profile-info {
  width: 100%;
  margin-bottom: 12px;
}
.profile-field {
  width: 100%;
  margin: 6px 0;
  font-size: 14px;
  color: #606266;
}
.profile-intro-label {
  font-weight: bold;
  font-size: 15px;
  margin-bottom: 6px;
  width: 100%;
}
.profile-intro {
  min-height: 40px;
  width: 100%;
  margin-bottom: 12px;
  color: #303133;
  line-height: 1.4;
}
.profile-btns {
  width: 100%;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
  margin-top: 10px;
}
.profile-avatar-skeleton {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  margin-bottom: 12px;
}
.skeleton-line {
  border-radius: 4px;
  margin: 6px 0;
}
.skeleton-line.short {
  width: 60%;
  height: 20px;
}
.skeleton-line.medium {
  width: 50%;
  height: 16px;
}
.skeleton-line.full {
  width: 100%;
  height: 14px;
}

</style>
