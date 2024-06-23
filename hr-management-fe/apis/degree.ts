import type { CreateDegreeParams, DegreeFilterParams } from '~/types/degree'

export const apiAllDegree = async () => {
    const config = useRuntimeConfig()
    return privateRequest(
        '/admin/degree/all',
        {
            method: 'get',
        },
        'apiAllDegree',
    )
}

export const apiListDegree = async (pageOpt: Pagination, params?: DegreeFilterParams) => {
    return privateRequest(
        '/admin/degree',
        {
            method: 'get',
            query: { ...pageOpt, ...params },
        },
        'apiListDegree',
    )
}

export const apiCreateDeparment = async (params: CreateDegreeParams) => {
    return privateRequest(
        '/admin/degree',
        {
            method: 'post',
            body: params,
        },
        'apiCreateDeparment',
    )
}

export const apiEditDegree = async (degreeId: string, params: CreateDegreeParams) => {
    return privateRequest(
        `/admin/degree/${degreeId}`,
        {
            method: 'patch',
            body: params,
        },
        'apiEditDegree',
    )
}
