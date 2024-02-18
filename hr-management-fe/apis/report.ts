import type {report , reprtParamFilterByEmpolyee, reprtParamFilterByDepartment, reportParamFilterByDate} from "~/types/report";
import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const apiGetReportByEmployee = async (pageOpt: Pagination, filter: reprtParamFilterByEmpolyee) => {
    return privateRequest(
        '/admin/report',
        {
            method: 'get',
            query: { ...pageOpt, ...filter },
        },
        'apiGetReportByEmployee',
    )
}
export const apiGetReportByDepartment = async (pageOpt: Pagination, filter: reprtParamFilterByDepartment) => {
    return privateRequest(
        '/admin/report',
        {
            method: 'get',
            query: { ...pageOpt, ...filter },
        },
        'apiGetReportByDepartment',
    )
}
export const apiGetReportByDate = async (pageOpt: Pagination, filter: reportParamFilterByDate) => {
    return privateRequest(
        '/admin/report',
        {
            method: 'get',
            query: { ...pageOpt, ...filter, startDate: moment(filter.startDate).startOf('day').utc().format(DATE_TIME_FORMAT), 
            endDate: moment(filter.endDate).endOf('day').utc().format(DATE_TIME_FORMAT) },
        },
        'apiGetReportByDate',
    )
}
export const apiGetReport = async () => {
    return privateRequest(
        '/admin/report',
        {
            method: 'get',
            
        },
        'apiGetReport',
    )
}