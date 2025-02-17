<template>
  <div class="bg-white rounded-lg shadow p-4 mb-4">
    <div class="flex justify-between items-start mb-2">
      <div>
        <span class="font-bold">{{ post.userId }}</span>
        <span class="text-gray-500 text-sm ml-2">
          {{ new Date(post.createdAt).toLocaleString() }}
        </span>
      </div>
      <div v-if="isOwner" class="flex space-x-2">
        <button
          @click="startEdit"
          class="text-blue-600 hover:text-blue-800"
        >
          編集
        </button>
        <button
          @click="deletePost"
          class="text-red-600 hover:text-red-800"
        >
          削除
        </button>
      </div>
    </div>

    <div v-if="isEditing" class="mb-4">
      <textarea
        v-model="editContent"
        class="w-full p-2 border rounded"
        rows="3"
      ></textarea>
      <div class="flex justify-end space-x-2 mt-2">
        <button
          @click="cancelEdit"
          class="px-3 py-1 text-gray-600 hover:text-gray-800"
        >
          キャンセル
        </button>
        <button
          @click="saveEdit"
          class="px-3 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          保存
        </button>
      </div>
    </div>
    <!-- 意図的な脆弱性: XSS - v-htmlによる生のHTML表示 -->
    <!-- Vulnerability: XSS through v-html directive -->
    <div v-else class="mb-4" v-html="post.content"></div>

    <div class="flex items-center">
      <button
        @click="toggleLike"
        class="flex items-center space-x-1"
        :class="{ 'text-red-500': post.isLiked, 'text-gray-500': !post.isLiked }"
      >
        <span class="text-xl">♡</span>
        <span>{{ post.likes }}</span>
      </button>
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
