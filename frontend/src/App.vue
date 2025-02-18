<template>
  <q-layout view="hHh LpR fFf" class="bg-black">
    <!-- Left Sidebar -->
    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered class="bg-black" :width="280">
      <div class="q-pa-md">
        <div class="text-h6 q-mb-lg">
          <i class="fas fa-feather text-primary text-h4"></i>
        </div>
        <q-list>
          <q-item clickable v-ripple to="/" class="q-py-md">
            <q-item-section avatar>
              <i class="fas fa-home"></i>
            </q-item-section>
            <q-item-section>
              <q-item-label class="text-h6">ホーム</q-item-label>
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple to="/search" class="q-py-md">
            <q-item-section avatar>
              <i class="fas fa-search"></i>
            </q-item-section>
            <q-item-section>
              <q-item-label class="text-h6">話題を検索</q-item-label>
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple to="/profile" class="q-py-md">
            <q-item-section avatar>
              <i class="fas fa-user"></i>
            </q-item-section>
            <q-item-section>
              <q-item-label class="text-h6">プロフィール</q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
        <q-btn 
          v-if="isAuthenticated"
          color="primary" 
          class="full-width q-mt-lg" 
          size="lg" 
          label="投稿する"
          @click="showPostDialog = true" 
        />
      </div>
    </q-drawer>

    <!-- Post Dialog -->
    <q-dialog v-model="showPostDialog">
      <q-card class="bg-black" style="width: 600px; max-width: 80vw;">
        <q-card-section class="row items-center">
          <div class="text-h6">新規投稿</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section>
          <q-input
            v-model="newPost"
            type="textarea"
            class="bg-black"
            outlined
            autogrow
            placeholder="いまどうしてる？"
          />
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="キャンセル" color="negative" v-close-popup />
          <q-btn 
            flat 
            label="投稿する" 
            color="primary" 
            @click="createPost" 
            :disable="!newPost.trim()" 
          />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Main Content -->
    <q-page-container class="bg-black">
      <div class="row">
        <div class="col-12 col-md-8 offset-md-3">
          <q-header elevated class="bg-black q-py-sm">
            <q-toolbar>
              <q-toolbar-title class="text-h6">ホーム</q-toolbar-title>
              <div v-if="!isAuthenticated">
                <q-btn flat to="/login" label="ログイン" color="primary" />
              </div>
              <div v-else>
                <q-btn flat @click="logout" label="ログアウト" color="negative" />
              </div>
            </q-toolbar>
          </q-header>
          <router-view />
        </div>

        <!-- Right Sidebar -->
        <div class="col-md-3 gt-sm">
          <div class="q-pa-md">
            <q-card flat bordered class="bg-dark q-pa-md">
              <div class="text-h6 q-mb-md">トレンド</div>
              <q-list>
                <q-item clickable v-ripple>
                  <q-item-section>
                    <q-item-label caption>日本のトレンド</q-item-label>
                    <q-item-label class="text-weight-bold">#セキュリティ</q-item-label>
                    <q-item-label caption>1,234 投稿</q-item-label>
                  </q-item-section>
                </q-item>
                <q-item clickable v-ripple>
                  <q-item-section>
                    <q-item-label caption>テクノロジー</q-item-label>
                    <q-item-label class="text-weight-bold">#脆弱性</q-item-label>
                    <q-item-label caption>891 投稿</q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-card>
          </div>
        </div>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from './stores/auth'
import { usePostStore } from './stores/posts'
import { storeToRefs } from 'pinia'

const leftDrawerOpen = ref(true)
const showPostDialog = ref(false)
const newPost = ref('')

const authStore = useAuthStore()
const postStore = usePostStore()
const { isAuthenticated } = storeToRefs(authStore)
const { logout } = authStore

const createPost = async () => {
  if (newPost.value.trim()) {
    await postStore.createPost(newPost.value)
    newPost.value = ''
    showPostDialog.value = false
    await postStore.fetchPosts() // Refresh posts after creating
  }
}
</script>

<style>
.q-drawer {
  border-color: #2F3336 !important;
}

.q-card {
  border-color: #2F3336 !important;
}

.q-toolbar {
  border-bottom: 1px solid #2F3336;
}

.q-item {
  min-height: 56px;
}

.q-item:hover {
  background: rgba(255, 255, 255, 0.03);
}

.q-dialog__inner {
  background: rgba(91, 112, 131, 0.4) !important;
}
</style>
