import moment from 'moment'
import { DATE_TIME_FORMAT } from '~/constants/time'
import type { CreateClockSetting } from '~/types/clockSetting'
import { privateRequest } from '~/utils/request'

export const apiGetClockSetting = async () => {
    return privateRequest(
        '/admin/clock-setting',
        {
            method: 'get',
        },
        'apiGetClockSetting',
    )
}

export const apiSaveClockSetting = async (createData: CreateClockSetting) => {
    return privateRequest(
        '/admin/clock-setting',
        {
            method: 'post',
            body: createData,
        },
        'apiSaveClockSetting',
    )
}
