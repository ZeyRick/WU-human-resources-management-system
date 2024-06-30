import type { Course } from './course'
import type { Degree } from './degree'
import type { Schedule } from './schedule'

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
    employeeType: EMPLOYEE_TYPE
    clockInTime: string
    clockOutTime: string
    courses?: Course[]
    courseIds?: string[]
    degrees?: Degree[]
    degreeIds?: string[]
    schedules: Schedule[]
    idNumber?: string
    idFileName?: string
    photoFileName?: string
    telegramId?: string
    telegramUsername?: string
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
    LECTURE = 'lecture',
}

export const BIND_STATUS_ENUM = {
    APPROVE: 'approved',
    PENDING: 'pending',
}
