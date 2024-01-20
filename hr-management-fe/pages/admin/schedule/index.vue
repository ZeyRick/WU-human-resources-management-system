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
            <ScheduleCreateModal
                :departmentOptions="departmentOptions"
                :departmentId="
                    departmentOptions.find((department) => department?.value === filterForm.departmentId)?.value ||
                    departmentOptions[0]?.value
                "
                :employeeOptions="employeeOptions"
            />
        </div>
        <Calendar @scopeChange="scopeChange" :schedules="scheduleDatas" />
    </n-layout>
</template>

<script setup lang="ts">
import { useLoadingBar } from 'naive-ui'
import { apiGetSchedule } from '~/apis/schedule'
import { apiAllEmployee } from '~/apis/employee'
import { apiAllDepartment } from '~/apis/department'
import type { Employee } from '~/types/employee'
import type { Department } from '~/types/department'
import type { ScheduleFilterParams, ScheduleInfo } from '~/types/schedule'
import ScheduleCreateModal from '~/components/SchedulePage/ScheduleCreateModal.vue'
import moment from 'moment'

const employeeOptions = ref<{ label: string; value: string }[]>([])
const departmentOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const scheduleDatas = ref<ScheduleInfo[]>([])
const loadingBar = useLoadingBar()
const filterForm = reactive<ScheduleFilterParams>({
    scope: moment().format('YYYY-MM'),
    departmentId: '',
    employeeId: '',
})

const scopeChange = (scope: string) => (filterForm.scope = scope)
const getDepartment = async () => {
    try {
        loadingBar.start()
        loading.value = true
        const res: any = await apiAllDepartment()
        const departments = JSON.parse(res).res as Department[]
        filterForm.departmentId = departments[0].id || ''
        departments.map((e) => {
            departmentOptions.value.push({
                label: `${e.id} - ${e.alias}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
        loadingBar.finish()
        loading.value = false
    }
}
const getEmployee = async () => {
    try {
        loadingBar.start()
        loading.value = true
        console.log(filterForm.departmentId)
        const res: any = await apiAllEmployee({ departmentId: filterForm.departmentId })
        const employees = JSON.parse(res).res as Employee[]
        employeeOptions.value = [{ label: 'All', value: '' }]
        employees.map((e) => {
            employeeOptions.value.push({
                label: `${e.id} - ${e.name}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
        loadingBar.finish()
        loading.value = false
    }
}

const fetchData = async () => {
    try {
        loadingBar.start()
        loading.value = true
        const res: any = await apiGetSchedule(filterForm)
        const jsonRes = JSON.parse(res).res
        scheduleDatas.value = jsonRes
    } catch (error) {
        console.error(error)
    } finally {
        loadingBar.finish()
        loading.value = false
    }
}

const onDepartmentChange = (value: any) =>  {
    filterForm.departmentId = value
    getEmployee() 
}
   
watch(filterForm, fetchData)

onMounted(() => {
    fetchData()
    getDepartment()
    getEmployee()
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
