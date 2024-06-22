<template>
    <n-modal :show="showModal" :mask-closable="false" :on-after-leave="() => (createFormData = defaultCreateData)">
        <n-card
            style="width: 600px"
            :title="isEdit ? 'Edit Employee' : 'Create New Employee'"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
            <n-form ref="formRef" :rules="CommonFormRules" :model="createFormData">
                <n-form-item path="employeeType" label="Employee Type">
                    <n-select
                        disabled
                        v-model:value="EMPLOYEE_TYPE.STAFF"
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
                <n-form-item v-if="employeeType != EMPLOYEE_TYPE.STAFF" path="courseIds" label="Course">
                    <n-select
                        multiple
                        :disable="loading"
                        v-model:value="createFormData.courseIds"
                        filterable
                        :placeholder="i18n.global.t('course')"
                        :options="courseOptions"
                    />
                </n-form-item>
                <n-form-item v-if="employeeType != EMPLOYEE_TYPE.STAFF" path="degreeIds" label="Degree">
                    <n-select
                        multiple
                        :disable="loading"
                        v-model:value="createFormData.degreeIds"
                        filterable
                        :placeholder="i18n.global.t('degree')"
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
                            :file-list="files.idFile"
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
                            :file-list="files.profileFile"
                            list-type="image-card"
                            :on-finish="(val: any) => onFinishUploadFile(val, 'photoFileName')"
                            :on-error="onErrorUploadFile"
                            max="1"
                        />
                    </n-form-item>
                </n-form-item>
            </n-form>
            <div style="display: flex; gap: 10px; justify-content: flex-end">
                <n-button :loading="loading" round @click="() => $emit('closeModal')"> Cancel </n-button>
                <n-button :loading="loading" round @click="onSubmit">
                    {{ isEdit ? 'Edit' : 'Create' }}
                </n-button>
            </div>
        </n-card> </n-modal
    >,
</template>

<script setup lang="ts">
import { darkTheme, createDiscreteApi, type FormInst, type FormValidationError, type UploadFileInfo } from 'naive-ui'
import { apiCreateEmployee, apiEditEmployee } from '~/apis/employee'
import { EMPLOYEE_TYPE, type CreateEmployeeType, type EmployeeWithFile } from '~/types/employee'
import { CommonFormRules } from '~/constants/formRules'
import { useAuthStore } from '~/store/auth'
import { apiAllCourse } from '~/apis/course'
import type { Course } from '~/types/course'

type props = {
    loading: boolean
    showModal: boolean
    isEdit?: boolean
    employeeType: EMPLOYEE_TYPE

    //for edit
    selectedEmployee?: EmployeeWithFile
}

const emit = defineEmits(['fetchData', 'loading', 'closeModal'])

const propsData = withDefaults(defineProps<props>(), {})

const defaultCreateData: CreateEmployeeType = {
    name: '',
    courseIds: [],
    degreeIds: [],
    salary: 0,
    employeeType: propsData.employeeType,
    idNumber: '',
    idFileName: '',
    photoFileName: '',
}
const createFormData = ref<CreateEmployeeType>(defaultCreateData)
const courseOptions = ref<{ label: string; value: string }[]>([])
const formRef = ref<FormInst>()
const files = reactive<{ idFile: UploadFileInfo[]; profileFile: UploadFileInfo[] }>({ idFile: [], profileFile: [] })
const { token } = useAuthStore()
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

const onSubmit = () => {
    formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                if (propsData.isEdit && !propsData.selectedEmployee?.id) {
                    return
                }
                propsData.loading = true
                if (propsData.isEdit) apiEditEmployee(propsData.selectedEmployee?.id || '', createFormData.value)
                else await apiCreateEmployee(createFormData.value)

                emit('closeModal')
                emit('fetchData')
            } catch (error) {
                console.error(error)
            } finally {
                propsData.loading = false
            }
        } else {
            console.log(errors)
        }
    })
}

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

const getCourse = async () => {
    try {
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
    }
}

watch(
    () => propsData.showModal,
    () => {
        if (propsData.showModal) onShow()
    },
)

const onShow = async () => {
    await Promise.all([getCourse()])
    if (propsData.isEdit) {
        files.idFile = []
        files.profileFile = []
        const config = useRuntimeConfig()
        createFormData.value = {
            name: propsData.selectedEmployee?.name || '',
            courseIds: propsData.selectedEmployee?.courses?.map((course) => course.id) || [],
            degreeIds: propsData.selectedEmployee?.courses?.map((course) => course.id) || [] || [],
            salary: propsData.selectedEmployee?.salary || 0,
            employeeType: propsData.selectedEmployee?.type || EMPLOYEE_TYPE.STAFF,
            idFileName: propsData.selectedEmployee?.idFileName || '',
            photoFileName: propsData.selectedEmployee?.photoFileName || '',
            idNumber: propsData.selectedEmployee?.idNumber,
        }
        if (propsData.selectedEmployee?.idFileName) {
            files.idFile = [
                {
                    id: 'idFile',
                    name: propsData.selectedEmployee?.idFileName,
                    status: 'finished',
                    url: `${config.public.apiURL}/public/images/employee/${propsData.selectedEmployee?.idFileName}`,
                },
            ]
        }
        if (propsData.selectedEmployee?.photoFileName) {
            files.profileFile = [
                {
                    id: 'photoFile',
                    name: propsData.selectedEmployee?.photoFileName,
                    status: 'finished',
                    url: `${config.public.apiURL}/public/images/employee/${propsData.selectedEmployee?.photoFileName}`,
                },
            ]
        }
    } else {
        createFormData.value = defaultCreateData
        files.idFile = []
        files.profileFile = []
    }
}
</script>
