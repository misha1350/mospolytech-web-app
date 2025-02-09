// store/index.js
import { createStore } from 'vuex'

export default createStore({
  state: {
    userDetails: null,
    isDarkMode: true // Default to dark mode
  },
  mutations: {
    setUserDetails(state, userDetails) {
      state.userDetails = userDetails
    },
    toggleDarkMode(state) {
      state.isDarkMode = !state.isDarkMode
      // Update class on html element
      if (state.isDarkMode) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      // Save preference to localStorage
      localStorage.setItem('darkMode', state.isDarkMode)
    },
    initializeDarkMode(state) {
      // Check localStorage first, otherwise use default (true)
      const darkMode = localStorage.getItem('darkMode')
      state.isDarkMode = darkMode !== null ? JSON.parse(darkMode) : true
      // Don't manipulate DOM here since main.js handles initial state
    }
  }
})
