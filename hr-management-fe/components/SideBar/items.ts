import type { MenuOption } from 'naive-ui'
import {
    CalendarNumberOutline,
    ManOutline,
    PeopleOutline,
    PersonAddOutline,
    TimerOutline,
    GridOutline,
    SettingsOutline,
    TodayOutline,
} from '@vicons/ionicons5'
import { useUserInfoStore } from '~/store/userInfo'

export const getMenuOptions = (): MenuOption[] => {
    const { hasSuperAdminPermission } = useUserInfoStore()

    return [
        {
            label: renderRoute('/admin/clock', i18n.global.t('clock_management')),
            icon: renderIcon(TimerOutline),
            key: 'clock',
        },
        {
            label: renderRoute('/admin/user', i18n.global.t('users')),
            icon: renderIcon(ManOutline),
            key: 'user',
            show: hasSuperAdminPermission(),
        },
        {
            label: renderRoute('/admin/schedule', i18n.global.t('schedules')),
            icon: renderIcon(CalendarNumberOutline),
            key: 'schedule',
        },
        {
            label: renderRoute('/admin/department', i18n.global.t('departments')),
            icon: renderIcon(GridOutline),
            key: 'deparments',
        },
        {
            label: renderRoute('/admin/employee', i18n.global.t('employee_list')),
            icon: renderIcon(PeopleOutline),
            key: 'employee',
        },
        {
            label: renderRoute('/admin/employee-request', i18n.global.t('telegram_requests')),
            icon: renderIcon(PersonAddOutline),
            key: 'telegram_request',
        },
        {
            label: renderRoute('/admin/clock-setting', i18n.global.t('clock_setting')),
            icon: renderIcon(SettingsOutline),
            key: 'clock_setting',
            show: hasSuperAdminPermission(),
        },
        {
            label: renderRoute('/admin/attendence', i18n.global.t('attendence')),
            icon: renderIcon(TodayOutline),
            key: 'attendence',
        },

        {
            label: renderRoute('/admin/report', i18n.global.t('employee_report')),
            key: 'hello',
            href: 'https://en.wikipedia.org/wiki/Hear_the_Wind_Sing',
        },
    ]
}
