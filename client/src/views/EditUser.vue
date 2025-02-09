<template>
  <ErrorBoundary>
    <div class="min-h-screen bg-white dark:bg-gray-900">
      <NavBar />
      <div class="container mx-auto mt-8 px-4">
        <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8">
          <div class="flex justify-between items-center mb-8">
            <h1 class="text-4xl font-bold text-gray-900 dark:text-white">
              {{ editingUserId ? 'Edit User' : 'Create User' }}
            </h1>
            <router-link 
              to="/client/admin/get_users"
              class="btn btn-primary">
              Back to Users
            </router-link>
          </div>

          <form @submit.prevent="submitForm" class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Email field -->
            <div class="field">
              <label class="block text-gray-700 dark:text-gray-300 mb-2">Email:</label>
              <input 
                v-model="form.email" 
                type="email" 
                class="form-input"
                :class="{ 'border-red-500': errors.email }"
                placeholder="Enter Email" 
                name="email" 
                required>
              <p v-if="errors.email" class="mt-1 text-sm text-red-500">{{ errors.email }}</p>
            </div>

            <!-- First name field -->
            <div class="field">
              <label class="block text-gray-700 dark:text-gray-300 mb-2">First name:</label>
              <input 
                v-model="form.firstname" 
                type="text" 
                class="form-input"
                :class="{ 'border-red-500': errors.firstname }"
                placeholder="Enter First Name" 
                name="firstname" 
                required>
              <p v-if="errors.firstname" class="mt-1 text-sm text-red-500">{{ errors.firstname }}</p>
            </div>

            <!-- Last name field -->
            <div class="field">
              <label class="block text-gray-700 dark:text-gray-300 mb-2">Last name:</label>
              <input 
                v-model="form.lastname" 
                type="text" 
                class="form-input"
                :class="{ 'border-red-500': errors.lastname }"
                placeholder="Enter Last Name" 
                name="lastname" 
                required>
              <p v-if="errors.lastname" class="mt-1 text-sm text-red-500">{{ errors.lastname }}</p>
            </div>

            <!-- Office field -->
            <div class="field">
              <label class="block text-gray-700 dark:text-gray-300 mb-2">Office:</label>
              <select 
                v-model="form.office" 
                class="form-select"
                :class="{ 'border-red-500': errors.office }"
                name="office" 
                required>
                <option value="">Select office</option>
                <option value="1">Office 1</option>
                <option value="2">Office 2</option>
              </select>
              <p v-if="errors.office" class="mt-1 text-sm text-red-500">{{ errors.office }}</p>
            </div>

            <!-- Role field -->
            <div class="field col-span-2">
              <label class="block text-gray-700 dark:text-gray-300 mb-2">Role:</label>
              <div class="control flex items-center space-x-6">
                <div class="flex items-center">
                  <input type="radio" 
                    :value="ROLES.USER.toString()" 
                    v-model="form.role" 
                    name="role"
                    class="form-radio">
                  <label class="ml-2 text-gray-700 dark:text-gray-300">{{ ROLE_NAMES[ROLES.USER] }}</label>
                </div>
                <div class="flex items-center">
                  <input type="radio" 
                    :value="ROLES.ADMIN.toString()" 
                    v-model="form.role" 
                    name="role"
                    class="form-radio">
                  <label class="ml-2 text-gray-700 dark:text-gray-300">{{ ROLE_NAMES[ROLES.ADMIN] }}</label>
                </div>
              </div>
              <p v-if="errors.role" class="mt-1 text-sm text-red-500">{{ errors.role }}</p>
            </div>

            <!-- Submit button -->
            <button type="submit" 
              :disabled="isSubmitting"
              class="col-span-2 btn"
              :class="isSubmitting ? 'btn-disabled' : 'btn-primary'">
              {{ isSubmitting ? 'Saving...' : (editingUserId ? 'Save Changes' : 'Create User') }}
            </button>
          </form>

          <!-- Message -->
          <div v-if="message" 
            class="mt-6 text-center text-lg p-4 rounded"
            :class="{
              'bg-green-100 text-green-700 dark:bg-green-900 dark:text-green-300': message.type === 'success',
              'bg-red-100 text-red-700 dark:bg-red-900 dark:text-red-300': message.type === 'error'
            }">
            {{ message.text }}
          </div>
        </div>
      </div>
    </div>

    <!-- Confirmation Modal -->
    <ConfirmModal
      v-model="showConfirmModal"
      title="Confirm Changes"
      confirm-text="Save Changes"
      @confirm="saveChanges">
      <p>Are you sure you want to save these changes for user {{ form.email }}?</p>
      <p class="mt-2 text-sm">This action cannot be undone.</p>
    </ConfirmModal>
  </ErrorBoundary>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import NavBar from './NavBar.vue'
import ErrorBoundary from '../components/ErrorBoundary.vue'
import ConfirmModal from '../components/ConfirmModal.vue'
import { ROLES, ROLE_NAMES } from '../constants'

const router = useRouter()
const route = useRoute()

const editingUserId = computed(() => route.query.id)
const isSubmitting = ref(false)
const showConfirmModal = ref(false)
const message = ref(null)
const errors = ref({})

const form = ref({
  email: '',
  firstname: '',
  lastname: '',
  office: '',
  role: ROLES.USER.toString()
})

const validateForm = () => {
  const newErrors = {}
  
  if (!form.value.email?.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/)) {
    newErrors.email = 'Please enter a valid email address'
  }
  if (!form.value.firstname?.trim()) {
    newErrors.firstname = 'First name is required'
  }
  if (!form.value.lastname?.trim()) {
    newErrors.lastname = 'Last name is required'
  }
  if (!form.value.office) {
    newErrors.office = 'Please select an office'
  }
  if (!form.value.role) {
    newErrors.role = 'Please select a role'
  }

  errors.value = newErrors
  return Object.keys(newErrors).length === 0
}

const submitForm = () => {
  if (!validateForm()) return
  showConfirmModal.value = true
}

const saveChanges = async () => {
  if (isSubmitting.value) return
  
  isSubmitting.value = true
  message.value = null
  
  try {
    const cookie = document.cookie
      .split('; ')
      .find(row => row.startsWith('Authorization='))
      ?.split('=')[1]

    if (!cookie) {
      throw new Error('Authentication required')
    }

    const response = await fetch('/api/server/user_edit', {
      method: editingUserId.value ? 'PUT' : 'POST',
      headers: { 
        'Content-Type': 'application/json',
        'Authorization': cookie
      },
      body: JSON.stringify({
        ...form.value,
        id: editingUserId.value
      }),
    })

    if (!response.ok) {
      throw new Error(response.status === 404 ? 'Cannot connect to server' : 'Failed to save user')
    }

    message.value = {
      type: 'success',
      text: `Successfully ${editingUserId.value ? 'updated' : 'created'} user information.`
    }

    // Clear form on success if creating new user
    if (!editingUserId.value) {
      form.value = {
        email: '',
        firstname: '',
        lastname: '',
        office: '',
        role: ROLES.USER.toString()
      }
    }
  } catch (err) {
    message.value = {
      type: 'error',
      text: err.message
    }
  } finally {
    isSubmitting.value = false
  }
}

// Load user data if editing
onMounted(async () => {
  if (editingUserId.value) {
    try {
      const cookie = document.cookie
        .split('; ')
        .find(row => row.startsWith('Authorization='))
        ?.split('=')[1]

      if (!cookie) {
        throw new Error('Authentication required')
      }

      const response = await fetch(`/api/server/get_users?id=${editingUserId.value}`, {
        headers: { 'Authorization': cookie }
      })

      if (!response.ok) {
        throw new Error('Failed to load user data')
      }

      const data = await response.json()
      if (data.users?.[0]) {
        const user = data.users[0]
        form.value = {
          email: user.email,
          firstname: user.firstname,
          lastname: user.lastname,
          office: user.office.toString(),
          role: user.role.toString()
        }
      }
    } catch (err) {
      message.value = {
        type: 'error',
        text: 'Failed to load user data: ' + err.message
      }
    }
  }
})
</script>
