import type { DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'

export const reportTableColumns: DataTableColumns<RowData> = [
    {
        title: 'Employee ID',
        key: 'employeeId',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Employee Name',
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
        title: 'Working Minutes',
        key: 'totalWorkMinute',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Total Late Minutes',
        key: 'totalLateMinute',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Total Early Minutes',
        key: 'totalEarlyMinute',
        titleAlign: 'center',
        align: 'center',
    },
]
