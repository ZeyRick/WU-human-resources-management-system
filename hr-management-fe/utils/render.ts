import { NIcon } from 'naive-ui'
import { RouterLink } from 'vue-router'

export function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) })
}

export function renderRoute(path: string, label: string) {
    return {
        label: () =>
            h(
                RouterLink,
                {
                    to: path,
                },
                { default: () => label || 'opsy... 😛' },
            ),
        key: path,
        text: label,
    }
}
