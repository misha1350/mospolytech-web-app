<script setup>
import { onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const store = useStore()
const router = useRouter()

onMounted(async () => {
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
  <div>
    <router-view></router-view>
  </div>
</template>

<!--TODO: Create buttons and CSS in the header for better navigation across the frontend-->
<!--TODO: Implement functionality to display SELECT query results-->
<!-- <template>
  <h1>Application</h1>
  <h3 v-if="returnCode">Return code: {{ returnCode }}</h3>
  <h3 v-if="errorMessage">{{ errorMessage }}</h3>
  <h3 v-if="userDetails">Welcome, {{ userDetails.firstname }} {{ userDetails.lastname }} from office no. {{ userDetails.office }}</h3>
  
  <button v-if="!userDetails" @click="login">Log In</button>
  <button v-if="userDetails" @click="endSession">End Session</button>
  <button id="show-modal" @click="showModal = true">Edit Role</button>

  <Teleport to="body">
    <modal :show="showModal" @close="showModal = false">
      <template #header>
        <h2>Edit Role for User</h2>
      </template>
    </modal>
  </Teleport>
</template> -->