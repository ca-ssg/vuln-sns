import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

interface User {
  id: string
  nickname?: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(JSON.parse(localStorage.getItem('user') || 'null'))
  const isAuthenticated = computed(() => !!token.value)

  const login = async (id: string, password: string): Promise<boolean> => {
    try {
      const response = await fetch('http://localhost:9090/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id, password }),
      })

      if (!response.ok) {
        const errorData = await response.json()
        console.error('Login failed:', errorData)
        return false
      }

      const data = await response.json()
      console.log('Login response:', data)
      
      token.value = data.token
      user.value = data.user
      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))
      return true
    } catch (error) {
      console.error('Login error:', error)
      return false
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    logout,
  }
})
