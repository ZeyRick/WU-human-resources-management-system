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
    DocumentOutline,
    SchoolOutline,
    MenuSharp,
    BookOutline,
} from '@vicons/ionicons5'
import { useUserInfoStore } from '~/store/userInfo'

export const getMenuOptions = (): MenuOption[] => {
    const { hasSuperAdminPermission } = useUserInfoStore()

    return [
        {
            label: renderRoute('/admin/dashboard', i18n.global.t('dashboard')),
            icon: renderIcon(MenuSharp),
            key: 'home',
        },
        {
            label: renderRoute('/admin/clock', i18n.global.t('clock_management')),
            icon: renderIcon(TimerOutline),
            key: 'clock',
        },

        {
            label: renderRoute('/admin/schedule', i18n.global.t('schedules')),
            icon: renderIcon(CalendarNumberOutline),
            key: 'schedule',
        },
        {
            label: renderRoute('/admin/course', i18n.global.t('courses')),
            icon: renderIcon(GridOutline),
            key: 'courses',
        },
        {
            label: renderRoute('/admin/degree', i18n.global.t('degree')),
            icon: renderIcon(SchoolOutline),
            key: 'degree',
        },
        {
            label: i18n.global.t('statistic'),
            icon: renderIcon(DocumentOutline),
            key: 'report-big',
            children: [
                {
                    label: renderRoute('/admin/report', i18n.global.t('employee_report')),
                    icon: renderIcon(DocumentOutline),
                    key: 'report',
                },
                {
                    label: renderRoute('/admin/attendence', i18n.global.t('attendence')),
                    icon: renderIcon(TodayOutline),
                    key: 'attendence',
                },
            ],
        },

        {
            label: 'Employee',
            icon: renderIcon(PeopleOutline),
            key: 'employee',
            children: [
                {
                    label: renderRoute('/admin/employee-staff', i18n.global.t('staff')),
                    icon: renderIcon(PeopleOutline),
                    key: 'employee-staff',
                },
                {
                    label: renderRoute('/admin/employee-teaching', i18n.global.t('teaching_staff')),
                    icon: renderIcon(SchoolOutline),
                    key: 'employee-teaching',
                },
                {
                    label: renderRoute('/admin/employee-lecture', i18n.global.t('lecture')),
                    icon: renderIcon(BookOutline),
                    key: 'employee-lecture',
                },
                {
                    label: renderRoute('/admin/employee-request', i18n.global.t('telegram_requests')),
                    icon: renderIcon(PersonAddOutline),
                    key: 'telegram_request',
                },
            ],
        },
        {
            label: 'Management',
            icon: renderIcon(SettingsOutline),
            key: 'management',
            children: [
                {
                    label: renderRoute('/admin/clock-setting', i18n.global.t('clock_setting')),
                    icon: renderIcon(SettingsOutline),
                    key: 'clock_setting',
                    show: hasSuperAdminPermission(),
                },
                {
                    label: renderRoute('/admin/user', i18n.global.t('users')),
                    icon: renderIcon(ManOutline),
                    key: 'user',
                    show: hasSuperAdminPermission(),
                },
            ],
        },
    ]
}
