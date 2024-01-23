<template>
    <n-button
        size="large"
        strong
        style="background-color: #409eff"
        color="#5cb85c"
        text-color="#000000"
        @click="openCreateModal"
        :disabled="props.isUpdate && !filterForm.employeeId"
    >
        <template #icon>
            <n-icon v-if="props.isUpdate" color="#000000">
                <ReloadSharp />
            </n-icon>
            <n-icon v-else color="#000000">
                <AddCircleOutline />
            </n-icon>
        </template>
        {{ props.isUpdate ? 'Update' : 'Create' }}
    </n-button>
    <n-modal
        :show="showCreateModal"
        :mask-closable="false"
        @negative-click="closeCreateModal"
        @positive-click="onSubmitCreate"
    >
        <n-card
            style="width: 600px"
            title="Add New Schedule"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
            <n-form ref="createFormRef" :rules="CommonFormRules" :model="createFormData">
                <n-form-item path="departmentId" :label="i18n.global.t('department')">
                    <n-select
                        :disabled="loading || props.isUpdate"
                        v-model:value="createFormData.departmentId"
                        filterable
                        :placeholder="i18n.global.t('department')"
                        :options="props.departmentOptions"
                        @update:value="getEmployee"
                    />
                </n-form-item>
                <n-form-item path="" :label="i18n.global.t('employee')">
                    <n-select
                        :disabled="loading || props.isUpdate"
                        v-model:value="createFormData.employeeId"
                        filterable
                        :placeholder="i18n.global.t('employee')"
                        :options="employeeOptions"
                        :default-value="props.isUpdate ? props.filterForm.employeeId : 'All'"
                    />
                </n-form-item>
                <n-form-item path="scope" :label="i18n.global.t('year_month')">
                    <VueDatePicker
                        ref="monthYearPicker"
                        :clearable="false"
                        v-model="scope"
                        month-picker
                        :disabled="loading || props.isUpdate"
                        required
                        dark
                    >
                        <template #input-icon>
                            <n-icon size="25px"><CalendarOutline /></n-icon>
                        </template>
                    </VueDatePicker>
                </n-form-item>
                <n-form-item path="dates" :label="i18n.global.t('days')">
                    <VueDatePicker
                        ref="datesPicker"
                        :format="dateFormat"
                        v-model="date"
                        disable-month-year-select
                        :enable-time-picker="false"
                        :disabled="loading"
                        required
                        dark
                        multi-dates
                        :month-change-on-arrows="false"
                        :clearable="false"
                        hide-offset-dates
                        :month-change-on-scroll="false"
                        focus-start-date
                        :start-date="datePickerStartDate"
                        week-start="0"
                    >
                        <template #input-icon>
                            <n-icon size="25px">
                                <CalendarNumberOutline />
                            </n-icon>
                        </template>
                    </VueDatePicker>
                </n-form-item>
                <n-form-item path="clockOutTime" :label="i18n.global.t('click_in_out_time')">
                    <VueDatePicker ref="timesPicker" :clearable="false" v-model="time" time-picker dark range>
                        <template #input-icon>
                            <n-icon size="25px"> <TimerOutline /></n-icon>
                        </template>
                    </VueDatePicker>
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
import { AddCircleOutline, CalendarNumberOutline, CalendarOutline, TimerOutline, ReloadSharp } from '@vicons/ionicons5'
import { useMessage, type FormInst, type FormValidationError } from 'naive-ui'
import { apiAllEmployee } from '~/apis/employee'
import { apiCreateSchedule, apiGetScheduleByEmployeeId, apiUpdateSchedule } from '~/apis/schedule'
import { CommonFormRules } from '~/constants/formRules'
import type { Employee } from '~/types/employee'
import type { CreateScheduleParams, ScheduleFilterParams, Schedule } from '~/types/schedule'
import VueDatePicker, { type DatePickerInstance } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import { parse } from 'date-fns'
import moment from 'moment'

const props = defineProps<{
    isUpdate: boolean
    employeeOptions: { label: string; value: string }[]
    departmentOptions: { label: string; value: string }[]
    filterForm: ScheduleFilterParams
}>()
const getMonthAndYear = (propsScope: string) => {
    const year: string = propsScope.split('-')[0]
    const month: string = propsScope.split('-')[1]
    return { month: Number(month) - 1, year: Number(year) }
}

const getTimesPickerValue = (dateString: string): { hours: number; minutes: number; seconds: number } => {
    // Create a moment object from the string
    const momentObject = moment.utc(dateString).local()

    // Extract hours, minutes, and seconds in local time
    const hours = momentObject.hours()
    const minutes = momentObject.minutes()
    const seconds = momentObject.seconds()

    return {
        hours,
        minutes,
        seconds,
    }
}

const emit = defineEmits<{
    (e: 'currentDateChange', newDate: Date): void
    (e: 'onDepartmentChange', departmentId: string): void
    (e: 'refreshData'): void
}>()
const createFormRef = ref<FormInst>()
const employeeOptions = ref<{ label: string; value: string | undefined }[]>(
    props.isUpdate ? props.employeeOptions.slice(1) : props.employeeOptions,
)
const showCreateModal = ref<boolean>(false)
const createFormData = ref<CreateScheduleParams>({
    scope: '',
    dates: '',
    clockInTime: '',
    clockOutTime: '',
    departmentId: props.filterForm.departmentId || props.departmentOptions[0]?.value,
})
const loading = ref<boolean>(false)
const time = ref<{ hours: number; minutes: number; seconds: number }[]>()
const date = ref()
const scope = ref<{ month: number; year: number }>(getMonthAndYear(props.filterForm.scope))
const monthYearPicker = ref<DatePickerInstance>()
const timesPicker = ref<DatePickerInstance>()
const datesPicker = ref<DatePickerInstance>()
const datePickerStartDate = ref(parse(`${scope.value?.year}-${scope.value?.month}`, 'yyyy-MM', new Date()))
const message = useMessage()
const closeCreateModal = () => {
    createFormData.value = {
        scope: '',
        dates: '',
        clockInTime: '',
        clockOutTime: '',
        departmentId: props.filterForm.departmentId,
    }
    monthYearPicker.value?.clearValue()
    datesPicker.value?.clearValue()
    timesPicker.value?.clearValue()
    showCreateModal.value = false
}
const employeeSchedule = ref<Schedule>()
const openCreateModal = async () => {
    createFormData.value.departmentId = props.filterForm.departmentId
    showCreateModal.value = true
    getEmployee()
    scope.value = getMonthAndYear(props.filterForm.scope)

    if (props.isUpdate) {
        const res: any = await apiGetScheduleByEmployeeId({
            ...props.filterForm,
            employeeId: props?.filterForm?.employeeId || employeeOptions?.value[0]?.value,
        })
        employeeSchedule.value = res as Schedule

        createFormData.value.employeeId = employeeSchedule.value.employeeId
        // convert date into date picker
        const datesArray = JSON.parse(employeeSchedule.value.dates.replace(/'/g, '"')) as number[]
        const curScop = getMonthAndYear(props.filterForm.scope)
        const dateObjects = datesArray.map((day) => new Date(curScop.year, curScop.month, day))
        date.value = dateObjects

        //convert time into time picker
        time.value = [
            getTimesPickerValue(employeeSchedule.value.clockInTime),
            getTimesPickerValue(employeeSchedule.value.clockOutTime),
        ]
    }
}

const onSubmitCreate = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                if (props.isUpdate) {
                    const res: any = await apiUpdateSchedule(createFormData.value)
                    message.success(res?.msg || 'Schedule Updated')
                } else {
                    const res: any = await apiCreateSchedule(createFormData.value)
                    message.success(res?.msg || 'Schedule Created')
                }
                emit('currentDateChange', new Date(scope.value.year, scope.value.month, 1))
                emit('onDepartmentChange', createFormData.value.departmentId)
                if (
                    props.filterForm.scope != createFormData.value.scope &&
                    props.filterForm.departmentId != createFormData.value.departmentId
                ) {
                    emit('refreshData')
                }
                closeCreateModal()
            } catch (error: any) {
                message.error(error.message)
            } finally {
                loading.value = false
            }
        } else {
            console.log(errors)
        }
    })
}

const getEmployee = async () => {
    try {
        loading.value = true
        const res: any = await apiAllEmployee({ departmentId: createFormData.value.departmentId })
        const employees = res as Employee[]
        employeeOptions.value = !props.isUpdate ? [{ label: 'All', value: undefined }] : []
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
    if (!dates) {
        return ''
    }
    dates.forEach((date) => {
        formatDates.push(date.getDate().toString())
    })

    return formatDates.toLocaleString()
}

const updateScope = (value: { month: number; year: number } | undefined) => {
    if (value) {
        createFormData.value.scope = moment(`${value.year}-${value.month + 1}`).format('YYYY-MM')
        if (scope.value?.year != undefined && scope.value?.month != undefined) {
            datePickerStartDate.value = new Date(parseInt(scope.value.year.toString()), scope.value.month)
        }
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
