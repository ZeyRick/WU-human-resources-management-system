type ClockParams = {
    employeeId: number
    clockType: 'in' | 'out'
}

type ClockFilter = {
    employeeId: string | null
    date: string
}