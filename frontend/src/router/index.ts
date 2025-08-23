import { createRouter, createWebHistory } from 'vue-router'
import StoreHome from '../components/StoreHome.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
      {
        path: '/',
        name: 'Home',
        component: StoreHome
      },
      {
        path: '/store',
        name: 'store',
        component: () => import('../components/HelloWorld.vue')
      },
      {
        path: '/about',
        name: 'about',
        component: () => import('../views/HelloView.vue')
      },
  ],
})

export default router
