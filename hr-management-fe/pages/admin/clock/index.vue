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
        <n-data-table :bordered="false" :loading="loading" :columns="clockColumns" :data="clockDatas"/>
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
import { clockColumns } from './table-columns'
import { CommonFormRules } from '../../../constants/formRules'
import { type FormInst, type FormValidationError } from 'naive-ui'
import { apiCreateUser } from '../../../apis/user'
import { apiGetClock } from '~/apis/clock'

const showCreateModal = ref<boolean>(false)
const createFormRef = ref<FormInst>()
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const clockDatas = ref([])

onMounted(async () => {
    try {
        const res: any = await apiGetClock(pageOption.value)
        const jsonRes = JSON.parse(res)
        clockDatas.value = jsonRes.Data.Data
    } catch (error) {
    } finally {
        loading.value = false
    }
}),
    definePageMeta({
        layout: 'main',
    })
</script>
