<script setup>
import { onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const store = useStore()
const router = useRouter()

onMounted(async () => {
  // Initialize store state without DOM manipulation
  store.commit('initializeDarkMode')

  // get cookie value
  const cookie = document.cookie
    .split('; ')
    .find(row => row.startsWith('Authorization='));

  if (!cookie) {
    login()
    return
  }

  const cookieValue = cookie.split('=')[1];

  // TODO: Prevent localStorage from being rewritten after logout
  localStorage.setItem('Authorization', cookieValue);
  const token = localStorage.getItem('Authorization')
  if (!token) {
    login()
    return
  }

  const response = await fetch('/api/server/check', {
    method: 'POST',
    headers: { 'Authorization': `${token}` },
  })
  const body = await response.json()

  if (response.status === 200) {
    store.commit('setUserDetails', body.userDetails)
    router.push('/client/home')
  } else {
    login()
  }
})

function login() {
  // Redirect to the backend's login page
  window.location.href = '/api/server/login'
}
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <router-view></router-view>
  </div>
</template>