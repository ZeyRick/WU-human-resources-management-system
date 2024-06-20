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
                <div style="margin-right: 15px; font-size: 16px; display: flex; align-items: center">
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
        <n-modal :show="showModal" :mask-closable="false" :on-after-leave="() => (createFormData = defaultCreateData)">
            <n-card
                style="width: 600px"
                :title="isEdit ? 'Edit Employee' : 'Create New Employee'"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <n-form ref="createFormRef" :rules="CommonFormRules" :model="createFormData">
                    <n-form-item path="employeeType" label="Employee Type">
                        <n-select
                            disabled
                            v-model:value="EMPLOYEE_TYPE.FULL_TIME"
                            filterable
                            :placeholder="i18n.global.t('employeeType')"
                        />
                    </n-form-item>
                    <n-form-item path="name" label="Full Name">
                        <n-input
                            :loading="loading"
                            :input-props="{ 'auto-complete': 'off' }"
                            v-model:value="createFormData.name"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="salary" label="Salary">
                        <n-input-number
                            style="width: 100%"
                            v-model:value="createFormData.salary"
                            :loading="loading"
                            :min="0"
                            :precision="2"
                            :input-props="{ 'auto-complete': 'off' }"
                            @keydown.enter.prevent
                        >
                            <template #suffix> USD </template>
                        </n-input-number>
                    </n-form-item>
                    <n-form-item path="courseId" label="Course">
                        <n-select
                            :disable="loading"
                            v-model:value="createFormData.courseId"
                            filterable
                            :placeholder="i18n.global.t('course')"
                            :options="courseOptions"
                        />
                    </n-form-item>
                    <n-form-item path="idNumber" label="ID Number">
                        <n-input
                            :loading="loading"
                            :input-props="{ 'auto-complete': 'off' }"
                            v-model:value="createFormData.idNumber"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item>
                        <n-form-item style="width: 50%" path="employeeIdFile" label="Employee ID Card">
                            <n-upload
                                :action="`${$config.public.apiURL}/admin/employee/uploadFiles`"
                                :headers="{
                                    Authorization: token ? `Bearer ${token}` : '',
                                }"
                                with-credentials
                                :default-file-list="defaultIdFileList"
                                list-type="image-card"
                                :on-finish="(val: any) => onFinishUploadFile(val, 'idFileName')"
                                :on-error="onErrorUploadFile"
                                max="1"
                            />
                        </n-form-item>
                        <n-form-item style="width: 50%" path="employeePhotoFile" label="Employee Photo">
                            <n-upload
                                :action="`${$config.public.apiURL}/admin/employee/uploadFiles`"
                                :headers="{
                                    Authorization: token ? `Bearer ${token}` : '',
                                }"
                                with-credentials
                                :default-file-list="defaultPhotoFileList"
                                list-type="image-card"
                                :on-finish="(val: any) => onFinishUploadFile(val, 'photoFileName')"
                                :on-error="onErrorUploadFile"
                                max="1"
                            />
                        </n-form-item>
                    </n-form-item>
                </n-form>
                <div style="display: flex; gap: 10px; justify-content: flex-end">
                    <n-button :loading="loading" round @click="() => (showModal = false)"> Cancel </n-button>
                    <n-button :loading="loading" round @click="() => (isEdit ? onSubmitEdit() : onSubmitCreate())">
                        {{ isEdit ? 'Edit' : 'Create' }}
                    </n-button>
                </div>
            </n-card>
        </n-modal>
    </n-layout>
</template>

<script setup lang="ts">
import { employeeColumns } from './table-columns'
import { apiListEmployee, apiDeleteEmployee, apiCreateEmployee, apiEditEmployee } from '~/apis/employee'
import type { Employee, EmployeeParams, CreateEmployeeType } from '~/types/employee'
import { EMPLOYEE_TYPE } from '~/types/employee'
import type { Course } from '~/types/course'
import { apiAllCourse } from '~/apis/course'
import OperateButton from '~/components/OperateButton/OperateButton.vue'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import {
    createDiscreteApi,
    darkTheme,
    jaJP,
    type DataTableColumns,
    type FormInst,
    type FormValidationError,
    type UploadFileInfo,
    type UploadInst,
} from 'naive-ui'
import { CommonFormRules } from '~/constants/formRules'
import { AddCircleOutline, CloseCircleOutline, Options } from '@vicons/ionicons5'
import NormalButton from '~/components/OperateButton/NormalButton.vue'
import { useAuthStore } from '~/store/auth'

const { token } = useAuthStore()
const defaultIdFileList = ref<UploadFileInfo[]>([])
const defaultPhotoFileList = ref<UploadFileInfo[]>([])
const courseOptions = ref<{ label: string; value: string }[]>([])
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
const defaultCreateData: CreateEmployeeType = {
    name: '',
    courseId: '',
    salary: 0,
    employeeType: EMPLOYEE_TYPE.FULL_TIME,
    idNumber: '',
    idFileName: '',
    photoFileName: '',
}
const createFormData = ref<CreateEmployeeType>(defaultCreateData)

const onErrorUploadFile = (options: { file: UploadFileInfo; event?: any }) => {
    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })

    const res: { code: number; msg: string } = JSON.parse(options.event?.target.response)
    if (res?.code === -1) {
        message.error(`Upload file failed: ${res.msg || 'Someting Went Wrong'}`)
    }
}
const onFinishUploadFile = (
    options: { file: UploadFileInfo; event?: any },
    fileKey: 'photoFileName' | 'idFileName',
) => {
    // msg is file name
    const res: { code: number; msg: string } = JSON.parse(options.event?.target.response)
    switch (fileKey) {
        case 'photoFileName':
            createFormData.value.photoFileName = res.msg
            break
        case 'idFileName':
            createFormData.value.idFileName = res.msg
            break
    }
}
const showModal = ref<boolean>(false)
const isEdit = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const totalPage = ref(0)
const selectedEmployee = ref<Employee | null>(null)
const columns: DataTableColumns<RowData> = [
    ...employeeColumns,
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
    employeeType: EMPLOYEE_TYPE.FULL_TIME,
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

const onSubmitCreate = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                await apiCreateEmployee(createFormData.value)
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

const onSubmitEdit = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                if (!selectedEmployee.value?.id) {
                    return
                }
                loading.value = true
                await apiEditEmployee(selectedEmployee.value?.id, createFormData.value)
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
    createFormData.value = defaultCreateData
    defaultIdFileList.value = []
    defaultPhotoFileList.value = []
    showModal.value = true
    isEdit.value = false
}

const resetFilter = () => Object.assign(filterForm, defaultFilterForm)

const showEditModal = (data: any) => {
    defaultIdFileList.value = []
    defaultPhotoFileList.value = []
    const config = useRuntimeConfig()
    createFormData.value = {
        name: data?.name,
        courseId: data?.courseId,
        salary: data?.salary,
        employeeType: data?.employeeType,
        idFileName: data?.idFileName,
        photoFileName: data?.photoFileName,
        idNumber: data?.idNumber,
    }
    if (data?.idFileName) {
        defaultIdFileList.value = [
            {
                id: 'idFile',
                name: data?.idFileName,
                status: 'finished',
                url: `${config.public.apiURL}/public/images/employee/${data?.idFileName}`,
            },
        ]
    }
    if (data?.photoFileName) {
        defaultPhotoFileList.value = [
            {
                id: 'photoFile',
                name: data?.photoFileName,
                status: 'finished',
                url: `${config.public.apiURL}/public/images/employee/${data?.photoFileName}`,
            },
        ]
    }
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
