<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="modelValue" class="fixed inset-0 z-50 overflow-y-auto" @click="closeOnOverlay">
        <div class="flex items-center justify-center min-h-screen p-4">
          <!-- Backdrop -->
          <div class="fixed inset-0 bg-black opacity-50"></div>
          
          <!-- Modal -->
          <div class="relative bg-white dark:bg-gray-800 rounded-lg w-full max-w-md p-6">
            <!-- Header -->
            <div class="mb-4">
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">
                {{ title }}
              </h3>
              <button 
                @click="close"
                class="absolute top-4 right-4 text-gray-400 hover:text-gray-500 dark:hover:text-gray-300"
                aria-label="Close">
                <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            
            <!-- Content -->
            <div class="mb-6 text-gray-600 dark:text-gray-300">
              <slot></slot>
            </div>
            
            <!-- Actions -->
            <div class="flex justify-end space-x-4">
              <button 
                @click="close"
                class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 dark:text-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 transition-colors duration-150">
                Cancel
              </button>
              <button 
                @click="confirm"
                :class="[
                  'px-4 py-2 rounded-md transition-colors duration-150',
                  confirmButtonClass || 'text-white bg-blue-500 hover:bg-blue-600'
                ]">
                {{ confirmText }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  confirmText: {
    type: String,
    default: 'Confirm'
  },
  confirmButtonClass: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'confirm'])

const close = () => {
  emit('update:modelValue', false)
}

const confirm = () => {
  emit('confirm')
  close()
}

const closeOnOverlay = (e) => {
  if (e.target === e.currentTarget) {
    close()
  }
}
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>