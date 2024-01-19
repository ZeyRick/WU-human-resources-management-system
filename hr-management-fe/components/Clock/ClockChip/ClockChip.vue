<template>
    <n-tag round :bordered="false" :type="props.text === 'in' ? 'success' : 'error'">
        Checked {{ i18n.global.t(props.text) }}
        <template #icon>
            <n-icon :component="props.text === 'in' ? CheckmarkCircle : CloseCircle" />
        </template>
    </n-tag>
</template>

<script setup lang="ts">
import { CheckmarkCircle, CloseCircle } from '@vicons/ionicons5'

const props = defineProps<{
    text: string
}>()
</script>
