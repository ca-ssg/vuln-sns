<template>
  <q-layout view="hHh LpR lFf" class="bg-black">
    <!-- Header with mobile menu button -->
    <q-header class="bg-black" style="border-bottom: 1px solid #2F3336">
      <q-toolbar style="min-height: 53px; padding: 0 16px;">
        <q-btn dense flat round icon="menu" @click="leftDrawerOpen = !leftDrawerOpen" class="lt-md" />
        <q-toolbar-title class="text-weight-bold" style="font-size: 20px;">
          ホーム
        </q-toolbar-title>
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
            <q-item-label>ホーム</q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable v-ripple to="/search">
          <q-item-section avatar>
            <q-icon name="search" size="md" />
          </q-item-section>
          <q-item-section>
            <q-item-label>話題を検索</q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable v-ripple to="/profile">
          <q-item-section avatar>
            <q-icon name="person" size="md" />
          </q-item-section>
          <q-item-section>
            <q-item-label>プロフィール</q-item-label>
          </q-item-section>
        </q-item>

        <q-btn
          color="primary"
          class="full-width-with-margin q-my-md"
          rounded
          size="lg"
          label="ポストする"
          @click="showPostDialog = true"
        />

        <q-separator class="q-my-md" />

        <q-item clickable v-ripple class="account-item">
          <q-item-section avatar>
            <q-avatar>
              <img :src="'https://ui-avatars.com/api/?name=' + (isLoggedIn ? authStore.user?.id : 'guest')" />
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>{{ isLoggedIn ? authStore.user?.id : 'ゲスト' }}</q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-btn flat round>
              <q-menu anchor="bottom right" self="top right">
                <q-list style="min-width: 200px">
                  <q-item v-if="!isLoggedIn" clickable v-ripple to="/login">
                    <q-item-section>ログイン</q-item-section>
                  </q-item>
                  <q-item v-else clickable v-ripple @click="logout">
                    <q-item-section>ログアウト</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
              <q-icon name="more_vert" />
            </q-btn>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container class="row">
      <div class="col-12 col-md-8">
        <router-view />
      </div>
      <div class="col-md-4 gt-sm">
        <div class="bg-black q-mt-md trends-section" style="border-left: 1px solid #2F3336">
          <div class="q-pa-md">
            <div class="text-h6 q-mb-md">トレンド</div>
            <q-list>
              <q-item clickable v-ripple @click="searchHashtag(tag)" v-for="tag in ['セキュリティ', '脆弱性']" :key="tag">
                <q-item-section>
                  <q-item-label caption>トレンド</q-item-label>
                  <q-item-label class="text-weight-bold">#{{ tag }}</q-item-label>
                  <q-item-label caption>{{ tag === 'セキュリティ' ? '1,234' : '891' }} 投稿</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </div>
        </div>
      </div>

      <!-- Mobile Trends Dialog -->
      <q-dialog v-model="showTrends" position="bottom">
        <div class="bg-black full-width">
          <div class="q-pa-md">
            <div class="text-h6 q-mb-md">トレンド</div>
            <q-list>
              <q-item clickable v-ripple @click="searchHashtagMobile(tag)" v-for="tag in ['セキュリティ', '脆弱性']" :key="tag">
                <q-item-section>
                  <q-item-label caption>トレンド</q-item-label>
                  <q-item-label class="text-weight-bold">#{{ tag }}</q-item-label>
                  <q-item-label caption>{{ tag === 'セキュリティ' ? '1,234' : '891' }} 投稿</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </div>
        </div>
      </q-dialog>

      <!-- Mobile Trends Button -->
      <q-page-sticky position="bottom-right" :offset="[18, 18]" class="lt-md">
        <q-btn
          round
          color="primary"
          icon="trending_up"
          @click="showTrends = true"
        />
      </q-page-sticky>

      <!-- Post Dialog -->
      <post-dialog
        v-model="showPostDialog"
      />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useQuasar } from 'quasar'
import PostDialog from '@/components/PostDialog.vue'

const $q = useQuasar()
const router = useRouter()
const authStore = useAuthStore()

const leftDrawerOpen = ref(false)
const showTrends = ref(false)
const showPostDialog = ref(false)
const isLoggedIn = computed(() => authStore.isAuthenticated)

const logout = () => {
  authStore.logout()
  router.push('/login')
  if ($q.screen.lt.md) {
    leftDrawerOpen.value = false
  }
}

const searchHashtag = (tag: string) => {
  router.push({ path: '/search', query: { tag } })
  if ($q.screen.lt.md) {
    leftDrawerOpen.value = false
  }
}

const searchHashtagMobile = (tag: string) => {
  router.push({ path: '/search', query: { tag } })
  showTrends.value = false
}
</script>

<style>
body {
  color: white;
}

.q-drawer {
  background-color: #000000 !important;
}

.q-item {
  color: var(--q-text);
  padding: 12px 16px;
  min-height: 50px;
}

.q-item__label--header {
  color: #71767B;
}

.q-item__label--caption {
  color: #71767B;
}

.q-separator {
  background: #2F3336;
}

.q-toolbar {
  min-height: 48px;
}

.q-page-container {
  background-color: #000000;
  padding: 0;
}

.q-card {
  border-radius: 0;
}

.q-item {
  border-radius: 0;
  margin: 0;
  &:hover {
    background-color: rgba(255, 255, 255, 0.03);
  }
}

.q-drawer {
  border-right: 1px solid #2F3336 !important;
}

/* カスタムクラス - full-widthを拡張し、左右マージンを追加 */
.full-width-with-margin {
  width: 100% !important;
  margin: 8px 16px !important;
}

/* Mobile-specific styles */
@media (max-width: 1023px) {
  .q-drawer {
    width: 100% !important;
    max-width: 280px !important;
  }
}
</style>
