<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px">
        <div
            style="
                flex-direction: row;
                display: flex;
                align-items: center;
                justify-content: space-between;
                overflow: hidden;
                margin-bottom: 20px;
            "
        >
            <div
                style="
                    flex-direction: row;
                    display: flex;
                    align-items: center;
                    justify-content: start;
                    overflow: hidden;
                "
            >
                <div style="font-size: 16px; display: flex; align-items: center; margin-left: 10px">
                    {{ i18n.global.t('staff') }}:
                    <n-select
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.employeeId"
                        filterable
                        :placeholder="i18n.global.t('employee')"
                        :options="employeeOptions"
                    />
                </div>
            </div>
            <div style="display: flex">
                <div style="margin-right: 10px">
                    <ScheduleCreateModal
                        :is-update="true"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @refresh-data="fetchData"
                        :disable="employees?.some(e => e.id === filterForm.employeeId && e.schedules.length < 1)"
                        :employees="employees"
                    />
                </div>
                <div>
                    <ScheduleCreateModal
                        :is-update="false"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @refresh-data="fetchData"
                         :employees="employees"
                    />
                </div>
            </div>
        </div>
        <Calendar :current-date="currentDate" @currentDateChange="currentDateChange" :schedules="scheduleDatas" />
    </n-layout>
</template>

<script setup lang="ts">
import { apiGetSchedule } from '~/apis/schedule'
import { apiAllEmployee } from '~/apis/employee'
import type { ScheduleFilterParams, ScheduleInfo } from '~/types/schedule'
import ScheduleCreateModal from '~/components/SchedulePage/ScheduleCreateModal.vue'
import moment from 'moment'
import { EMPLOYEE_TYPE, type Employee } from '~/types/employee'

const currentDate = ref<Date>(new Date())
const employeeOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const scheduleDatas = ref<ScheduleInfo[]>([])
const disableUpdate = ref<boolean>(true)
const employees = ref<Employee[]>([])

const filterForm = reactive<ScheduleFilterParams>({
    scope: moment().format('YYYY-MM'),
    employeeId: '',
})

const currentDateChange = (newDate: Date) => {
    currentDate.value = newDate
    filterForm.scope = moment(currentDate.value).format('YYYY-MM')
}


const getEmployee = async () => {
    try {
        loading.value = true
        employees.value = (await apiAllEmployee({
            employeeType: EMPLOYEE_TYPE.STAFF,
        })) as Employee[]
        employeeOptions.value = [{ label: 'All', value: '' }]
        filterForm.employeeId = ''
        employees.value.map((e) => {
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

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiGetSchedule(filterForm)
        scheduleDatas.value = res
        disableUpdate.value = true
    } catch (error) {
        console.error(error)
    } finally {
        loading.value = false
    }
}

watch(filterForm, fetchData)

onMounted(async () => {
    await getEmployee()
    await fetchData()
}),
    definePageMeta({
        layout: 'main',
    })
</script>

<style>
.n-data-table {
    border-top: 220px !important;
    overflow: hidden;
}
</style>
