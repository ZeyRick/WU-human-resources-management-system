import type { CreateCourseParams, CourseFilterParams } from '~/types/course'

export const apiAllCourse = async () => {
    const config = useRuntimeConfig()
    return privateRequest(
        '/admin/course/all',
        {
            method: 'get',
        },
        'apiAllCourse',
    )
}
export const apiCourseEmployee = async (employeeId: number) => {
    const config = useRuntimeConfig()
    return privateRequest(
        `/admin/course/employee/${employeeId}`,
        {
            method: 'get',
        },
        'apiCourseEmployee',
    )
}

export const apiListCourse = async (pageOpt: Pagination, params?: CourseFilterParams) => {
    return privateRequest(
        '/admin/course',
        {
            method: 'get',
            query: { ...pageOpt, ...params },
        },
        'apiListCourse',
    )
}

export const apiCreateDeparment = async (params: CreateCourseParams) => {
    return privateRequest(
        '/admin/course',
        {
            method: 'post',
            body: params,
        },
        'apiCreateDeparment',
    )
}

export const apiEditCourse = async (courseId: string, params: CreateCourseParams) => {
    return privateRequest(
        `/admin/course/${courseId}`,
        {
            method: 'patch',
            body: params,
        },
        'apiEditCourse',
    )
}
