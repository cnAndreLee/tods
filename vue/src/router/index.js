import { createRouter, createWebHistory } from "vue-router"
import Home from "../views/Home/index.vue"
import Login from "../views/Login/index.vue"
import Manage from "../views/Manage/index.vue"
import Users from "../views/Manage/users.vue"
import Files from "../views/Manage/files.vue"
import Categories1 from "../views/Manage/categories1.vue"

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
      component: Manage,
      children:[
        {
          path: 'users',
          component: Users
        },
        {
          path: 'files',
          component: Files
        },
        {
          path: 'categories1',
          component: Categories1
        },
      ],
    },
]
export const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
