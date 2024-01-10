import { NP, type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { useI18n } from 'vue-i18n'
import ClockChipVue from '~/components/Clock/ClockChip/ClockChip.vue'
import i18n from '~/utils/i18n'

export const clockColumns: DataTableColumns<RowData> = [
    {
        title: i18n.global.t('employee_name'),
        key: 'Name',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.Employee.Name || '-'
        },
    },
    {
        title: i18n.global.t('clock_time'),
        key: 'CreatedAt',
        titleAlign: 'center',
        align: 'center',
    },

    {
        title: i18n.global.t('clock_type'),
        key: 'ClockType',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return h(ClockChipVue, {
                text: data.ClockType || '-',
            })
        },
    },
]
