import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

export const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Add auth header to all requests
axiosInstance.interceptors.request.use(config => {
  const storedToken = localStorage.getItem('token')
  if (storedToken) {
    config.headers.Authorization = `Bearer ${storedToken}`
  }
  return config
})

interface User {
  id: string
  nickname?: string
  avatar_data?: string
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
      token.value = storedToken
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
      const response = await axiosInstance.post('/login', { user_id: id, password })

      const data = response.data
      console.log('Login response:', data)
      
      if (!data.token || !data.user) {
        console.error('Invalid login response:', data)
        return false
      }

      const tokenValue = data.token
      token.value = tokenValue
      user.value = data.user
      localStorage.setItem('token', tokenValue)
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
      const response = await axiosInstance.put('/profile', { nickname })

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

  const uploadAvatar = async (fileId: string, imageData: string): Promise<boolean> => {
    try {
      const response = await axiosInstance.post('/profile/avatar', {
        file_id: fileId,
        image_data: imageData
      })

      const data = response.data

      if (user.value && data.avatar_data) {
        const updatedUser: User = { ...user.value, avatar_data: data.avatar_data }
        user.value = updatedUser
        localStorage.setItem('user', JSON.stringify(updatedUser))
        return true
      }
      return false
    } catch (error) {
      console.error('Error uploading avatar:', error)
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
    uploadAvatar,
  }
})
