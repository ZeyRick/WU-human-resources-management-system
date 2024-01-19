import type { EmployeeParams } from '~/types/employee'

export const apiAllEmployee = async (params?: EmployeeParams) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/employee/all', {
        method: 'get',
        baseURL: String(config.public.apiURL),
        query: { ...params },
    })
}
