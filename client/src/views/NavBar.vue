<template>
  <nav class="bg-blue-800 dark:bg-gray-900 py-4">
    <ul class="flex items-center">
      <li>
        <router-link class="text-white hover:text-blue-300 mx-4" to="/client/home">Home</router-link>
      </li>
      <li v-if="store.getters.isAdmin">
        <router-link class="text-white hover:text-blue-300 mx-4" to="/client/admin/get_users">Display Users</router-link>
      </li>
      <li>
        <router-link class="text-white hover:text-blue-300 mx-4" to="/client/settings">Settings</router-link>
      </li>
      <li class="ml-auto flex items-center">
        <span class="text-white mr-4">{{ store.getters.userFullName }} ({{ roleName }})</span>
        <button @click="logout" 
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mr-4 transition-colors duration-150">
          Logout
        </button>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { ROLES, ROLE_NAMES } from '../constants'

const store = useStore()

const roleName = computed(() => 
  ROLE_NAMES[parseInt(store.state.userDetails?.role)] || 'Unknown'
)

const logout = async () => {
  // Clear auth cookie with proper path and domain
  document.cookie = 'Authorization=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; domain=localhost;'
  
  // Clear user details from store
  store.commit('setUserDetails', null)
  
  // Redirect to login
  window.location.href = '/api/server/login'
}
</script>
