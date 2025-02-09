// store/index.js
import { createStore } from 'vuex'

export default createStore({
  state: {
    userDetails: null,
    isDarkMode: true
  },
  mutations: {
    setUserDetails(state, userDetails) {
      state.userDetails = userDetails
    },
    toggleDarkMode(state) {
      state.isDarkMode = !state.isDarkMode
      // Update class on html element
      document.documentElement.classList.toggle('dark', state.isDarkMode)
      // Save preference to localStorage
      localStorage.setItem('darkMode', String(state.isDarkMode))
    },
    initializeDarkMode(state) {
      // Check localStorage first, otherwise use default (true)
      const darkMode = localStorage.getItem('darkMode')
      state.isDarkMode = darkMode === null ? true : darkMode === 'true'
      // Update class to match state
      document.documentElement.classList.toggle('dark', state.isDarkMode)
    }
  },
  getters: {
    isAuthenticated: state => !!state.userDetails,
    isAdmin: state => state.userDetails?.role === '2',
    userFullName: state => state.userDetails ? `${state.userDetails.firstname} ${state.userDetails.lastname}` : ''
  }
})
