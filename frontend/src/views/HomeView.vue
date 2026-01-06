<template>
  <div class="home-view">
    <div class="welcome-section">
      <h1 class="welcome-title">欢迎使用工作计划管理系统</h1>
      <p class="welcome-subtitle">高效管理您的工作计划,提升工作效率</p>
    </div>

    <div class="features-grid">
      <FeatureCard
        v-for="feature in features"
        :key="feature.name"
        :title="feature.title"
        :description="feature.description"
        :icon="feature.icon"
        :route="feature.route"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import FeatureCard from '@/components/FeatureCard.vue'

interface Feature {
  name: string
  title: string
  description: string
  icon: string
  route: string
}

const features: Feature[] = [
  {
    name: 'plans',
    title: '计划管理',
    description: '管理和跟踪工作计划',
    icon: '📋',
    route: '/plans',
  },
  {
    name: 'statistics',
    title: '统计分析',
    description: '查看工作计划的统计数据',
    icon: '📊',
    route: '/statistics',
  },
  {
    name: 'progress',
    title: '进度跟踪',
    description: '跟踪计划进度和完成情况',
    icon: '📈',
    route: '/progress',
  },
  {
    name: 'history',
    title: '历史记录',
    description: '查看操作历史和变更记录',
    icon: '📜',
    route: '/history',
  },
]
</script>

<style scoped>
.home-view {
  padding: 40px 20px;
}

.welcome-section {
  text-align: center;
  margin-bottom: 60px;
}

.welcome-title {
  font-size: 32px;
  font-weight: 600;
  color: #303133;
  margin: 0 0 16px 0;
}

.welcome-subtitle {
  font-size: 16px;
  color: #909399;
  margin: 0;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
  justify-items: center;
}

@media (max-width: 768px) {
  .welcome-title {
    font-size: 24px;
  }

  .welcome-subtitle {
    font-size: 14px;
  }

  .features-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
