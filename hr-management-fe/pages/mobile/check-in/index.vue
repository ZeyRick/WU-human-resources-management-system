<template>
    <div class="centerWrapper">
        <n-card content-style="padding: 10px;" class="centerCard">
            <n-input
                v-model:value="employeeId"
                class="idInput"
                type="text"
                :allow-input="onlyAllowNumber"
                placeholder="ID"
                :disabled="isLoading"
            />
            <n-space vertical />
            <div class="radioWrapper">
                <n-radio :disabled="isLoading" :checked="clockTypeRef === 'out'" value="out" @change="handleClockType">
                    Check Out
                </n-radio>
                <n-radio :disabled="isLoading" :checked="clockTypeRef === 'in'" value="in" @change="handleClockType">
                    Check In
                </n-radio>
            </div>
            <n-space vertical />
            <div class="radioWrapper">
                <n-button :loading="isLoading" @click="handleClock" style="background-color: green" type="primary">
                    Check {{ clockTypeRef.toUpperCase() }}
                </n-button>
            </div>
        </n-card>
    </div>
</template>

<script setup lang="ts">
import { apiClock } from '../../../apis/clock'
const clockTypeRef = ref<'in' | 'out'>('in')
const employeeId = ref<number | null>(null)
const isLoading = ref<boolean>(false)
const onlyAllowNumber = (value: string) => !value || /^\d+$/.test(value)

const handleClockType = (e: Event) => {
    clockTypeRef.value = (e.target as HTMLInputElement).value as 'in' | 'out'
}

const handleClock = async () => {
    try {
        if (!employeeId.value || isLoading.value) return
        isLoading.value = true
        const res = await apiClock({
            employeeId: employeeId.value,
            clockType: clockTypeRef.value,
        })
    } catch (error) {
        console.error(error)
    } finally {
        isLoading.value = false
    }
}

definePageMeta({
    layout: 'moble',
})
</script>

<style>
.radioWrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;
}
.idInput {
    margin-bottom: 20px;
    margin-top: 20px;
}
.centerWrapper {
    width: 100%;
    display: flex;
    overflow: hidden;
    justify-content: center;
    align-items: center;
    padding-left: 20px;
    padding-right: 20px;
}
.centerCard {
    height: fit-content;
}
</style>
