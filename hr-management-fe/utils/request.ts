export const privateRequest = async (path: string, object: Object) => {
    const config = useRuntimeConfig()
    return $fetch(path, {
        headers: {
            Accept: '*/*',
        },
        credentials: 'include',
        baseURL: String(config.public.apiURL),
        onResponse({ response }) {
            if (response.status === 401) {
                const authCookie = useCookie('auth')
                authCookie.value = null
                navigateTo('/admin/login')
            }
            response._data = JSON.parse(response._data)
            if (response._data?.code === -1) {
                throw new Error(response._data?.msg || 'Someting went wrong')
            }
        },
        ...object,
    })
}
