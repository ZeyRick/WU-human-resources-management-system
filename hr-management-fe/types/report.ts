import type { EMPLOYEE_TYPE } from './employee'

export type Report = {
    employeeId: number
    employeeName: string
    courseAlias: string
    totalWorkMinute: number
    attandance: number
}

export type ReportFilter = {
    startDate: string
    endDate: string
    isTeaching: boolean
    employeeId: string
}

export type EmployeeCount = {
    totalCount: number
    employeeType: EMPLOYEE_TYPE
}

export type DashboardSummary = {
    employeeCounts: EmployeeCount[]
    degreeCount: number
    courseCount: number
    userCount: number
}
