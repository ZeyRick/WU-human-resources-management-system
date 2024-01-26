<template>
    <n-button :style="style" :loading="loadingRef" @click="onClick">{{ text }}</n-button>
</template>

<script setup lang="ts">
interface props {
    text?: string
    loading?: boolean
    style?: string
    onClick: Function
}

const props = withDefaults(defineProps<props>(), {
    text: 'Button',
})

const loadingRef = toRef(props.loading)
</script>
