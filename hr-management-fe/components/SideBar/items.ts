import type { MenuOption } from 'naive-ui'
import { RouterLink } from 'vue-router'
import { Timer, Man, Calendar } from '@vicons/ionicons5'

export const menuOptions: MenuOption[] = [
    {
        label: renderRoute('/admin/clock', i18n.global.t('clock_management')),
        icon: renderIcon(Timer),
        key: 'clock',
    },
    {
        label: renderRoute('/admin/user', i18n.global.t('user_management')),
        icon: renderIcon(Man),
        key: 'user',
    },
    {
        label: renderRoute('/admin/schedule', i18n.global.t('schedule_management')),
        icon: renderIcon(Calendar),
        key: 'schedule',
    },
    {
        label: renderRoute('/admin/employee', i18n.global.t('employee_list')),
        icon: renderIcon(Calendar),
        key: 'employee',
    },
    {
        label: renderRoute('/admin/employee-request', i18n.global.t('employee_request')),
        icon: renderIcon(Calendar),
        key: 'employee_request',
    },
    {
        label: renderRoute('/admin/report', i18n.global.t('employee_report')),
        key: 'hello',
        href: 'https://en.wikipedia.org/wiki/Hear_the_Wind_Sing',
    },
]
