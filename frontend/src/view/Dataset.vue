<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElAvatar, ElCard, ElTag, ElMenu, ElMenuItem, ElInput } from 'element-plus'

const router = useRouter()

// 顶部导航菜单
const navMenus = [
  { index: '/home', label: '模型库', icon: 'i-ep-house' },
  { index: '/dataset', label: '数据集', icon: 'i-ep-user' },
  { index: '/online', label: '在线分级', icon: 'i-ep-document' },
  { index: '/history', label: '历史记录', icon: 'i-ep-setting' },
]
const activeMenu = ref('/dataset')

// 标签分类
const categories = [
  '全部',
  '甲状腺',
  '乳腺',
  '肝脏',
  '心脏',
  '肺部',
  '妇科',
  '其他'
]
const selectedCategory = ref('全部')

// 热门超声影像数据集
const datasets = [
  { id: 1, name: 'Thyroid Ultrasound Dataset', desc: '包含甲状腺结节超声图像及分级标签，适用于甲状腺良恶性分类与分级任务。', category: '甲状腺', source: '中国医学影像AI大赛', avatar: 'https://img.icons8.com/color/96/thyroid.png', link: 'https://aistudio.baidu.com/aistudio/competition/detail/152/0/task-definition' },
  { id: 2, name: 'BUSI 乳腺超声图像数据集', desc: '公开乳腺肿块超声图像，含良恶性标签，常用于乳腺癌检测与分割。', category: '乳腺', source: 'University of the Basque Country', avatar: 'https://img.icons8.com/color/96/breast-cancer.png', link: 'https://scholar.cu.edu.eg/?q=afahmy/pages/dataset' },
  { id: 3, name: 'CAMUS 心脏超声数据集', desc: '包含2D心脏超声图像及左心室分割标签，适用于心脏结构分割与功能评估。', category: '心脏', source: '法国雷恩大学', avatar: 'https://img.icons8.com/color/96/heart-with-pulse.png', link: 'https://www.creatis.insa-lyon.fr/Challenge/camus/' },
  { id: 4, name: 'Liver Ultrasound Dataset', desc: '肝脏肿瘤超声图像及分割标签，适用于肝脏肿瘤检测与分割。', category: '肝脏', source: 'Kaggle', avatar: 'https://img.icons8.com/color/96/liver.png', link: 'https://www.kaggle.com/datasets/andrewmvd/liver-tumor-segmentation' },
  { id: 5, name: 'Lung Ultrasound COVID Dataset', desc: '新冠肺炎相关肺部超声图像，含分级标签，适用于肺部病变检测。', category: '肺部', source: 'POCUS', avatar: 'https://img.icons8.com/color/96/lungs.png', link: 'https://pocus-covid19.org/en/dataset/' },
  { id: 6, name: 'OB/GYN Ultrasound Dataset', desc: '妇科及产科超声图像，适用于子宫、卵巢等结构识别。', category: '妇科', source: 'Kaggle', avatar: 'https://img.icons8.com/color/96/pregnant.png', link: 'https://www.kaggle.com/datasets/andrewmvd/ultrasound-nerve-segmentation' },
]

const filteredDatasets = computed(() => {
  if (selectedCategory.value === '全部') return datasets
  return datasets.filter(d => d.category === selectedCategory.value)
})

function handleMenuSelect(index) {
  activeMenu.value = index
  router.push(index)
}
</script>

<template>
  <el-container style="height: 100vh;width: 100%;">
    <!-- 顶部导航栏 -->
    <el-header height="64px" style="background: #409EFF; color: #fff; display: flex; align-items: center; padding: 0 32px;">
      <div style="font-size: 20px; font-weight: bold; margin-right: 32px;">疾病智能分级系统</div>
      <el-input placeholder="搜索数据集..." clearable style="width: 300px; margin-right: 32px;" prefix-icon="el-icon-search" />
      <el-menu :default-active="activeMenu" mode="horizontal" background-color="#409EFF" text-color="#fff" active-text-color="#ffd04b" @select="handleMenuSelect" style="flex: 1; min-width: 500px; border-bottom: none;">
        <el-menu-item v-for="item in navMenus" :key="item.index" :index="item.index">
          <el-icon v-if="item.icon"><component :is="item.icon" /></el-icon>
          {{ item.label }}
        </el-menu-item>
      </el-menu>
      <el-avatar size="40" src="https://element-plus.org/images/element-plus-logo.svg" style="margin-left: 32px; background: #fff;" />
    </el-header>
    <el-main style="background: #fff; padding: 24px 0; height: calc(100vh - 64px);">
      <div style="display: flex; gap: 24px; height: 100%;">
        <!-- 左侧数据集卡片区 -->
        <div style="flex: 2; display: flex; flex-direction: column; gap: 20px; overflow-y: auto; margin-left: 32px;">
          <el-card v-for="ds in filteredDatasets" :key="ds.id" shadow="hover" style="margin-bottom: 0;">
            <div style="display: flex; align-items: center;">
              <el-avatar :src="ds.avatar" size="large" style="margin-right: 16px;" />
              <div style="flex: 1;">
                <div style="font-size: 18px; font-weight: bold;">{{ ds.name }}</div>
                <div style="color: #888; margin: 4px 0;">{{ ds.desc }}</div>
                <el-tag size="small" type="info">{{ ds.category }}</el-tag>
                <el-tag size="small" type="success" style="margin-left: 8px;">{{ ds.source }}</el-tag>
              </div>
              <a :href="ds.link" target="_blank" style="margin-left: 16px; color: #409EFF; text-decoration: underline;">数据集链接</a>
            </div>
          </el-card>
        </div>
        <!-- 右侧标签分类区 -->
        <div style="flex: 1; min-width: 200px;">
          <div style="font-weight: bold; margin-bottom: 12px;">标签分类</div>
          <div style="display: flex; flex-wrap: wrap; gap: 10px;">
            <el-tag
              v-for="cat in categories"
              :key="cat"
              :type="selectedCategory === cat ? 'primary' : 'info'"
              @click="selectedCategory = cat"
              style="cursor: pointer; user-select: none;"
            >
              {{ cat }}
            </el-tag>
          </div>
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