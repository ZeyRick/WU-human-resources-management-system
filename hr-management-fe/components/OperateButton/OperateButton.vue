<template>
    <n-popconfirm
        :negative-text="negativeText"
        :positive-text="positiveText"
        @positive-click="positiveClick"
        @negative-click="negativeClick"
        :loading="loadingRef"
         
    >
        <template #trigger>
            <n-button  :style="style" :type="type ? type : 'error'" quaternary :loading="loadingRef">{{ text }}</n-button>
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
    loading?: boolean
    style?: string
    type?: string
}

const props = withDefaults(defineProps<props>(), {
    text: 'Button',
    description: 'Are you sure ?',
    negativeText: 'Cancel',
    positiveText: 'Confirm',
})

const loadingRef = toRef(props.loading)
</script>
