<template>
    <n-modal
        size="small"
        preset="card"
        :show="show"
        :style="{ width: '600px' }"
        :title="i18n.global.t('employee')"
        @esc="closeModal"
        @close="closeModal"
        @mask-click="closeModal"
    >
        <n-table v-if="employee" :bordered="false">
            <tbody>
                <tr>
                    <td>{{ i18n.global.t('id') }}</td>
                    <td>{{ employee.id }}</td>
                </tr>
                <tr>
                    <td>{{ i18n.global.t('type') }}</td>
                    <td>{{ i18n.global.t(employee?.employeeType || '-') }}</td>
                </tr>
                <tr>
                    <td>{{ i18n.global.t('full_name') }}</td>
                    <td>{{ employee.name }}</td>
                </tr>
                <tr>
                    <td>{{ i18n.global.t('id_number') }}</td>
                    <td>
                        {{ employee.idNumber }}
                    </td>
                </tr>
                <tr
                    v-if="
                        employee.employeeType === EMPLOYEE_TYPE.STAFF ||
                        employee.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
                    "
                >
                    <td>{{ i18n.global.t('salary') }}</td>
                    <td>{{ employee.salary }}</td>
                </tr>
                <tr
                    v-if="
                        employee.employeeType === EMPLOYEE_TYPE.LECTURE ||
                        employee.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
                    "
                >
                    <td>{{ i18n.global.t('course') }}</td>
                    <td>
                        <n-button v-for="course in employee.courses" size="small" style="margin-right: 5px">
                            {{ course.alias }}</n-button
                        >
                    </td>
                </tr>
                <tr
                    v-if="
                        employee.employeeType === EMPLOYEE_TYPE.LECTURE ||
                        employee.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
                    "
                >
                    <td>{{ i18n.global.t('degree') }}</td>
                    <td>
                        <n-button v-for="degree in employee.degrees" size="small" style="margin-right: 5px">
                            {{ degree.alias }}</n-button
                        >
                    </td>
                </tr>
                <tr
                    v-if="
                        employee.employeeType === EMPLOYEE_TYPE.STAFF ||
                        employee.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
                    "
                >
                    <td>{{ i18n.global.t('telegram_id') }}</td>
                    <td>
                        {{ employee.telegramId || '-' }}
                    </td>
                </tr>
                <tr
                    v-if="
                        employee.employeeType === EMPLOYEE_TYPE.STAFF ||
                        employee.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
                    "
                >
                    <td>{{ i18n.global.t('telegram_username') }}</td>
                    <td>
                        {{ employee.telegramUsername || '-' }}
                    </td>
                </tr>
            </tbody>
        </n-table>
    </n-modal>
</template>

<script setup lang="ts">
import { NModal, NTable } from 'naive-ui'
import { EMPLOYEE_TYPE, type Employee } from '~/types/employee'

defineProps<{
    show: boolean
    employee?: Employee
}>()

const emit = defineEmits<{
    (e: 'update:show', value: boolean): void
}>()

function closeModal() {
    emit('update:show', false)
}
</script>
