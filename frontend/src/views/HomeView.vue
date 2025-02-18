<template>
  <div class="q-pa-md">
    <div class="posts-container">
      <post-card v-for="post in postStore.posts" :key="post.id" :post="post" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { usePostStore } from '../stores/posts'
import PostCard from '../components/PostCard.vue'

const postStore = usePostStore()

onMounted(async () => {
  await postStore.fetchPosts()
})
</script>
