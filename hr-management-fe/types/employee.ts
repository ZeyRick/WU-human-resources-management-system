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
    employeeName: string
    clockInTime: string
    clockOutTime: string
}

export type CreateEmployeeType = {
    name: string
    courseId: string
    salary: number
    employeeType: EMPLOYEE_TYPE
    idNumber?: string
    idFileName?: string
    photoFileName?: string
}

export type UploadEmployeeFile = {
    idFile?: File
    photoFile?: File
}

export enum EMPLOYEE_TYPE {
    FULL_TIME = 'Fulltime',
    PART_TIME = 'Parttime',
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
