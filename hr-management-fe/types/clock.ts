import type { Employee } from './employee'
import type { Schedule } from './schedule'

export type ClockParams = {
    employeeId: number
    clockType: 'in' | 'out'
}

export type ClockFilter = {
    employeeId: string | null
    date: string
}

export type Clock = {
    id: number
    createdAt: string // Adjust type based on your use case (Date or string with specific format)
    updatedAt: string
    deletedAt: null | Date // Allow null or Date
    employeeId: number
    clockType: string // "in" or "out"
    clockInId: number | null // Allow null for clockOut entries
    ClockIn: Clock | null // Recursive type reference for nested Clock
    clockOutMinute: number
    employee: Employee
    schedule: Schedule
}

export type AttendenceFilter = {
    startDate: string
    endDate: string
    employeeName: string
    employeeId: string
    courseId: string
}

export type EditClock = {
    clockTime: string | null
}
