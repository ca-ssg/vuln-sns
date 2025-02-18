<template>
  <q-page class="flex flex-center bg-black">
    <q-card flat bordered class="login-card bg-black q-pa-xl" style="width: 400px">
      <q-card-section class="text-center q-pb-xl">
        <i class="fas fa-feather text-primary text-h2 q-mb-md"></i>
        <div class="text-h4 text-weight-bold q-mb-sm">VulnAppにログイン</div>
        <div class="text-grey-6">セキュリティの学習用アプリケーション</div>
      </q-card-section>

      <q-card-section>
        <form @submit.prevent="handleLogin">
          <div class="q-gutter-y-md">
            <q-input
              v-model="id"
              label="ユーザーID"
              class="bg-black"
              outlined
              :rules="[val => !!val || 'IDを入力してください']"
            >
              <template v-slot:prepend>
                <q-icon name="fas fa-user" />
              </template>
            </q-input>

            <q-input
              v-model="password"
              type="password"
              label="パスワード"
              class="bg-black"
              outlined
              :rules="[val => !!val || 'パスワードを入力してください']"
            >
              <template v-slot:prepend>
                <q-icon name="fas fa-lock" />
              </template>
            </q-input>

            <div v-if="error" class="text-negative text-center q-mb-md">{{ error }}</div>

            <q-btn
              type="submit"
              color="primary"
              class="full-width q-py-sm text-h6"
              label="ログイン"
              unelevated
            />
          </div>
        </form>
      </q-card-section>

      <q-card-section class="text-center text-grey-6 text-caption">
        初期アカウントについてはREADMEをご確認ください
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const id = ref('')
const password = ref('')
const error = ref('')

const handleLogin = async () => {
  error.value = ''
  try {
    const success = await authStore.login(id.value, password.value)
    if (success) {
      router.push('/')
    } else {
      error.value = 'ログインに失敗しました'
    }
  } catch (e) {
    console.error('Login error:', e)
    error.value = 'ログインに失敗しました'
  }
}
</script>

<style scoped>
.q-field {
  border-color: #2F3336;
}
.q-field--outlined .q-field__control:before {
  border-color: #2F3336;
}
</style>
