<template>
  <div>
    <div v-if="authStore.user" class="mb-8">
      <textarea
        v-model="newPostContent"
        class="w-full p-4 border rounded-lg"
        rows="3"
        placeholder="投稿内容を入力..."
      ></textarea>
      <div class="flex justify-end mt-2">
        <button
          @click="createPost"
          class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
        >
          投稿する
        </button>
      </div>
    </div>

    <div class="space-y-4">
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
