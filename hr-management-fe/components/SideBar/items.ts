import type { MenuOption } from 'naive-ui'
import { RouterLink } from 'vue-router'
import { Timer } from '@vicons/ionicons5'

export const menuOptions: MenuOption[] = [
    {
        label: renderRoute('/admin/clock', i18n.global.t('clock_management')),
        icon: renderIcon(Timer),
        key: 'clock',
    },
    {
        label: 'Hear the Wind Sing',
        key: 'hear-the-wind-sing1',
        href: 'https://en.wikipedia.org/wiki/Hear_the_Wind_Sing',
    },
]
