<template>
    <n-icon-wrapper :color="!status && 'red'" :size="24" :border-radius="10">
        <n-icon v-if="status" :size="18" :component="CheckmarkSharp" />
        <n-icon v-else :size="18" :component="CloseSharp" />
    </n-icon-wrapper>
</template>

<script setup lang="ts">
import { CheckmarkSharp, CloseSharp } from '@vicons/ionicons5';


defineProps<{
    status: boolean
}>()

</script>
