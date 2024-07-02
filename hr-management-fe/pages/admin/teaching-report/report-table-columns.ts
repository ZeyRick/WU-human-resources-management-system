import type { DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'

export const reportTableColumns: DataTableColumns<RowData> = [
    {
        title: 'Staff ID',
        key: 'employeeId',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Staff Name',
        key: 'employeeName',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Working Minutes',
        key: 'totalWorkMinute',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Degree',
        key: 'degreeAlias',
        titleAlign: 'center',
        align: 'center',
    },
]
