<template>
  <div class="min-h-screen bg-black text-white">
    <div class="flex">
      <!-- Left Sidebar -->
      <div class="w-64 fixed h-screen border-r border-gray-800">
        <div class="p-4">
          <div class="text-xl font-bold mb-8">
            <i class="fas fa-feather text-blue-400 text-2xl"></i>
          </div>
          <nav class="space-y-4">
            <a href="/" class="flex items-center space-x-4 text-xl hover:bg-gray-900 p-3 rounded-full">
              <i class="fas fa-home"></i>
              <span>ホーム</span>
            </a>
            <a href="/search" class="flex items-center space-x-4 text-xl hover:bg-gray-900 p-3 rounded-full">
              <i class="fas fa-search"></i>
              <span>話題を検索</span>
            </a>
            <a href="/notifications" class="flex items-center space-x-4 text-xl hover:bg-gray-900 p-3 rounded-full">
              <i class="fas fa-bell"></i>
              <span>通知</span>
            </a>
            <a href="/messages" class="flex items-center space-x-4 text-xl hover:bg-gray-900 p-3 rounded-full">
              <i class="fas fa-envelope"></i>
              <span>メッセージ</span>
            </a>
            <a href="/profile" class="flex items-center space-x-4 text-xl hover:bg-gray-900 p-3 rounded-full">
              <i class="fas fa-user"></i>
              <span>プロフィール</span>
            </a>
          </nav>
          <button class="w-full bg-blue-500 hover:bg-blue-600 text-white rounded-full py-3 mt-4 font-bold">
            投稿する
          </button>
        </div>
      </div>

      <!-- Main Content -->
      <div class="flex-1 ml-64">
        <header class="sticky top-0 z-50 bg-black bg-opacity-70 backdrop-blur-md border-b border-gray-800">
          <div class="flex justify-between items-center px-4 py-3">
            <h1 class="text-xl font-bold">ホーム</h1>
            <div v-if="!isAuthenticated" class="flex space-x-4">
              <a href="/login" class="text-blue-400 hover:text-blue-500">ログイン</a>
            </div>
            <div v-else class="flex items-center space-x-4">
              <button @click="logout" class="text-red-400 hover:text-red-500">ログアウト</button>
            </div>
          </div>
        </header>
        <main class="max-w-2xl mx-auto">
          <router-view></router-view>
        </main>
      </div>

      <!-- Right Sidebar -->
      <div class="w-80 fixed right-0 h-screen border-l border-gray-800 p-4">
        <div class="bg-gray-900 rounded-2xl p-4 mb-4">
          <h2 class="text-xl font-bold mb-4">トレンド</h2>
          <div class="space-y-4">
            <div class="hover:bg-gray-800 p-2 rounded-lg cursor-pointer">
              <p class="text-gray-500 text-sm">日本のトレンド</p>
              <p class="font-bold">#セキュリティ</p>
              <p class="text-gray-500 text-sm">1,234 投稿</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from './stores/auth'
import { storeToRefs } from 'pinia'

const authStore = useAuthStore()
const { isAuthenticated } = storeToRefs(authStore)
const { logout } = authStore
</script>
