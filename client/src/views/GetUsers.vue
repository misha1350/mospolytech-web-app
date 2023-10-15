<template>
    <div id="app">
        <NavBar />
        <div class="container mx-auto mt-8 px-4">
            <h1 class="text-4xl font-bold mb-4">Get users from database</h1>
            <button @click="debug = !debug" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mr-4">
                            Toggle Debug
            </button>
            <div v-if="debug">
                            <p>Status Code: {{ statusCode }}</p>
                            <p>Status Text: {{ statusText }}</p>
            </div>
            <button @click="fetchUsers" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Refresh
            </button>
            
            <!-- If loading, show loading message -->
            <div v-if="loading">Loading...</div>
            
            <!-- If error, show error message -->
            <div v-else-if="error">{{ error }}</div>
            
            <!-- If no users found, show message -->
            <div v-else-if="users.length === 0">No users found.</div>
            
      <!-- Otherwise, show users -->
      <div v-else class="overflow-x-auto mt-6">
        <table class="min-w-full leading-normal">
          <thead>
            <tr>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                ID
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                Role ID
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                    Email
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                    First Name
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                    Last Name
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id" class="text-gray-700">
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">{{ user.id }}</td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">{{ user.role }}</td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">{{ user.email }}</td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">{{ user.firstname }}</td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">{{ user.lastname }}</td>
                            <!-- Add more fields as needed -->
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
  
<script setup>
import NavBar from './NavBar.vue'
import { ref, onMounted } from 'vue'

const users = ref([])
const loading = ref(false)
const error = ref(null)

const debug = ref(false)
const statusCode = ref(null)
const statusText = ref(null)

const fetchUsers = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await fetch('/api/server/get_users')
    statusCode.value = response.status
    statusText.value = response.statusText
    if (response.ok) {
  const data = await response.json()
  users.value = data.users
}
 else {
      throw new Error(`Failed to fetch users.`)
    }
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchUsers)
</script>

  
<style scoped>
</style>
  


  