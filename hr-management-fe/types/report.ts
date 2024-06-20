export type Report = {
    employeeId: number
    employeeName: string
    courseAlias: string
    totalWorkMinute: number
    attandance: number
}

export type ReportFilter = {
    courseId: string
    employeeName: string
    startDate: string
    endDate: string
}
