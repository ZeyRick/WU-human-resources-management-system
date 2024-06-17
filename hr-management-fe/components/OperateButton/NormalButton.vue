<template>
    <n-button :disabled="disabled" :style="style" quaternary type="primary" :loading="loadingRef" @click="onClick">{{ text }}</n-button>
</template>

<script setup lang="ts">
interface props {
    text?: string
    loading?: boolean
    style?: string
    onClick: Function
    disabled?: boolean
}

const props = withDefaults(defineProps<props>(), {
    text: 'Button',
})

const loadingRef = toRef(props.loading)
</script>
