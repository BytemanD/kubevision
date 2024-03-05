// Composables
import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "home" */ '@/views/Home.vue'),
      },
      {
        path: 'welcome',
        name: 'Welcome',
        component: () => import('@/views/Welcome.vue'),
      }
    ],
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
  },
  {
    path: '/welcome',
    name: 'Welcome',
    component: () => import('@/views/Welcome.vue'),
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/DashboardPage.vue'),
    children: [
      {
        path: 'home',
        component: () => import('@/components/dashboard/containers/overview/HomePage.vue'),
      },
      {
        path: 'namespace',
        component: () => import('@/components/dashboard/containers/overview/NamespacePage.vue'),
      },
      {
        path: 'node',
        component: () => import('@/components/dashboard/containers/overview/NodePage.vue'),
      },
      {
        path: 'workload',
        component: () => import('@/components/dashboard/containers/application/WorkloadPage.vue'),
      },
      {
        path: 'pod',
        component: () => import('@/components/dashboard/containers/application/PodPage.vue'),
      },
      {
        path: 'service',
        component: () => import('@/components/dashboard/containers/application/ServicePage.vue'),
      },
      {
        path: 'configmap',
        component: () => import('@/components/dashboard/containers/config/ConfigmapPage.vue'),
      },
      {
        path: 'secret',
        component: () => import('@/components/dashboard/containers/config/SecretPage.vue'),
      },
    ]
  },
  {
    path: '/config.json',
    name: 'ConfigJson',
    component: () => import('/config.json?url'),
  }
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
})

export default router
