<!--
https://eugenkiss.github.io/7guis/tasks/#flight
-->

<script setup>
import Modal from './Modal.vue'
import { onMounted, ref } from 'vue'

const showModal = ref(false)
const returnCode = ref(null)
const errorMessage = ref(null)
const userDetails = ref(null)
onMounted(async () => {
  // get cookie value
  const cookie = document.cookie
    .split('; ')
    .find(row => row.startsWith('Authorization='));

  if (!cookie) {
    returnCode.value = 401
    errorMessage.value = 'There is no Authorization cookie. Try logging in again.'
    return
  }

  const cookieValue = cookie.split('=')[1];

  // TODO: Prevent localStorage from being rewritten after logout
  localStorage.setItem('Authorization', cookieValue);
  const token = localStorage.getItem('Authorization')
  if (!token) {
    returnCode.value = 401
    errorMessage.value = 'There is no JWT token in localStorage. Try logging in again.'
    return
  }

  const response = await fetch('/api/server/check', {
    method: 'POST',
    headers: { 'Authorization': `${token}` },
  })
  const body = await response.json()

  returnCode.value = response.status
  if (response.status === 200) {
    userDetails.value = body.userDetails
  } else {
    errorMessage.value = body.error.message
  }
})

function endSession() {
  localStorage.removeItem('Authorization')
  // Clear user details
  userDetails.value = null
}

function logIn() {
  // Redirect to the backend's login page
  window.location.href = '/api/server/login'
}
</script>

<!--TODO: Create buttons and CSS in the header for better navigation across the frontend-->
<template>
  <h1>Application</h1>
  <h3 v-if="returnCode">Return code: {{ returnCode }}</h3>
  <h3 v-if="errorMessage">{{ errorMessage }}</h3>
  <h3 v-if="userDetails">Welcome, {{ userDetails.firstname }} {{ userDetails.lastname }} from office no. {{ userDetails.office }}</h3>
  
  <button v-if="!userDetails" @click="logIn">Log In</button>
  <button v-if="userDetails" @click="endSession">End Session</button>
  <button id="show-modal" @click="showModal = true">Edit Role</button>

  <Teleport to="body">
    <!-- use the modal component, pass in the prop -->
    <modal :show="showModal" @close="showModal = false">
      <template #header>
        <h2>Edit Role for User</h2>
      </template>
    </modal>
  </Teleport>
</template>