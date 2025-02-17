<template>
  <div class="max-w-md mx-auto">
    <h1 class="text-2xl font-bold mb-6">プロフィール設定</h1>
    
    <form @submit.prevent="updateProfile" class="space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">ニックネーム</label>
        <input
          v-model="nickname"
          type="text"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          required
        />
      </div>

      <div v-if="message" class="text-green-600">
        {{ message }}
      </div>

      <button
        type="submit"
        class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        更新
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const nickname = ref(authStore.user?.nickname || '')
const message = ref('')

const updateProfile = async () => {
  if (await authStore.updateNickname(nickname.value)) {
    message.value = 'プロフィールを更新しました'
  } else {
    message.value = '更新に失敗しました'
  }
}
</script>
