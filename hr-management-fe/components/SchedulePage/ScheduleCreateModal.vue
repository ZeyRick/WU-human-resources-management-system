<template>
    <n-button
        size="large"
        strong
        style="background-color: #409eff"
        color="#5cb85c"
        text-color="#000000"
        @click="openCreateModal"
        :disabled="(props.isUpdate && !filterForm.employeeId) || disable"
    >
        <template #icon>
            <n-icon v-if="props.isUpdate" color="#000000">
                <ReloadSharp />
            </n-icon>
            <n-icon v-else color="#000000">
                <AddCircleOutline />
            </n-icon>
        </template>
        {{ props.isUpdate ? 'Update' : 'Add' }}
    </n-button>
    <n-modal
        :show="showCreateModal"
        :mask-closable="false"
        @negative-click="closeCreateModal"
        @positive-click="onSubmitCreate"
    >
        <div style="display: flex">
            <n-card
                style="width: 600px; height: 750px"
                :title="isUpdate ? 'Update Schedule' : 'Add New Schedule'"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <n-form ref="createFormRef" :rules="CommonFormRules" :model="createFormData">
                    <n-tabs
                        style="height: 590px"
                        pane-wrapper-style="height: 100%"
                        pane-style="height: 100%;"
                        type="segment"
                        animated
                        :on-update:value="(value: string) => (curTab = value)"
                    >
                        <n-tab-pane name="employees" tab="Employees">
                            <n-form-item :label="i18n.global.t('select_department')">
                                <n-select
                                    :disabled="loading || props.isUpdate"
                                    v-model:value="createFormData.departmentId"
                                    filterable
                                    :placeholder="i18n.global.t('department')"
                                    :options="[{ label: 'All', value: '' }, ...props.departmentOptions]"
                                    @update:value="getEmployee"
                                />
                            </n-form-item>
                            <n-form-item
                                style="height: 100%"
                                path="employeeId"
                                :label="i18n.global.t('select_employee')"
                            >
                                <n-transfer
                                    v-model:value="createFormData.employeeId"
                                    style="height: 100%"
                                    ref="transfer"
                                    virtual-scroll
                                    :options="employeeOptions"
                                    target-filterable
                                    source-filterable
                                />
                            </n-form-item>
                        </n-tab-pane>
                        <n-tab-pane name="schedule" tab="Schedule">
                            <div style="display: flex; flex-direction: row; gap: 50px">
                                <n-form-item
                                    style="flex: 1"
                                    path="clockOutTime"
                                    :label="i18n.global.t('click_in_out_time')"
                                >
                                    <VueDatePicker
                                        class="timePicker"
                                        ref="timesPicker"
                                        :clearable="false"
                                        v-model="time"
                                        time-picker
                                        dark
                                        range
                                        @open="() => (timePickerOpen = true)"
                                        @closed="() => (timePickerOpen = false)"
                                    >
                                        <template #input-icon>
                                            <n-icon size="25px"> <TimerOutline /></n-icon>
                                        </template>
                                    </VueDatePicker>
                                </n-form-item>
                                <n-form-item
                                    style="flex: 1"
                                    path="minuteBreakTime"
                                    :label="i18n.global.t('break_time_minute')"
                                >
                                    <n-input-number
                                        size="large"
                                        default-value="1"
                                        :loading="loading"
                                        :input-props="{ 'auto-complete': 'off' }"
                                        v-model:value="breakTime"
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
                                                @select="handleBreakTimeOptionsSelect"
                                            >
                                                <n-button>{{ i18n.global.t(selectedBreakTimeOption) }}</n-button>
                                            </n-dropdown>
                                        </template></n-input-number
                                    >
                                </n-form-item>
                            </div>
                            <n-form-item
                                path="dates"
                                :label="`${i18n.global.t('employee_work_days')}${
                                    isUpdate ? ` ${createFormData.scope}` : ''
                                }`"
                            >
                                <VueDatePicker
                                    :style="timePickerOpen ? 'display: none;' : ''"
                                    class="datePicker"
                                    ref="datesPicker"
                                    :format="dateFormat"
                                    v-model="date"
                                    :enable-time-picker="false"
                                    required
                                    dark
                                    inline
                                    :disable-month-year-select="isUpdate"
                                    auto-apply
                                    multi-dates
                                    :month-change-on-arrows="false"
                                    :clearable="false"
                                    hide-offset-dates
                                    :month-change-on-scroll="false"
                                    focus-start-date
                                    :start-date="createFormData.scope"
                                    week-start="0"
                                    @update-month-year="updateScope"
                                >
                                </VueDatePicker>
                            </n-form-item>
                        </n-tab-pane>
                    </n-tabs>
                    <div style="display: flex; gap: 10px; justify-content: flex-end; margin-top: 20px">
                        <n-button v-if="curTab == 'schedule'" size="large" round @click="onSelectAllDates">
                            Select all
                        </n-button>
                        <n-button size="large" round @click="closeCreateModal"> Cancel </n-button>
                        <n-button size="large" round @click="onSubmitCreate">
                            {{ isUpdate ? 'Update' : 'Add' }}
                        </n-button>
                    </div>
                </n-form>
            </n-card>
        </div>
    </n-modal>
</template>

<script setup lang="ts">
import { AddCircleOutline, CalendarNumberOutline, CalendarOutline, TimerOutline, ReloadSharp } from '@vicons/ionicons5'
import { useMessage, type FormInst, type FormValidationError } from 'naive-ui'
import { apiAllEmployee } from '~/apis/employee'
import { apiCreateSchedule, apiGetScheduleByEmployeeId, apiUpdateSchedule } from '~/apis/schedule'
import { CommonFormRules } from '~/constants/formRules'
import type { EmployeeWithSchedule } from '~/types/employee'
import type { CreateScheduleParams, ScheduleFilterParams, Schedule } from '~/types/schedule'
import VueDatePicker, { type DatePickerInstance } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import moment from 'moment'

const timePickerOpen = ref<boolean>(false)
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
const props = defineProps<{
    isUpdate: boolean
    employeeOptions: { label: string; value: string }[]
    departmentOptions: { label: string; value: string }[]
    filterForm: ScheduleFilterParams
    disable?: boolean
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
const employeeOptions = ref<{ label: string; value: string | undefined; disabled?: boolean }[]>(
    props.isUpdate ? props.employeeOptions.slice(1) : [],
)
const showCreateModal = ref<boolean>(false)
const createFormData = ref<CreateScheduleParams>({
    scope: '',
    employeeId: [],
    minuteBreakTime: 1,
    dates: '',
    clockInTime: '',
    clockOutTime: '',
    departmentId: props.filterForm.departmentId || props.departmentOptions[0]?.value,
})
const loading = ref<boolean>(false)
const curTab = ref<string>('employees')
const time = ref<{ hours: number; minutes: number; seconds: number }[]>()
const date = ref()
const breakTime = ref<number>(1)
const monthYearPicker = ref<DatePickerInstance>()
const timesPicker = ref<DatePickerInstance>()
const datesPicker = ref<DatePickerInstance>()
const message = useMessage()
const closeCreateModal = () => {
    createFormData.value = {
        scope: '',
        dates: '',
        employeeId: [],
        minuteBreakTime: breakTime.value,
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

const handleBreakTimeOptionsSelect = (key: string) => {
    if (selectedBreakTimeOption.value != key) {
        switch (key) {
            case 'minute':
                breakTime.value = Math.ceil(breakTime.value * 60)
                break
            case 'hour':
                breakTime.value = Math.ceil(breakTime.value / 60)
                break
        }
    }
    selectedBreakTimeOption.value = key
}
const openCreateModal = async () => {
    createFormData.value.departmentId = props.filterForm.departmentId
    createFormData.value.scope = moment().format('YYYY-MM')
    showCreateModal.value = true
    getEmployee()

    if (props.isUpdate) {
        try {
            const res: any = await apiGetScheduleByEmployeeId({
                ...props.filterForm,
                employeeId: props?.filterForm?.employeeId || employeeOptions?.value[0]?.value,
            })
            employeeSchedule.value = res as Schedule

            createFormData.value.employeeId = [employeeSchedule.value.employeeId]
            createFormData.value.scope = props.filterForm.scope
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
        } catch (error) {
            closeCreateModal()
            console.error(error)
        }
    }
}

const onSubmitCreate = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                createFormData.value.minuteBreakTime =
                    selectedBreakTimeOption.value == 'hours' ? breakTime.value * 60 : breakTime.value
                let res: any
                if (props.isUpdate) {
                    res = await apiUpdateSchedule(createFormData.value)
                } else {
                    res = await apiCreateSchedule(createFormData.value)
                }

                if (res) {
                    const curScop = getMonthAndYear(createFormData.value.scope)
                    emit('currentDateChange', new Date(curScop.year, curScop.month, 1))
                    emit('onDepartmentChange', createFormData.value.departmentId)
                    if (
                        props.filterForm.scope == createFormData.value.scope &&
                        props.filterForm.departmentId == createFormData.value.departmentId
                    ) {
                        emit('refreshData')
                    }
                    closeCreateModal()
                }
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

const onSelectAllDates = () => {
    const curScop = getMonthAndYear(createFormData.value.scope)
    const daysInMonth = new Date(curScop.year, curScop.month + 1, 0).getDate()
    let dateObjects: Date[] = []
    for (let i = 0; i < daysInMonth; i++) {
        dateObjects.push(new Date(curScop.year, curScop.month, i + 1))
    }
    date.value = dateObjects
}

const getEmployee = async () => {
    try {
        loading.value = true
        const res: any = await apiAllEmployee({
            departmentId: createFormData.value.departmentId,
            scope: createFormData.value.scope,
        })
        const employees = res as EmployeeWithSchedule[]
        employeeOptions.value = []
        employees.map((e) => {
            employeeOptions.value.push({
                label: `${e.id} - ${e.name}`,
                value: e.id,
                disabled: e.schedule.scheduleId != 0,
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

const updateScope = ({ instance, month, year }: { instance: number; month: number; year: number }) => {
    datesPicker.value?.clearValue()
    createFormData.value.scope = moment(`${year}-${month + 1}`).format('YYYY-MM')
    getEmployee()
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
</script>

<style>
.datePicker {
    --dp-menu-min-width: 430px; /*Adjust the min width of the menu*/
    --dp-cell-size: 50px; /*Width and height of calendar cell*/
    --dp-font-size: 1rem; /*Default font-size*/
}

.timePicker {
    --dp-menu-min-width: 230px; /*Adjust the min width of the menu*/
    --dp-font-size: 1.1rem; /*Default font-size*/
}

.dp__input {
    text-align: center;
}
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
