<script setup>
import { ref, computed,watch,onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Delete,ZoomIn, ArrowDown,MoreFilled} from '@element-plus/icons-vue'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput, ElTour, ElUpload, ElSelect, ElOption, ElTable, ElTableColumn, ElMessage } from 'element-plus'
import logoImg from '@/assets/logo.png'
import axios from 'axios'
import UserInfo from '@/view/components/userInfo.vue'
import { useUserStore } from '@/store/user'
const router = useRouter()
const userStore = useUserStore()
const username = userStore.userInfo.UserName
const open = ref(true)
const ref1 = ref()
const ref2 = ref()
const ref3 = ref()
const ref4 = ref()
const userInfoRef = ref(null)
// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型广场' },
  { index: '/aichat', label: 'ai助手' },
  { index: '/my-model', label: '在线训练' },
]
const activeMenu = ref('/my-model')
function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}

onMounted(() => {
  if (!localStorage.getItem('tour_shown')) {
    open.value = true
  } else {
    open.value = false
  }
})

// 监听 open 变化，关闭时设置标记
watch(open, (val) => {
  if (!val) {
    localStorage.setItem('tour_shown', '1')
  }
})

// 可选模型列表
const models = [
  { value: '1', label: '随机森林+逻辑回归', desc: '随机森林搭配逻辑回归模型的综合性解决方案' },
  { value: '2', label: 'GoogleNet', desc: '高准确率甲状腺分级AI模型' },
  { value: '3', label: '乳腺癌辅助诊断', desc: '乳腺癌影像分级智能模型' },
  { value: '4', label: '肝脏肿瘤分级', desc: '肝脏肿瘤分级AI模型' },
  { value: '5', label: '肺结节检测', desc: '肺结节分级与检测一体化模型' },
]
const selectedModel = ref(models[0].value)
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
// 上传文件
const imageList = ref([])
const maskList = ref([])

function getMaskNameFromImageName(imageName) {
  const dotIndex = imageName.lastIndexOf('.')
  if (dotIndex === -1) return imageName + '_mask' // 没有扩展名
  const nameWithoutExt = imageName.substring(0, dotIndex)
  const ext = imageName.substring(dotIndex)
  return nameWithoutExt + '_mask' + ext
}

const handleUploadSuccess = async () => {
  if (imageList.value.length === 0) {
    ElMessage.warning('请先上传至少一张图片')
    return
  }

  const formData = new FormData()

  const pairs = []

  imageList.value.forEach(imageFile => {
    const imageName = imageFile.name
    const expectedMaskName = getMaskNameFromImageName(imageName)

    // 在 maskList 中查找匹配的掩膜
    const matchedMask = maskList.value.find(maskFile => maskFile.name === expectedMaskName)

    if (!matchedMask) {
      console.warn(`未找到 ${imageName} 对应的掩膜 ${expectedMaskName}`)
      // 你可以选择跳过这一对，或者报错提醒
      ElMessage.warning(`未找到 ${imageName} 对应的掩膜`)
      return
    }

    // 收集 pair，供上传
    pairs.push({ image: imageFile.raw, mask: matchedMask.raw })
  })

  if (pairs.length === 0) {
    ElMessage.warning('没有匹配成功的图像-掩膜对')
    return
  }

  // 将每个 pair 添加到 formData 中（可按序号命名）
  pairs.forEach((pair, index) => {
    formData.append(`images`, pair.image)
    formData.append(`masks`, pair.mask)
    // 或者你想用更结构化的字段名：
    // formData.append(`pair_${index}_image`, pair.image)
    // formData.append(`pair_${index}_mask`, pair.mask)
  })

  try {
    const response = await axios.post('http://127.0.0.1:8888/predict', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })

    // 
    resultTable.value = response.data.results  // 假设后端返回的是 { results: [{ name, result, score }] }
    ElMessage.success('预测完成！')

  } catch (error) {
    console.error('上传或预测失败:', error)
    ElMessage.error('上传失败，请重试')
  }
}


const handleRemove = (file, uploadFiles) => {
  console.log(file, uploadFiles);
  // 你可以在这里添加自定义删除逻辑
  // 例如：
  // const index = uploadFiles.findIndex(f => f.uid === file.uid);
  // if (index !== -1) {
  //   uploadFiles.splice(index, 1);
  // }
}

// 模型简介
const modelDesc = computed(() => {
  const m = models.find(m => m.value === selectedModel.value)
  return m ? m.desc : ''
})
const modelTitle = computed(() => {
  const m = models.find(m => m.value === selectedModel.value)
  return m ? m.label : ''
})

// 预测结果表格（示例数据）
const resultTable = ref([
  { id: 1, name: '样本1', result: '良性', score: 0.92 },
  { id: 2, name: '样本2', result: '恶性', score: 0.81 },
])
</script>

<template>
  <el-container style="height: 100vh;width: 100%;">
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

    <el-main style="background: #fff; padding: 24px 0; height: calc(100vh - 64px);">
      <div style="display: flex; gap: 24px; height: 100%; padding: 0 32px;">
        <!-- 左侧区域 (1/2宽度) -->
        <div style="flex: 1; display: flex; flex-direction: column; gap: 24px;">
          <!-- 模型选择卡片 -->
          <el-card shadow="hover" style="flex: 1; min-height: 250px;">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">选择模型</div>
            <el-select v-model="selectedModel" placeholder="请选择模型" style="width: 100%;">
              <el-option v-for="m in models" :key="m.value" :label="m.label" :value="m.value" />
            </el-select>
            <div style="margin-top: 16px;">
              <div style="font-size: 16px; font-weight: bold;">{{ modelTitle }}</div>
              <div style="color: #888; margin-top: 4px;">{{ modelDesc }}</div>
            </div>
            <el-button type="primary" style="margin-top: 15px; margin-left:500px;"  @click="handleUploadSuccess" ref="ref3">开始预测</el-button>
          </el-card>
          
          <!-- 结果预览卡片 -->
          <el-card shadow="hover" style="flex: 2;" ref="ref4">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">模型预测结果</div>
            <el-table :data="resultTable" style="width: 100%; height: calc(100% - 40px);">
              <el-table-column prop="id" label="#" width="60" />
              <el-table-column prop="name" label="样本名" />
              <el-table-column prop="result" label="预测分级" />
              <el-table-column prop="score" label="置信度" />
            </el-table>
          </el-card>
        </div>
        
        <!-- 右侧上传区域 (1/2宽度) -->
        <div style="flex: 1;">
          
          <el-card shadow="hover" style="height: 50%;">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">上传超声影像</div>
            <div style="max-height: 500px; overflow-y: auto; padding-right: 4px;">
            
              <el-upload
              ref="ref1"
              v-model:file-list="imageList"
              action="#"
              :auto-upload="false"
              :multiple="true"
              list-type="picture-card"
              :on-change="handleChange"
              :on-remove="handleRemove"
              class="custom-uploader"
              style="height: calc(100% - 60px);"
            >           
              <span style="font-size:60px;font-weight: lighter; color:gray">+</span>             
              <template #file="{ file }">
                <div class="uploaded-file">
                  <img
                    v-if="file.raw.type.includes('image')"
                    :src="file.url ? file.url : URL.createObjectURL(file.raw)"
                    class="uploaded-image"
                    alt=""
                  />
                  <el-icon v-else class="uploaded-file-icon"><Document /></el-icon>
                  <span class="uploaded-file-name">{{ file.name }}</span>
                  <span class="uploaded-file-actions">
                    <el-icon @click="handlePreview(file)"><ZoomIn /></el-icon>
                    <el-icon @click="handleRemove(file)"><Delete /></el-icon>
                  </span>
                </div>
              </template>
            </el-upload>
            </div>
          </el-card>
          <el-card shadow="hover" style="height: 50%;">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">上传掩膜</div>
            <div style="max-height: 500px; overflow-y: auto; padding-right: 4px;">
            
              <el-upload
              ref="ref2"
              v-model:file-list="maskList"
              action="#"
              :auto-upload="false"
              :multiple="true"
              list-type="picture-card"
              :on-change="handleChange"
              :on-remove="handleRemove"
              class="custom-uploader"
              style="height: calc(100% - 60px);"
            >           
              <span style="font-size:60px;font-weight: lighter; color:gray">+</span>             
              <template #file="{ file }">
                <div class="uploaded-file">
                  <img
                    v-if="file.raw.type.includes('image')"
                    :src="file.url ? file.url : URL.createObjectURL(file.raw)"
                    class="uploaded-image"
                    alt=""
                  />
                  <el-icon v-else class="uploaded-file-icon"><Document /></el-icon>
                  <span class="uploaded-file-name">{{ file.name }}</span>
                  <span class="uploaded-file-actions">
                    <el-icon @click="handlePreview(file)"><ZoomIn /></el-icon>
                    <el-icon @click="handleRemove(file)"><Delete /></el-icon>
                  </span>
                </div>
              </template>
            </el-upload>
            </div>
          </el-card>
        </div>
      </div>
    </el-main>
  </el-container>
  <el-tour v-model="open">
    <el-tour-step :target="ref1?.$el" title="在此上传原始超声影像,仅支持jpg格式">
    </el-tour-step>
    <el-tour-step
      :target="ref2?.$el"
      title="在此上传掩膜，注意命名格式需为原图名加上后缀 _mask，仅支持jpg格式"
    />
    <el-tour-step
      :target="ref3?.$el"
      title="点击此按钮开始预测分级结果"
    />
    <el-tour-step
      :target="ref4?.$el"
      title="在此查看预测结果"
      description="结果以表格形式呈现"
    />
  </el-tour>
  <UserInfo ref="userInfoRef" />
</template>

<style scoped>
.user-info-button{
  color:#cf5454;
  font-size:15px;
  border: none;          /* 移除边框 */
  background: none;      /* 移除背景 */
  padding: 0;            /* 移除内边距 */
  cursor: pointer;       /* 保持手型指针 */
  outline: none;         /* 移除聚焦时的轮廓线 */
}
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
  font-size: 16px;
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

/* 修改上传框 */
.upload-container {
  height: calc(100% - 60px); /* 减去标题高度 */
  overflow-y: auto; /* 垂直滚动 */
  padding-right: 8px; /* 为滚动条留出空间 */
}

/* 调整上传列表布局 */
:deep(.el-upload-list--picture-card) {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 10px;
}

/* 调整单个上传卡片大小 */
:deep(.el-upload-list--picture-card .el-upload-list__item),
:deep(.el-upload--picture-card) {
  width: 160px;
  height: 120px;
  margin: 0; /* 移除默认margin */
}

/* 自定义滚动条样式 */
.upload-container::-webkit-scrollbar {
  width: 6px;
}

.upload-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.upload-container::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.upload-container::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 其他原有样式保持不变 */
.upload-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 160px;
  color: var(--el-text-color-secondary);
  cursor: pointer;
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
}

.upload-area:hover {
  border-color: var(--el-color-primary);
}
.uploaded-image {
  width: 100%;
  height: 100%;
  object-fit: contain;  /* 关键属性 - 保持比例完整显示 */
  object-position: center; /* 居中显示 */
  background-color: #f5f5f5; /* 添加背景色填充空白区域 */
}
</style>