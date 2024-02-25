import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'
import type { AttendenceFilter, ClockFilter, ClockParams } from '~/types/clock'
import { privateRequest } from '~/utils/request'

export const apiClock = async (params: ClockParams) => {
    return privateRequest(
        '/mobile/clock',
        {
            method: 'post',
            body: params,
        },
        'createClock',
    )
}

export const apiGetClock = async (pageOpt: Pagination, filter: ClockFilter) => {
    return privateRequest(
        '/admin/clock',
        {
            method: 'get',
            query: { ...pageOpt, ...filter, date: moment(filter.date).startOf('day').utc().format(DATE_TIME_FORMAT) },
        },
        'getClock',
    )
}

export const apiExportAttendence = async (filter: AttendenceFilter) => {
    return privateRequest(
        '/admin/clock/attendence/export',
        {
            method: 'get',
            query: {
                ...filter,
                startDate: moment(parseInt(filter.startDate)).utc().format(DATE_TIME_FORMAT),
                endDate: moment(parseInt(filter.endDate)).utc().format(DATE_TIME_FORMAT),
            },
        },
        'apiExportAttendence',
    )
}

export const apiGetAttendence = async (pageOpt: Pagination, filter: AttendenceFilter) => {
    return privateRequest(
        '/admin/clock/attendence',
        {
            method: 'get',
            query: {
                ...pageOpt,
                ...filter,
                startDate: moment(parseInt(filter.startDate)).utc().format(DATE_TIME_FORMAT),
                endDate: moment(parseInt(filter.endDate)).utc().format(DATE_TIME_FORMAT),
            },
        },
        'apiGetAttendence',
    )
}
