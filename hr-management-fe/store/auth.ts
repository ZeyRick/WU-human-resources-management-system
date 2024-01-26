import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: (): { token: string | null | undefined } => ({
        token: null,
    }),
    actions: {
        storeToken(value: string | null | undefined) {
            this.token = value
        },
        removeToken() {
            this.token = ''
        },
    },
})
