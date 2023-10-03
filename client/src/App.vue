<!--
https://eugenkiss.github.io/7guis/tasks/#flight
-->

<script setup>
import Modal from './Modal.vue'
import { ref } from 'vue'

const showModal = ref(false)
</script>

<template>
  <button id="show-modal" @click="showModal = true">Show Modal</button>

  <Teleport to="body">
    <!-- use the modal component, pass in the prop -->
    <modal :show="showModal" @close="showModal = false">
      <template #header>
        <h2>Sign Up</h2>
      </template>
    </modal>
  </Teleport>
</template>

<!-- <style>
select,
input,
button {
  display: block;
  margin: 0.5em 0;
  font-size: 15px;
}

input[disabled] {
  color: #999;
}

p {
  color: red;
}
</style> -->