import type { DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'

export const reportTableColumns: DataTableColumns<RowData> = [
    {
        title: 'ID',
        key: 'employeeId',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Name',
        key: 'employeeName',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Department',
        key: 'departmentAlias',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Working Hours',
        key: 'totalWorkMinute',
        titleAlign: 'center',
        align: 'center',
    },
    // You don't have an equivalent for 'Attandance' in the response
    // {
    //     title: 'Attandance',
    //     key: 'attandance',
    //     titleAlign: 'center',
    //     align: 'center',
]
