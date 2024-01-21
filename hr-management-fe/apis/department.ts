export const apiAllDepartment = async () => {
    const config = useRuntimeConfig()
    return privateRequest('/admin/department/all', {
        method: 'get',
    })
}
