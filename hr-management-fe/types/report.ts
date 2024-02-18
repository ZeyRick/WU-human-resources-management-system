export type report = {
    employeeId: number;
    employeeName: string;
    departmentAlias: string;
    totalWorkMinute: number;
    attandance: number;
}
export type reprtParamFilterByEmpolyee = {
    name: string;

}
export type reprtParamFilterByDepartment = {
    department: string;
}
export type reportParamFilterByDate = {
    startDate: string;
    endDate: string;
}