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
            <div>
                <n-select
                    :disable="loading"
                    v-model:value="filterForm.employeeId"
                    filterable
                    :placeholder="i18n.global.t('employee')"
                    :options="employeeOptions"
                />
            </div>
            <div style="display: flex">
                <n-button
                    size="large"
                    strong
                    style="background-color: #409eff"
                    color="#5cb85c"
                    text-color="#000000"
                    type="primary"
                    :disabled="loading"
                    @click="handlePreviousClick"
                    >Previous</n-button
                >
                <n-date-picker
                    size="large"
                    :disabled="loading"
                    style="width: 130px; margin: 0 10px; border-radius: 900px"
                    :on-update:value="handlerTimeChange"
                    format="yyyy-MM-dd"
                    :v-model:value="filterForm.date"
                    v-model:formatted-value="filterForm.date"
                    :is-date-disabled="(ts: number) => ts > Date.now()"
                    type="date"
                >
                    <template #date-icon> <span></span> </template>
                </n-date-picker>
                <n-button
                    size="large"
                    strong
                    style="background-color: #409eff"
                    color="#5cb85c"
                    text-color="#000000"
                    type="primary"
                    :disabled="filterForm.date == moment().format(DATE_FORMAT) || loading"
                    @click="handleNextClick"
                    >Next</n-button
                >
            </div>
        </div>
        <n-data-table size="large" :bordered="false" :loading="loading" :columns="clockColumns" :data="clockDatas" />
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
import { useLoadingBar } from 'naive-ui'
import { clockColumns } from './table-columns'
import { apiGetClock } from '~/apis/clock'
import { apiAllEmployee } from '~/apis/employee'
import { getNowLocal } from '~/utils/time'
import { DATE_FORMAT } from '~/constants/time'
import moment from 'moment'
import type { Employee } from '~/types/employee'

const employeeOptions = ref<{ label: string; value: string }[]>([])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const clockDatas = ref([])
const totalPage = ref(0)
const loadingBar = useLoadingBar()
const filterForm = reactive<ClockFilter>({
    employeeId: null,
    date: getNowLocal(DATE_FORMAT),
})

const handlePreviousClick = () =>
    (filterForm.date = moment(filterForm.date, DATE_FORMAT).add(-1, 'days').format(DATE_FORMAT))

const handleNextClick = () =>
    (filterForm.date = moment(filterForm.date, DATE_FORMAT).add(1, 'days').format(DATE_FORMAT))
const handlerTimeChange = (value: string, formatValue: string) => (filterForm.date = formatValue)

const getEmployee = async () => {
    try {
        loadingBar.start()
        loading.value = true
        const res: any = await apiAllEmployee()
        const employees = JSON.parse(res).res as Employee[]
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
        const res: any = await apiGetClock(pageOption.value, filterForm)
        const jsonRes = JSON.parse(res).res
        totalPage.value = jsonRes.pageOpt.totalPage
        clockDatas.value = jsonRes.data
        pageOption.value = {
            size: jsonRes.pageOpt.pageSize,
            page: jsonRes.pageOpt.curPage,
        }
    } catch (error) {
    } finally {
        loadingBar.finish()
        loading.value = false
    }
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

onMounted(() => {
    fetchData()
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
