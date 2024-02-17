import type { CreateEmployeeType, EmployeeParams } from '~/types/employee'

export const apiAllEmployee = async (params?: EmployeeParams) => {
    return privateRequest(
        '/admin/employee/all',
        {
            method: 'get',
            query: { ...params },
        },
        'allEmployee',
    )
}

export const apiCreateEmployee = async (params: CreateEmployeeType) => {
    return privateRequest(
        '/admin/employee',
        {
            method: 'post',
            body: params,
        },
        'apiCreateEmployee',
    )
}

export const apiEditEmployee = async (employeeId: string, params: CreateEmployeeType) => {
    return privateRequest(
        `/admin/employee/${employeeId}`,
        {
            method: 'patch',
            body: params,
        },
        'apiEditEmployee',
    )
}

export const apiListEmployee = async (pageOpt: Pagination, params?: EmployeeParams) => {
    return privateRequest(
        '/admin/employee',
        {
            method: 'get',
            query: { ...pageOpt, ...params },
        },
        'listEmployee',
    )
}

export const apiDeleteEmployee = async (employeeId: string) => {
    return privateRequest(
        `/admin/employee/${employeeId}`,
        {
            method: 'delete',
        },
        'delEmployee',
    )
}
