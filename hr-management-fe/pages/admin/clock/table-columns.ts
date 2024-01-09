import type { DataTableColumns } from "naive-ui";
import type { RowData } from "naive-ui/es/data-table/src/interface";
import { useI18n } from 'vue-i18n'

export const clockColumns: DataTableColumns<RowData> = [
    {
        title: 'employee_name',
        key: 'Name',
        render: (data, index) => {
            return data.Employee.Name || '-'
        },
    },
    {
        title: 'clock_time',
        key: 'CreatedAt',
    },

    {
        title: 'clock_type',
        key: 'ClockType',
        render: (data, index) => {
            return useI18n().t(data.ClockType) || '-'
        },
    },
]
