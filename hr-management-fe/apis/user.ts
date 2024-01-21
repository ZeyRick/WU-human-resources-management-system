import { useMessage } from 'naive-ui'

export const apiCreateUser = () => {
    $fetch('/ggwellplay')
}

export const apiLogin = async (params: LoginParams) => {
    const config = useRuntimeConfig()
    await $fetch('/admin/user/login', {
        method: 'post',
        headers: {
            Accept: '*/*',
        },
        credentials: 'include',
        baseURL: String(config.public.apiURL),
        body: params,
        onResponse({ response }) {
            if (response.status === 200) {
                navigateTo('/admin/schedule')
                return
            }
            const body = JSON.parse(response._data)
            throw new Error(body?.msg || 'Somthing went wrong')
        },
    })
}

export const apiLogout = async () => {
    const config = useRuntimeConfig()
    const res: any = await $fetch('/admin/user/logout', {
        method: 'get',
        headers: {
            Accept: '*/*',
        },
        credentials: 'include',
        baseURL: String(config.public.apiURL),
        onResponse({ response }) {
            if (response.status === 200) {
                navigateTo('/admin/login')
                return res
            }
        },
    })
    return res
}
