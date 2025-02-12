import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base: '/client', // necessary when a reverse proxy serves the client from a subpath other than '/'
  plugins: [vue()],
  server: {
    port: 8087,
    proxy: {
      '/api': {
        target: 'http://localhost:8086',
        // changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api\/server/, '')
      }
    }
  }
})
