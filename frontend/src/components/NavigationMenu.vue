<template>
  <el-aside class="navigation-menu" :width="isCollapsed ? '64px' : '200px'" role="navigation" aria-label="主导航">
    <div
      class="menu-toggle"
      @click="toggleCollapse"
      @keydown.enter="toggleCollapse"
      @keydown.space.prevent="toggleCollapse"
      v-if="isMobile"
      tabindex="0"
      role="button"
      :aria-label="isCollapsed ? '展开菜单' : '折叠菜单'"
      :aria-expanded="!isCollapsed"
    >
      <el-icon><component :is="isCollapsed ? 'Expand' : 'Fold'" /></el-icon>
    </div>

    <el-menu
      :default-active="activeMenu"
      :collapse="isCollapsed"
      :collapse-transition="false"
      background-color="#304156"
      text-color="#bfcbd9"
      active-text-color="#ffffff"
      @select="handleSelect"
      role="menubar"
    >
      <el-menu-item
        v-for="item in menuItems"
        :key="item.path"
        :index="item.path"
        :aria-label="item.name"
        role="menuitem"
        :aria-current="isActive(item.path) ? 'page' : undefined"
      >
        <el-icon>
          <span class="menu-icon" :aria-hidden="true">{{ item.icon }}</span>
        </el-icon>
        <template #title>
          <span class="menu-title">{{ item.name }}</span>
        </template>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

interface NavigationItem {
  name: string
  path: string
  icon: string
  description: string
}

const route = useRoute()
const router = useRouter()

const menuItems: NavigationItem[] = [
  {
    name: '计划管理',
    path: '/plans',
    icon: '📋',
    description: '管理和跟踪工作计划',
  },
  {
    name: '统计分析',
    path: '/statistics',
    icon: '📊',
    description: '查看工作计划的统计数据',
  },
  {
    name: '进度跟踪',
    path: '/progress',
    icon: '📈',
    description: '跟踪计划进度和完成情况',
  },
  {
    name: '历史记录',
    path: '/history',
    icon: '📜',
    description: '查看操作历史和变更记录',
  },
]

const isCollapsed = ref(false)
const isMobile = ref(false)

const activeMenu = computed(() => {
  return route.path
})

const isActive = (path: string) => {
  return route.path === path
}

const handleSelect = (path: string) => {
  router.push(path)
}

const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}

const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
  if (isMobile.value) {
    isCollapsed.value = true
  }
}

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
})
</script>

<style scoped>
.navigation-menu {
  background-color: #304156;
  transition: width 0.3s ease;
  overflow: hidden;
}

.menu-toggle {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;
  cursor: pointer;
  color: #bfcbd9;
  transition: background-color 0.3s ease;
}

.menu-toggle:hover {
  background-color: #263445;
}

.menu-icon {
  font-size: 20px;
}

.menu-title {
  margin-left: 8px;
}

:deep(.el-menu) {
  border-right: none;
}

:deep(.el-menu-item) {
  min-height: 56px;
  line-height: 56px;
}

:deep(.el-menu-item.is-active) {
  background-color: #409eff;
}

:deep(.el-menu-item:hover) {
  background-color: #263445;
}

:deep(.el-menu-item.is-active:hover) {
  background-color: #409eff;
}
</style>
