<template>
    <n-tag round :size="size || 'medium'" :bordered="false" :type="props.text === type || 'error'">
        {{ i18n.global.t(props.text) }}
    </n-tag>
</template>

<script setup lang="ts">
import type { DefineComponent } from 'vue'

const props = defineProps<{
    text: string
    type?: 'default' | 'primary' | 'info' | 'success' | 'warning' | 'error'
    size?: 'small' | 'medium' | 'large'
}>()
</script>
