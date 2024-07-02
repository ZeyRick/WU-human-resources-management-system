import type { FormItemRule, FormRules } from 'naive-ui'

export const CommonFormRules: FormRules = {
    userName: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    name: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    alias: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    password: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    courseId: {
        type: 'number',
        required: true,
        message: 'Course ID Must Input',
        trigger: 'blur',
    },
    degreeId: {
        type: 'number',
        required: true,
        message: 'Degree ID Must Input',
        trigger: 'blur',
    },
    employeeType: {
        type: 'string',
        required: true,
        message: 'Employee Type Must Input',
        trigger: 'blur',
    },
    salary: {
        type: 'number',
        required: true,
        message: 'Salary Must Input',
        trigger: 'blur',
    },
    employeeId: {
        type: 'number',
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    minuteBreakTime: {
        type: 'number',
        required: true,
        message: 'Break Time Must Input',
        trigger: 'blur',
    },
    dates: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    clockDate: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    clockInTime: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    clockOutTime: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    totalMinute: {
        type: 'number',
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    scope: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    idNumber: {
        required: true,
        trigger: ['blur', 'input'],
        validator(rule: FormItemRule, value: string) {
            if (!value) {
                return new Error('ID number is required')
            } else if (!/^\d*$/.test(value)) {
                return new Error('ID number should be number')
            } else if (value.length != 10) {
                return new Error('ID number should be 10 digit')
            }
            return true
        },
    },
}
