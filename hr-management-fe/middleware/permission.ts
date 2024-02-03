import { createDiscreteApi, darkTheme } from 'naive-ui'
import { USER_LEVEL } from '~/constants/userLevel'
import { useAuthStore } from '~/store/auth'
import { useUserInfoStore } from '~/store/userInfo'

export default defineNuxtRouteMiddleware((to) => {
    const { userInfo } = useUserInfoStore()

    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })
    console.log(userInfo)
    if (userInfo?.userLevel != USER_LEVEL.SUPER_ADMIN && userInfo?.userLevel != USER_LEVEL.ROOT) {
        message.error('No Permission')
        abortNavigation()
        return navigateTo('/admin/schedule')
    }
    // if token doesn't exist redirect to login
})
