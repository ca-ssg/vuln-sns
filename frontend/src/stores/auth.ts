import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

interface User {
  id: string
  nickname?: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(null)
  const user = ref<User | null>(null)
  
  // Initialize from localStorage
  try {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    if (storedToken && storedUser) {
      // Ensure token has Bearer prefix
      token.value = storedToken.startsWith('Bearer ') ? storedToken : `Bearer ${storedToken}`
      user.value = JSON.parse(storedUser)
    }
  } catch (e) {
    console.error('Failed to parse stored user:', e)
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const login = async (id: string, password: string): Promise<boolean> => {
    try {
      console.log('Attempting login with:', { id, password })
      const response = await fetch(`${import.meta.env.VITE_API_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id, password }),
        credentials: 'include'
      })

      if (!response.ok) {
        console.error('Login failed:', await response.text())
        return false
      }

      const data = await response.json()
      console.log('Login response:', data)
      
      if (!data.token || !data.user) {
        console.error('Invalid login response:', data)
        return false
      }

      // Store token with Bearer prefix for consistency
      const bearerToken = `Bearer ${data.token}`
      token.value = bearerToken
      user.value = data.user
      localStorage.setItem('token', bearerToken)
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

  const updateNickname = async (nickname: string): Promise<boolean> => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/profile`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token.value || ''
        },
        body: JSON.stringify({ nickname }),
      })

      if (!response.ok) {
        console.error('Failed to update nickname:', await response.text())
        return false
      }

      if (user.value) {
        const updatedUser: User = { ...user.value, nickname }
        user.value = updatedUser
        localStorage.setItem('user', JSON.stringify(updatedUser))
        return true
      }
      return false
    } catch (error) {
      console.error('Error updating nickname:', error)
      return false
    }
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    logout,
    updateNickname,
  }
})
