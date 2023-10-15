<script setup>
import { useStore } from 'vuex'
import NavBar from './NavBar.vue'

const store = useStore()
</script>

<template>
  <div id="app">
      <NavBar />
      <div class="container mx-auto mt-8 px-4">
          <h1 class="text-4xl font-bold mb-4">Welcome to the Home Page, {{ store.state.userDetails.firstname }}</h1>
      </div>
  </div>
</template>