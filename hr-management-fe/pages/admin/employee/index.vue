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
                    flex-wrap: wrap;
                    row-gap: 15px;
                "
            >
                <!-- <div style="margin-right: 15px; font-size: 16px; display: flex; align-items: center">
                    Course:
                    <n-select
                        @update:value="onCourseChange"
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.courseId"
                        filterable
                        :placeholder="i18n.global.t('course')"
                        :options="[{ label: 'All', value: '' }, ...courseOptions]"
                    />
                </div> -->
                <!-- <div style="margin-right: 15px; font-size: 16px; display: flex; align-items: center">
                    Employee Type:
                    <n-select
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.employeeType"
                        filterable
                        :placeholder="i18n.global.t('employeeType')"
                        :options="[{ label: 'All', value: '' }, ...employeeTypeOptions]"
                    />
                </div> -->
                <div
                    style="margin-right: 15px; font-size: 16px; display: flex; align-items: center; white-space: nowrap"
                >
                    <div>Employee Name:</div>
                    <n-input
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.employeeName"
                        :placeholder="i18n.global.t('employee_name')"
                    />
                </div>

                <div
                    style="margin-right: 15px; font-size: 16px; display: flex; align-items: center; white-space: nowrap"
                >
                    <div>Salary:</div>
                    <n-input-number
                        style="margin-left: 10px"
                        v-model:value="filterForm.startSalary"
                        :loading="loading"
                        placeholder="From"
                        :min="0"
                        :precision="2"
                        :input-props="{ 'auto-complete': 'off' }"
                        @keydown.enter.prevent
                    >
                        <template #suffix> USD </template>
                    </n-input-number>
                    <div style="margin: 0px 10px">-</div>
                    <n-input-number
                        v-model:value="filterForm.endSalary"
                        :loading="loading"
                        placeholder="To"
                        :min="0"
                        :precision="2"
                        :input-props="{ 'auto-complete': 'off' }"
                        @keydown.enter.prevent
                    >
                        <template #suffix> USD </template>
                    </n-input-number>
                </div>
            </div>
            <div>
                <n-button
                    :loading="loading"
                    size="large"
                    strong
                    style="background-color: #409eff; width: 100%"
                    color="#5cb85c"
                    text-color="#000000"
                    @click="showCreateModal"
                    clearable
                >
                    <template #icon>
                        <n-icon color="#000000">
                            <AddCircleOutline />
                        </n-icon>
                    </template>
                    Create
                </n-button>
                <n-button
                    :loading="loading"
                    size="large"
                    strong
                    style="background-color: #ffa140; width: 100%; margin-top: 15px"
                    color="#5cb85c"
                    text-color="#000000"
                    @click="resetFilter"
                    clearable
                >
                    <template #icon>
                        <n-icon color="#000000">
                            <CloseCircleOutline />
                        </n-icon>
                    </template>
                    Clear
                </n-button>
            </div>
        </div>
        <n-data-table size="large" :bordered="false" :loading="loading" :columns="columns" :data="employeeData" />
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

        <CreateEmployee
            :loading="loading"
            :show-modal="showModal"
            :is-edit="isEdit"
            :employee-type="EMPLOYEE_TYPE.STAFF"
            :selected-employee="selectedEmployee"
            @close-modal="() => (showModal = false)"
            @fetch-data="fetchData"
        />
    </n-layout>
</template>

<script setup lang="ts">
import { employeeColumns } from './table-columns'
import { apiListEmployee, apiDeleteEmployee, apiCreateEmployee, apiEditEmployee } from '~/apis/employee'
import type { Employee, EmployeeParams, CreateEmployeeType, EmployeeWithFile } from '~/types/employee'
import { EMPLOYEE_TYPE } from '~/types/employee'
import OperateButton from '~/components/OperateButton/OperateButton.vue'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { type DataTableColumns } from 'naive-ui'
import { AddCircleOutline, CloseCircleOutline, Options } from '@vicons/ionicons5'
import NormalButton from '~/components/OperateButton/NormalButton.vue'

// const employeeTypeOptions = ref<{ label: string; value: string }[]>([
//     {
//         label: `Full Time`,
//         value: EMPLOYEE_TYPE.FULL_TIME,
//     },
//     {
//         label: `Part Time`,
//         value: EMPLOYEE_TYPE.PART_TIME,
//     },
// ])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const employeeData = ref<Employee[]>([])

const showModal = ref<boolean>(false)
const isEdit = ref<boolean>(false)

const totalPage = ref(0)
const selectedEmployee = ref<EmployeeWithFile | undefined>()
const columns: DataTableColumns<RowData> = [
    ...employeeColumns(EMPLOYEE_TYPE.STAFF),
    {
        title: 'Operate',
        key: 'operate',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: any) => {
            return [
                h(OperateButton, {
                    text: 'Remove',
                    loading: loading.value,
                    positiveClick: () => handleDelete(data.id),
                }),
                h(NormalButton, {
                    text: 'Edit',
                    loading: loading.value,
                    style: 'margin-left: 10px;',
                    onClick: () => showEditModal(data),
                }),
            ]
        },
    },
]

const defaultFilterForm: EmployeeParams = {
    employeeName: '',
    courseId: '',
    employeeType: EMPLOYEE_TYPE.STAFF,
    startSalary: null,
    endSalary: null,
    scope: '',
    id: undefined,
}
const filterForm = reactive<EmployeeParams>({ ...defaultFilterForm })
const handleDelete = async (employeeId: string) => {
    try {
        loading.value = true
        const res: any = await apiDeleteEmployee(employeeId)
        await fetchData()
    } catch (error) {
    } finally {
        loading.value = false
    }
}

const showCreateModal = () => {
    showModal.value = true
    isEdit.value = false
}

const resetFilter = () => Object.assign(filterForm, defaultFilterForm)

const showEditModal = (data: any) => {
    selectedEmployee.value = data
    showModal.value = true
    isEdit.value = true
}

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiListEmployee(pageOption.value, filterForm)
        totalPage.value = res.pageOpt.totalPage
        employeeData.value = res.data as Employee[]
        pageOption.value = {
            size: res.pageOpt.pageSize,
            page: res.pageOpt.curPage,
        }
    } catch (error) {
    } finally {
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

onMounted(async () => {
    await fetchData()
}),
    definePageMeta({
        layout: 'main',
    })
</script>

<style>
.n-data-table {
    overflow: hidden;
}
</style>
