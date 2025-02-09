import { createApp } from 'vue'
import App from './App.vue'
import router from './routes.js'
import store from './store/index.js'
import './index.css'

// Global error handler
const handleError = (error, instance, info) => {
  console.error('Global error:', error)
  console.error('Vue instance:', instance)
  console.error('Error info:', info)
  
  // Could add error reporting service here
  if (error.response?.status === 401) {
    // Handle authentication errors
    store.commit('setUserDetails', null)
    window.location.href = '/api/server/login'
  }
}

// Initialize dark mode before creating the app
const darkMode = localStorage.getItem('darkMode')
if (darkMode === 'true' || darkMode === null) { // null means default to dark mode
    document.documentElement.classList.add('dark')
}

// Create app
const app = createApp(App)

// Add error handler
app.config.errorHandler = handleError
window.onerror = (msg, src, line, col, error) => {
  handleError(error || msg)
}

// Add plugins
app.use(router)
app.use(store)

// Mount app
app.mount('#app')