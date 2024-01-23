import { createDiscreteApi, darkTheme, useLoadingBar } from 'naive-ui'
import { useRequestId } from '~/store/requestId'

export const privateRequest = async (path: string, object: Object) => {
    const requestId = useRequestId()
    if (requestId.requests[path]) {
        return
    }
    requestId.addRequest(path, '1')
    const loadingBar = useLoadingBar()
    if (loadingBar) {
        loadingBar.start()
    }
    const config = useRuntimeConfig()
    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })

    return $fetch(path, {
        headers: {
            Accept: '*/*',
        },
        credentials: 'include',
        baseURL: String(config.public.apiURL),
        onResponse({ response }) {
            requestId.removeRequest(path)
            if (response.status === 401) {
                const authCookie = useCookie('auth')
                authCookie.value = null
                message.error('Please Login First')
                navigateTo('/admin/login')
                return
            }
            response._data = JSON.parse(response._data)
            if (response._data?.code === -1) {
                if (loadingBar) {
                    loadingBar.error()
                }
                message.error(response._data?.msg || 'Someting Went Wrong')
                throw new Error(response._data?.msg || 'Someting Went Wrong')
            }
            if (response._data?.res) {
                response._data = response._data.res
            }
            if (loadingBar) {
                loadingBar.finish()
            }
        },
        ...object,
    })
}
