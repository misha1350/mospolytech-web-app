import { createApp } from 'vue'
import App from './App.vue'
import router from './routes.js'
import store from './store/index.js'
import './index.css'

// Initialize dark mode before creating the app
const darkMode = localStorage.getItem('darkMode')
if (darkMode === 'true' || darkMode === null) { // null means default to dark mode
    document.documentElement.classList.add('dark')
}

const app = createApp(App)
app.use(router)
app.use(store)
app.mount('#app')