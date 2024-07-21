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
    DocumentAttachOutline,
} from '@vicons/ionicons5'
import { useUserInfoStore } from '~/store/userInfo'

export const getMenuOptions = (): MenuOption[] => {
    const { hasSuperAdminPermission } = useUserInfoStore()

    return [
        {
            ...renderRoute('/admin/dashboard', i18n.global.t('dashboard')),
            icon: renderIcon(MenuSharp),
        },
        {
            ...renderRoute('/admin/clock', i18n.global.t('staff_clock')),
            icon: renderIcon(TimerOutline),
        },
        {
            ...renderRoute('/admin/schedule', i18n.global.t('schedules')),
            icon: renderIcon(CalendarNumberOutline),
        },
        {
            ...renderRoute('/admin/course', i18n.global.t('courses')),
            icon: renderIcon(GridOutline),
        },
        {
            ...renderRoute('/admin/degree', i18n.global.t('degree')),
            icon: renderIcon(SchoolOutline),
        },
        {
            label: i18n.global.t('statistic'),
            icon: renderIcon(DocumentOutline),
            key: 'report-big',
            children: [
                {
                    ...renderRoute('/admin/report', i18n.global.t('employee_report')),
                    icon: renderIcon(DocumentOutline),
                },
                {
                    ...renderRoute('/admin/teaching-report', i18n.global.t('teaching_report')),
                    icon: renderIcon(DocumentAttachOutline),
                },
                {
                    ...renderRoute('/admin/attendance', i18n.global.t('staff_attendance')),
                    icon: renderIcon(TodayOutline),
                },
                {
                    ...renderRoute('/admin/teaching-attendance', i18n.global.t('teaching_attendance')),
                    icon: renderIcon(TimerOutline),
                },
            ],
        },

        {
            label: 'Employee',
            icon: renderIcon(PeopleOutline),
            key: 'employee',
            children: [
                {
                    ...renderRoute('/admin/employee-staff', i18n.global.t('staff')),
                    icon: renderIcon(PeopleOutline),
                },
                {
                    ...renderRoute('/admin/employee-teaching', i18n.global.t('teaching_staff')),
                    icon: renderIcon(SchoolOutline),
                },
                {
                    ...renderRoute('/admin/employee-lecture', i18n.global.t('lecture')),
                    icon: renderIcon(BookOutline),
                },
                {
                    ...renderRoute('/admin/employee-request', i18n.global.t('telegram_requests')),
                    icon: renderIcon(PersonAddOutline),
                },
            ],
        },
        {
            label: 'Management',
            icon: renderIcon(SettingsOutline),
            key: 'management',
            children: [
                {
                    ...renderRoute('/admin/clock-setting', i18n.global.t('clock_setting')),
                    icon: renderIcon(SettingsOutline),
                    show: hasSuperAdminPermission(),
                },
                {
                    ...renderRoute('/admin/user', i18n.global.t('users')),
                    icon: renderIcon(ManOutline),
                    show: hasSuperAdminPermission(),
                },
            ],
        },
    ]
}
