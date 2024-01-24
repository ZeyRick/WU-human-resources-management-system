import type { EmployeeParams } from '~/types/employee'

export const apiAllEmployee = async (params?: EmployeeParams) => {
    return privateRequest('/admin/employee/all', {
        method: 'get',
        query: { ...params },
    })
}


export const apiListEmployee = async (pageOpt: Pagination, params?: EmployeeParams) => {
    return privateRequest('/admin/employee', {
        method: 'get',
        query: { ...pageOpt, ...params },
    })
}

export const apiDeleteEmployee = async (employeeId: string) => {
    return privateRequest(`/admin/employee/${employeeId}`, {
        method: 'delete',
    })
}
