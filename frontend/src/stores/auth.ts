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
      const parsedUser = JSON.parse(storedUser)
      // Ensure token has userID and proper format
      const cleanToken = storedToken.startsWith('Bearer ') ? storedToken.substring(7) : storedToken
      token.value = cleanToken.includes(parsedUser.id) ? cleanToken : `${parsedUser.id}_token`
      user.value = parsedUser
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
        body: JSON.stringify({ user_id: id, password }),
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

      // Store token in userID_token format
      const tokenValue = data.token.endsWith('_token') ? data.token : `${data.user.id}_token`
      // Ensure token includes userID
      token.value = tokenValue.includes(data.user.id) ? tokenValue : `${data.user.id}_token`
      user.value = data.user
      localStorage.setItem('token', token.value)
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
          'Authorization': `Bearer ${token.value || ''}`
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
