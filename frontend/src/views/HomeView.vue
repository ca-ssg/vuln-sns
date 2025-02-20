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

<script setup>
import { onMounted } from 'vue'
import { usePostsStore } from '../stores/posts'
import PostCard from '../components/PostCard.vue'

const postsStore = usePostsStore()

onMounted(async () => {
  await postsStore.fetchPosts()
})
</script>

<style scoped>
.home-container {
  max-width: 600px;
  margin: 0 auto;
}
</style>
