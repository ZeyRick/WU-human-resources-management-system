<template>
    <n-layout-header
        bordered
        style="padding: 10px; height: 60px; display: flex; justify-content: space-between; align-items: center"
    >
        <n-breadcrumb :clickable="false">
            <n-breadcrumb-item v-if="routeObject" style="font-size: 20px">
                <div style="display: flex; align-items: center; gap: 4px">
                    <!-- <n-icon :component="routeObject && routeObject.icon" /> -->
                    {{ i18n.global.t(routeObject.text as string) }}
                </div>
            </n-breadcrumb-item>
        </n-breadcrumb>
        <div style="display: flex; align-items: center; gap: 4px">
            <!-- <n-switch
        @update:value="switchTheme"
        v-model:value="isDark"
        style="margin-right: 8px"
        size="medium"
      >
        <template #icon> 🤔 </template>
      </n-switch> -->
            <n-avatar v-if="profileFileName" round size="small" :src="profileFileName" />
            <n-avatar v-else round :style="avatarFallbackStyle">
                {{ userInfo?.username.charAt(0).toUpperCase() }}
            </n-avatar>
            <n-dropdown :options="options">
                <n-button :onmouseover="onProfileHover" :onmouseout="onProfileOut" :bordered="false">
                    {{ userInfo?.name }} / {{ userInfo?.username }}
                    <n-icon style="margin-left: 10px"> <CaretUp v-if="isProfileHover" /> <CaretDown v-else /> </n-icon>
                </n-button>
            </n-dropdown>
        </div>
    </n-layout-header>

    <n-modal
        v-model:show="showLogoutModal"
        preset="dialog"
        title="Logout"
        content="Do you want to logout?"
        positive-text="Yes"
        negative-text="No"
        @negative-click="closelogoutModal"
        @positive-click="apiLogout"
    />
    <n-modal v-model:show="showChangeProfileModal" transform-origin="center">
        <n-card
            style="width: 300px; display: flex; align-items: center"
            title="Change Profile"
            :bordered="false"
            role="dialog"
            aria-modal="true"
        >
            <n-avatar
                v-if="profileFileName"
                round
                :size="100"
                style="margin-bottom: 30px"
                :src="profileFileName"
            />
            <n-avatar
                v-else
                round
                :size="100"
                :style="{ ...avatarFallbackStyle, 'margin-bottom': '30px', 'font-size': '30px' }"
            >
                {{ userInfo?.username.charAt(0).toUpperCase() }}
            </n-avatar>
            <n-upload
                :show-file-list="false"
                :loading="loading"
                :action="`${$config.public.apiURL}/admin/user/uploadFiles`"
                :headers="{
                    Authorization: token ? `Bearer ${token}` : '',
                }"
                with-credentials
                :on-finish="(val: any) => onFinishUploadFile(val)"
                :on-error="onErrorUploadFile"
                :on-before-upload="() => (loading = true)"
            >
                <n-button :loading="loading">Upload File</n-button>
            </n-upload>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { CaretDown, CaretUp, LogOutOutline, Pencil } from '@vicons/ionicons5'
import { createDiscreteApi, darkTheme, NAvatar, type UploadFileInfo } from 'naive-ui'
import { useRoute } from 'vue-router'
import { apiLogout } from '~/apis/user'
import { Routes } from '~/constants/routes'
import { useAuthStore } from '~/store/auth'
import { useDarkThemeStore } from '~/store/theme'
import { useUserInfoStore } from '~/store/userInfo'

const route = useRoute()
const config = useRuntimeConfig()
const themeStore = useDarkThemeStore()

// const isDark = computed(() => themeStore.isDarkTheme);

const findRoute = () => {
    for (let i = 0; i < Routes.length; i++) {
        const routeChildren = Routes[i].children
        if (routeChildren && routeChildren.length > 0) {
            for (let j = 0; j < routeChildren.length; j++) {
                if (routeChildren[j].key === route.path) return routeChildren[j]
            }
        } else {
            if (Routes[i].key === route.path) return Routes[i]
        }
    }
}

const routeObject = ref(findRoute())
const { userInfo } = useUserInfoStore()
const isProfileHover: Ref<boolean> = ref(false)
const showLogoutModal = ref<boolean>(false)
const showChangeProfileModal = ref<boolean>(false)
const loading = ref<boolean>(false)
const { token } = useAuthStore()
const profileFileName = ref<string | null>(
    userInfo?.profilePic ? `${config.public.apiURL}/public/user/${userInfo?.profilePic}` : null,
)
const colorStyle = generateColorAndText(userInfo?.username || 'R')
const avatarFallbackStyle = {
    color: colorStyle.textColor,
    backgroundColor: colorStyle.backgroundColor,
}

const closelogoutModal = () => {
    showLogoutModal.value = false
}
const openLogoutModal = () => {
    showLogoutModal.value = true
}

const onFinishUploadFile = (options: { file: UploadFileInfo; event?: any }) => {
    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })

    const res: { code: number; msg: string } = JSON.parse(options.event?.target.response)

    loading.value = false
    profileFileName.value = ''
    profileFileName.value = `${config.public.apiURL}/public/user/${res.msg}`
    message.success(`Success`)
}

const onErrorUploadFile = (options: { file: UploadFileInfo; event?: any }) => {
    const { message } = createDiscreteApi(['message'], {
        configProviderProps: {
            theme: darkTheme,
        },
    })

    const res: { code: number; msg: string } = JSON.parse(options.event?.target.response)
    if (res?.code === -1) {
        message.error(`Upload file failed: ${res.msg || 'Someting Went Wrong'}`)
    }
    loading.value = false
}

watch(
    () => route.path,
    // @ts-ignore
    () => (routeObject.value = findRoute()),
)

const renderFallbackAvatar = () => {
    return h(
        NAvatar,
        {
            style: `{
      color: 'yellow',
      backgroundColor: 'red',
    }`,
        },
        [userInfo?.username.charAt(0).toLocaleUpperCase()],
    )
}

const options = [
    // {
    //     label: 'Profile',
    //     key: 'profile',
    //     icon: renderIcon(PersonCircleOutline),
    // },
    {
        label: 'Edit Profile',
        key: 'editProfile',
        icon: renderIcon(Pencil),
        props: {
            onClick: () => (showChangeProfileModal.value = true),
        },
    },
    {
        label: 'Logout',
        key: 'logout',
        icon: renderIcon(LogOutOutline),
        props: {
            onClick: () => openLogoutModal(),
        },
    },
]

const switchTheme = (value: boolean) => themeStore.setDarkTheme(value)

const onProfileHover = () => (isProfileHover.value = true)

const onProfileOut = () => (isProfileHover.value = false)
</script>
