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
import { onMounted, watch } from 'vue'
import { usePostsStore } from '../stores/posts'
import { useRoute, useRouter } from 'vue-router'
import PostCard from '../components/PostCard.vue'

const props = defineProps<{
  tag?: string
}>()

const postsStore = usePostsStore()
const route = useRoute()
const router = useRouter()

const handleHashtagClick = (tag: string) => {
  router.push({ path: '/search', query: { tag: tag.slice(1) } })
}

// Watch for route query changes to update posts
watch(() => route.query.tag, async (newTag) => {
  if (newTag) {
    await postsStore.searchByHashtag(newTag as string)
  } else {
    await postsStore.fetchPosts()
  }
}, { immediate: true })

// Watch for route query changes to update posts
watch(() => route.query.tag, async (newTag) => {
  if (newTag) {
    await postsStore.searchByHashtag(newTag as string)
  } else {
    await postsStore.fetchPosts()
  }
}, { immediate: true })
</script>

<style scoped>
.home-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 0 16px;
}
</style>
