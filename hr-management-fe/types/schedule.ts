import type { Employee } from './employee'

export type CreateScheduleParams = {
    employeeId: number[]
    scope: string
    dates: string
    clockInTime: string
    clockOutTime: string
    courseId: string
    minuteBreakTime: number
}

export type ScheduleFilterParams = {
    employeeId?: string
    scope: string
    courseId: string
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
    minuteBreakPerDay: number
    minuteWorkPerDay: number
    createdAt: string
    updatedAt: string
}
