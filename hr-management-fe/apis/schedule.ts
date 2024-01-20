import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'
import type { CreateScheduleParams, ScheduleFilterParams } from '~/types/schedule'

export const apiCreateSchedule = async (params: CreateScheduleParams, departmentId: string) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/schedule', {
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: {
            departmentId,
            ...params,
            clockInTime: moment(params.clockInTime, 'HH:mm:ss').utc().format(DATE_TIME_FORMAT),
            clockOutTime: moment(params.clockOutTime, 'HH:mm:ss').utc().format(DATE_TIME_FORMAT),
        },
    })
}

export const apiGetSchedule = async (filter: ScheduleFilterParams) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/schedule', {
        method: 'get',
        headers: {
            Accept: '*/*',
        },
        baseURL: String(config.public.apiURL),
        query: { ...filter },
        onResponse({ response }) {
            if (response.status === 401) {
                const authCookie = useCookie('auth')
                authCookie.value = null
                navigateTo('/login')
            }
        },
    })
}
