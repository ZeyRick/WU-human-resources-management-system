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
            <n-avatar round size="small" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
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
</template>
<script setup lang="ts">
import { CaretDown, CaretUp, LogOutOutline, Pencil, PersonCircleOutline } from '@vicons/ionicons5'
import { useRoute } from 'vue-router'
import { apiLogout } from '~/apis/user'
import { Routes } from '~/constants/routes'
import { useDarkThemeStore } from '~/store/theme'
import { useUserInfoStore } from '~/store/userInfo'

const route = useRoute()
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
const loading = ref<boolean>(false)

const closelogoutModal = () => {
    showLogoutModal.value = false
}
const openLogoutModal = () => {
    showLogoutModal.value = true
}

watch(
    () => route.path,
    // @ts-ignore
    () => (routeObject.value = findRoute()),
)

const options = [
    // {
    //     label: 'Profile',
    //     key: 'profile',
    //     icon: renderIcon(PersonCircleOutline),
    // },
    // {
    //     label: 'Edit Profile',
    //     key: 'editProfile',
    //     icon: renderIcon(Pencil),
    // },
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
