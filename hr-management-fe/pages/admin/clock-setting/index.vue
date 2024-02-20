<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px">
        <n-form inline ref="createFormRef" :model="clockSettingForm">
            <n-form-item path="coordinate" label="Clock Coordinate">
                <n-input
                    size="large"
                    :loading="loading"
                    v-model:value="clockSettingForm.coordinate"
                    @keydown.enter.prevent
                />
            </n-form-item>
            <n-form-item path="range" label="Clock Range">
                <n-input-number
                    size="large"
                    v-model:value="clockSettingForm.clockRange"
                    :loading="loading"
                    :min="10"
                    :precision="0"
                    :input-props="{ 'auto-complete': 'off' }"
                    @keydown.enter.prevent
                >
                    <template #suffix> Meters </template>
                </n-input-number>
            </n-form-item>
            <n-form-item path="allowTime" label="Allow Time">
                <n-input-number
                    size="large"
                    default-value="1"
                    :loading="loading"
                    :input-props="{ 'auto-complete': 'off' }"
                    v-model:value="allowTime"
                    @keydown.enter.prevent
                    :min="0"
                    :precision="0"
                    :max="selectedBreakTimeOption == 'hour' ? 24 : 24 * 60"
                >
                    <template #suffix>
                        <n-dropdown
                            trigger="hover"
                            placement="bottom-start"
                            :options="breakTimeOptions"
                            @select="handleAllowTimeOptionsSelect"
                        >
                            <n-button>{{ i18n.global.t(selectedBreakTimeOption) }}</n-button>
                        </n-dropdown>
                    </template></n-input-number
                >
            </n-form-item>
            <n-form-item>
                <n-button :loading="loading" size="large" @click="onSaveSettingClick"> Save </n-button>
            </n-form-item>
        </n-form>
        <div style="position: relative; height: 550px; overflow: hidden">
            <iframe
                :src="`http://maps.google.com/maps?q=${clockSettingForm.coordinate}&output=embed`"
                width="100%"
                height="550px"
                style="border: 0; pointer-events: none !important; position: absolute"
                loading="lazy"
                referrerpolicy="no-referrer-when-downgrade"
            >
            </iframe>
            <!-- 1px = 1 meter -->
            <div
                v-show="!loading"
                :style="`width: ${clockSettingForm.clockRange}px; height: ${clockSettingForm.clockRange}px;`"
                class="transparent-circle"
            />
        </div>
    </n-layout>
</template>

<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import { apiGetClockSetting, apiSaveClockSetting } from '~/apis/clockSetting'
import type { CreateClockSetting } from '~/types/clockSetting'
const createFormRef = ref<FormInst>()
const loading = ref<boolean>(true)
const allowTime = ref<number>(1)
const breakTimeOptions = [
    {
        label: 'minutes',
        key: 'minutes',
    },
    {
        label: 'hours',
        key: 'hours',
    },
]
const selectedBreakTimeOption = ref<string>('minutes')
const clockSettingForm = ref<CreateClockSetting>({
    coordinate: '',
    allowTime: 0,
    clockRange: 0,
})

const handleAllowTimeOptionsSelect = (key: string) => {
    if (selectedBreakTimeOption.value != key) {
        switch (key) {
            case 'minutes':
                allowTime.value = Math.ceil(allowTime.value * 60)
                break
            case 'hours':
                allowTime.value = Math.ceil(allowTime.value / 60)
                break
        }
    }
    selectedBreakTimeOption.value = key
}

const onSaveSettingClick = async () => {
    if (!loading.value) {
        try {
            clockSettingForm.value.allowTime =
                selectedBreakTimeOption.value == 'hours' ? allowTime.value * 60 : allowTime.value
            const res: any = await apiSaveClockSetting(clockSettingForm.value)
        } catch (error) {
            console.error(error)
        }
    }
}

onMounted(async () => {
    const res: CreateClockSetting = (await apiGetClockSetting()) as CreateClockSetting
    clockSettingForm.value = res
    loading.value = false
})

definePageMeta({
    layout: 'main',
})
</script>

<style>
.transparent-circle {
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    margin: auto;
    border: 1px solid #ccc; /* Choose your border color */
    border-radius: 50%;
    background: rgba(232, 1, 1, 0.5);
    position: absolute;
}
</style>
