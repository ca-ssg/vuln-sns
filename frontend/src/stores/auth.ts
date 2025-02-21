import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_URL + '/api',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Basic ' + btoa('user:3d0b26c76947dc404912e2110babeac0')
  }
})

const API_URL = import.meta.env.VITE_API_URL

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
      const response = await axiosInstance.post('/api/login', { user_id: id, password })

      const data = response.data
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
      if (token.value) {
        localStorage.setItem('token', token.value)
      }
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
      const response = await axios.put(`${API_URL}/profile`, 
        { nickname },
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token.value || ''}`
          }
        })

      const data = response.data

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
