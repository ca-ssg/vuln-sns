<template>
  <div class="max-w-md mx-auto">
    <h1 class="text-2xl font-bold mb-6">ログイン</h1>
    
    <form @submit.prevent="handleLogin" class="space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">ユーザーID</label>
        <input
          v-model="id"
          type="text"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">パスワード</label>
        <input
          v-model="password"
          type="password"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          required
        />
      </div>

      <div v-if="error" class="text-red-600">
        {{ error }}
      </div>

      <button
        type="submit"
        class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        ログイン
      </button>
    </form>
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
