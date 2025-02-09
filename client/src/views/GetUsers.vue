<template>
  <ErrorBoundary>
    <div class="min-h-screen bg-white dark:bg-gray-900">
      <NavBar />
      <div class="container mx-auto mt-8 px-4">
        <div class="flex items-center justify-between mb-6">
          <h1 class="text-4xl font-bold dark:text-white">User Management</h1>
          <div class="flex gap-4">
            <router-link to="/client/admin/edit_user" class="btn btn-primary">
              Add User
            </router-link>
            <button @click="debug = !debug" class="btn" :class="debug ? 'btn-primary' : 'bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-300'">
              {{ debug ? 'Hide' : 'Show' }} Debug Info
            </button>
            <button @click="fetchUsers" :disabled="loading" class="btn" :class="loading ? 'btn-disabled' : 'btn-primary'">
              {{ loading ? 'Loading...' : 'Refresh' }}
            </button>
          </div>
        </div>

        <div v-if="debug" class="mb-4 p-4 bg-white dark:bg-gray-800 rounded-lg">
          <p class="dark:text-white">Status Code: {{ statusCode }}</p>
          <p class="dark:text-white">Status Text: {{ statusText }}</p>
        </div>
        
        <div v-if="loading" class="animate-pulse">
          <div v-for="i in 3" :key="i" class="mb-4 h-20 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
        </div>
        
        <div v-else-if="error" 
          class="p-4 bg-red-100 dark:bg-red-900 border-l-4 border-red-500 text-red-700 dark:text-red-300">
          <p class="font-medium">Error</p>
          <p>{{ error }}</p>
        </div>
        
        <div v-else-if="users.length === 0" 
          class="p-4 bg-yellow-100 dark:bg-yellow-900 border-l-4 border-yellow-500 text-yellow-700 dark:text-yellow-300">
          No users found.
        </div>
        
        <div v-else class="overflow-x-auto bg-white dark:bg-gray-800 rounded-lg shadow-lg">
          <table class="min-w-full leading-normal">
            <thead>
              <tr>
                <th v-for="header in tableHeaders" :key="header.key"
                  class="px-5 py-3 border-b-2 border-gray-200 dark:border-gray-700 bg-gray-100 dark:bg-gray-800 
                  text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                  {{ header.label }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in users" :key="user.id" 
                class="hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-150">
                <td class="table-cell">{{ user.id }}</td>
                <td class="table-cell">{{ ROLE_NAMES[user.role] || 'Unknown' }}</td>
                <td class="table-cell">{{ user.email }}</td>
                <td class="table-cell">{{ user.firstname }}</td>
                <td class="table-cell">{{ user.lastname }}</td>
                <td class="table-cell">
                  <router-link :to="{ path: '/client/admin/edit_user', query: { id: user.id }}"
                    class="btn btn-primary">
                    Edit
                  </router-link>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </ErrorBoundary>
</template>
  
<script setup>
import { ref, onMounted } from 'vue'
import NavBar from './NavBar.vue'
import ErrorBoundary from '../components/ErrorBoundary.vue'
import { ROLE_NAMES } from '../constants'

const users = ref([])
const loading = ref(false)
const error = ref(null)
const debug = ref(false)
const statusCode = ref(null)
const statusText = ref(null)

const tableHeaders = [
  { key: 'id', label: 'ID' },
  { key: 'role', label: 'Role' },
  { key: 'email', label: 'Email' },
  { key: 'firstname', label: 'First Name' },
  { key: 'lastname', label: 'Last Name' },
  { key: 'actions', label: '' }
]

const fetchUsers = async () => {
  loading.value = true
  error.value = null
  
  try {
    const cookie = document.cookie
      .split('; ')
      .find(row => row.startsWith('Authorization='))
      ?.split('=')[1]

    if (!cookie) {
      throw new Error('Authentication required')
    }

    const response = await fetch('/api/server/get_users', {
      headers: { 'Authorization': cookie }
    })
    
    statusCode.value = response.status
    statusText.value = response.statusText

    if (!response.ok) {
      throw new Error(response.status === 404 ? 'Cannot connect to server' : 'Failed to fetch users')
    }

    const data = await response.json()
    users.value = data.users
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.table-cell {
  @apply px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm text-gray-700 dark:text-gray-300;
}
</style>



