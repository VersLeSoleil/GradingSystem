<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput, ElUpload, ElSelect, ElOption, ElTable, ElTableColumn, ElMessage } from 'element-plus'

const router = useRouter()

// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型库', icon: 'i-ep-house' },
  { index: '/dataset', label: '数据集', icon: 'i-ep-user' },
  { index: '/online', label: '在线分级', icon: 'i-ep-document' },
  { index: '/history', label: '历史记录', icon: 'i-ep-setting' },
]
const activeMenu = ref('/online')

// 可选模型列表
const models = [
  { value: 'thyroid', label: '甲状腺AI分级模型', desc: '高准确率甲状腺分级AI模型' },
  { value: 'breast', label: '乳腺癌辅助诊断', desc: '乳腺癌影像分级智能模型' },
  { value: 'liver', label: '肝脏肿瘤分级', desc: '肝脏肿瘤分级AI模型' },
  { value: 'lung', label: '肺结节检测', desc: '肺结节分级与检测一体化模型' },
]
const selectedModel = ref(models[0].value)

// 上传文件
const fileList = ref([])
const handleUploadSuccess = () => {
  ElMessage.success('上传成功！')
  // 这里可触发后端推理请求
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
    <el-header height="64px" style="background: #409EFF; color: #fff; display: flex; align-items: center; padding: 0 32px;">
      <div style="font-size: 20px; font-weight: bold; margin-right: 32px;">疾病智能分级系统</div>
      <el-input placeholder="搜索..." clearable style="width: 300px; margin-right: 32px;" prefix-icon="el-icon-search" />
      <el-menu :default-active="activeMenu" mode="horizontal" background-color="#409EFF" text-color="#fff" active-text-color="#ffd04b" style="flex: 1; min-width: 500px; border-bottom: none;">
        <el-menu-item v-for="item in navMenus" :key="item.index" :index="item.index" @click="router.push(item.index)">
          <el-icon v-if="item.icon"><component :is="item.icon" /></el-icon>
          {{ item.label }}
        </el-menu-item>
      </el-menu>
      <el-avatar size="40" src="https://element-plus.org/images/element-plus-logo.svg" style="margin-left: 32px; background: #fff;" />
    </el-header>
    <el-main style="background: #fff; padding: 24px 0; height: calc(100vh - 64px);">
      <div style="display: flex; gap: 24px; height: 100%;">
        <!-- 左侧上传与模型选择区 -->
        <div style="flex: 1.2; display: flex; flex-direction: column; gap: 24px; margin-left: 32px;">
          <el-card shadow="hover">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">上传超声影像</div>
            <el-upload
              v-model:file-list="fileList"
              action="#"
              :auto-upload="false"
              :show-file-list="true"
              style="margin-bottom: 16px;"
              drag
            >
              <el-button type="primary">选择文件</el-button>
              <template #tip>
                <div style="color: #888;">支持jpg/png/dcm等格式，单次可多选</div>
              </template>
            </el-upload>
          </el-card>
          <el-card shadow="hover">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">选择模型</div>
            <el-select v-model="selectedModel" placeholder="请选择模型" style="width: 100%;">
              <el-option v-for="m in models" :key="m.value" :label="m.label" :value="m.value" />
            </el-select>
            <div style="margin-top: 16px;">
              <div style="font-size: 16px; font-weight: bold;">{{ modelTitle }}</div>
              <div style="color: #888; margin-top: 4px;">{{ modelDesc }}</div>
            </div>
          </el-card>
        </div>
        <!-- 右侧预测结果表格区 -->
        <div style="flex: 2; min-width: 300px;">
          <el-card shadow="hover">
            <div style="font-size: 18px; font-weight: bold; margin-bottom: 16px;">模型预测结果</div>
            <el-table :data="resultTable" style="width: 100%;">
              <el-table-column prop="id" label="#" width="60" />
              <el-table-column prop="name" label="样本名" />
              <el-table-column prop="result" label="预测分级" />
              <el-table-column prop="score" label="置信度" />
            </el-table>
          </el-card>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>
.el-header {
  box-shadow: 0 2px 8px #f0f1f2;
}
</style>