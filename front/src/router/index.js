import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: "/booth/amount",
    name: "BoothAmount",
    component: () => import("@/views/BoothAmount.vue")
  },
  {
    path: "/booth/scan/:amount",
    name: "BoothScan",
    component: () => import ("@/views/BoothScan.vue")
  },
  {
    path: "/booth/records/:id",
    name: "BoothRecordDetail",
    component: () => import ("@/views/BoothRecordDetail.vue")
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
