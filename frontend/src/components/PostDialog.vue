<template>
  <q-dialog v-model="modelValue">
    <q-card class="bg-black" style="min-width: 500px">
      <q-card-section>
        <div class="text-h6">新規投稿</div>
      </q-card-section>

      <q-card-section>
        <q-input
          v-model="content"
          type="textarea"
          class="bg-black"
          dark
          autofocus
          placeholder="いまどうしてる？"
        />
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat label="キャンセル" v-close-popup />
        <q-btn color="primary" label="投稿する" @click="handlePost" :loading="posting" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { usePostsStore } from '../stores/posts'

const props = defineProps<{
  modelValue: boolean
}>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const postsStore = usePostsStore()
const content = ref('')
const posting = ref(false)

const handlePost = async () => {
  if (!content.value.trim()) return
  
  posting.value = true
  try {
    await postsStore.createPost(content.value)
    emit('update:modelValue', false)
    content.value = ''
  } catch (error) {
    console.error('Error creating post:', error)
  } finally {
    posting.value = false
  }
}
</script>
