<template>
  <div class="border-b border-gray-800 p-4 hover:bg-gray-900/50">
    <div class="flex space-x-4">
      <!-- アバター -->
      <div class="avatar">
        <i class="fas fa-user avatar-icon"></i>
      </div>

      <!-- 投稿内容 -->
      <div class="flex-1">
        <div class="flex justify-between items-start">
          <div class="flex items-center space-x-2">
            <span class="font-bold">{{ post.userId }}</span>
            <span class="text-gray-500">・{{ formatDate(post.createdAt) }}</span>
          </div>
          <div v-if="isOwner" class="flex space-x-2">
            <button
              v-if="!isEditing"
              @click="startEdit"
              class="text-gray-500 hover:text-blue-400"
            >
              <i class="fas fa-edit"></i>
            </button>
            <button
              @click="deletePost"
              class="text-gray-500 hover:text-red-400"
            >
              <i class="fas fa-trash"></i>
            </button>
          </div>
        </div>

        <div v-if="isEditing" class="mt-2">
          <textarea
            v-model="editContent"
            class="w-full p-2 bg-transparent border border-gray-800 rounded-lg focus:border-blue-400 focus:ring-0 text-white"
            rows="3"
          ></textarea>
          <div class="flex justify-end space-x-2 mt-2">
            <button
              @click="cancelEdit"
              class="btn-secondary text-sm"
            >
              キャンセル
            </button>
            <button
              @click="saveEdit"
              class="btn-primary text-sm"
            >
              保存
            </button>
          </div>
        </div>
        <!-- 意図的な脆弱性: XSS - v-htmlによる生のHTML表示 -->
        <!-- Vulnerability: XSS through v-html directive -->
        <div v-else class="mt-1" v-html="post.content"></div>

        <!-- アクションボタン -->
        <div class="flex items-center space-x-8 mt-3">
          <button
            @click="toggleLike"
            class="flex items-center space-x-2 text-gray-500 hover:text-pink-500 group"
          >
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
  if (confirm('本当に削除しますか？')) {
    await postStore.deletePost(props.post.id)
  }
}

const toggleLike = () => {
  if (authStore.token) {
    postStore.toggleLike(props.post.id)
  }
}
</script>
