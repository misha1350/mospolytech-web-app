<template>
  <div v-if="error" class="min-h-screen flex items-center justify-center bg-white dark:bg-gray-900">
    <div class="max-w-md w-full p-6 bg-white dark:bg-gray-800 rounded-lg shadow-xl">
      <div class="text-red-600 dark:text-red-400 text-xl font-bold mb-4">
        Something went wrong
      </div>
      <pre class="text-gray-600 dark:text-gray-400 bg-gray-100 dark:bg-gray-700 p-4 rounded overflow-x-auto">
        {{ error }}
      </pre>
      <button @click="reset" 
        class="mt-4 w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded transition-colors duration-150">
        Try Again
      </button>
    </div>
  </div>
  <slot v-else></slot>
</template>

<script setup>
import { ref, onErrorCaptured } from 'vue'

const error = ref(null)

const reset = () => {
  error.value = null
}

onErrorCaptured((err) => {
  error.value = err
  return false // Prevent error from propagating
})
</script>