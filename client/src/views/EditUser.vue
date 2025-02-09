<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <NavBar />
    <div class="container mx-auto mt-8 px-4">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8">
        <h1 class="text-4xl font-bold mb-8 text-gray-900 dark:text-white">Edit User</h1>
        <form @submit.prevent="submitForm" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Email field -->
          <div class="field">
            <label class="block text-gray-700 dark:text-gray-300 mb-2">Email:</label>
            <div class="control">
              <input v-model="user.email" type="email" placeholder="Enter Email" name="email" required 
                class="w-full form-input border-2 border-gray-300 p-2 rounded-md bg-white dark:bg-gray-700 dark:border-gray-600 dark:text-white placeholder-gray-400 dark:placeholder-gray-500">
            </div>
          </div>

          <!-- First name field -->
          <div class="field">
            <label class="block text-gray-700 dark:text-gray-300 mb-2">First name:</label>
            <div class="control">
              <input v-model="user.firstname" type="text" placeholder="Enter First Name" name="firstname" required 
                class="w-full form-input border-2 border-gray-300 p-2 rounded-md bg-white dark:bg-gray-700 dark:border-gray-600 dark:text-white placeholder-gray-400 dark:placeholder-gray-500">
            </div>
          </div>

          <!-- Last name field -->
          <div class="field">
            <label class="block text-gray-700 dark:text-gray-300 mb-2">Last name:</label>
            <div class="control">
              <input v-model="user.lastname" type="text" placeholder="Enter Last Name" name="lastname" required 
                class="w-full form-input border-2 border-gray-300 p-2 rounded-md bg-white dark:bg-gray-700 dark:border-gray-600 dark:text-white placeholder-gray-400 dark:placeholder-gray-500">
            </div>
          </div>

          <!-- Office field -->
          <div class="field">
            <label class="block text-gray-700 dark:text-gray-300 mb-2">Office:</label>
            <div class="control">
              <select v-model="user.office" name="office" 
                class="w-full form-select border-2 border-gray-300 p-2 rounded-md bg-white dark:bg-gray-700 dark:border-gray-600 dark:text-white">
                <option value="">Select your office</option>
                <option value="1">Office 1</option>
                <option value="2">Office 2</option>
              </select>
            </div>
          </div>

          <!-- Role field -->
          <div class="field col-span-2">
            <label class="block text-gray-700 dark:text-gray-300 mb-2">Role:</label>
            <div class="control flex items-center space-x-6">
              <div class="flex items-center">
                <input type="radio" id="user" value="1" v-model="user.role" name="role"
                  class="form-radio text-blue-600 dark:text-blue-400">
                <label for="user" class="ml-2 text-gray-700 dark:text-gray-300">User</label>
              </div>
              <div class="flex items-center">
                <input type="radio" id="admin" value="2" v-model="user.role" name="role"
                  class="form-radio text-blue-600 dark:text-blue-400">
                <label for="admin" class="ml-2 text-gray-700 dark:text-gray-300">Administrator</label>
              </div>
            </div>
          </div>

          <!-- Submit button -->
          <button type="submit" class="
              col-span-2
              bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-lg
              w-full md:w-auto mx-auto block mt-6
              transition duration-150 ease-in-out
          ">
            Apply Changes
          </button>
        </form>

        <!-- Message -->
        <div v-if="message" class="mt-6 text-center text-lg" :class="{
          'text-green-600 dark:text-green-400': message.includes('Success'),
          'text-red-600 dark:text-red-400': !message.includes('Success')
        }">{{ message }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import NavBar from './NavBar.vue'

// User data
const user = ref({
  email: '',
  firstname: '',
  lastname: '',
  office: '',
  role: '1',
})

// Message on screen
const message = ref('')

// // Fetch user data when component is created
// onMounted(async () => {
//   const response = await fetch(`api/server/user_edit/`)
//   if (response.ok) {
//     user.value = await response.json()
//   } else {
//     message.value = 'Failed to load user data.'
//   }
// })

// Submit form
const submitForm = async () => {
  const response = await fetch('/api/server/user_edit', {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(user.value),
  })

  if (response.ok) {
    message.value = 'Successfully updated user information.'
  } else {
    message.value = 'Failed to update user information.'
  }
}
</script>

<style scoped>
/* Add your styles here */
</style>
