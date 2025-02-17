<template>
  <div class="max-w-2xl mx-auto border-x border-gray-800 min-h-screen">
    <!-- 新規投稿フォーム -->
    <div v-if="authStore.user" class="p-4 border-b border-gray-800">
      <div class="flex space-x-4">
        <div class="w-12 h-12 rounded-full bg-gray-700 flex items-center justify-center">
          <i class="fas fa-user text-gray-400"></i>
        </div>
        <div class="flex-1">
          <textarea
            v-model="newPostContent"
            class="w-full p-2 bg-transparent border-b border-gray-800 focus:border-blue-400 focus:ring-0 resize-none text-white"
            rows="3"
            placeholder="いまどうしてる？"
          ></textarea>
          <div class="flex justify-end mt-2">
            <button
              @click="createPost"
              class="px-6 py-2 bg-blue-500 text-white rounded-full hover:bg-blue-600 font-bold"
              :disabled="!newPostContent.trim()"
            >
              ツイートする
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 投稿一覧 -->
    <div>
      <PostCard
        v-for="post in postStore.posts"
        :key="post.id"
        :post="post"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { usePostStore } from '../stores/posts'
import PostCard from '../components/PostCard.vue'

const authStore = useAuthStore()
const postStore = usePostStore()
const newPostContent = ref('')

const createPost = async () => {
  if (newPostContent.value.trim()) {
    if (await postStore.createPost(newPostContent.value)) {
      newPostContent.value = ''
    }
  }
}

onMounted(() => {
  postStore.fetchPosts()
})
</script>
