<template>
  <q-dialog v-model="isOpen" persistent>
    <q-card class="bg-blue-grey-9 post-dialog" style="min-width: 500px">
      <q-card-section class="q-pb-none">
        <div class="row items-center">
          <q-btn flat round icon="close" @click="closeDialog" />
          <div class="text-h6 text-white q-ml-md">新規投稿</div>
        </div>
      </q-card-section>

      <q-card-section>
        <q-input
          v-model="content"
          type="textarea"
          class="q-mb-md bg-blue-grey-8"
          dark
          outlined
          autofocus
          placeholder="いまどうしてる？"
          :rules="[val => val.length <= 140 || '140文字以内で入力してください']"
          maxlength="140"
          counter
        />
      </q-card-section>

      <q-card-actions align="right" class="q-pa-md">
        <q-btn 
          color="primary" 
          class="full-width-with-margin"
          label="投稿する" 
          @click="handlePost" 
          :loading="posting"
          :disable="!content.trim() || content.length > 140"
          rounded
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
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

<style scoped>
:deep(.q-field--outlined .q-field__control:before) {
  border-color: var(--q-blue-grey-7);
}

:deep(.q-dialog__backdrop) {
  background: rgba(91, 112, 131, 0.4);
}
</style>
