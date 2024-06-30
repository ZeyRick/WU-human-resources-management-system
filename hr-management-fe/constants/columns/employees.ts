import StatusChip from '~/components/StatusChip/StatusChip.vue'
import { type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { BIND_STATUS_ENUM, EMPLOYEE_TYPE } from '~/types/employee'
import i18n from '~/utils/i18n'
import StatusMark from '~/components/StatusMark/StatusMark.vue'

export const employeeColumns = (employeeType: EMPLOYEE_TYPE): any => [
    {
        title: i18n.global.t('id'),
        key: 'ID',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: number) => {
            return data.id || '-'
        },
    },
    {
        title: i18n.global.t('employee_name'),
        key: 'Name',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: number) => {
            return data.name || '-'
        },
    },
    ...(employeeType !== EMPLOYEE_TYPE.LECTURE
        ? [
              {
                  title: i18n.global.t('salary'),
                  key: 'Salary',
                  titleAlign: 'center',
                  align: 'center',
                  render: (data: any, index: number) => {
                      return `${data?.salary} $` || '-'
                  },
              },
          ]
        : []),
    // {
    //     title: i18n.global.t('employeeType'),
    //     key: 'EmployeeType',
    //     titleAlign: 'center',
    //     align: 'center',
    //     render: (data: any, index: number) => {
    //         return data?.employeeType || '-'
    //     },
    // },

    {
        title: 'ID Number',
        key: 'idNumber',
        titleAlign: 'center',
        align: 'center',
        render: (data: any, index: number) => {
            return data?.idNumber || '-'
        },
    },
    ...(employeeType === EMPLOYEE_TYPE.STAFF || employeeType === EMPLOYEE_TYPE.TEACHING_STAFF
        ? [
              //   {
              //       title: i18n.global.t('telegram_id'),
              //       key: 'TelegramId',
              //       titleAlign: 'center',
              //       align: 'center',
              //       render: (data: any, index: number) => {
              //           return data?.telegramId || '-'
              //       },
              //   },
              //   {
              //       title: i18n.global.t('telegram_username'),
              //       key: 'TelegramUsername',
              //       titleAlign: 'center',
              //       align: 'center',
              //       render: (data: any, index: number) => {
              //           return data?.telegramUsername || '-'
              //       },
              //   },
              {
                  title: i18n.global.t('telegram'),
                  key: 'telegram',
                  titleAlign: 'center',
                  align: 'center',
                  render: (data: any, index: number) => {
                      return h(StatusMark, { status: data?.telegramId })
                  },
              },
          ]
        : []),
]
