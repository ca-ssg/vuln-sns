<template>
  <div class="q-pa-md">
    <div class="posts-container">
      <q-list separator>
        <q-item v-for="post in postStore.posts" :key="post.id" class="q-py-md post-item">
          <q-item-section>
            <q-item-label class="text-subtitle1 q-mb-sm">{{ post.content }}</q-item-label>
            <q-item-label caption>
              {{ formatDate(post.createdAt) }}
            </q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-btn 
              flat 
              round 
              color="primary" 
              icon="favorite_border"
              @click="postStore.toggleLike(post.id)"
            >
              <q-badge color="primary" floating>{{ post.likes }}</q-badge>
            </q-btn>
          </q-item-section>
        </q-item>
      </q-list>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { usePostStore } from '../stores/posts'
import PostCard from '../components/PostCard.vue'

const postStore = usePostStore()

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  if (isNaN(date.getTime())) {
    return '無効な日付'
  }
  return new Intl.DateTimeFormat('ja-JP', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false
  }).format(date)
}

onMounted(async () => {
  await postStore.fetchPosts()
})
</script>

<style>
.posts-container {
  max-width: 600px;
  margin: 0 auto;
}

.post-item {
  border-color: #2F3336 !important;
  transition: background-color 0.2s;
}

.post-item:hover {
  background: rgba(255, 255, 255, 0.03);
}

.q-item__section--side {
  padding-right: 8px;
}

.q-badge {
  font-size: 12px;
  padding: 2px 4px;
}
</style>
