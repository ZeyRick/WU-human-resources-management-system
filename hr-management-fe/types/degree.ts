export type Degree = {
    id: string
    alias: string
    rate: number
}

export type DegreeFilterParams = {
    alias: string
}

export type CreateDegreeParams = {
    alias: string
    rate: number
}
