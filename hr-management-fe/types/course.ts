export type Course = {
    id: string
    alias: string
}

export type CourseFilterParams = {
    alias: string
    employeeId: number
}

export type CreateCourseParams = {
    alias: string
}
