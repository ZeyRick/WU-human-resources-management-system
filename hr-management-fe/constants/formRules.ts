export const CommonFormRules = {
    userName: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    password: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    departmentId: {
        type: 'number',
        required: true,
        message: 'Department ID Must Input',
        trigger: 'blur',
    },
    employeeId: {
        type: 'number',
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    dates: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    scope: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
    clockOutTime: {
        required: true,
        message: 'Must Input',
        trigger: 'blur',
    },
}
