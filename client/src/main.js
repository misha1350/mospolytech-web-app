import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import router from './routes.js'
import store from './store/index.js'

createApp(App).use(router).use(store).mount('#app')