<template>
  <div class="min-h-screen bg-black">
    <div class="max-w-2xl mx-auto border-x border-gray-800 min-h-screen">
      <div class="max-w-md mx-auto py-12 px-4">
        <div class="text-center mb-8">
          <h1 class="text-3xl font-bold mb-2 text-white">VulnAppにログイン</h1>
          <p class="text-gray-500">セキュリティの学習用アプリケーション</p>
        </div>
      
      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-400">ユーザーID</label>
          <div class="mt-1 relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-500">
              <i class="fas fa-user"></i>
            </span>
            <input
              v-model="id"
              type="text"
              class="block w-full pl-10 rounded-lg bg-gray-900 border-gray-800 text-white shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-400">パスワード</label>
          <div class="mt-1 relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-500">
              <i class="fas fa-lock"></i>
            </span>
            <input
              v-model="password"
              type="password"
              class="block w-full pl-10 rounded-lg bg-gray-900 border-gray-800 text-white shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>
        </div>

        <div v-if="error" class="text-red-500 text-center">
          {{ error }}
        </div>

        <button
          type="submit"
          class="w-full flex justify-center py-3 px-4 rounded-full text-white bg-blue-500 hover:bg-blue-600 font-bold text-lg"
        >
          ログイン
        </button>

        <p class="text-center text-sm text-gray-500">
          初期アカウントについてはREADMEをご確認ください
        </p>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const id = ref('')
const password = ref('')
const error = ref('')

const handleLogin = async () => {
  error.value = ''
  if (await authStore.login(id.value, password.value)) {
    router.push('/')
  } else {
    error.value = 'ログインに失敗しました'
  }
}
</script>
