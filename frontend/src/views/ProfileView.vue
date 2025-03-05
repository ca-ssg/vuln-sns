<template>
  <div class="profile-container">
    <!-- Profile Header -->
    <div class="profile-header q-px-md q-py-lg">
      <div class="row justify-between items-center">
        <div>
          <!-- XSS脆弱性: ユーザー入力を適切にエスケープせずに表示 -->
          <div class="text-h4 text-weight-bold q-mb-sm" v-html="authStore.user?.nickname"></div>
          <div class="text-grey-6">@{{ authStore.user?.id }}</div>
        </div>
        <q-btn
          color="primary"
          label="プロフィールを編集"
          rounded
          unelevated
          @click="showEditDialog = true"
        />
      </div>
    </div>

    <!-- Edit Profile Dialog -->
    <base-dialog
      v-model="showEditDialog"
      title="プロフィールを編集"
      min-width="350px"
      @close="showEditDialog = false"
    >
      <template #content>
        <q-input
          v-model="newNickname"
          label="ニックネーム"
          dark
          outlined
          class="q-mb-md"
          :rules="[val => !!val || 'ニックネームを入力してください']"
        />
      </template>
      <template #actions>
        <q-btn flat label="キャンセル" @click="showEditDialog = false" />
        <q-btn
          color="primary"
          label="保存"
          rounded
          unelevated
          @click="updateProfile"
          :loading="loading"
        />
      </template>
    </base-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import BaseDialog from '@/components/BaseDialog.vue'

const authStore = useAuthStore()
const showEditDialog = ref(false)
const newNickname = ref('')
const loading = ref(false)

const updateProfile = async () => {
  if (!newNickname.value) return

  loading.value = true
  try {
    const success = await authStore.updateNickname(newNickname.value)
    if (success) {
      showEditDialog.value = false
    }
  } catch (error) {
    console.error('Failed to update profile:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style>
.profile-container {
  max-width: 600px;
  margin: 0 auto;
}

.profile-header {
  border-bottom: 1px solid #38444d;
}

:deep(.q-field__control) {
  background: #253341 !important;
}

:deep(.q-field__label) {
  color: #8899a6 !important;
}

:deep(.q-field--outlined .q-field__control:before) {
  border-color: #38444d !important;
}

:deep(.q-dialog__backdrop) {
  background: rgba(91, 112, 131, 0.4) !important;
}

:deep(.q-card) {
  background-color: #15202b !important;
}
</style>
