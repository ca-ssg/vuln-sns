import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

interface User {
  id: string
  nickname: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)

  const login = async (id: string, password: string) => {
    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/auth/login`, { id, password })
      user.value = response.data.user
      token.value = response.data.token
      localStorage.setItem('token', response.data.token)
      return true
    } catch (error) {
      console.error('Login failed:', error)
      return false
    }
  }

  const logout = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
  }

  const updateNickname = async (nickname: string) => {
    if (!token.value) return false

    try {
      const response = await axios.put(
        `${import.meta.env.VITE_API_URL}/profile/nickname`,
        { nickname },
        {
          headers: { Authorization: `Bearer ${token.value}` }
        }
      )
      if (user.value) {
        user.value.nickname = nickname
      }
      return true
    } catch (error) {
      console.error('Failed to update nickname:', error)
      return false
    }
  }

  // 保存されたトークンからユーザー情報を復元
  const initializeAuth = () => {
    const savedToken = localStorage.getItem('token')
    if (savedToken) {
      token.value = savedToken
      // TODO: トークンの検証とユーザー情報の取得
    }
  }

  return {
    user,
    token,
    login,
    logout,
    updateNickname,
    initializeAuth
  }
})
