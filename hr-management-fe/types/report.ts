export type Report = {
    employeeId: number
    employeeName: string
    departmentAlias: string
    totalWorkMinute: number
    attandance: number
}

export type ReportFilter = {
    departmentId: string
    employeeName: string
    startDate: string
    endDate: string
}
