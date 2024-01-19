export const apiAllDepartment = async () => {
    const config = useRuntimeConfig()
    return $fetch('/admin/department/all', {
        method: 'get',
        baseURL: String(config.public.apiURL),
    })
}
