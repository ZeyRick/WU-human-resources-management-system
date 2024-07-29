<template>
    <n-button :disabled="disabled" :style="style" quaternary :type="type ? type : 'primary'" :loading="loadingRef" @click="onClick">{{ text }}</n-button>
</template>

<script setup lang="ts">
interface props {
    text?: string
    loading?: boolean
    style?: string
    onClick: Function
    disabled?: boolean
    type?: string
}

const props = withDefaults(defineProps<props>(), {
    text: 'Button',
})

const loadingRef = toRef(props.loading)
</script>
