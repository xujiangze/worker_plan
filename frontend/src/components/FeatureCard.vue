<template>
  <div class="feature-card" @click="handleClick" tabindex="0" role="button" :aria-label="`点击进入${title}`">
    <div class="icon-wrapper">
      <span class="icon">{{ icon }}</span>
    </div>
    <h3 class="title">{{ title }}</h3>
    <p class="description">{{ description }}</p>
    <div class="hint">点击进入</div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

interface Props {
  title: string
  description: string
  icon: string
  route: string
}

const props = defineProps<Props>()
const router = useRouter()

const handleClick = () => {
  router.push(props.route)
}
</script>

<style scoped>
.feature-card {
  width: 300px;
  height: 200px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.feature-card:hover {
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.15);
  transform: translateY(-4px);
}

.feature-card:focus {
  outline: 2px solid #409eff;
  outline-offset: 2px;
}

.icon-wrapper {
  margin-bottom: 16px;
}

.icon {
  font-size: 48px;
  display: block;
}

.title {
  margin: 0 0 12px 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  text-align: center;
}

.description {
  margin: 0;
  font-size: 14px;
  color: #909399;
  text-align: center;
  line-height: 1.5;
}

.hint {
  position: absolute;
  bottom: 12px;
  font-size: 12px;
  color: #909399;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card:hover .hint {
  opacity: 1;
}

@media (max-width: 768px) {
  .feature-card {
    width: 100%;
    max-width: 300px;
  }
}
</style>
