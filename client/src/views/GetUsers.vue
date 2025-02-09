<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <NavBar />
    <div class="container mx-auto mt-8 px-4">
      <h1 class="text-4xl font-bold mb-4 dark:text-white">Get users from database</h1>
      <div class="flex gap-4 mb-4">
        <button @click="debug = !debug" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          Toggle Debug
        </button>
        <button @click="fetchUsers" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          Refresh
        </button>
      </div>

      <div v-if="debug" class="mb-4 p-4 bg-white dark:bg-gray-800 rounded-lg">
        <p class="dark:text-white">Status Code: {{ statusCode }}</p>
        <p class="dark:text-white">Status Text: {{ statusText }}</p>
      </div>
      
      <!-- If loading, show loading message -->
      <div v-if="loading" class="p-4 bg-white dark:bg-gray-800 rounded-lg dark:text-white">Loading...</div>
      
      <!-- If error, show error message -->
      <div v-else-if="error" class="p-4 bg-white dark:bg-gray-800 rounded-lg text-red-600 dark:text-red-400">{{ error }}</div>
      
      <!-- If no users found, show message -->
      <div v-else-if="users.length === 0" class="p-4 bg-white dark:bg-gray-800 rounded-lg dark:text-white">No users found.</div>
      
      <!-- Otherwise, show users -->
      <div v-else class="overflow-x-auto mt-6 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
        <table class="min-w-full leading-normal">
          <thead>
            <tr>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                ID
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                Role ID
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                Email
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                First Name
              </th>
              <th class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                Last Name
              </th>
              <th class="px-2 py-3 border-b-2 border-gray-200 bg-gray-100 dark:bg-gray-800 dark:border-gray-700 text-left text-xs font-semibold text-gray-600 dark:text-gray-300 uppercase tracking-wider">
                Edit
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id" class="text-gray-700 dark:text-gray-300">
              <td class="px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">{{ user.id }}</td>
              <td class="px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">{{ user.role }}</td>
              <td class="px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">{{ user.email }}</td>
              <td class="px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">{{ user.firstname }}</td>
              <td class="px-5 py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">{{ user.lastname }}</td>
              <td class="py-2 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-sm">
                <router-link :to="`/client/admin/edit_user`"
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-3 rounded">
                  Edit
                </router-link>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
  
<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import NavBar from './NavBar.vue'

const users = ref([])
const loading = ref(false)
const error = ref(null)

const debug = ref(false)
const statusCode = ref(null)
const statusText = ref(null)

// Add new data properties
const showModal = ref(false)
const currentUser = ref({
  email: '',
  firstname: '',
  lastname: '',
  office: '',
  role: '1',
})

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
    } else {
      throw new Error(`Failed to fetch users.`)
    }
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

// Add new methods
const openEditModal = (user) => {
  currentUser.value = user;
  showModal.value = true;
}

// const submitForm = async () => {
//   this.processing = true;
//   this.message = 'Processing request...';

//   const response = await fetch('/api/server/user_edit', {
//     method: 'POST',
//     headers: {
//       'Content-Type': 'application/json',
//     },
//     body: JSON.stringify(currentUser.value),
//   });

//   if (response.ok) {
//     this.message = 'Successfully updated user information.';
//   } else if (response.status === 404) {
//     this.message = 'Cannot establish connection with the backend. (HTTP return code 404)';
//   } else {
//     this.message = `Cannot edit user data in the database. (HTTP return code ${response.status})`;
//   }

//   this.processing = false;
// }

onMounted(fetchUsers)
</script>



  
<style scoped>
</style>



