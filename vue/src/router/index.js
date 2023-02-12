import { createRouter, createWebHashHistory,createWebHistory } from "vue-router"
import Home from "../views/Home/index.vue"
import Login from "../views/Login/index.vue"
import Manage from "../views/Manage/index.vue"

const routes = [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/manage',
      name: 'manage',
      component: Manage
    },
]
export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

export default router
