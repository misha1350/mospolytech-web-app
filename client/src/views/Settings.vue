<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <NavBar />
    <div class="container mx-auto mt-8 px-4">
      <div class="max-w-2xl mx-auto bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
        <h1 class="text-3xl font-bold mb-8 text-gray-900 dark:text-white">Settings</h1>
        
        <div class="flex items-center justify-between py-4 border-b dark:border-gray-700">
          <div>
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white">Dark Mode</h2>
            <p class="text-gray-600 dark:text-gray-400">Switch between light and dark theme</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" 
                   :checked="store.state.isDarkMode"
                   @change="toggleDarkMode" 
                   class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
          </label>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useStore } from 'vuex'
import NavBar from './NavBar.vue'

const store = useStore()

const toggleDarkMode = () => {
  store.commit('toggleDarkMode')
}

// Initialize dark mode when component mounts
store.commit('initializeDarkMode')
</script>