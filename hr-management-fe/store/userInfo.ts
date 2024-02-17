import { defineStore } from 'pinia'
import { USER_LEVEL } from '~/constants/userLevel'
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
        hasSuperAdminPermission(): boolean {
            return this.userInfo?.userLevel == USER_LEVEL.SUPER_ADMIN || this.userInfo?.userLevel == USER_LEVEL.ROOT
        },
    },
})
