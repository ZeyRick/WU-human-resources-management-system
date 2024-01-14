import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const aslocalTime = (utcTime: string, format = DATE_TIME_FORMAT) => {
    return moment.utc(utcTime).local().format(format)
}

export const getNowLocal = (format = DATE_TIME_FORMAT) => {
    return moment.utc().local().format(format)
}
