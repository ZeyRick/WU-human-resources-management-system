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
        <n-data-table :bordered="false" :loading="loading" :columns="clockColumns" :data="clockDatas" />
        <n-card
            content-style="padding: 10px;"
            style="display: flex; align-items: center; height: 50px; overflow: hidden"
        >
            <n-pagination
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
import { useLoadingBar } from 'naive-ui';
import { clockColumns } from './table-columns'
import { apiGetClock } from '~/apis/clock'

const pageOption = ref<Pagination>({ page: 1, size: 10 })
const loading = ref<boolean>(true)
const clockDatas = ref([])
const totalPage = ref(0)
const loadingBar = useLoadingBar()

const fetchData = async () => {
    try {
        loadingBar.start()
        loading.value = true    
        const res: any = await apiGetClock(pageOption.value)
        const jsonRes = JSON.parse(res).Data
        totalPage.value = jsonRes.PageOpt.TotalPage
        clockDatas.value = jsonRes.Data
        pageOption.value = {
            size: jsonRes.PageOpt.PageSize,
            page: jsonRes.PageOpt.CurPage,
        }
    } catch (error) {
    } finally {
        loadingBar.finish()
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

onMounted(() => {
    fetchData()
}),
    definePageMeta({
        layout: 'main',
    })
</script>
