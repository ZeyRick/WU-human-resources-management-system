<template>
    <n-modal :show="showModal" :mask-closable="false">
        <n-card
            style="width: 600px"
            :title="isEdit ? 'Edit Employee' : 'Create New Employee'"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
            <n-form ref="formRef" :rules="CommonFormRules" :model="createFormData">
                <n-form-item path="employeeId" label="Employee Type">
                    <n-select
                        :disable="loading"
                        v-model:value="createFormData.employeeId"
                        filterable
                        :placeholder="i18n.global.t('employee')"
                        :options="employeeOptions"
                    />
                </n-form-item>
                <n-form-item path="courseId" label="Course">
                    <n-select
                        :disabled="loading || !createFormData.employeeId"
                        v-model:value="createFormData.courseId"
                        filterable
                        :placeholder="i18n.global.t('course')"
                        :options="courseOptions"
                    />
                </n-form-item>
                <n-form-item path="degreeId" label="Degree">
                    <n-select
                        :disabled="loading || !createFormData.employeeId"
                        v-model:value="createFormData.degreeId"
                        filterable
                        :placeholder="i18n.global.t('degree')"
                        :options="degreeOptions"
                    />
                </n-form-item>
                <n-form-item path="clockDate" label="Date">
                    <n-date-picker
                        style="width: 100%"
                        v-model:formatted-value="createFormData.clockDate"
                        value-format="yyyy-MM-dd"
                        type="date"
                    />
                </n-form-item>
                <n-form-item path="clockInTime" label="Start Time">
                    <n-time-picker
                        style="width: 100%"
                        v-model:formatted-value="createFormData.clockInTime"
                        value-format="HH:mm:ss"
                    />
                </n-form-item>
                <n-form-item path="clockOutTime" label="End Time">
                    <n-time-picker
                        style="width: 100%"
                        v-model:formatted-value="createFormData.clockOutTime"
                        value-format="HH:mm:ss"
                    />
                </n-form-item>
                <n-form-item path="totalMinute" label="End Time">
                    <n-input-number
                        style="width: 100%"
                        v-model:value="createFormData.totalMinute"
                        :loading="loading"
                        :min="1"
                        :precision="0"
                        :input-props="{ 'auto-complete': 'off' }"
                        @keydown.enter.prevent
                    >
                        <template #suffix> Minute </template>
                    </n-input-number>
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
import { type FormInst, type FormValidationError, type UploadFileInfo } from 'naive-ui'
import { CommonFormRules } from '~/constants/formRules'
import { apiAllCourse, apiCourseEmployee } from '~/apis/course'
import type { Course } from '~/types/course'
import { apiAllDegree, apiDegreeEmployee } from '~/apis/degree'
import type { CreateAttendance } from '~/types/attendance'
import moment from 'moment'
import { apiEditManualClock, apiManualClock } from '~/apis/clock'

type props = {
    loading: boolean
    showModal: boolean
    isEdit?: boolean
    employeeOptions: { label: string; value: string }[]
}

const emit = defineEmits(['fetchData', 'loading', 'closeModal'])
const curTime = moment()
const propsData = withDefaults(defineProps<props>(), {})
const defaultCreateData: CreateAttendance = {
    employeeId: null,
    degreeId: null,
    courseId: null,
    clockDate: curTime.format('YYYY-MM-DD'),
    clockInTime: curTime.format('HH:mm:ss'),
    clockOutTime: curTime.format('HH:mm:ss'),
    totalMinute: 0,
}
const createFormData = ref<CreateAttendance>({...defaultCreateData})
const courseOptions = ref<{ label: string; value: string }[]>([])
const degreeOptions = ref<{ label: string; value: string }[]>([])
const formRef = ref<FormInst>()
const onSubmit = () => {
    formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                propsData.loading = true
                if (propsData.isEdit) await apiEditManualClock(createFormData.value, 0)
                else await apiManualClock(createFormData.value)

                emit('closeModal')
                emit('fetchData')
            } catch (error) {
                console.error(error)
            } finally {
                propsData.loading = false
            }
        } else {
            console.error(errors)
        }
    })
}

const getDegree = async () => {
    try {
        if (!createFormData.value.employeeId) return
        const res: any = await apiDegreeEmployee(createFormData.value.employeeId)
        const degrees = res as Course[]
        degreeOptions.value = []
        degrees.map((e) => {
            degreeOptions.value.push({
                label: `${e.id} - ${e.alias}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
    }
}

const getCourse = async () => {
    try {
        if (!createFormData.value.employeeId) return
        const res: any = await apiCourseEmployee(createFormData.value.employeeId)
        const courses = res as Course[]
        courseOptions.value = []
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
    () => createFormData.value.employeeId,
    async () => {
        try {
            propsData.loading = true
            await Promise.all([getCourse(), getDegree()])
        } catch (error) {
        } finally {
            propsData.loading = false
        }
    },
)

watch(
    () => propsData.showModal,
    () => {
        createFormData.value = {...defaultCreateData}
    },
)
</script>
