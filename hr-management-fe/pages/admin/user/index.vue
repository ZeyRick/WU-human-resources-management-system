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
        <n-data-table style="margin-top: 20px;" :columns="tableColumns" :data="userData" />

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
                <n-form ref="createFormRef" :rules="CommonFormRules" :model="createFormData">
                    <n-form-item path="userName" label="UserName">
                        <n-input :input-props="{ 'auto-complete': 'off' }"  v-model:value="createFormData.userName" @keydown.enter.prevent />
                    </n-form-item>
                    <n-form-item path="name" label="Name">
                        <n-input v-model:value="createFormData.name" @keydown.enter.prevent />
                    </n-form-item>
                    <n-form-item path="password" label="Password">
                        <n-input type="password" v-model:value="createFormData.password" @keydown.enter.prevent />
                    </n-form-item>
                </n-form>
                <div style="display: flex; gap: 10px; justify-content: flex-end">
                    <n-button round @click="closeCreateModal"> Cancel </n-button>
                    <n-button round @click="onSubmitCreate"> Create </n-button>
                </div>
            </n-card>
        </n-modal>
    </n-layout>
</template>

<script setup lang="ts">
import { AddCircleOutline } from '@vicons/ionicons5'
import { tableColumns } from './table-columns'
import { CommonFormRules } from '../../../constants/formRules'
import { type FormInst, type FormValidationError } from 'naive-ui'
import { apiCreateUser, apiGetUser } from '../../../apis/user'
import type { User, CreateUserType } from '~/types/user'
const showCreateModal = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const userData = ref<User[]>()

const defaultCreateData: CreateUserType = {
    userName: '',
    password: '',
    name: '',
}
const createFormData = ref<CreateUserType>(defaultCreateData)

const fetchData = async () => {
    try {
        const res: any = await apiGetUser()
        userData.value = res
    } catch (error) {
    } finally {
    }
}

const onSubmitCreate = () => {
    createFormRef.value?.validate((errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                apiCreateUser(createFormData.value)
                clearModalValue()
                closeCreateModal()
                fetchData()
            } catch (error) {
                console.error(error)
            }
        } else {
            console.log(errors)
        }
    })
}

onMounted(() => {
    fetchData()
})

const clearModalValue = () => (createFormData.value = defaultCreateData)

const closeCreateModal = () => (showCreateModal.value = false)

const openCreateModal = () => (showCreateModal.value = true)

definePageMeta({
    layout: 'main',
})
</script>
