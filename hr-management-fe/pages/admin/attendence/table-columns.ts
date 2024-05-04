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
        render: (data, index) => aslocalTime(data.createdAt, DATE_FORMAT),
    },
    {
        title: i18n.global.t('employee_name'),
        key: 'EmployeeName',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => data.employee.name || '-',
    },
    {
        title: i18n.global.t('clock_in_time'),
        key: 'ClockInTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => aslocalTime(data.clockIn.createdAt, TIME_FORMAT),
    },
    {
        title: i18n.global.t('clock_out_time'),
        key: 'ClockOutTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => aslocalTime(data.createdAt, TIME_FORMAT),
    },
    {
        title: i18n.global.t('total_work_minute'),
        key: 'WorkMinute',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => data.clockOutMinute || 0,
    },
    {
        title: i18n.global.t('work_time'),
        key: 'WorkTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) =>
            `${aslocalTime(data.schedule.clockInTime, TIME_FORMAT)}-${aslocalTime(
                data.schedule.clockOutTime,
                TIME_FORMAT,
            )}`,
    },
    {
        title: i18n.global.t('status'),
        key: 'Status',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            const isLate = data.clockIn.status == 'late'
            const isEarly = data.status == 'early'
            if (isLate && isEarly) {
                return h('div', {
                    innerHTML: 'Late-Early',
                    style: 'color: #730000',
                })
            } else if (isLate) {
                return h('div', {
                    innerHTML: 'Late',
                    style: 'color: red',
                }) 
            } else if (isEarly) {
                return h('div', {
                    innerHTML: 'Early',
                    style: 'color: blue',
                }) 
            } else {
                return h('div', {
                    innerHTML: 'On Time',
                    style: 'color: green',
                }) 
            }
        },
    },
]
