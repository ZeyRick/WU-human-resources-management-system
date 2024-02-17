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

const { token } = useAuthStore()

if (token) {
    const config = useRuntimeConfig()
    const { storeUserInfo } = useUserInfoStore()
    try {
        const { data } = await useAsyncData('userinfo', async () => {
            return await $fetch('/admin/user/userInfo', {
                headers: {
                    Accept: '*/*',
                    Authorization: token ? `Bearer ${token}` : '',
                },
                baseURL: String(config.public.apiURL),
            })
        })
        const userInfo = JSON.parse(data.value as string)?.res
        storeUserInfo(userInfo as User)
    } catch (error) {}
}
</script>
