import moment from 'moment'
import { DATE_TIME_FORMAT, TIME_FORMAT } from '~/constants/time'

export const aslocalTime = (utcTime: string, format = DATE_TIME_FORMAT) => {
    return utcTime ? moment.utc(utcTime).local().format(format) : '-'
}

export const getNowLocal = (format = DATE_TIME_FORMAT) => {
    return moment.utc().local().format(format)
}

export const isTimeBefore = (referenceTimeStr: string, timeStr: string) => {
    console.log(213, 'Before', 'Current', timeStr, 'Ref', referenceTimeStr)
    const referenceTime = moment.utc(referenceTimeStr).format(TIME_FORMAT)
    const currentTime = moment.utc(timeStr).format(TIME_FORMAT)
    const referenceMoment = moment.utc(referenceTime, TIME_FORMAT)
    const currentMoment = moment.utc(currentTime, TIME_FORMAT)
    const timeOnlyComparison = moment.duration(currentMoment.diff(referenceMoment)).asMinutes()
    // Determine if the current time is before the reference time (ignoring date)
    console.log(
        'Before',
        'Current',
        currentMoment.format(DATE_TIME_FORMAT),
        'Ref',
        referenceMoment.format(DATE_TIME_FORMAT),
        timeOnlyComparison,
        Number(timeOnlyComparison) < 0,
    )
    return Number(timeOnlyComparison) < 0
}

export const isTimeAfter = (referenceTimeStr: string, timeStr: string) => {
    const referenceTime = moment.utc(referenceTimeStr).format(TIME_FORMAT)
    const currentTime = moment.utc(timeStr).format(TIME_FORMAT)
    const referenceMoment = moment.utc(referenceTime, TIME_FORMAT)
    const currentMoment = moment.utc(currentTime, TIME_FORMAT)
    const timeOnlyComparison = moment.duration(currentMoment.diff(referenceMoment)).asMinutes()
    // Determine if the current time is before the reference time (ignoring date)
    console.log(
        'After',
        'Current',
        currentMoment.format(DATE_TIME_FORMAT),
        'ref',
        referenceMoment.format(DATE_TIME_FORMAT),
        timeOnlyComparison,
        Number(timeOnlyComparison) > 0,
    )
    return Number(timeOnlyComparison) > 0
}
