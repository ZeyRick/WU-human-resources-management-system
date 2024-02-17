<template>
    <div :style="`background-color: ${DARK_THEME.SIDE_BAR_COLOR}; padding: 10px; border-radius: 10px`">
        <div style="margin-bottom: 10px; display: flex; justify-content: space-between">
            <div style="font-size: 22px">{{ MONTH[month - 1] }} {{ year }}</div>
            <n-text type="primary" style="font-size: 24px">Employees Work Schedule</n-text>
            <div style="font-size: 22px; display: flex; display: flex; justify-content: center; align-items: center">
                <n-button @click="() => updateMonthClick(-1)">
                    <n-icon><ChevronBackOutline /></n-icon>
                </n-button>
                <n-button @click="() => (currentDate = new Date())"> Today </n-button>
                <n-button @click="() => updateMonthClick(1)">
                    <n-icon>
                        <ChevronForwardOutline />
                    </n-icon>
                </n-button>
            </div>
        </div>
        <n-grid :x-gap="10" :y-gap="10" :cols="7">
            <n-grid-item v-for="week in weeks" :key="`wk-${week}`">
                <div
                    :style="`background-color: ${BUTTON_COLOR}; text-overflow: ellipsis; overflow: hidden; white-space: nowrap; padding: 10px;`"
                    :class="'week-cell'"
                >
                    {{ i18n.global.t(week) }}
                </div>
            </n-grid-item>
            <n-grid-item v-for="index in firstDayOfMonth" :key="`fodm-${index}`">
                <div :class="'disabled-date-cell'" />
            </n-grid-item>
            <n-grid-item v-for="(schdule, index) in schedules" :key="index">
                <div
                    :onclick="schdule.employees?.length > 0 && (() => onDateCellClick(index + 1, schdule))"
                    :class="`transform ${getDateClass(index + 1)}`"
                >
                    <div>{{ index + 1 }}</div>

                    <div
                        v-if="schdule.employees && schdule.employees.length <= 2"
                        style="margin-bottom: 5px"
                        v-for="employee in schdule.employees"
                    >
                        <n-tag strong size="small" :bordered="false" type="success"> {{ employee.name }} </n-tag>
                    </div>
                    <div v-else-if="schdule.employees && schdule.employees.length > 2">
                        <div style="margin-bottom: 5px">
                            <n-tag strong size="small" :bordered="false" type="success">
                                {{ schdule.employees[0].name }}
                            </n-tag>
                        </div>
                        <div style="margin-bottom: 5px">
                            <n-tag strong size="small" :bordered="false" type="success">
                                {{ schdule.employees[1].name }}
                            </n-tag>
                        </div>
                        <div style="margin-bottom: 5px">
                            <n-tag strong size="small" :bordered="false" type="success">
                                {{ `${schdule.employees.length - 1} Others...` }}
                            </n-tag>
                        </div>
                    </div>
                </div>
            </n-grid-item>
            <n-grid-item
                v-for="index in () => {
                    return 35 - (daysInMonth || 0) - (firstDayOfMonth || 0)
                }"
                :key="`nm-${index}`"
            >
                <div :class="'disabled-date-cell'" />
            </n-grid-item>
        </n-grid>
        <n-modal v-model:show="showDetails">
            <n-card
                style="width: 600px"
                title="Details"
                header-style="font-size: 20px;"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <div style="font-size: 18px; display: flex; justify-content: space-between">
                    <div>Date :</div>
                    <div>{{ `${selectedCell?.date} ${MONTH[month - 1]} ${year}` }}</div>
                </div>
                <div style="font-size: 18px; display: flex; justify-content: space-between">
                    <div>Total Employees Count :</div>
                    <div>{{ selectedCell?.schedule.employees.length }} Employees</div>
                </div>
                <n-divider />
                <div style="font-size: 18px; display: flex; justify-content: space-between">
                    <div style="flex: 1">Employee's Name</div>
                    <div style="flex: 1; text-align: center">Clock In Time</div>
                    <div style="flex: 1; text-align: end">Clock Out Time</div>
                </div>
                <div
                    style="font-size: 18px; display: flex; justify-content: space-between"
                    v-for="(employee, index) in selectedCell?.schedule.employees"
                >
                    <div style="flex: 1">{{ index + 1 }}. {{ employee.name }}</div>
                    <div style="flex: 1; text-align: center">
                        {{ format(new Date(employee.clockInTime), 'H:mm') }}
                    </div>
                    <div style="flex: 1; text-align: end">{{ format(new Date(employee.clockOutTime), 'H:mm') }}</div>
                </div>
            </n-card>
        </n-modal>
    </div>
</template>

<script setup lang="ts">
import { BUTTON_COLOR, DARK_THEME } from '~/constants/theme'
import { MONTH } from '~/constants/month'
import type { ScheduleInfo } from '~/types/schedule'
import { format } from 'date-fns'
import { ChevronBackOutline, ChevronForwardOutline } from '@vicons/ionicons5'

const props = defineProps<{
    schedules: ScheduleInfo[]
    currentDate: Date
}>()
const { schedules, currentDate } = toRefs(props)
const year = ref<number>(currentDate.value.getFullYear())
const month = ref<number>(currentDate.value.getMonth() + 1)
const daysInMonth = ref<number>(new Date(year.value, month.value, 0).getDate())
const firstDayOfMonth = ref<number>(new Date(year.value, month.value - 1, 1).getDay())

const showDetails = ref<boolean>(false)
const selectedCell = ref<{ date: number; schedule: ScheduleInfo } | null>(null)
const emit = defineEmits<{
    (e: 'currentDateChange', newDate: Date): void
}>()

const onDateCellClick = (date: number, schedule: ScheduleInfo) => {
    selectedCell.value = { date, schedule }
    showDetails.value = true
}

const updateMonthClick = (value: number) => {
    const newDate = new Date(currentDate.value)
    newDate.setMonth(newDate.getMonth() + value)
    emit('currentDateChange', newDate)
}

watch(currentDate, () => {
    year.value = currentDate.value.getFullYear()
    month.value = currentDate.value.getMonth() + 1
    daysInMonth.value = new Date(year.value, month.value, 0).getDate()
    firstDayOfMonth.value = new Date(year.value, month.value - 1, 1).getDay()
    const date = new Date(2000, month.value - 1, 1) // Subtract 1 because months are 0-based in JavaScript
})

const weeks = ['sun', 'mon', 'tue', 'wed', 'thu', 'fri', 'sat']

const getDateClass = (date: number) => {
    const curDate = new Date(year.value, month.value - 1, date).getDay()
    return curDate === 0 || curDate === 6 ? 'weekend-cell' : 'date-cell'
}
</script>

<style scoped>
.week-cell {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 20px;
    height: 40px;
}
.date-cell {
    overflow: hidden;
    border: 1px solid rgba(0, 0, 0, 0.12);
    height: 130px;
    font-size: 15px;
    background-color: rgba(108, 177, 246, 0.12);
}
.disabled-date-cell {
    border: 1px solid rgba(0, 0, 0, 0.12);
    height: 130px;
    font-size: 15px;
    background-color: rgba(205, 225, 246, 0.12);
}

.employee-data {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 5px;
    font-size: 13px;
}

.transform {
    padding: 7px;
    transition: transform 0.3s ease-in-out;
}
.transform:hover {
    transform: scale(1.2); /* Increase the scale factor as needed */
    background-color: rgba(108, 177, 246, 0.12);
    opacity: 1;
}

.weekend-cell {
    border: 1px solid rgba(0, 0, 0, 0.12);
    height: 130px;
    font-size: 15px;
    overflow: hidden;
    background-color: rgba(46, 145, 243, 0.12);
}
</style>
