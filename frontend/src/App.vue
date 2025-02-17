<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  router.push('/login')
}

onMounted(() => {
  authStore.initializeAuth()
})
</script>

<template>
  <div class="min-h-screen bg-black text-white">
    <!-- サイドバー -->
    <div class="fixed h-screen w-64 border-r border-gray-800 bg-black">
      <div class="p-4">
        <RouterLink to="/" class="text-xl font-bold text-white hover:text-blue-400">
          VulnApp
        </RouterLink>
      </div>
      
      <nav class="mt-8 space-y-1">
        <RouterLink to="/" class="nav-link">
          <i class="fas fa-home mr-4"></i>
          ホーム
        </RouterLink>
        <template v-if="authStore.user">
          <RouterLink to="/profile" class="nav-link">
            <i class="fas fa-user mr-4"></i>
            プロフィール
          </RouterLink>
          <button @click="logout" class="nav-link w-full text-left text-red-400 hover:text-red-500">
            <i class="fas fa-sign-out-alt mr-4"></i>
            ログアウト
          </button>
        </template>
        <template v-else>
          <RouterLink to="/login" class="nav-link">
            <i class="fas fa-sign-in-alt mr-4"></i>
            ログイン
          </RouterLink>
        </template>
      </nav>

      <div v-if="authStore.user" class="absolute bottom-4 left-4 right-4">
        <div class="flex items-center space-x-3 p-3 rounded-full hover:bg-gray-900 cursor-pointer">
          <div class="avatar">
            <i class="fas fa-user avatar-icon"></i>
          </div>
          <div class="flex-1 truncate">
            <div class="font-bold">{{ authStore.user.nickname || authStore.user.id }}</div>
            <div class="text-sm text-gray-500">@{{ authStore.user.id }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- メインコンテンツ -->
    <div class="ml-64">
      <RouterView />
    </div>
  </div>
</template>

<style>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css');
</style>
