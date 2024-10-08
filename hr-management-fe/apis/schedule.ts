import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'
import type { CreateScheduleParams, ScheduleFilterParams } from '~/types/schedule'

export const apiCreateSchedule = async (params: CreateScheduleParams) => {
    const clockInTime = `${params.scope}-01 ${params.clockInTime}`
    const clockOutTime = `${params.scope}-01 ${params.clockOutTime}`
    try {
        const res = await privateRequest(
            '/admin/schedule',
            {
                method: 'post',
                body: {
                    ...params,
                    clockInTime: moment(clockInTime, 'YYYY-MM-DD HH:mm:ss').utc().format(DATE_TIME_FORMAT),
                    clockOutTime: moment(clockOutTime, 'YYYY-MM-DD HH:mm:ss').utc().format(DATE_TIME_FORMAT),
                },
            },
            'apiCreateSchedule',
        )
        return res
    } catch (error) {
        return error
    }
}

export const apiUpdateSchedule = async (params: CreateScheduleParams) => {
    try {
        const clockInTime = `${params.scope}-01 ${params.clockInTime}`
        const clockOutTime = `${params.scope}-01 ${params.clockOutTime}`
        const res = await privateRequest(
            '/admin/schedule',
            {
                method: 'patch',
                body: {
                    ...params,
                    clockInTime: moment(clockInTime, 'YYYY-MM-DD HH:mm:ss').utc().format(DATE_TIME_FORMAT),
                    clockOutTime: moment(clockOutTime, 'YYYY-MM-DD HH:mm:ss').utc().format(DATE_TIME_FORMAT),
                },
            },
            'updateSchedule',
        )
        return res
    } catch (error) {
        return error
    }
}

export const apiGetSchedule = async (filter: ScheduleFilterParams) => {
    return privateRequest(
        '/admin/schedule',
        {
            method: 'get',
            query: { ...filter },
        },
        'apiGetSchedule',
    )
}

export const apiGetScheduleByEmployeeId = async (filter: ScheduleFilterParams) => {
    try {
        return privateRequest(
            `/admin/schedule/${filter.employeeId}`,
            {
                method: 'get',
                query: { ...filter },
            },
            'getScheduleByEmployeeId',
        )
    } catch (error) {
        return error
    }
}
