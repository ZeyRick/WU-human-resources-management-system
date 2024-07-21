import type { DataTableColumns } from "naive-ui";
import type { RowData } from "naive-ui/es/data-table/src/interface";

export const tableColumns: DataTableColumns<RowData> = [
    {
        title: 'Username',
        key: 'username',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Name',
        key: 'name',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Type',
        key: 'userLevel',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return i18n.global.t(data.userLevel) || '-'
        },
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
