import moment from 'moment'
import { DATE_TIME_FORMAT, TIME_FORMAT } from '~/constants/time'
import type { CreateAttendance } from '~/types/attendance'
import type { AttendenceFilter, ClockFilter, ClockParams, EditClock } from '~/types/clock'
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

export const apiEditClock = async (payload: EditClock, id: number, scope: string) => {
    const newClockTime = `${scope}-01 ${payload.clockTime}`
    return privateRequest(
        `/admin/clock/${id}`,
        {
            method: 'patch',
            body: {
                clockTime: moment(newClockTime, 'YYYY-MM-DD HH:mm:ss').utc().format(TIME_FORMAT),
            },
        },
        'apiEditClock',
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

export const apiManualClock = async (payload: CreateAttendance) => {
    return privateRequest(
        `/admin/clock/manualClock`,
        {
            method: 'post',
            body: {
                ...payload,
                clockInTime: moment(`${payload.clockDate} ${payload.clockInTime}`, DATE_TIME_FORMAT)
                    .utc()
                    .format(DATE_TIME_FORMAT),
                clockOutTime: moment(`${payload.clockDate} ${payload.clockOutTime}`, DATE_TIME_FORMAT)
                    .utc()
                    .format(DATE_TIME_FORMAT),
            },
        },
        'apiManualClock',
    )
}

export const apiEditManualClock = async (payload: CreateAttendance, id: number) => {
    return privateRequest(
        `/admin/clock/manualClock/${id}`,
        {
            method: 'patch',
            body: payload,
        },
        'apiManualClock',
    )
}
