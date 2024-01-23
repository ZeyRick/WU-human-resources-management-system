import { defineStore } from 'pinia'
interface RequestStoreState {
    requests: Record<string, string>
}

export const useRequestId = defineStore({
    id: 'requestId',
    state: (): RequestStoreState => ({
        requests: {},
    }),
    actions: {
        addRequest(key: string, value: string) {
            this.requests[key] = value
        },
        removeRequest(key: string) {
            delete this.requests[key]
        },
    },
})
