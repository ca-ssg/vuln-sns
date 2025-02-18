<template>
  <q-card flat bordered class="post-card q-mb-md bg-black" style="border-color: #2F3336">
    <q-card-section>
      <div class="row items-center">
        <q-avatar size="48px">
          <img :src="'https://ui-avatars.com/api/?name=' + post.user_id" />
        </q-avatar>
        <div class="q-ml-md">
          <div class="text-weight-bold">{{ post.user_id }}</div>
          <div class="text-grey-6">{{ formatDate(post.created_at) }}</div>
        </div>
      </div>
      
      <div class="q-mt-md" v-html="post.content"></div>

      <div class="row q-mt-lg justify-between items-center">
        <q-btn flat round color="grey-6" icon="far fa-comment">
          <q-tooltip>返信</q-tooltip>
        </q-btn>
        
        <q-btn flat round color="grey-6" icon="fas fa-retweet">
          <q-tooltip>リツイート</q-tooltip>
        </q-btn>
        
        <q-btn
          flat
          round
          :color="isLiked ? 'pink-6' : 'grey-6'"
          :icon="isLiked ? 'fas fa-heart' : 'far fa-heart'"
          @click="toggleLike"
        >
          <q-tooltip>いいね</q-tooltip>
        </q-btn>
        
        <q-btn flat round color="grey-6" icon="far fa-share-square">
          <q-tooltip>共有</q-tooltip>
        </q-btn>
      </div>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '../stores/auth'

const props = defineProps<{
  post: {
    id: number
    user_id: string
    content: string
    created_at: string
  }
}>()

const authStore = useAuthStore()
const isLiked = ref(false)
const likesCount = ref(0)

const toggleLike = async () => {
  if (!authStore.isAuthenticated) return

  try {
    const response = await fetch(`http://localhost:9090/api/posts/${props.post.id}/like`, {
      method: 'POST',
      headers: {
        'Authorization': authStore.token || '',
      },
    })

    if (response.ok) {
      isLiked.value = !isLiked.value
      likesCount.value += isLiked.value ? 1 : -1
    }
  } catch (error) {
    console.error('Error toggling like:', error)
  }
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ja-JP')
}
</script>

<style scoped>
.post-card {
  transition: background-color 0.2s;
}
.post-card:hover {
  background-color: rgba(255, 255, 255, 0.03);
}
</style>
