import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

// API URL for the backend
const apiUrl = '/api'

export const axiosInstance = axios.create({
  baseURL: apiUrl,
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

export const useAuthStore = defineStore('auth', () => {
  const user = ref<string | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  const login = async (id: string, password: string): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      console.log('Attempting login with:', { id })
      
      const response = await axiosInstance.post('/login', { 
        user_id: id, 
        password: password 
      })
      console.log('Login response:', response.data)
      
      token.value = response.data.token
      user.value = id
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
    } else {
      // Set guest token for anonymous access
      token.value = 'guest_token'
      user.value = 'guest'
      axiosInstance.defaults.headers.common['Authorization'] = `Bearer guest_token`
    }
  }

  // Initialize auth on store creation
  initAuth()

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
