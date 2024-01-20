import type { Employee } from './employee'

export type CreateScheduleParams = {
    employeeId?: number | undefined
    scope: string
    dates: string
    clockInTime: string
    clockOutTime: string
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
