export const apiClock = async (params: ClockParams) => {
    const config = useRuntimeConfig()
    return $fetch('/mobile/clock', {
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: params,
    })
}

export const apiGetClock = async (params: Pagination) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/clock/list', {
        method: 'get',
        headers: {
            Accept: '*/*',
        },
        baseURL: String(config.public.apiURL),
        query: params,
        onResponse({ response }) {
            if (response.status === 401) {
                const authCookie = useCookie('auth')
                authCookie.value = null
                navigateTo('/login')
            }
        },
    })
}
