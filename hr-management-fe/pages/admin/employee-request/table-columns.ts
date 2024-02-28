import StatusChip from '~/components/StatusChip/StatusChip.vue'
import { type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { BIND_STATUS_ENUM } from '~/types/employee'
import i18n from '~/utils/i18n'

export const clockColumns: DataTableColumns<RowData> = [
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
    {
        title: i18n.global.t('employee_id'),
        key: 'ID',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.employee.id || '-'
        },
    },
    {
        title: i18n.global.t('employee_name'),
        key: 'Name',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.employee.name || '-'
        },
    },
]
