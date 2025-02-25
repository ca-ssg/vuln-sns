<template>
  <q-dialog v-model="isOpen" persistent>
    <q-card :style="cardStyle">
      <q-card-section class="q-pb-none">
        <div class="row items-center">
          <q-btn flat round icon="close" @click="closeDialog" />
          <div class="text-h6 text-white q-ml-md">{{ title }}</div>
        </div>
      </q-card-section>

      <q-card-section>
        <slot name="content"></slot>
      </q-card-section>

      <q-card-actions align="right" class="q-pa-md">
        <slot name="actions"></slot>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'

const props = defineProps<{
  modelValue: boolean
  title: string
  minWidth?: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
}>()

const isOpen = ref(false)
const cardStyle = computed(() => ({
  minWidth: props.minWidth || '500px'
}))

watch(() => props.modelValue, (val) => {
  isOpen.value = val
})

watch(isOpen, (val) => {
  emit('update:modelValue', val)
})

const closeDialog = () => {
  emit('close')
  emit('update:modelValue', false)
}
</script>

<style>
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
