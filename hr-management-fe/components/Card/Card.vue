<template>
    <n-card content-style="padding: 8px;">
        <div class="card-content">
            <div class="icon-container">
                <n-icon color="black" size="40px" class="left-icon">
                    <slot><AccessibilitySharp /></slot>
                    <!-- slot outlet -->
                </n-icon>
            </div>
            <div class="data-content">
                <p>{{ i18n.global.t(label) }}</p>
                <h3>{{ value }}</h3>
            </div>
        </div>
    </n-card>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import { NCard, NIcon } from 'naive-ui'
import { AccessibilitySharp } from '@vicons/ionicons5'

// Define the props for the component
const props = defineProps({
    Icon: {
        type: Object,
        required: false,
        default: AccessibilitySharp,
    },
    label: {
        type: String,
        required: true,
        default: 'User', // or 'Admin'
    },
    value: {
        type: Number,
        required: true,
        default: 0,
    },
})
</script>

<style scoped>
.card-content {
    display: flex;
    align-items: center;
}

.icon-container {
    width: 60px;
    height: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: cyan; /* Tailwind Gray 200 */
    border-radius: 5%;
    margin-right: 16px;
}

.left-icon {
    color: #374151; /* Tailwind Gray 700 */
}

.data-content p {
    margin: 0;
    font-size: 16px;
    color: #6b7280; /* Tailwind Gray 500 */
}

.data-content h3 {
    margin: 0;
    font-size: 24px;
    font-weight: bold;
}
</style>
