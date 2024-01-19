import type { Employee } from './employee'

export type CreateScheduleParams = {
    employeeId: number
    scope: string
    dates: string
}

export type ScheduleFilterParams = {
    employeeId?: string
    scope: string
    departmentId: string
    date?: string
}

export type ScheduleInfo = {
    scope: string
    employees: Employee[]
}
