import { NP, type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { useI18n } from 'vue-i18n'
import ClockChipVue from '~/components/Clock/ClockChip/ClockChip.vue'
import { DATE_FORMAT, TIME_FORMAT } from '~/constants/time'
import i18n from '~/utils/i18n'
import { aslocalTime, isTimeAfter, isTimeBefore } from '~/utils/time'

export const attendenceColumns: DataTableColumns<RowData> = [
    {
        title: i18n.global.t('date'),
        key: 'Date',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            console.log(data)
            return aslocalTime(data.createdAt, DATE_FORMAT)
        },
    },
    {
        title: i18n.global.t('employee_name'),
        key: 'EmployeeName',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.employee.name || '-'
        },
    },

    {
        title: i18n.global.t('clock_in_time'),
        key: 'ClockInTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return aslocalTime(data.clockIn.createdAt, TIME_FORMAT)
        },
    },
    {
        title: i18n.global.t('clock_out_time'),
        key: 'ClockOutTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return aslocalTime(data.createdAt, TIME_FORMAT)
        },
    },
    {
        title: i18n.global.t('total_work_minute'),
        key: 'WorkMinute',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return data.clockOutMinute || 0
        },
    },
    {
        title: i18n.global.t('work_time'),
        key: 'WorkTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return `${aslocalTime(data.schedule.clockInTime, TIME_FORMAT)}-${aslocalTime(
                data.schedule.clockOutTime,
                TIME_FORMAT,
            )}`
        },
    },
    {
        title: i18n.global.t('status'),
        key: 'Status',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            if (isTimeAfter(data.schedule.clockInTime, data.clockIn.createdAt)) {
                return 'Late'
            } else if (isTimeBefore(data.schedule.clockOutTime, data.createdAt)) {
                return 'Early'
            } else {
                return 'On Time'
            }
        },
    },
]
