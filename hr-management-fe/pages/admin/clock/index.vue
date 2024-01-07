<template>
    <n-layout style="flex-grow: 1; display: flex; flex-direction: column">
        <n-card content-style="padding: 10px;" style="height: 50px; overflow: hidden">
            <div
                style="
                    flex-direction: row;
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    overflow: hidden;
                "
            >
                <n-text type="primary" style="font-size: 18px">Time Clock</n-text>
            </div></n-card
        >
        <!-- <n-data-table :loading="loading" :columns="tableColumns" :data="null" /> -->
        <n-card
            content-style="padding: 10px;"
            style="display: flex; align-items: center; height: 50px; overflow: hidden"
        >
            <n-pagination
                v-model:page-size="pageOption.size"
                v-model:page="pageOption.page"
                :page-count="100"
                show-size-picker
                :page-sizes="[10, 20, 30, 40]"
            />
        </n-card>
    </n-layout>
</template>

<script setup lang="ts">
import { AddCircleOutline } from '@vicons/ionicons5'
import { tableColumns } from './table-columns'
import { CommonFormRules } from '../../../constants/formRules'
import { type FormInst, type FormValidationError } from 'naive-ui'
import { apiCreateUser } from '../../../apis/user'
import { apiGetClock } from '~/apis/clock'

const showCreateModal = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(false)

onMounted( async () => {
    try {
        loading.value = true
        const res = await apiGetClock(pageOption.value)
        console.log(res)
    } catch (error) {} finally {
      loading.value = false
    }
}),
    definePageMeta({
        layout: 'main',
    })
</script>
