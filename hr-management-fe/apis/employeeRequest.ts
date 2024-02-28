import type { CreateEmployeeType, EmployeeParams } from '~/types/employee'

export const apiListEmployeeRequest = async (pageOpt: Pagination, params?: EmployeeParams) => {
    return privateRequest(
        '/admin/employee_request',
        {
            method: 'get',
            query: { ...pageOpt, ...params },
        },
        'apiListEmployeeRequest',
    )
}

export const apiDenyEmployeeRequest = async (id: string) => {
    return privateRequest(
        '/admin/employee_request/confirmation',
        {
            method: 'post',
            body: { requestId: id, confirmation: 'rejected' },
        },
        'apiDenyEmployeeRequest',
    )
}

export const apiApproveEmployeeRequest = async (id: string) => {
    return privateRequest(
        '/admin/employee_request/confirmation',
        {
            method: 'post',
            body: { requestId: id, confirmation: 'confirm' },
        },
        'apiApproveEmployeeRequest',
    )
}
