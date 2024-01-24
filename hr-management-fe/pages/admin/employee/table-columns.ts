import StatusChip from '~/components/StatusChip/StatusChip.vue'
import { type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { BIND_STATUS_ENUM } from '~/types/employee'
import i18n from '~/utils/i18n'

export const clockColumns: DataTableColumns<RowData> = [
    {
        title: i18n.global.t('id'),
        key: 'ID',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.id || '-'
        },
    },
    {
        title: i18n.global.t('employee_name'),
        key: 'Name',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            console.log(data)
            return data.name || '-'
        },
    },
    {
        title: i18n.global.t('department'),
        key: 'Department',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data?.department?.alias || '-'
        },
    },

    {
        title: i18n.global.t('status'),
        key: 'Stauts',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data?.bindingStatus
                ? h(StatusChip, {
                      text: data?.bindingStatus,
                      type: data?.bindingStatus === BIND_STATUS_ENUM.APPROVE ? 'success' : 'error',
                  })
                : '-'
        },
    },
    {
        title: i18n.global.t('telegram_id'),
        key: 'TelegramId',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data?.telegramId || '-'
        },
    },
    {
        title: i18n.global.t('telegram_username'),
        key: 'TelegramUsername',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data?.telegramUsername || '-'
        },
    },
]
