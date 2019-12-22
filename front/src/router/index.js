import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import auth from '@/common/auth.service'

function checkAdmin(to, from, next) {
  if (localStorage.getItem("is_admin") !== "true") {
    next({name: "Home"})
    return
  }
  next()
}

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
  },
  {
    path: "/admin/scan/:amount",
    name: "AdminScan",
    component: () => import ("@/views/AdminScan.vue"),
    beforeEnter: checkAdmin
  },
  {
    path: "/admin/charge",
    name: "AdminCharge",
    component: () => import ("@/views/AdminCharge.vue"),
    beforeEnter: checkAdmin
  },
  {
    path: "/admin/records/:id",
    name: "AdminChargeDetail",
    component: () => import ("@/views/AdminChargeDetail.vue"),
    beforeEnter: checkAdmin
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
