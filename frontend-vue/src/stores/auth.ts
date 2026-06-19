import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const registeredUsername = ref<string | null>(null)
  const justRegistered = ref(false)

  function setRegistered(username: string) {
    registeredUsername.value = username
    justRegistered.value = true
  }

  function clearRegistered() {
    registeredUsername.value = null
    justRegistered.value = false
  }

  return { registeredUsername, justRegistered, setRegistered, clearRegistered }
})
