<template>
    <n-popconfirm
        :negative-text="negativeText"
        :positive-text="positiveText"
        @positive-click="positiveClick"
        @negative-click="negativeClick"
    >
        <template #trigger>
            <n-button>{{ text }}</n-button>
        </template>
        {{ description }}
    </n-popconfirm>
</template>

<script setup lang="ts">
interface props {
    text?: string
    variant?: string
    size?: string
    positiveClick?: Function
    negativeClick?: Function
    positiveText?: string
    negativeText?: string
    description?: string
}

withDefaults(defineProps<props>(), {
    text: 'Button',
    description: 'Are you sure ?',
    negativeText: 'Cancel',
    positiveText: 'Confirm',
})
</script>
