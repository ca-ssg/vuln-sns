<template>
  <div class="profile-container">
    <!-- Profile Header -->
    <div class="profile-header q-px-md q-py-lg">
      <div class="row justify-between items-center">
        <div>
          <div class="text-h4 text-weight-bold q-mb-sm">{{ authStore.user?.nickname || authStore.user?.id }}</div>
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
    <q-dialog v-model="showEditDialog" persistent>
      <q-card class="bg-dark text-white" style="min-width: 350px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">プロフィールを編集</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section class="q-pt-lg">
          <q-input
            v-model="newNickname"
            label="ニックネーム"
            dark
            outlined
            class="q-mb-md"
            :rules="[val => !!val || 'ニックネームを入力してください']"
          />
        </q-card-section>

        <q-card-actions align="right" class="bg-dark text-white">
          <q-btn flat label="キャンセル" v-close-popup />
          <q-btn
            color="primary"
            label="保存"
            rounded
            unelevated
            @click="updateProfile"
            :loading="loading"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const showEditDialog = ref(false)
const newNickname = ref('')
const loading = ref(false)

const updateProfile = async () => {
  if (!newNickname.value) return

  loading.value = true
  try {
    await authStore.updateProfile(newNickname.value)
    showEditDialog.value = false
  } catch (error) {
    console.error('Failed to update profile:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
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
