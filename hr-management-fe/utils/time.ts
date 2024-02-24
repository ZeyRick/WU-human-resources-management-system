import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'

export const aslocalTime = (utcTime: string, format = DATE_TIME_FORMAT) => {
    return utcTime ? moment.utc(utcTime).local().format(format) : '-'
}

export const getNowLocal = (format = DATE_TIME_FORMAT) => {
    return moment.utc().local().format(format)
}

export const isTimeBefore = (referenceTimeStr: string, timeStr: string) => {
    const referenceTime = moment.utc(referenceTimeStr)
    const currentMoment = moment.utc(timeStr)
    const timeOnlyComparison = moment.duration(currentMoment.diff(referenceTime)).minutes()
    // Determine if the current time is before the reference time (ignoring date)
    return timeOnlyComparison < 0
}

export const isTimeAfter = (referenceTimeStr: string, timeStr: string) => {
    const referenceTime = moment.utc(referenceTimeStr)
    const currentMoment = moment.utc(timeStr)
    const timeOnlyComparison = moment.duration(currentMoment.diff(referenceTime)).minutes()
    // Determine if the current time is before the reference time (ignoring date)
    return timeOnlyComparison > 0
}

