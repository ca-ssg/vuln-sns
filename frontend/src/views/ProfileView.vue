<template>
  <div class="max-w-2xl mx-auto border-x border-gray-800 min-h-screen">
    <!-- プロフィールヘッダー -->
    <div class="relative">
      <div class="h-48 bg-gradient-to-r from-blue-900 to-blue-800"></div>
      <div class="absolute -bottom-16 left-4">
        <div class="w-32 h-32 rounded-full bg-gray-700 border-4 border-black flex items-center justify-center overflow-hidden">
          <i class="fas fa-user text-4xl text-gray-400"></i>
        </div>
      </div>
    </div>
    
    <!-- プロフィール情報 -->
    <div class="px-4 pt-20 pb-4 border-b border-gray-800">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-2xl font-bold">{{ authStore.user?.nickname || authStore.user?.id }}</h1>
          <p class="text-gray-500">@{{ authStore.user?.id }}</p>
        </div>
        <button class="btn-secondary">
          プロフィールを編集
        </button>
      </div>
      
      <div class="mt-4 flex space-x-6 text-gray-500">
        <div class="flex items-center space-x-1">
          <i class="fas fa-calendar"></i>
          <span>2025年2月から利用</span>
        </div>
      </div>
    </div>

    <!-- プロフィール編集フォーム -->
    <div class="p-4">
      <div class="max-w-md">
        <h2 class="text-xl font-bold mb-4">プロフィール設定</h2>
        <form @submit.prevent="updateProfile" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-400">ニックネーム</label>
            <input
              v-model="nickname"
              type="text"
              class="mt-1 block w-full rounded-lg bg-gray-900 border-gray-800 text-white shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>

          <div v-if="message" class="text-green-500">
            {{ message }}
          </div>

          <button
            type="submit"
            class="btn-primary w-full"
          >
            保存
          </button>
        </form>
      </div>
    </div>
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
