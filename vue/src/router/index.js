import { createRouter, createWebHashHistory,createWebHistory } from "vue-router"
import Home from "../views/Home/index.vue"
import Login from "../views/Login/index.vue"

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
]
export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

export default router
