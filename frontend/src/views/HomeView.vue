<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { usePostsStore } from '@/stores/posts'

const route = useRoute()
const postsStore = usePostsStore()

onMounted(async () => {
  await postsStore.fetchPosts()
})

// Watch for route query changes to update posts
watch(() => route.query.tag, async (newTag) => {
  if (newTag) {
    await postsStore.searchByHashtag(newTag as string)
  } else {
    await postsStore.fetchPosts()
  }
}, { immediate: true })
</script>

<template>
  <main>
    <div v-if="postsStore.loading" class="loading">
      Loading...
    </div>
    <div v-else-if="postsStore.error" class="error">
      {{ postsStore.error }}
    </div>
    <div v-else>
      <div v-for="post in postsStore.posts" :key="post.id" class="post">
        <div class="post-header">
          <div class="post-avatar">{{ post.user_id.substring(0, 2).toUpperCase() }}</div>
          <div class="post-user">{{ post.user_id }}</div>
          <div class="post-actions" v-if="post.user_id === 'alice'">
            <button @click="postsStore.showEditModal(post)">Edit</button>
            <button @click="postsStore.deletePost(post.id)">Delete</button>
          </div>
        </div>
        <div class="post-content">{{ post.content }}</div>
        <div class="post-footer">
          <button 
            class="like-button" 
            :class="{ 'liked': post.isLiked }"
            @click="post.isLiked ? postsStore.unlikePost(post.id) : postsStore.likePost(post.id)"
          >
            ❤️ {{ post.likes }}
          </button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.post {
  border: 1px solid #ccc;
  margin-bottom: 1rem;
  padding: 1rem;
}

.post-header {
  display: flex;
  align-items: center;
  margin-bottom: 0.5rem;
}

.post-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #ccc;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 0.5rem;
}

.post-user {
  font-weight: bold;
  flex-grow: 1;
}

.post-actions {
  display: flex;
  gap: 0.5rem;
}

.post-content {
  margin-bottom: 0.5rem;
}

.post-footer {
  display: flex;
  justify-content: flex-start;
}

.like-button {
  background: none;
  border: none;
  cursor: pointer;
  color: #888;
}

.like-button.liked {
  color: #ff69b4;
}

.loading, .error {
  padding: 1rem;
  text-align: center;
}

.error {
  color: red;
}
</style>
