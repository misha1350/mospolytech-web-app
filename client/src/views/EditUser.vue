<template>
  <div id="app">
    <NavBar />
    <div class="container mx-auto mt-8 px-4">
  <h1 class="text-4xl font-bold mb-4">Edit User</h1>
  <form @submit.prevent="submitForm" class="grid grid-cols-1 md:grid-cols-2 gap-4">
    <!-- Email field -->
    <div class="field">
      <label class="form-label">Email:</label>
      <div class="control">
        <input v-model="user.email" type="email" placeholder="Enter Email" name="email" required class="form-input border-2 border-gray-300 p-2 rounded-md">
      </div>
    </div>

    <!-- First name field -->
    <div class="field">
      <label class="form-label">First name:</label>
      <div class="control">
        <input v-model="user.firstname" type="text" placeholder="Enter First Name" name="firstname" required class="form-input border-2 border-gray-300 p-2 rounded-md">
      </div>
    </div>

    <!-- Last name field -->
    <div class="field">
      <label class="form-label">Last name:</label>
      <div class="control">
        <input v-model="user.lastname" type="text" placeholder="Enter Last Name" name="lastname" required class="form-input border-2 border-gray-300 p-2 rounded-md">
      </div>
    </div>

    <!-- Office field -->
    <div class="field">
      <label class="form-label">Office:</label>
      <div class="control">
        <select v-model="user.office" name="office" class="form-select border-2 border-gray-300 p-2 rounded-md">
          <option value="">Select your office</option>
          <!-- Add more options as needed -->
          <option value="1">Office 1</option>
          <option value="2">Office 2</option>
          <!-- ... -->
        </select>
      </div>
    </div>

    <!-- Role field -->
    <div class="field">
      <label class="form-label">Role:</label>
      <div class="control flex items-center space-x-4">
        <input type="radio" id="user" value="1" v-model="user.role" name="role">
        <label for="user">User</label><br>
        <input type="radio" id="admin" value="2" v-model="user.role" name="role">
        <label for="admin">Administrator</label><br>
      </div>
    </div>

    <!-- Submit button -->
    <button type="submit" class="
        bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-3 rounded
        w-full md:w-auto mx-auto block mt-4
    ">
      Apply
    </button>
  </form>

  <!-- Message -->
  <div v-if="$message" class="$message">{{ $message }}</div>
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
    method: 'POST',
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
