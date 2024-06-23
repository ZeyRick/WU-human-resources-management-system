import type { Course } from './course'

export type EmployeeParams = {
    id?: number
    employeeName?: string
    courseId?: string | null
    scope?: string
    employeeType: EMPLOYEE_TYPE | ''
    startSalary: number | null
    endSalary: number | null
}

export type Employee = {
    id: string
    name: string
    salary: number
    employeeName: string
    type: EMPLOYEE_TYPE
    clockInTime: string
    clockOutTime: string
    courses?: Course[]
    courseIds?: string[]
    degrees?: Course[]
    degreeIds?: string[]
}

export type EmployeeWithFile = Employee & {
    idFileName?: string
    photoFileName?: string
    idNumber?: string
}

export type CreateEmployeeType = {
    name: string
    courseIds: string[]
    degreeIds: string[]
    salary: number
    employeeType?: EMPLOYEE_TYPE
    idNumber?: string
    idFileName?: string
    photoFileName?: string
}

export type UploadEmployeeFile = {
    idFile?: File
    photoFile?: File
}

export enum EMPLOYEE_TYPE {
    STAFF = 'staff',
    TEACHING_STAFF = 'teaching_staff',
}

export const BIND_STATUS_ENUM = {
    APPROVE: 'approved',
    PENDING: 'pending',
}

export type EmployeeWithSchedule = {
    id: string
    name: string
    courseId: number | null
    course: {
        courseId: number
        alias: string
        createdAt: string // You may need to adjust the date format based on your requirements
        updatedAt: string // You may need to adjust the date format based on your requirements
    }
    profilePic: string
    schedule: {
        scheduleId: number
        employeeId: number
        scope: string
        dates: string | null
        clockInTime: string // You may need to adjust the date format based on your requirements
        clockOutTime: string // You may need to adjust the date format based on your requirements
        createdAt: string // You may need to adjust the date format based on your requirements
        updatedAt: string // You may need to adjust the date format based on your requirements
    }
}
