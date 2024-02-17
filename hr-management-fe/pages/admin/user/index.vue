<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px">
        <div
            style="
                flex-direction: row;
                display: flex;
                align-items: center;
                justify-content: space-between;
                overflow: hidden;
                padding: 0px 10px;
            "
        >
            <n-text type="primary" style="font-size: 24px">User Table</n-text>
            <n-button
                :loading="loading"
                size="large"
                strong
                style="background-color: #409eff"
                color="#5cb85c"
                text-color="#000000"
                @click="openCreateModal"
            >
                <template #icon>
                    <n-icon color="#000000">
                        <AddCircleOutline />
                    </n-icon>
                </template>
                Create
            </n-button>
        </div>
        <n-data-table :loading="loading" size="large" style="margin-top: 20px" :columns="columns" :data="userData" />

        <n-modal
            :show="showCreateModal"
            :mask-closable="false"
            @negative-click="closeCreateModal"
            @positive-click="onSubmitCreate"
        >
            <n-card
                style="width: 600px"
                title="Create New User"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <n-form ref="createFormRef" :rules="CommonFormRules" :model="createForm">
                    <n-form-item path="userName" label="UserName">
                        <n-input
                            :loading="loading"
                            :input-props="{ 'auto-complete': 'off' }"
                            v-model:value="createForm.userName"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="name" label="Name">
                        <n-input :loading="loading" v-model:value="createForm.name" @keydown.enter.prevent />
                    </n-form-item>
                    <n-form-item path="password" label="Password">
                        <n-input
                            :loading="loading"
                            type="password"
                            v-model:value="createForm.password"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="userLeval" label="User Leval">
                        <n-select
                            :disable="loading"
                            v-model:value="createForm.userLevel"
                            filterable
                            :placeholder="i18n.global.t('department')"
                            :options="userLevalOptions"
                        />
                    </n-form-item>
                </n-form>
                <div style="display: flex; gap: 10px; justify-content: flex-end">
                    <n-button :loading="loading" round @click="closeCreateModal"> Cancel </n-button>
                    <n-button :loading="loading" round @click="onSubmitCreate"> Create </n-button>
                </div>
            </n-card>
        </n-modal>
        <n-modal
            :show="showResetPwModal"
            :mask-closable="false"
            @negative-click="closeCreateModal"
            @positive-click="onSubmitCreate"
        >
            <n-card
                style="width: 600px"
                title="Reset Password"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <n-form ref="createFormRef" :rules="CommonFormRules" :model="resetPwForm">
                    <n-form-item path="password" label="New Password">
                        <n-input
                            :loading="loading"
                            :input-props="{ 'auto-complete': 'off' }"
                            v-model:value="resetPwForm.password"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                </n-form>
                <div style="display: flex; gap: 10px; justify-content: flex-end">
                    <n-button :loading="loading" round @click="closeResetPwModal"> Cancel </n-button>
                    <n-button :loading="loading" round @click="onSubmitResetPW"> Create </n-button>
                </div>
            </n-card>
        </n-modal>
    </n-layout>
</template>

<script setup lang="ts">
import { AddCircleOutline } from '@vicons/ionicons5'
import { tableColumns } from './table-columns'
import { CommonFormRules } from '../../../constants/formRules'
import { type DataTableColumns, type FormInst, type FormValidationError } from 'naive-ui'
import { apiCreateUser, apiGetUser, apiDelUser, apiUserResetPW } from '../../../apis/user'
import type { User, CreateUserType } from '~/types/user'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import OperateButton from '~/components/OperateButton/OperateButton.vue'
import NormalButton from '~/components/OperateButton/NormalButton.vue'
import { USER_LEVEL } from '~/constants/userLevel'
const showCreateModal = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const userData = ref<User[]>()
const loading = ref<boolean>(false)
const defaultCreateData: CreateUserType = {
    userName: '',
    password: '',
    name: '',
    userLevel: USER_LEVEL.ADMIN,
}
const showResetPwModal = ref<boolean>(false)
const resetPwForm = ref<{ password: string }>({
    password: '',
})
const createForm = ref<CreateUserType>(defaultCreateData)
const selectedUser = ref<any>(null)
const columns: DataTableColumns<RowData> = [
    ...tableColumns,
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
                    text: 'Reset Password',
                    loading: loading.value,
                    style: 'margin-left: 10px;',
                    onClick: () => {
                        selectedUser.value = data
                        showResetPwModal.value = true
                    },
                }),
            ]
        },
    },
]
const userLevalOptions: { label: string; value: string }[] = [
    { label: 'Admin', value: USER_LEVEL.ADMIN },
    { label: 'Super Admin', value: USER_LEVEL.SUPER_ADMIN },
]

const handleDelete = async (userId: string) => {
    try {
        loading.value = true
        const res: any = await apiDelUser(userId)
        await fetchData()
    } catch (error) {
    } finally {
        loading.value = false
    }
}

const closeResetPwModal = () => (showResetPwModal.value = false)
const onSubmitResetPW = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                await apiUserResetPW(selectedUser?.value?.id, resetPwForm.value.password)
                resetPwForm.value.password = ''
                closeResetPwModal()
                await fetchData()
                loading.value = false
            } catch (error) {
                console.error(error)
            }
        } else {
            console.log(errors)
        }
    })
}

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiGetUser()
        userData.value = res
        loading.value = false
    } catch (error) {
    } finally {
    }
}

const onSubmitCreate = () => {
    createFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                await apiCreateUser(createForm.value)
                createForm.value = defaultCreateData
                closeCreateModal()
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

onMounted(() => {
    fetchData()
})

const closeCreateModal = () => (showCreateModal.value = false)

const openCreateModal = () => (showCreateModal.value = true)

definePageMeta({
    layout: 'main',
    middleware: ['permission']
})
</script>
