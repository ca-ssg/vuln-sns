<template>
  <div class="border-b border-gray-800 p-4 hover:bg-gray-900/50">
    <div class="flex space-x-4">
      <div class="avatar w-12 h-12 rounded-full bg-gray-800 flex items-center justify-center">
        <i class="fas fa-user text-gray-500 text-xl"></i>
      </div>
      <div class="flex-1">
        <div class="flex justify-between items-start">
          <div class="flex items-center space-x-2">
            <span class="font-bold text-white">{{ post.userId }}</span>
            <span class="text-gray-500">・{{ formatDate(post.createdAt) }}</span>
          </div>
          <div v-if="isOwner" class="flex space-x-2">
            <button v-if="!isEditing" @click="startEdit" class="text-gray-500 hover:text-blue-400">
              <i class="fas fa-edit"></i>
            </button>
            <button @click="deletePost" class="text-gray-500 hover:text-red-400">
              <i class="fas fa-trash"></i>
            </button>
          </div>
        </div>
        <div v-if="isEditing" class="mt-2">
          <textarea v-model="editContent" class="w-full p-2 bg-transparent border border-gray-800 rounded-lg focus:border-blue-400 focus:ring-0 text-white" rows="3"></textarea>
          <div class="flex justify-end space-x-2 mt-2">
            <button @click="cancelEdit" class="btn-secondary text-sm">キャンセル</button>
            <button @click="saveEdit" class="btn-primary text-sm">保存</button>
          </div>
        </div>
        <div v-else class="mt-1 text-white" v-html="post.content"></div>
        <div class="flex items-center space-x-8 mt-3">
          <button @click="toggleLike" class="flex items-center space-x-2 text-gray-500 hover:text-pink-500 group">
            <i class="fas fa-heart" :class="{ 'text-pink-500': post.isLiked }"></i>
            <span>{{ post.likes }}</span>
          </button>
          <button class="flex items-center space-x-2 text-gray-500 hover:text-green-400 group">
            <i class="fas fa-retweet"></i>
            <span>0</span>
          </button>
          <button class="flex items-center space-x-2 text-gray-500 hover:text-blue-400 group">
            <i class="fas fa-comment"></i>
            <span>0</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '../stores/auth'
import { usePostStore } from '../stores/posts'

const props = defineProps<{
  post: {
    id: number
    userId: string
    content: string
    createdAt: string
    likes: number
    isLiked?: boolean
  }
}>()

const authStore = useAuthStore()
const postStore = usePostStore()

const isOwner = computed(() => authStore.user?.id === props.post.userId)
const isEditing = ref(false)
const editContent = ref(props.post.content)

const startEdit = () => {
  isEditing.value = true
  editContent.value = props.post.content
}

const cancelEdit = () => {
  isEditing.value = false
}

const saveEdit = async () => {
  if (await postStore.updatePost(props.post.id, editContent.value)) {
    isEditing.value = false
  }
}

const deletePost = async () => {
  if (confirm('この投稿を削除しますか？')) {
    await postStore.deletePost(props.post.id)
  }
}

const toggleLike = () => {
  if (authStore.token) {
    postStore.toggleLike(props.post.id)
  }
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return date.toLocaleDateString('ja-JP')
}
</script>

<style scoped>
.btn-primary {
  @apply px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600;
}

.btn-secondary {
  @apply px-4 py-2 bg-gray-800 text-gray-300 rounded-lg hover:bg-gray-700;
}
</style>
