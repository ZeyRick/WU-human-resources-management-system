<template>
  <n-config-provider class="fullScreen" :theme="isDarkTheme ? darkTheme : undefined" :theme-overrides="isDarkTheme ? darkThemeOverrides : lightThemeOverrides">
    <n-layout has-sider style="height: 100%; display: flex">
      <SideBar />
      <div style="flex-grow: 1; display: flex; flex-direction: column;">
          <HeadBar />
          <slot style="flex-grow: 1;" />
      </div>
    </n-layout>
    </n-config-provider>
</template>

<script setup lang="ts">
import { darkTheme, type GlobalThemeOverrides } from 'naive-ui';
import { PRIMARY_COLOR } from '~/constants/theme';
import { useDarkThemeStore } from '~/store/theme';

const isDarkTheme = computed(() => useDarkThemeStore().isDarkTheme)

const darkThemeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: PRIMARY_COLOR,
    bodyColor: '#303030',
    iconColor: '#2A947D',
  },
}

const lightThemeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: PRIMARY_COLOR,
    bodyColor: '#F0F2F5',
    iconColor: '#2A947D',
  }
}
</script>

<style>
.fullScreen,
html,
body,
#__nuxt,
#__layout {
  height: 100% !important;
  width: 100% !important;
  overflow: hidden !important;
}
</style>
