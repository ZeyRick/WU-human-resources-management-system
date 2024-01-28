import type { CreateDepartmentParams, DepartmentFilterParams } from '~/types/department'

export const apiAllDepartment = async () => {
    const config = useRuntimeConfig()
    return privateRequest(
        '/admin/department/all',
        {
            method: 'get',
        },
        'apiAllDepartment',
    )
}

export const apiListDepartment = async (pageOpt: Pagination, params?: DepartmentFilterParams) => {
    return privateRequest(
        '/admin/department',
        {
            method: 'get',
            query: { ...pageOpt, ...params },
        },
        'apiListDepartment',
    )
}

export const apiCreateDeparment = async (params: CreateDepartmentParams) => {
    return privateRequest(
        '/admin/department',
        {
            method: 'post',
            body: params,
        },
        'apiCreateDeparment',
    )
}

export const apiEditDepartment = async (departmentId: string, params: CreateDepartmentParams) => {
    return privateRequest(
        `/admin/department/${departmentId}`,
        {
            method: 'patch',
            body: params,
        },
        'apiEditDepartment',
    )
}
