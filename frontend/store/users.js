import { defineStore } from 'pinia'

export const useUsersStore = defineStore({
  id: 'users-store',
  state: () => {
    return {
      isLogged: true,
    }
  },
  actions: {    
    login() {
      this.isLogged = true
    },
    logout() {
      this.isLogged = false
    },
  },
  getters: {
    isLogged: (state) => state.isLogged,
  },
})