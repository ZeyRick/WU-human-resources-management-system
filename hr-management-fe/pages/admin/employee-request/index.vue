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
                        :options="[{ label: 'All', value: '' }, ...departmentOptions]"
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
import type { Department } from '~/types/department'
import { apiAllDepartment } from '~/apis/department'
import OperateButton from '~/components/OperateButton/OperateButton.vue'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import type { DataTableColumns, FormInst, FormValidationError } from 'naive-ui'
import { CommonFormRules } from '~/constants/formRules'
import { AddCircleOutline } from '@vicons/ionicons5'
import NormalButton from '~/components/OperateButton/NormalButton.vue'

const departmentOptions = ref<{ label: string; value: string }[]>([])
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const employeeData = ref<Employee[]>([])
const defaultCreateData: CreateEmployeeType = {
    name: '',
    departmentId: '',
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
    departmentId: '',
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

const getDepartment = async () => {
    try {
        loading.value = true
        const res: any = await apiAllDepartment()
        const departments = res as Department[]
        departmentOptions.value
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

const onDepartmentChange = (value: any) => {
    filterForm.departmentId = value
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
    await getDepartment()
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
