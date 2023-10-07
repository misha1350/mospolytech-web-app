import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base: '/client/', // necessary when a reverse proxy serves the client from a subpath other than '/'
  plugins: [vue()],
})
