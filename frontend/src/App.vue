<template>
  <q-layout view="hHh LpR fFf" class="bg-black">
    <!-- Header with mobile menu button -->
    <q-header elevated class="bg-black q-py-sm">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="leftDrawerOpen = !leftDrawerOpen" class="lt-md" />
        <q-toolbar-title class="text-h6">ホーム</q-toolbar-title>
      </q-toolbar>
    </q-header>

    <!-- Responsive drawer - behavior changes on mobile -->
    <q-drawer
      v-model="leftDrawerOpen"
      :breakpoint="1024"
      :behavior="$q.screen.lt.md ? 'mobile' : 'desktop'"
      bordered
      class="bg-black"
      show-if-above
    >
      <q-list>
        <q-item clickable v-ripple to="/" exact>
          <q-item-section avatar>
            <q-icon name="home" size="md" />
          </q-item-section>
          <q-item-section>
            <q-item-label class="text-h6">ホーム</q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable v-ripple to="/profile">
          <q-item-section avatar>
            <q-icon name="person" size="md" />
          </q-item-section>
          <q-item-section>
            <q-item-label class="text-h6">プロフィール</q-item-label>
          </q-item-section>
        </q-item>

        <q-separator class="q-my-md" />

        <q-item-label header>トレンド</q-item-label>
        <q-item clickable v-ripple @click="searchHashtag(tag)" v-for="tag in ['セキュリティ', '脆弱性']" :key="tag">
          <q-item-section>
            <q-item-label caption>トレンド</q-item-label>
            <q-item-label class="text-weight-bold">#{{ tag }}</q-item-label>
            <q-item-label caption>{{ tag === 'セキュリティ' ? '1,234' : '891' }} 投稿</q-item-label>
          </q-item-section>
        </q-item>

        <q-separator class="q-my-md" />

        <q-item v-if="!isLoggedIn" clickable v-ripple to="/login">
          <q-item-section>
            <q-item-label class="text-h6">ログイン</q-item-label>
          </q-item-section>
        </q-item>

        <q-item v-else clickable v-ripple @click="logout">
          <q-item-section>
            <q-item-label class="text-h6">ログアウト</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePostsStore } from '@/stores/posts'
import { useQuasar } from 'quasar'

const $q = useQuasar()
const router = useRouter()
const authStore = useAuthStore()
const postsStore = usePostsStore()
const leftDrawerOpen = ref(false)
const isLoggedIn = computed(() => authStore.isLoggedIn)

const logout = () => {
  authStore.logout()
  router.push('/login')
  if ($q.screen.lt.md) {
    leftDrawerOpen.value = false
  }
}

const searchHashtag = async (tag) => {
  try {
    await postsStore.searchByHashtag(tag)
    router.push('/')
    if ($q.screen.lt.md) {
      leftDrawerOpen.value = false
    }
  } catch (error) {
    console.error('Failed to search hashtag:', error)
  }
}
</script>

<style>
body {
  color: white;
}

.q-drawer {
  background-color: #15202b !important;
}

.q-item {
  color: white;
}

.q-item__label--header {
  color: #8899a6;
}

.q-item__label--caption {
  color: #8899a6;
}

.q-separator {
  background: #38444d;
}

.q-toolbar {
  min-height: 48px;
}

.q-page-container {
  background-color: #15202b;
}

/* Mobile-specific styles */
@media (max-width: 1023px) {
  .q-drawer {
    width: 100% !important;
    max-width: 280px !important;
  }
}
</style>
