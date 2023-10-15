<template>
  <nav class="bg-blue-800 py-4">
    <ul class="flex">
      <li>
        <router-link class="text-white hover:text-blue-300 mx-4" to="/client/home">Home</router-link>
      </li>
      <li v-if="isAdmin">
        <router-link class="text-white hover:text-blue-300 mx-4" to="/client/admin/get_users">Display Users</router-link>
      </li>
      <li class="ml-auto pr-4">
        <span class="text-white">{{store.state.userDetails.firstname}} {{store.state.userDetails.lastname}} ({{userRole}})</span>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'

const store = useStore()

const isAdmin = computed(() => store.state.userDetails.role == '1')
const userRole = computed(() => store.state.userDetails.role == '1' ? 'User' : 'Administrator')
</script>
