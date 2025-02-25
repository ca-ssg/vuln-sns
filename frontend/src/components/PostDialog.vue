<template>
  <base-dialog
    v-model="isOpen"
    title="新規投稿"
    @close="closeDialog"
  >
    <template #content>
      <q-input
        v-model="content"
        type="textarea"
        class="q-mb-md"
        dark
        outlined
        autofocus
        placeholder="いまどうしてる？"
        :rules="[val => val.length <= 140 || '140文字以内で入力してください']"
        maxlength="140"
        counter
      />
    </template>
    <template #actions>
      <q-btn 
        color="primary" 
        class="q-px-md"
        label="投稿する" 
        @click="handlePost" 
        :loading="posting"
        :disable="!content.trim() || content.length > 140"
        rounded
      />
    </template>
  </base-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { usePostsStore } from '../stores/posts'
import BaseDialog from './BaseDialog.vue'

const props = defineProps<{
  modelValue: boolean
}>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const postsStore = usePostsStore()
const content = ref('')
const posting = ref(false)

const isOpen = ref(false)

watch(() => props.modelValue, (val) => {
  isOpen.value = val
})

watch(isOpen, (val) => {
  emit('update:modelValue', val)
})

const closeDialog = () => {
  content.value = ''
  emit('update:modelValue', false)
}

const handlePost = async () => {
  if (!content.value.trim() || content.value.length > 140) return
  
  posting.value = true
  try {
    await postsStore.createPost(content.value)
    closeDialog()
  } catch (error) {
    console.error('Error creating post:', error)
  } finally {
    posting.value = false
  }
}
</script>
