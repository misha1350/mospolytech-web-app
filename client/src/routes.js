import { createRouter, createWebHistory } from 'vue-router'
import Homepage from './views/Homepage.vue'
import GetUsers from './views/GetUsers.vue'
import EditUser from './views/EditUser.vue'
import Settings from './views/Settings.vue'
import store from './store'
import { ROLES } from './constants'

const routes = [
  {
    path: '/client/home',
    name: 'Home',
    component: Homepage
  },
  {
    path: '/client/admin/get_users',
    name: 'Get Users',
    component: GetUsers,
    meta: { requiresAdmin: true }
  },
  { 
    path: '/client/admin/edit_user',
    name: 'Edit Users',
    component: EditUser,
    meta: { requiresAdmin: true }
  },
  {
    path: '/client/settings',
    name: 'Settings',
    component: Settings
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/client/home'
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_APP_BASE_URL),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const userDetails = store.state.userDetails
  
  // Check if route requires admin
  if (to.meta.requiresAdmin) {
    if (!userDetails || parseInt(userDetails.role) !== ROLES.ADMIN) {
      // Not admin, redirect to home
      return next('/client/home')
    }
  }

  // No auth token in cookie, redirect to login
  const hasAuthCookie = document.cookie.split('; ').some(row => row.startsWith('Authorization='))
  if (!hasAuthCookie) {
    window.location.href = '/api/server/login'
    return
  }

  next()
})

export default router
