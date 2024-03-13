import type { ReportFilter } from '~/types/report'
import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const apiGetReport = async (pageOpt: Pagination, filter: ReportFilter) => {
    return privateRequest(
        '/admin/report',
        {
            method: 'get',
            query: {
                ...filter,
                startDate: moment(parseInt(filter.startDate)).utc().format(DATE_TIME_FORMAT),
                endDate: moment(parseInt(filter.endDate)).utc().format(DATE_TIME_FORMAT),
            },
        },
        'apiGetReport',
    )
}
