<template>
    <n-tag v-for="employee in props.employees" round :bordered="false" type="success">
        {{ employee.name }}
    </n-tag>
</template>

<script setup lang="ts">
import type { Employee } from '~/types/employee'
const props = defineProps<{ employees: Employee[] }>()
</script>
