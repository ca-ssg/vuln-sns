<template>
  <q-card flat bordered class="post-card q-mb-md bg-black" style="border-color: #2F3336">
    <q-card-section>
      <div class="row items-center">
        <q-avatar size="40px" class="q-mr-md">
          <img :src="'https://ui-avatars.com/api/?name=' + post.userId" />
        </q-avatar>
        <div>
          <div class="text-weight-bold">{{ post.userId }}</div>
          <div class="text-caption text-grey">{{ formattedDate }}</div>
        </div>
      </div>
      <div class="q-mt-sm" v-html="post.content" @click="handleHashtagClick"></div>
      <div class="row q-mt-md">
        <q-btn flat round :color="post.isLiked ? 'pink' : 'grey'" icon="far fa-heart" @click="toggleLike">
          <div class="like-count q-ml-sm">{{ post.likes }}</div>
        </q-btn>
      </div>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { usePostsStore } from '../stores/posts'
import { useRouter } from 'vue-router'

const router = useRouter()
const props = defineProps<{
  post: {
    id: number
    userId: string
    content: string
    createdAt: string
    likes: number
  }
}>()

const postsStore = usePostsStore()

const formattedDate = computed(() => {
  const date = new Date(props.post.createdAt)
  if (isNaN(date.getTime())) return ''
  
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(minutes / 60)
  const days = Math.floor(hours / 24)

  if (minutes < 1) return '今'
  if (minutes < 60) return `${minutes}分`
  if (hours < 24) return `${hours}時間`
  if (days < 7) return `${days}日`

  return new Intl.DateTimeFormat('ja-JP', {
    month: 'short',
    day: 'numeric'
  }).format(date)
})

const toggleLike = async () => {
  try {
    if (props.post.isLiked) {
      await postsStore.unlikePost(props.post.id)
    } else {
      await postsStore.likePost(props.post.id)
    }
  } catch (error) {
    console.error('Error toggling like:', error)
    // Error is already handled by the store, no need to show UI error
  }
}

const handleHashtagClick = (event: MouseEvent) => {
  const target = event.target as HTMLElement
  if (target.tagName === 'A' && target.textContent?.startsWith('#')) {
    event.preventDefault()
    const tag = target.textContent.slice(1)
    router.push({ path: '/search', query: { tag } })
  }
}
</script>

<style scoped>
.post-card {
  transition: background-color 0.2s;
}
.post-card:hover {
  background-color: rgba(255, 255, 255, 0.03);
}
.like-count {
  font-size: 0.9rem;
  color: inherit;
  min-width: 1rem;
  text-align: left;
}
</style>
