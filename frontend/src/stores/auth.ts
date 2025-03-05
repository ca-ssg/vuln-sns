import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

// Hardcoded API URL for now
// In production, this would be set from environment variables
const apiUrl = '/api'

export const axiosInstance = axios.create({
  baseURL: apiUrl,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const useAuthStore = defineStore('auth', () => {
  const user = ref<string | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  const login = async (username: string, password: string): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      const response = await axiosInstance.post('/login', { username, password })
      token.value = response.data.token
      user.value = username
      if (token.value) localStorage.setItem('token', token.value)
      if (user.value) localStorage.setItem('user', user.value)
      axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token.value}`
    } catch (err) {
      console.error('Login error:', err)
      error.value = 'Invalid username or password'
      token.value = null
      user.value = null
    } finally {
      loading.value = false
    }
  }

  const logout = (): void => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    delete axiosInstance.defaults.headers.common['Authorization']
  }

  const initAuth = (): void => {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    if (storedToken && storedUser) {
      token.value = storedToken
      user.value = storedUser
      axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token.value}`
    }
  }

  return {
    user,
    token,
    loading,
    error,
    isAuthenticated,
    login,
    logout,
    initAuth
  }
})
