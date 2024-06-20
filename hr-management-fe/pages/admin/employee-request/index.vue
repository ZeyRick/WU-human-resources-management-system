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
                        :options="[{ label: 'All', value: '' }, ...courseOptions]"
                    />
                </div>
                <div
                    style="margin-left: 15px; font-size: 16px; display: flex; align-items: center; white-space: nowrap"
                >
                    <div>Employee Name:</div>
                    <n-input
                        style="margin-left: 10px"
                        :disable="loading"
                        v-model:value="filterForm.employeeName"
                        :placeholder="i18n.global.t('employee_name')"
                    />
                </div>
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
    </n-layout>
</template>

<script setup lang="ts">
import { clockColumns } from './table-columns'
import { apiDeleteEmployee, apiEditEmployee } from '~/apis/employee'
import { apiListEmployeeRequest, apiDenyEmployeeRequest, apiApproveEmployeeRequest } from '~/apis/employeeRequest'
import type { Employee, EmployeeParams, CreateEmployeeType } from '~/types/employee'
import type { Course } from '~/types/course'
import { apiAllCourse } from '~/apis/course'
import OperateButton from '~/components/OperateButton/OperateButton.vue'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import type { DataTableColumns, FormInst, FormValidationError } from 'naive-ui'
import { CommonFormRules } from '~/constants/formRules'
import { AddCircleOutline } from '@vicons/ionicons5'
import NormalButton from '~/components/OperateButton/NormalButton.vue'

const courseOptions = ref<{ label: string; value: string }[]>([])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const employeeData = ref<Employee[]>([])
const defaultCreateData: CreateEmployeeType = {
    name: '',
    courseId: '',
}
const createFormData = ref<CreateEmployeeType>(defaultCreateData)
const showModal = ref<boolean>(false)
const isEdit = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const totalPage = ref(0)
const selectedEmployee = ref<Employee | null>(null)
const columns: DataTableColumns<RowData> = [
    ...clockColumns,
    {
        title: 'Operate',
        key: 'operate',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: any) => {
            return [
                h(OperateButton, {
                    text: 'Approve',
                    loading: loading.value,
                    positiveClick: async () => {
                        await apiApproveEmployeeRequest(data.id)
                        await fetchData()
                    },
                }),
                h(NormalButton, {
                    text: 'Deny',
                    loading: loading.value,
                    style: 'margin-left: 10px;',
                    onClick: () => {
                        onDenyClick(data.id)
                    },
                }),
            ]
        },
    },
]

const filterForm = reactive<EmployeeParams>({
    employeeName: '',
    courseId: '',
})

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

const getCourse = async () => {
    try {
        loading.value = true
        const res: any = await apiAllCourse()
        const courses = res as Course[]
        courseOptions.value
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

const onCourseChange = (value: any) => {
    filterForm.courseId = value
}

const onDenyClick = async (requestId: string) => {
    try {
        await apiDenyEmployeeRequest(requestId)
        await fetchData()
    } catch (error) {
        console.error(error)
    }
}

const onSubmitEdit = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                if (!selectedEmployee.value?.id) {
                    return
                }
                loading.value = true
                await apiEditEmployee(selectedEmployee.value?.id, createFormData.value)
                createFormData.value = defaultCreateData
                showModal.value = false
                await fetchData()
            } catch (error) {
                console.error(error)
            } finally {
                loading.value = false
            }
        } else {
            console.log(errors)
        }
    })
}

const showCreateModal = () => {
    showModal.value = true
    isEdit.value = false
}

const showEditModal = () => {
    showModal.value = true
    isEdit.value = true
}

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiListEmployeeRequest(pageOption.value, filterForm)
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
    await getCourse()
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
