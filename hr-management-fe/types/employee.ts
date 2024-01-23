export type EmployeeParams = {
    id?: number
    employeeName?: string
    departmentId?: string | null
}

export type Employee = {
    id: string
    name: string
    employeeName: string
    clockInTime: string
    clockOutTime: string
}
