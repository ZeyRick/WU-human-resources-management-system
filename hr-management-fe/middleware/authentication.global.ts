import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
    const { token, storeToken } = useAuthStore()
    if (!token) {
        storeToken(useCookie('lin').value)
    }

    if (token && to?.name === 'admin-login') {
        return navigateTo('/admin/schedule')
    }

    // if token doesn't exist redirect to login
    if (!token && to?.name !== 'admin-login') {
        abortNavigation()
        return navigateTo('/admin/login')
    }
})
