<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px">
        <div style="flex-direction: row; display: flex; align-items: center; justify-content: start; overflow: hidden">
            <!-- <div style="font-size: 16px; display: flex; align-items: center">
                Course:
                <n-select
                    @update:value="onCourseChange"
                    style="margin-left: 10px"
                    :disable="loading"
                    v-model:value="filterForm.courseId"
                    filterable
                    :placeholder="i18n.global.t('course')"
                    :options="courseOptions"
                />
            </div> -->
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
            <div style="font-size: 16px; display: flex; align-items: center; margin-left: 10px">
                {{ i18n.global.t('range') }}::
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
            :disabled="attendenceData.length < 1"
            :loading="loading"
            size="large"
            strong
            style="background-color: #409eff; margin-top: 20px"
            color="#5cb85c"
            text-color="#000000"
            @click="onExportClick"
        >
            {{ i18n.global.t('export') }}
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
import { apiAllCourse } from '~/apis/course'
import type { Course } from '~/types/course'
import type { AttendenceFilter, Clock } from '~/types/clock'
import type { ScheduleInfo } from '~/types/schedule'
import moment from 'moment'
import { attendenceColumns } from './table-columns'
import { DATE_TIME_FORMAT } from '~/constants/time'
import { EMPLOYEE_TYPE, type Employee } from '~/types/employee'

const employeeOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const attendenceData = ref<Clock[]>([])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const totalPage = ref(0)
const range = ref<number[]>([moment().startOf('days').valueOf(), moment().endOf('days').valueOf()])
const filterForm = reactive<AttendenceFilter>({
    startDate: range.value[0].toString(),
    endDate: range.value[1].toString(),
    employeeName: '',
    employeeId: '',
    employeeType: [EMPLOYEE_TYPE.STAFF, EMPLOYEE_TYPE.TEACHING_STAFF],
    isTeaching: false
})

watch(range, () => {
    filterForm.startDate = range.value[0].toString()
    filterForm.endDate = range.value[1].toString()
})

// const getCourse = async () => {
//     try {
//         loading.value = true
//         const res: any = await apiAllCourse()
//         const courses = res as Course[]
//         courseOptions.value = [{ label: 'All', value: '' }]
//         filterForm.courseId = courses[0].id || ''
//         courses.map((e) => {
//             courseOptions.value.push({
//                 label: `${e.id} - ${e.alias}`,
//                 value: e.id,
//             })
//         })
//     } catch (error) {
//     } finally {
//         loading.value = false
//     }
// }

const onExportClick = () => {
    try {
        loading.value = true
        const config = useRuntimeConfig()
        const searchParams = new URLSearchParams({
            employeeName: filterForm.employeeName,
            employeeId: filterForm.employeeId,
            startDate: moment(parseInt(filterForm.startDate)).utc().format(DATE_TIME_FORMAT),
            endDate: moment(parseInt(filterForm.endDate)).utc().format(DATE_TIME_FORMAT),
        })
        for (let i = 0; i < filterForm.employeeType.length; i++) {
            searchParams.append('employeeType', filterForm.employeeType[i])
        }
        const exportUrl = `${String(config.public.apiURL)}/admin/clock/attendence/export?${searchParams.toString()}`
        window.open(exportUrl, '_self')
    } catch (error) {
        console.error(error)
    } finally {
        loading.value = false
    }
}

const getEmployee = async () => {
    try {
        const employees: Employee[] = (await apiAllEmployee({
            employeeType: [EMPLOYEE_TYPE.STAFF, EMPLOYEE_TYPE.TEACHING_STAFF],
        })) as Employee[]
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

const fetchData = async () => {
    try {
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
    }
}

// const onCourseChange = (value: any) => {
//     filterForm.courseId = value
//     getEmployee()
// }

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
    try {
        loading.value = true
        await Promise.all([getEmployee(), fetchData()])
    } catch (error) {
    } finally {
        loading.value = false
    }
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
