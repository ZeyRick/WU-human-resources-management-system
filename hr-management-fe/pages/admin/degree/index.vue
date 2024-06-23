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
            <div style="font-size: 16px; display: flex; align-items: center; white-space: nowrap">
                <div>{{ i18n.global.t('degree') }} Name:</div>
                <n-input
                    style="margin-left: 10px"
                    :disable="loading"
                    v-model:value="filterForm.alias"
                    :placeholder="i18n.global.t('degree_alias')"
                />
            </div>
            <n-button
                :loading="loading"
                size="large"
                strong
                style="background-color: #409eff"
                color="#5cb85c"
                text-color="#000000"
                @click="showCreateModal"
            >
                <template #icon>
                    <n-icon color="#000000">
                        <AddCircleOutline />
                    </n-icon>
                </template>
                Create
            </n-button>
        </div>
        <n-data-table :loading="loading" size="large" style="margin-top: 20px" :columns="columns" :data="degreeData" />
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
        <n-modal
            :show="showModal"
            :mask-closable="false"
            :on-after-leave="
                () =>
                    (createForm = {
                        alias: '',
                        rate: 0,
                    })
            "
        >
            <n-card
                style="width: 600px"
                :title="isEdit ? 'Edit Degree' : 'Create New Degree'"
                :bordered="false"
                size="huge"
                role="dialog"
                aria-modal="true"
            >
                <n-form ref="formRef" :rules="CommonFormRules" :model="createForm">
                    <n-form-item path="alias" label="Alias">
                        <n-input
                            :loading="loading"
                            :input-props="{ 'auto-complete': 'off' }"
                            v-model:value="createForm.alias"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="rate" label="Rate">
                        <n-input-number
                            style="width: 100%"
                            v-model:value="createForm.rate"
                            :loading="loading"
                            placeholder="From"
                            :min="0"
                            :precision="2"
                            :input-props="{ 'auto-complete': 'off' }"
                            @keydown.enter.prevent
                        >
                            <template #suffix> USD </template>
                        </n-input-number>
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
import { AddCircleOutline } from '@vicons/ionicons5'
import { CommonFormRules } from '../../../constants/formRules'
import { type DataTableColumns, type FormInst, type FormValidationError } from 'naive-ui'
import { apiDelUser } from '../../../apis/user'
import { apiCreateDeparment, apiListDegree, apiEditDegree } from '../../../apis/degree'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import NormalButton from '~/components/OperateButton/NormalButton.vue'
import type { CreateDegreeParams, Degree } from '~/types/degree'
import { degreeColumns } from '~/constants/columns/degree'
const showModal = ref<boolean>(false)
const isEdit = ref<boolean>(false)
const formRef = ref<FormInst>()
const degreeData = ref<Degree[]>()
const pageOption = ref<Pagination>({ page: 1, size: 10 })
const totalPage = ref(0)
const loading = ref<boolean>(false)
const defaultCreateData: CreateDegreeParams = {
    alias: '',
    rate: 0,
}
const createForm = ref<CreateDegreeParams>(defaultCreateData)
const filterForm = reactive({
    alias: '',
})
const selectedDegree = ref<Degree | null>(null)
const columns: DataTableColumns<RowData> = [
    ...degreeColumns,
    {
        title: 'Operate',
        key: 'operate',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: any) => {
            return [
                h(NormalButton, {
                    text: 'Edit',
                    loading: loading.value,
                    style: 'margin-left: 10px;',
                    onClick: () => {
                        createForm.value = { alias: data?.alias, rate: data?.rate || 0 }
                        selectedDegree.value = data
                        showEditModal()
                    },
                }),
            ]
        },
    },
]

// const handleDelete = async (userId: string) => {
//     try {
//         loading.value = true
//         const res: any = await apiDelUser(userId)
//         await fetchData()
//     } catch (error) {
//     } finally {
//         loading.value = false
//     }
// }

const fetchData = async () => {
    try {
        loading.value = true
        const res: any = await apiListDegree(pageOption.value, filterForm)
        totalPage.value = res.pageOpt.totalPage
        degreeData.value = res.data as Degree[]
        pageOption.value = {
            size: res.pageOpt.pageSize,
            page: res.pageOpt.curPage,
        }
    } catch (error) {
    } finally {
        loading.value = false
    }
}

const onSubmitCreate = () => {
    formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                await apiCreateDeparment(createForm.value)
                createForm.value = defaultCreateData
                showModal.value = false
                await fetchData()
            } catch (error) {
                console.error(error)
            } 
        } else {
            console.log(errors)
            loading.value = false
        }
    })
}

const onSubmitEdit = () => {
    formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                if (!selectedDegree.value?.id) {
                    return
                }
                loading.value = true
                await apiEditDegree(selectedDegree.value?.id, createForm.value)
                createForm.value = defaultCreateData
                showModal.value = false
                await fetchData()
            } catch (error) {
                console.error(error)
                loading.value = false
            } 
        } else {
            console.log(errors)
        }
    })
}

onMounted(async () => {
    await fetchData()
})

const showCreateModal = () => {
    showModal.value = true
    isEdit.value = false
}

const showEditModal = () => {
    showModal.value = true
    isEdit.value = true
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

definePageMeta({
    layout: 'main',
})
</script>
