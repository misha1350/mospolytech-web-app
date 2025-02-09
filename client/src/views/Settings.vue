<template>
  <ErrorBoundary>
    <div class="min-h-screen bg-white dark:bg-gray-900">
      <NavBar />
      <div class="container mx-auto mt-8 px-4">
        <div class="max-w-2xl mx-auto bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
          <h1 class="text-3xl font-bold mb-8 text-gray-900 dark:text-white">Settings</h1>
          
          <!-- Theme Settings -->
          <div class="flex items-center justify-between py-4 border-b dark:border-gray-700">
            <div>
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">Theme Preferences</h2>
              <p class="text-gray-600 dark:text-gray-400">Choose between light and dark theme</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" 
                     :checked="isDarkMode"
                     @change="toggleDarkMode" 
                     class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
              <span class="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300">
                {{ isDarkMode ? 'Dark Mode' : 'Light Mode' }}
              </span>
            </label>
          </div>

          <!-- User Information -->
          <div class="mt-8">
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">Account Information</h2>
            <div class="grid grid-cols-2 gap-6">
              <div class="space-y-2">
                <p class="text-sm text-gray-500 dark:text-gray-400">Full Name</p>
                <p class="text-gray-900 dark:text-white">{{ store.getters.userFullName }}</p>
              </div>
              <div class="space-y-2">
                <p class="text-sm text-gray-500 dark:text-gray-400">Role</p>
                <p class="text-gray-900 dark:text-white">{{ roleName }}</p>
              </div>
              <div class="space-y-2">
                <p class="text-sm text-gray-500 dark:text-gray-400">Email</p>
                <p class="text-gray-900 dark:text-white">{{ store.state.userDetails?.email }}</p>
              </div>
              <div class="space-y-2">
                <p class="text-sm text-gray-500 dark:text-gray-400">Office</p>
                <p class="text-gray-900 dark:text-white">Office {{ store.state.userDetails?.office }}</p>
              </div>
            </div>
          </div>

          <!-- Version Information -->
          <div class="mt-8 pt-4 border-t dark:border-gray-700">
            <p class="text-sm text-gray-500 dark:text-gray-400">
              Application Version: 1.0.0
            </p>
          </div>
        </div>
      </div>
    </div>
  </ErrorBoundary>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import NavBar from './NavBar.vue'
import ErrorBoundary from '../components/ErrorBoundary.vue'
import { ROLE_NAMES } from '../constants'

const store = useStore()

const isDarkMode = computed(() => store.state.isDarkMode)
const roleName = computed(() => 
  ROLE_NAMES[parseInt(store.state.userDetails?.role)] || 'Unknown'
)

const toggleDarkMode = () => {
  store.commit('toggleDarkMode')
}
</script>