import { createDiscreteApi, darkTheme } from 'naive-ui'
import { useUserInfoStore } from '~/store/userInfo'

export default defineNuxtRouteMiddleware((to) => {
    const { hasSuperAdminPermission } = useUserInfoStore()

    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })
    if (!hasSuperAdminPermission()) {
        message.error('No Permission')
        abortNavigation()
        return navigateTo('/admin/schedule')
    }
    // if token doesn't exist redirect to login
})
