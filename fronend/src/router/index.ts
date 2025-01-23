import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import LoginView from '../views/login/Index.vue'
import LayoutView from "@/layout/layout.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'login',
    component: LoginView
  }, {
    path:"/admin",
    name:"admin",
    component:LayoutView,
    children:[{
      path:"/",
      name:"dash",
      component:import('../views/dash/Index.vue')
    }]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
