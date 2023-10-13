<script setup>
const props = defineProps({
  show: Boolean
})
</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask">
      <div class="modal-container">
        <div class="modal-header">
          <slot name="header">Edit Role</slot>
        </div>

        <div class="modal-body">
          <form @submit.prevent="submitForm">
            <div class="field">
              <label class="label">Email address:</label>
              <div class="control">
                <input v-model="user.email" type="email" placeholder="Enter Email" name="email" required>
              </div>
            </div>

            <div class="field">
              <label class="label">First name:</label>
              <div class="control">
                <input v-model="user.firstname" type="text" placeholder="Enter First Name" name="firstname" required>
              </div>
            </div>

            <div class="field">
              <label class="label">Last name:</label>
              <div class="control">
                <input v-model="user.lastname" type="text" placeholder="Enter Last Name" name="lastname" required>
              </div>
            </div>

            <div class="field">
              <label class="label">Office:</label>
              <div class="control">
                <select v-model="user.office" name="office">
                  <option value="">Select your office</option>
                  <!-- Add more options as needed -->
                  <option value="1">Office 1</option>
                  <option value="2">Office 2</option>
                  <!-- ... -->
                </select>
              </div>
            </div>

            <div class="field">
              <label class="label">Role:</label>
              <div class="control">
                <input type="radio" id="user" value="1" v-model="user.role" name="role">
                <label for="user">User</label><br>
                <input type="radio" id="admin" value="2" v-model="user.role" name="role">
                <label for="admin">Administrator</label><br>
              </div>
            </div>
            <div class="modal-footer">
              <slot name="footer">
                <button type="submit">Apply</button>
                <button
                  class="modal-default-button"
                  @click="$emit('close')"
                >OK</button>
              </slot>
            </div>
          </form>
          <div v-if="message" class="message">{{ message }}</div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script>
export default {
  data() {
    return {
      user: {
        email: '',
        firstname: '',
        lastname: '',
        office: '',
        role: '1',
      },
      processing: false,
      message: '',
    };
  },
  methods: {
    async submitForm() {
      this.processing = true;
      this.message = 'Processing request...';

      const response = await fetch('/api/server/user_edit', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.user),
      });

      if (response.ok) {
        this.message = 'Successfully updated user information.';
      } else if (response.status === 404) {
        this.message = 'Cannot establish connection with the backend. (HTTP return code 404)';
      } else {
        this.message = `Cannot edit user data in the database. (HTTP return code ${response.status})`;
      }

      this.processing = false;
    },
  },
};
</script>


<style>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  transition: opacity 0.3s ease;
}

.modal-container {
  width: 300px;
  margin: auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

.modal-header h2 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter-from {
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>