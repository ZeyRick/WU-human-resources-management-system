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
                        :courseOptions="courseOptions"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @on-course-change="onCourseChange"
                        @refresh-data="fetchData"
                    />
                </div>
                <div>
                    <ScheduleCreateModal
                        :is-update="false"
                        :courseOptions="courseOptions"
                        :filter-form="filterForm"
                        :employeeOptions="employeeOptions"
                        @currentDateChange="currentDateChange"
                        @on-course-change="onCourseChange"
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
import { apiAllCourse } from '~/apis/course'
import type { EmployeeWithSchedule } from '~/types/employee'
import type { Course } from '~/types/course'
import type { ScheduleFilterParams, ScheduleInfo } from '~/types/schedule'
import ScheduleCreateModal from '~/components/SchedulePage/ScheduleCreateModal.vue'
import moment from 'moment'

const currentDate = ref<Date>(new Date())
const employeeOptions = ref<{ label: string; value: string }[]>([])
const courseOptions = ref<{ label: string; value: string }[]>([])
const loading = ref<boolean>(true)
const scheduleDatas = ref<ScheduleInfo[]>([])
const diableUpdate = ref<boolean>(true)

const filterForm = reactive<ScheduleFilterParams>({
    scope: moment().format('YYYY-MM'),
    courseId: '',
    employeeId: '',
})

const currentDateChange = (newDate: Date) => {
    currentDate.value = newDate
    filterForm.scope = moment(currentDate.value).format('YYYY-MM')
}

const getCourse = async () => {
    try {
        loading.value = true
        const res: any = await apiAllCourse()
        const courses = res as Course[]
        courseOptions.value = [{ label: 'All', value: '' }]
        filterForm.courseId = courses[0].id || ''
        courses.map((e) => {
            courseOptions.value.push({
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
        const res: any = await apiAllEmployee({ courseId: filterForm.courseId })
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

const onCourseChange = (value: any) => {
    filterForm.courseId = value
    getEmployee()
}

watch(filterForm, fetchData)

onMounted(async () => {
    await getCourse()
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
