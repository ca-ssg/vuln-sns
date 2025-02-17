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
  <header class="bg-white border-b">
    <div class="container mx-auto px-4 py-3 flex justify-between items-center">
      <router-link to="/" class="text-xl font-bold text-blue-500">VulnApp</router-link>
      <nav v-if="authStore.user">
        <router-link to="/profile" class="text-gray-600 hover:text-gray-900 mr-4">
          {{ authStore.user.nickname }}
        </router-link>
        <button @click="logout" class="text-red-600 hover:text-red-800">
          ログアウト
        </button>
      </nav>
      <router-link v-else to="/login" class="text-blue-600 hover:text-blue-800">
        ログイン
      </router-link>
    </div>
  </header>

  <main class="container mx-auto px-4 py-8">
    <router-view />
  </main>
</template>
