import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const apiClock = async (params: ClockParams) => {
    const config = useRuntimeConfig()
    return $fetch('/mobile/clock', {
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: params,
    })
}

export const apiGetClock = async (pageOpt: Pagination, filter: ClockFilter) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/clock', {
        method: 'get',
        headers: {
            Accept: '*/*',
        },
        baseURL: String(config.public.apiURL),
        query: { ...pageOpt, ...filter, date: moment(filter.date).startOf('day').utc().format(DATE_TIME_FORMAT) },
        onResponse({ response }) {
            if (response.status === 401) {
                const authCookie = useCookie('auth')
                authCookie.value = null
                navigateTo('/login')
            }
        },
    })
}
