import { useAuthStore } from '~/store/auth.ts'

export default defineNuxtRouteMiddleware((to) => {
    const auth = useAuthStore()
    console.log(`2123x`)
    console.log(auth.authenticated)
    if (auth.authenticated && to?.name === 'admin-login') {
        return navigateTo('/admin/schedule')
    }

    if (!auth.authenticated && to?.name !== 'admin-login') {
        abortNavigation()
        return navigateTo('/admin/login')
    }
})
