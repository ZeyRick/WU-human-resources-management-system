import type { Employee } from './employee'

export type CreateScheduleParams = {
    employeeId: number[]
    scope: string
    dates: string
    clockInTime: string
    clockOutTime: string
    departmentId: string
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

export type Schedule = {
    id: number
    employeeId: number
    scope: string
    dates: string
    clockInTime: string
    clockOutTime: string
    employee: Employee
    createdAt: string
    updatedAt: string
}