import { useAuthStore } from '~/store/auth'
import type { CreateUserType } from '~/types/user'

export const apiGetUser = async () => {
    return privateRequest(
        '/admin/user',
        {
            method: 'get',
        },
        'getUser',
    )
}

export const apiDelUser = async (userId: string) => {
    return privateRequest(
        `/admin/user/${userId}`,
        {
            method: 'delete',
        },
        'delUser',
    )
}

export const apiUserResetPW = async (userId: string, pw: string) => {
    return privateRequest(
        `/admin/user/${userId}`,
        {
            method: 'patch',
            body: { password: pw },
        },
        'apiUserResetPW',
    )
}

export const apiCreateUser = async (params: CreateUserType) => {
    return privateRequest(
        '/admin/user',
        {
            method: 'post',
            body: params,
        },
        'createUser',
    )
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
            const body = JSON.parse(response._data)
            if (response.status === 200 && body?.code === 0 && body?.res) {
                const { storeToken } = useAuthStore()
                storeToken(body.res)
                if (params.rememberMe) {
                    const cookie = useCookie('lin')
                    cookie.value = body.res
                }
                navigateTo('/admin/schedule')
                return
            }
            throw new Error(body?.msg || 'Somthing went wrong')
        },
    })
}

export const apiLogout = async () => {
    const cookie = useCookie('lin')
    const { storeToken } = useAuthStore()
    cookie.value = null
    storeToken(null)
    navigateTo('/admin/login')
    return
}
