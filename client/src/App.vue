<!--
https://eugenkiss.github.io/7guis/tasks/#flight
-->

<script setup>
import Modal from './Modal.vue'
import { onMounted, ref } from 'vue'

const showModal = ref(false)
const errorCode = ref(null)
const errorMessage = ref(null)
onMounted(async () => {
  // get cookie value
  const cookie = document.cookie
    .split('; ')
    .find(row => row.startsWith('Authorization='));

  if (!cookie) {
    errorCode.value = 401
    errorMessage.value = 'There is no Authorization cookie. Try logging in again.'
    return
  }

  const cookieValue = cookie.split('=')[1];

  // store cookie value in local storage
  localStorage.setItem('Authorization', cookieValue);
  const token = localStorage.getItem('Authorization')
  if (!token) {
    errorCode.value = 401
    errorMessage.value = 'There is no JWT token in localStorage. Try logging in again.'
    return
  }

  const response = await fetch('http://localhost:8086/api/server/check', {
    method: 'GET',
    headers: { 'Authorization': `${token}` },
  })
  const body = await response.json()

  errorCode.value = response.status
  errorMessage.value = body.error.message
})

</script>

<!--TODO: Create buttons and CSS in the header for better navigation across the frontend-->
<template>
  <h1>Application</h1>
  <h3 v-if="errorCode">Error code: {{ errorCode }}</h3>
  <h3 v-if="errorMessage">{{ errorMessage }}</h3>
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