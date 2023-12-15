export const apiClock = async (params: clock) => {
    const config = useRuntimeConfig()
    return $fetch('/clock', {
        mode: 'no-cors',
        method: 'post',
        baseURL: config.public.baseURL,
        body: params,
    })
}
