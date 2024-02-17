import type { USER_LEVEL } from '~/constants/userLevel'

export type User = {
    id: number
    createdAt: string
    updatedAt: string
    deletedAt: string | null
    username: string
    name: string
    userLevel: string
    profilePic: string
}

export type CreateUserType = {
    userName: string
    password: string
    name: string
    userLevel: USER_LEVEL
}
