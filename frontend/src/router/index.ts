import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import MainLayout from '@/layouts/MainLayout.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '',
        redirect: '/plans',
      },
      {
        path: 'plans',
        name: 'Plans',
        component: () => import('@/views/PlansView.vue'),
      },
      {
        path: 'plans/create',
        name: 'PlanCreate',
        component: () => import('@/views/PlansView.vue'),
      },
      {
        path: 'plans/:id/edit',
        name: 'PlanEdit',
        component: () => import('@/views/PlansView.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
