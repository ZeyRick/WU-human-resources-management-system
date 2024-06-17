<template>
    <n-loading-bar-provider>
        <n-message-provider>
            <NuxtLayout>
                <NuxtPage />
            </NuxtLayout>
        </n-message-provider>
    </n-loading-bar-provider>
</template>

<script setup lang="ts">
import type { User } from './types/user'
import { useAuthStore } from './store/auth'
import { useUserInfoStore } from './store/userInfo'
import './assets/vfonts-main/FiraCode.css'

const { token } = useAuthStore()

if (token) {
    const config = useRuntimeConfig()
    const { storeUserInfo } = useUserInfoStore()
    try {
        const userData: any = JSON.parse(
            await $fetch('/admin/user/userInfo', {
                headers: {
                    Accept: '*/*',
                    Authorization: `Bearer ${token}`,
                },
                baseURL: String(config.public.apiURL),
            }),
        )
        const userInfo = JSON.parse(decrypteData(userData.res as string))
        storeUserInfo(userInfo as User)
    } catch (error) {}
}
</script>
