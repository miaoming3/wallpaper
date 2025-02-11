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
    beforeEnter: (to, from, next) => {
      if (!localStorage.getItem("token")) {
        next('/');
      } else {
        next();
      }
    },
    children:[{
      path:"",
      name:"dash",
      component:import('../views/dash/Index.vue')
    },{
      path:"info",
      name:"adminInfo",
      component:import('../views/admin/Info.vue')
    },{
      path:"index",
      name:"adminIndex",
      component:import('../views/admin/Index.vue')
    },{
      path:"update/:id",
      name:"adminUpdate",
      component:import('../views/admin/Edit.vue')
    },{
      path:"changePassword",
      name:"changePassword",
      component:import('../views/admin/Change-password.vue')
    },{
      path:"setting",
      name:"setting",
      component:import('../views/setting/Index.vue')
    },{
      path:"user",
      name:"user",
      children:[{
        path:"index",
        name:"UserIndex",
        component:import('../views/user/Index.vue')
      }]
    },{
      path:"category",
      name:"category",
      children:[{
        path:"index",
        name:"categoryIndex",
        component:import('../views/category/Index.vue')
      }]
    },{
      path:"img",
      name:"img",
      children:[{
        path:"index",
        name:"imgIndex",
        component:import('../views/images/Index.vue')
      },{
        path:"tags",
        name:"imgTags",
        component:import('../views/tags/Index.vue')
      }]
    },{
      path:"menu",
      name:"menu",
      children:[{
        path:"index",
        name:"MenuIndex",
        component:import('../views/menu/Index.vue')
      }]
    }]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
