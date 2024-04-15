import { createDiscreteApi, darkTheme, useLoadingBar } from 'naive-ui'
import { useAuthStore } from '~/store/auth'
import { useRequestId } from '~/store/requestId'

export const privateRequest = async (path: string, object: any, key: string) => {
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
    if (object?.method && object.method.toLowerCase() === 'post' && object.body) {
        object.body = encrypteData(JSON.stringify(object.body))
    }
    return $fetch(path, {
        headers: {
            Accept: '*/*',
            Authorization: token ? `Bearer ${token}` : '',
        },
        baseURL: String(config.public.apiURL),
        onResponse({ response }: any) {
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
            try {
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
                const decypt = decrypteData(response._data.res)

                if (decypt) {
                    response._data = JSON.parse(decypt)
                }

                if (loadingBar) {
                    loadingBar.finish()
                }
            } catch (error) {}
        },
        ...object,
    })
}
