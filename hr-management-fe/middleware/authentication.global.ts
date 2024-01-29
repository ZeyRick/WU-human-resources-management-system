import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
    const { token, storeToken } = useAuthStore()
    let curToken = token
    if (!token) {
        curToken = useCookie('lin').value
        storeToken(useCookie('lin').value)
    }

    if (curToken && to?.name === 'admin-login') {
        return navigateTo('/admin/schedule')
    }

    // if token doesn't exist redirect to login
    // if (!curToken && to?.name !== 'admin-login') {
    //     abortNavigation()
    //     return navigateTo('/admin/login')
    // }
})
