export const apiClock = async (params: ClockParams) => {
    const config = useRuntimeConfig()
    return $fetch('/mobile/clock', {
        mode: 'no-cors',
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: params,
    })
}

export const apiGetClock = async (params: Pagination) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/clock/list', {
        mode: 'no-cors',
        method: 'get',
        baseURL: String(config.public.apiURL),
        query: params,
    })
}
