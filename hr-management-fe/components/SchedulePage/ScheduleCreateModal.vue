<template>
    <n-button
        size="large"
        strong
        style="background-color: #409eff"
        color="#5cb85c"
        text-color="#000000"
        @click="openCreateModal"
    >
        <template #icon>
            <n-icon color="#000000">
                <AddCircleOutline />
            </n-icon>
        </template>
        Create
    </n-button>
    <n-modal
        :show="showCreateModal"
        :mask-closable="false"
        @negative-click="closeCreateModal"
        @positive-click="onSubmitCreate"
    >
        <n-card
            style="width: 600px"
            title="Create New User"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
            <n-form ref="createFormRef" :rules="CommonFormRules" :model="createFormData">
                <n-form-item path="scope" :label="i18n.global.t('department')">
                    <n-select
                        :disable="loading"
                        v-model:value="departmentId"
                        filterable
                        :placeholder="i18n.global.t('department')"
                        :options="props.departmentOptions"
                    />
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('employee')">
                    <n-select
                        :disable="loading"
                        v-model:value="createFormData.employeeId"
                        filterable
                        :placeholder="i18n.global.t('employee')"
                        :options="employeeOptions"
                        :default-value="'All'"
                    />
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('year_month')">
                    <VueDatePicker v-model="scope" month-picker :disabled="loading" required dark>
                        <template #input-icon>
                            <n-icon size="25px"><CalendarOutline /></n-icon>
                        </template>
                    </VueDatePicker>
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('days')">
                    <VueDatePicker
                        :format="dateFormat"
                        v-model="date"
                        disable-month-year-select
                        :enable-time-picker="false"
                        :disabled="loading"
                        required
                        dark
                        multi-dates
                        hide-offset-dates
                    >
                        <template #input-icon>
                            <n-icon size="25px">
                                <CalendarNumberOutline />
                            </n-icon>
                        </template>
                    </VueDatePicker>
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('click_in_out_time')">
                    <VueDatePicker v-model="time" time-picker dark range>
                        <template #input-icon>
                            <n-icon size="25px"> <TimerOutline /></n-icon>
                        </template>
                    </VueDatePicker>
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('scope')">
                    <n-input
                        :input-props="{ autocomplete: 'off' }"
                        v-model:value="createFormData.scope"
                        @keydown.enter.prevent
                    />
                </n-form-item>
                <n-form-item path="date" :label="i18n.global.t('date')">
                    <n-input
                        :input-props="{ autocomplete: 'off' }"
                        v-model:value="createFormData.dates"
                        @keydown.enter.prevent
                    />
                </n-form-item>
            </n-form>
            <div style="display: flex; gap: 10px; justify-content: flex-end">
                <n-button round @click="closeCreateModal"> Cancel </n-button>
                <n-button round @click="onSubmitCreate"> Create </n-button>
            </div>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { AddCircleOutline, CalendarNumberOutline, CalendarOutline, TimerOutline } from '@vicons/ionicons5'
import type { FormInst, FormValidationError } from 'naive-ui'
import { apiAllEmployee } from '~/apis/employee'
import { apiCreateSchedule } from '~/apis/schedule'
import { CommonFormRules } from '~/constants/formRules'
import type { Employee } from '~/types/employee'
import type { CreateScheduleParams } from '~/types/schedule'
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import moment from 'moment'

const props = defineProps<{
    employeeOptions: { label: string; value: string }[]
    departmentOptions: { label: string; value: string }[]
    departmentId: string
}>()
const createFormRef = ref<FormInst>()
const employeeOptions = ref<{ label: string; value: string | undefined }[]>(props.employeeOptions)
const showCreateModal = ref<boolean>(false)
const createFormData = ref<CreateScheduleParams>({
    scope: '',
    dates: '',
    clockInTime: '',
    clockOutTime: '',
})
const departmentId = ref<string>(props.departmentId)
const loading = ref<boolean>(false)
const time = ref<{ hours: number; minutes: number; seconds: number }[]>()
const date = ref()
const scope = ref<{ month: number; year: number }>()

const closeCreateModal = () => {
    createFormData.value = {
        scope: '',
        dates: '',
        clockInTime: '',
        clockOutTime: '',
    }
    showCreateModal.value = false
}
const openCreateModal = () => {
    departmentId.value = props.departmentId
    showCreateModal.value = true
}

const onSubmitCreate = () => {
    createFormRef.value?.validate((errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            apiCreateSchedule(createFormData.value, departmentId.value)
        } else {
            console.log(errors)
        }
    })
}

const getEmployee = async () => {
    try {
        loading.value = true
        const res: any = await apiAllEmployee({ departmentId: departmentId.value })
        const employees = JSON.parse(res).res as Employee[]
        employeeOptions.value = [{ label: 'All', value: undefined }]
        employees.map((e) => {
            employeeOptions.value.push({
                label: `${e.id} - ${e.name}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
        loading.value = false
    }
}

const dateFormat = (dates: Date[]): string => {
    let formatDates: string[] = []
    dates.forEach((date) => {
        formatDates.push(date.getDate().toString())
    })

    return formatDates.toLocaleString()
}

const updateScope = (value: { month: number; year: number } | undefined) => {
    if (value) {
        createFormData.value.scope = moment( `${value.year}-${value.month}`).format('YYYY-MM')
    }
}

const updateDates = (dates: Date[]) => (createFormData.value.dates = dateFormat(dates))

const upateClockInOutTime = (times: { hours: number; minutes: number; seconds: number }[] | undefined) => {
    if (times) {
        createFormData.value.clockInTime = `${times[0].hours}:${times[0].minutes}:${times[0].seconds}`
        createFormData.value.clockOutTime = `${times[1].hours}:${times[1].minutes}:${times[1].seconds}`
    }
}

watch(time, () => upateClockInOutTime(time.value))
watch(date, () => updateDates(date.value))
watch(scope, () => updateScope(scope.value))
watch(departmentId, getEmployee)
</script>

<style>
.dp__theme_dark {
    --dp-background-color: #424246;
    --dp-text-color: #fff;
    --dp-hover-color: #484848;
    --dp-hover-text-color: #fff;
    --dp-hover-icon-color: #959595;
    --dp-primary-color: #3f9eff;
    --dp-primary-disabled-color: #61a8ea;
    --dp-primary-text-color: #fff;
    --dp-secondary-color: #a9a9a9;
    --dp-border-color: #2d2d2d;
    --dp-menu-border-color: #2d2d2d;
    --dp-border-color-hover: #63e2b7;
    --dp-disabled-color: #737373;
    --dp-disabled-color-text: #d0d0d0;
    --dp-scroll-bar-background: #212121;
    --dp-scroll-bar-color: #484848;
    --dp-success-color: #00701a;
    --dp-success-color-disabled: #428f59;
    --dp-icon-color: #959595;
    --dp-danger-color: #e53935;
    --dp-marker-color: #e53935;
    --dp-tooltip-color: #3e3e3e;
    --dp-highlight-color: rgb(0 92 178 / 20%);
    --dp-range-between-dates-background-color: var(--dp-hover-color, #484848);
    --dp-range-between-dates-text-color: var(--dp-hover-text-color, #fff);
    --dp-range-between-border-color: var(--dp-hover-color, #fff);
}
.dp__action_button {
    width: 60px;
}
.dp__input_icon {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
