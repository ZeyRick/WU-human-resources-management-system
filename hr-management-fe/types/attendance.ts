export type CreateAttendance = {
    employeeId: number | null
    courseId: number | null
    degreeId: number | null
    clockDate: string
    clockInTime: string
    clockOutTime: string
}
