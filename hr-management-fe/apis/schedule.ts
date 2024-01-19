import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const apiCreateSchedule = async (params: CreateScheduleParams) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/schedule', {
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: params,
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
