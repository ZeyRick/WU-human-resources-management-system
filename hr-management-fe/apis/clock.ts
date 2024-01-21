import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'
import { privateRequest } from '~/utils/request'

export const apiClock = async (params: ClockParams) => {
    return privateRequest('/mobile/clock', {
        method: 'post',
        body: params,
    })
}

export const apiGetClock = async (pageOpt: Pagination, filter: ClockFilter) => {
    return privateRequest('/admin/clock', {
        method: 'get',
        query: { ...pageOpt, ...filter, date: moment(filter.date).startOf('day').utc().format(DATE_TIME_FORMAT) },
    })
}
