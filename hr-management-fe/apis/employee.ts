import type { EmployeeParams } from '~/types/employee'

export const apiAllEmployee = async (params?: EmployeeParams) => {
    return privateRequest('/admin/employee/all', {
        method: 'get',
        query: { ...params },
    })
}
