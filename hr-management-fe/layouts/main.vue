<template>
  <n-config-provider class="fullScreen" :theme="isDarkTheme ? darkTheme : undefined" :theme-overrides="isDarkTheme ? darkThemeOverrides : undefined">
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
import { useDarkThemeStore } from '~/store/theme';

const isDarkTheme = computed(() => useDarkThemeStore().isDarkTheme)

const darkThemeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#63E2B7',
    bodyColor: '#303030',
    iconColor: '63E2B7',
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
