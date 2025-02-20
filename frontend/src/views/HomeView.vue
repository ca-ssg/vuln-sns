<template>
  <div class="home-container">
    <div v-if="postsStore.loading" class="text-center q-pa-md">
      <q-spinner color="primary" size="3em" />
    </div>
    <div v-else>
      <post-card
        v-for="post in postsStore.posts"
        :key="post.id"
        :post="post"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { usePostsStore } from '../stores/posts'
import PostCard from '../components/PostCard.vue'
import { useRoute } from 'vue-router'
import { watch } from 'vue'

const postsStore = usePostsStore()
const route = useRoute()

onMounted(async () => {
  await postsStore.fetchPosts()
})

// Refresh posts when returning to home
watch(() => route.path, async (newPath) => {
  if (newPath === '/') {
    await postsStore.fetchPosts()
  }
})
</script>

<style scoped>
.home-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 0 16px;
}
</style>
