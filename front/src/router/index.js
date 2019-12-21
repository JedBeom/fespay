import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import auth from '@/common/auth.service'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    beforeEnter: auth
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/Login.vue'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: "/booth/amount",
    name: "BoothAmount",
    component: () => import("@/views/BoothAmount.vue"),
    beforeEnter: auth
  },
  {
    path: "/booth/scan/:amount",
    name: "BoothScan",
    component: () => import ("@/views/BoothScan.vue"),
    beforeEnter: auth
  },
  {
    path: "/booth/records/:id",
    name: "BoothRecordDetail",
    component: () => import ("@/views/BoothRecordDetail.vue"),
    beforeEnter: auth
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
