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
                <div style="font-size: 16px; display: flex; align-items: center">
                    Department:
                    <n-select
                        @update:value="onDepartmentChange"
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.departmentId"
                        filterable
                        :placeholder="i18n.global.t('department')"
                        :options="departmentOptions"
                    />
                </div>
                <div style="font-size: 16px; display: flex; align-items: center; margin-left: 10px">
                    Employee:
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
                        :departmentOptions="departmentOptions"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @on-department-change="onDepartmentChange"
                        @refresh-data="fetchData"
                    />
                </div>
                <div>
                    <ScheduleCreateModal
                        :is-update="false"
                        :departmentOptions="departmentOptions"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @on-department-change="onDepartmentChange"
                        @refresh-data="fetchData"
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
import { apiAllDepartment } from '~/apis/department'
import type { EmployeeWithSchedule } from '~/types/employee'
import type { Department } from '~/types/department'
import type { ScheduleFilterParams, ScheduleInfo } from '~/types/schedule'
import ScheduleCreateModal from '~/components/SchedulePage/ScheduleCreateModal.vue'
import moment from 'moment'

const currentDate = ref<Date>(new Date())
const employeeOptions = ref<{ label: string; value: string }[]>([])
const departmentOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const scheduleDatas = ref<ScheduleInfo[]>([])
const diableUpdate = ref<boolean>(true)

const filterForm = reactive<ScheduleFilterParams>({
    scope: moment().format('YYYY-MM'),
    departmentId: '',
    employeeId: '',
})

const currentDateChange = (newDate: Date) => {
    currentDate.value = newDate
    filterForm.scope = moment(currentDate.value).format('YYYY-MM')
}

const getDepartment = async () => {
    try {
        loading.value = true
        const res: any = await apiAllDepartment()
        const departments = res as Department[]
        filterForm.departmentId = departments[0].id || ''
        departments.map((e) => {
            departmentOptions.value.push({
                label: `${e.id} - ${e.alias}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
        loading.value = false
    }
}
const getEmployee = async () => {
    try {
        loading.value = true
        const res: any = await apiAllEmployee({ departmentId: filterForm.departmentId })
        const employees = res as EmployeeWithSchedule[]
        employeeOptions.value = [{ label: 'All', value: '' }]
        filterForm.employeeId = ''
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

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiGetSchedule(filterForm)
        scheduleDatas.value = res
        diableUpdate.value = true
    } catch (error) {
        console.error(error)
    } finally {
        loading.value = false
    }
}

const onDepartmentChange = (value: any) => {
    filterForm.departmentId = value
    getEmployee()
}

watch(filterForm, fetchData)

onMounted(async () => {
    await getDepartment()
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
