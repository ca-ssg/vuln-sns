<template>
  <div class="home-container">
    <div class="post-button-container">
      <q-btn color="primary" rounded class="post-button" @click="showPostDialog = true">
        ポストする
      </q-btn>
    </div>
    <div v-if="postsStore.loading" class="text-center q-pa-md">
      <q-spinner color="primary" size="3em" />
    </div>
    <div v-else>
      <post-card
        v-for="post in postsStore.posts"
        :key="post.id"
        :post="post"
      />
    </div>

    <!-- Post Dialog -->
    <q-dialog v-model="showPostDialog">
      <q-card class="bg-black" style="min-width: 500px">
        <q-card-section>
          <div class="text-h6">新規投稿</div>
        </q-card-section>

        <q-card-section>
          <q-input
            v-model="newPostContent"
            type="textarea"
            class="bg-black"
            dark
            autofocus
            placeholder="いまどうしてる？"
          />
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="キャンセル" v-close-popup />
          <q-btn color="primary" label="投稿する" @click="createPost" :loading="posting" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { usePostsStore } from '../stores/posts'
import { useRoute } from 'vue-router'
import PostCard from '../components/PostCard.vue'

const postsStore = usePostsStore()
const route = useRoute()

const showPostDialog = ref(false)
const newPostContent = ref('')
const posting = ref(false)

const createPost = async () => {
  if (!newPostContent.value.trim()) return
  
  posting.value = true
  try {
    await postsStore.createPost(newPostContent.value)
    showPostDialog.value = false
    newPostContent.value = ''
  } catch (error) {
    console.error('Error creating post:', error)
  } finally {
    posting.value = false
  }
}

// Watch for route query changes to update posts
watch(() => route.query.tag, async (newTag) => {
  if (newTag) {
    await postsStore.searchByHashtag(newTag as string)
  } else {
    await postsStore.fetchPosts()
  }
}, { immediate: true })

// Watch for route query changes to update posts
watch(() => route.query.tag, async (newTag) => {
  if (newTag) {
    await postsStore.searchByHashtag(newTag as string)
  } else {
    await postsStore.fetchPosts()
  }
}, { immediate: true })
</script>

<style scoped>
.home-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 0;
  border-left: 1px solid #2F3336;
  border-right: 1px solid #2F3336;
}

.post-button-container {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 1000;
}

.post-button {
  padding: 12px 24px;
  font-weight: bold;
}
</style>
