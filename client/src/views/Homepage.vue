<script setup>
import { useStore } from 'vuex'
import NavBar from './NavBar.vue'

const store = useStore()
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <NavBar />
    <div class="container mx-auto mt-8 px-4">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8">
        <h1 class="text-4xl font-bold mb-4 text-gray-900 dark:text-white">Welcome to the Home Page, {{ store.state.userDetails.firstname }}</h1>
      </div>
    </div>
  </div>
</template>