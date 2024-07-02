import { NP, type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { useI18n } from 'vue-i18n'
import ClockChipVue from '~/components/Clock/ClockChip/ClockChip.vue'
import { DATE_FORMAT, TIME_FORMAT } from '~/constants/time'
import i18n from '~/utils/i18n'
import { aslocalTime, isTimeAfter, isTimeBefore } from '~/utils/time'

export const teachingAttendanceColumns: DataTableColumns<RowData> = [
    {
        title: i18n.global.t('date'),
        key: 'Date',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => aslocalTime(data.createdAt, DATE_FORMAT),
    },
    {
        title: i18n.global.t('staff_name'),
        key: 'EmployeeName',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => data.employee.name || '-',
    },
    {
        title: i18n.global.t('start_time'),
        key: 'ClockInTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => aslocalTime(data.clockIn.clockTime, TIME_FORMAT),
    },
    {
        title: i18n.global.t('end_time'),
        key: 'ClockOutTime',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => aslocalTime(data.clockTime, TIME_FORMAT),
    },
    {
        title: i18n.global.t('total_minute'),
        key: 'WorkMinute',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => data.clockOutMinute || 0,
    },
]
