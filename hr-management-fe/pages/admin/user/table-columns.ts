import { NAvatar, type DataTableColumns } from 'naive-ui'
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { generateColorAndText } from '~/utils/format'

export const tableColumns: DataTableColumns<RowData> = [
    {
        title: 'Profile',
        titleAlign: 'center',
        key: 'profile',
        align: 'center',
        render: (data, index) => {
            const config = useRuntimeConfig()
            if (!data.profilePic) {
                const colorStyle = generateColorAndText(data.username)
                return h(
                    NAvatar,
                    {
                        round: true,
                        style: `color: ${colorStyle.textColor}; background-color: ${colorStyle.backgroundColor};`,
                    },
                    [data?.username.charAt(0).toUpperCase()],
                )
            }
            return h(NAvatar, {
                round: true,
                size: 'small',
                src: `${config.public.apiURL}/public/user/${data.profilePic}`,
            })
        },
    },
    {
        title: 'Username',
        key: 'username',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Name',
        key: 'name',
        titleAlign: 'center',
        align: 'center',
    },
    {
        title: 'Type',
        key: 'userLevel',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return i18n.global.t(data.userLevel) || '-'
        },
    },
    {
        title: 'Created At',
        key: 'createdAt',
        titleAlign: 'center',
        align: 'center',
        render: (data, index) => {
            return aslocalTime(data.createdAt) || '-'
        },
    },
]
