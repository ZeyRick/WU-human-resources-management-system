import type { DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'

export const degreeColumns: DataTableColumns<RowData> = [
    {
        title: 'ID',
        key: 'id',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Alias',
        key: 'alias',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Rate',
        key: 'rate',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Created At',
        key: 'createdAt',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return aslocalTime(data.createdAt) || '-'
        },
    },
]
