import { createDiscreteApi, darkTheme, useLoadingBar } from 'naive-ui'
import { useAuthStore } from '~/store/auth'
import { useRequestId } from '~/store/requestId'

export const privateRequest = async (path: string, object: Object, key: string) => {
    const requestId = useRequestId()
    const { token } = useAuthStore()
    if (requestId.requests[key]) {
        return
    }
    requestId.addRequest(key, '1')
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
            Authorization: token ? `Bearer ${token}` : '',
        },
        baseURL: String(config.public.apiURL),
        onResponse({ response }) {
            requestId.removeRequest(key)
            if (response.status === 401) {
                const { storeToken } = useAuthStore()
                const authCookie = useCookie('auth')
                authCookie.value = null
                storeToken(null)
                message.error('Please Login First')
                navigateTo('/admin/login')
                if (loadingBar) {
                    loadingBar.error()
                }
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
            if (response._data?.msg) {
                message.success(response._data?.msg)
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
