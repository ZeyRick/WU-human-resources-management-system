<template>
    <n-layout>
        <n-card>
            <div style="display: flex; align-items: center; gap: 10px">
                <n-date-picker
                    v-model:value="range"
                    type="datetimerange"
                    :is-date-disabled="(timeStamp: number) => timeStamp > moment().endOf('days').valueOf()"
                    :default-time="['00:00:00', '23:59:59']"
                    placeholder="Select a date"
                />
                <div style="font-size: 16px; display: flex; align-items: center; white-space: nowrap">
                    Staff :
                    <n-select
                        :disable="loading"
                        v-model:value="filterForm.employeeId"
                        filterable
                        :placeholder="i18n.global.t('employee')"
                        :options="employeeOptions"
                    />
                </div>
            </div>
            <n-button
                :loading="loading"
                size="large"
                strong
                style="background-color: #409eff; margin-top: 20px; margin-bottom: 20px"
                color="#5cb85c"
                text-color="#000000"
                @click="onExportClick"
            >
                Export
            </n-button>
            <n-data-table :columns="columns" :loading="loading" :data="reportDatas" :bordered="false" />
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
        </n-card>
    </n-layout>
</template>

<script setup lang="ts">
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { NLayout, NInput, NSelect, NCard, NText, type DataTableColumns } from 'naive-ui'
import { type Employee, type EmployeeParams, type CreateEmployeeType, EMPLOYEE_TYPE } from '~/types/employee'
import type { Course } from '~/types/course'
import { apiAllCourse } from '~/apis/course'
import { apiGetReport } from '~/apis/report'
import './index.css'
import moment from 'moment'
import { reportTableColumns } from './report-table-columns'
import type { Report, ReportFilter } from '~/types/report'
import { DATE_TIME_FORMAT } from '~/constants/time'
import { apiAllEmployee } from '~/apis/employee'

const totalPage = ref(0)
const loading = ref<boolean>(true)
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const courseOptions = ref<{ label: string; value: string }[]>([])

const reportDatas = ref<Report[]>([])

const employeeOptions = ref<{ label: string; value: string }[]>([])
const columns: DataTableColumns<RowData> = [...reportTableColumns]

const getEmployee = async () => {
    try {
        const res: any = await apiAllEmployee({ employeeType: [EMPLOYEE_TYPE.LECTURE, EMPLOYEE_TYPE.TEACHING_STAFF] })
        const employees = res as Employee[]
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
    }
}

const range = ref<number[]>([moment().startOf('days').valueOf(), moment().endOf('days').valueOf()])
const filterForm = reactive<ReportFilter>({
    employeeId: '',
    isTeaching: true,
    startDate: range.value[0].toString(),
    endDate: range.value[1].toString(),
})

watch(range, () => {
    filterForm.startDate = range.value[0].toString()
    filterForm.endDate = range.value[1].toString()
})

const fetchReport = async () => {
    try {
        reportDatas.value = []
        const res: any = await apiGetReport(pageOption.value, filterForm)
        if (res) {
            totalPage.value = res.pageOpt.totalPage
            reportDatas.value = (res.data as Report[]).filter((report) => report.totalWorkMinute > 0) || []
            pageOption.value = {
                size: res.pageOpt.pageSize,
                page: res.pageOpt.curPage,
            }
        }
    } catch (error) {
    } finally {
    }
}

const onPageChange = (page: number) => {
    pageOption.value.page = page
    fetchReport()
}

const onPageSizeChange = (pageSize: number) => {
    pageOption.value.size = pageSize
    fetchReport()
}

const onExportClick = () => {
    try {
        loading.value = true
        const config = useRuntimeConfig()
        const params = new URLSearchParams({
            employeeId: filterForm.employeeId,
            startDate: moment(parseInt(filterForm.startDate)).utc().format(DATE_TIME_FORMAT),
            endDate: moment(parseInt(filterForm.endDate)).utc().format(DATE_TIME_FORMAT),
        })

        params.append('isTeaching', String(filterForm.isTeaching))

        const exportUrl = `${String(config.public.apiURL)}/admin/clock/report/export?${params}`
        window.open(exportUrl, '_self')
    } catch (error) {
        console.error(error)
    } finally {
        loading.value = false
    }
}

onMounted(async () => {
    try {
        loading.value = true
        await Promise.all([getEmployee(), fetchReport()])
    } catch (error) {
    } finally {
        loading.value = false
    }
}),
    watch(filterForm, fetchReport)
definePageMeta({
    layout: 'main',
})
</script>
