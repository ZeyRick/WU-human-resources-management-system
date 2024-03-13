<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px">
        <div style="flex-direction: row; display: flex; align-items: center; justify-content: start; overflow: hidden">
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
            <div style="font-size: 16px; display: flex; align-items: center; margin-left: 10px">
                Range:
                <n-date-picker
                    v-model:value="range"
                    style="margin-left: 10px"
                    type="datetimerange"
                    :is-date-disabled="(timeStamp: number) => timeStamp > moment().endOf('days').valueOf()"
                    :default-time="['00:00:00', '23:59:59']"
                />
            </div>
        </div>
        <n-button
            :loading="loading"
            size="large"
            strong
            style="background-color: #409eff; margin-top: 20px"
            color="#5cb85c"
            text-color="#000000"
            @click="onExportClick"
        >
            Export
        </n-button>
        <n-data-table
            :loading="loading"
            size="large"
            style="margin-top: 20px"
            :columns="attendenceColumns"
            :data="attendenceData"
        />
        <n-card
            content-style="padding: 10px;"
            style="display: flex; align-items: center; height: 50px; overflow: hidden"
        >
            <n-pagination
                size="large"
                :disabled="loading"
                v-model:page-size="pageOption.size"
                v-model:page="pageOption.page"
                :page-count="totalPage"
                show-size-picker
                :page-sizes="[10, 4, 30, 40]"
                :on-update:page-size="onPageSizeChange"
                :on-update:page="onPageChange"
            />
        </n-card>
    </n-layout>
</template>

<script setup lang="ts">
import { apiGetAttendence, apiExportAttendence } from '~/apis/clock'
import { apiAllEmployee } from '~/apis/employee'
import { apiAllDepartment } from '~/apis/department'
import type { EmployeeWithSchedule } from '~/types/employee'
import type { Department } from '~/types/department'
import type { AttendenceFilter, Clock } from '~/types/clock'
import type { ScheduleInfo } from '~/types/schedule'
import moment from 'moment'
import { attendenceColumns } from './table-columns'
import { DATE_TIME_FORMAT } from '~/constants/time'

const currentDate = ref<Date>(new Date())
const employeeOptions = ref<{ label: string; value: string }[]>([])
const departmentOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const scheduleDatas = ref<ScheduleInfo[]>([])
const diableUpdate = ref<boolean>(true)
const attendenceData = ref<Clock[]>([])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const totalPage = ref(0)
const range = ref<number[]>([moment().startOf('days').valueOf(), moment().endOf('days').valueOf()])
const filterForm = reactive<AttendenceFilter>({
    startDate: range.value[0].toString(),
    endDate: range.value[1].toString(),
    employeeName: '',
    employeeId: '',
    departmentId: '',
})

watch(range, () => {
    filterForm.startDate = range.value[0].toString()
    filterForm.endDate = range.value[1].toString()
})

const getDepartment = async () => {
    try {
        loading.value = true
        const res: any = await apiAllDepartment()
        const departments = res as Department[]
        departmentOptions.value = [{ label: 'All', value: '' }]
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

const onExportClick = () => {
    try {
        loading.value = true
        const config = useRuntimeConfig()
        const params = {
            ...filterForm,
            startDate: moment(parseInt(filterForm.startDate)).utc().format(DATE_TIME_FORMAT),
            endDate: moment(parseInt(filterForm.endDate)).utc().format(DATE_TIME_FORMAT),
        }
        const exportUrl = `${String(config.public.apiURL)}/admin/clock/attendence/export?${new URLSearchParams(
            params,
        ).toString()}`
        window.open(exportUrl, '_self')
    } catch (error) {
        console.error(error)
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
        const res: any = await apiGetAttendence(pageOption.value, filterForm)
        if (res) {
            totalPage.value = res.pageOpt.totalPage
            attendenceData.value = res.data as Clock[]
            pageOption.value = {
                size: res.pageOpt.pageSize,
                page: res.pageOpt.curPage,
            }
        }
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

const onPageChange = (page: number) => {
    pageOption.value.page = page
    fetchData()
}

const onPageSizeChange = (pageSize: number) => {
    pageOption.value.size = pageSize
    fetchData()
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
