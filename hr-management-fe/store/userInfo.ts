import { defineStore } from 'pinia'
import type { User } from '~/types/user'

export const useUserInfoStore = defineStore('userInfo', {
    state: (): { userInfo: User | null | undefined } => ({
        userInfo: null,
    }),
    actions: {
        storeUserInfo(value: User | null | undefined) {
            this.userInfo = value
        },
        clearUserInfo() {
            this.userInfo = null
        },
    },
})
