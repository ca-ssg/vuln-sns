<template>
  <div class="min-h-screen bg-black">
    <header class="bg-black border-b border-gray-800">
      <nav class="max-w-7xl mx-auto px-4">
        <div class="flex justify-between h-16">
          <div class="flex">
            <router-link to="/" class="flex items-center text-white font-bold text-xl">
              VulnApp
            </router-link>
          </div>
          <div class="flex items-center">
            <template v-if="authStore.user">
              <router-link to="/profile" class="text-gray-300 hover:text-white px-3 py-2">
                プロフィール
              </router-link>
              <button @click="logout" class="text-gray-300 hover:text-white px-3 py-2">
                ログアウト
              </button>
            </template>
            <template v-else>
              <router-link to="/login" class="text-gray-300 hover:text-white px-3 py-2">
                ログイン
              </router-link>
            </template>
          </div>
        </div>
      </nav>
    </header>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from './stores/auth'

const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
}
</script>
