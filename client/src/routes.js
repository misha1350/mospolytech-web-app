import { createRouter, createWebHistory } from 'vue-router'
import Homepage from './views/Homepage.vue'
import GetUsers from './views/GetUsers.vue'
import EditUser from './views/EditUser.vue'

const routes = [
  {
    path: '/client/home',
    name: 'Home',
    component: Homepage
  },
  {
    path: '/client/admin/get_users',
    name: 'Get Users',
    component: GetUsers
  },
  { 
    path: '/client/admin/edit_user',
    name: 'Edit Users',
    component: EditUser },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_APP_BASE_URL),
  routes
})

export default router
