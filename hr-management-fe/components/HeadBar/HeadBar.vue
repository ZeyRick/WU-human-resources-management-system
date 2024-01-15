<template>
  <n-layout-header
    bordered
    style="
      padding: 10px;
      height: 60px;
      display: flex;
      justify-content: space-between;
      align-items: center;
    "
  >
    <n-breadcrumb :clickable="false">
      <n-breadcrumb-item v-if="routeObject" style="font-size: 20px">
        <div style="display: flex; align-items: center; gap: 4px;">
          /
          <!-- <n-icon :component="routeObject && routeObject.icon" /> -->
          {{ i18n.global.t(routeObject.label) }}
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
      <n-avatar
        round
        size="small"
        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
      <n-dropdown :options="options">
        <n-button
          :onmouseover="onProfileHover"
          :onmouseout="onProfileOut"
          :bordered="false"
        >
          User Profile
          <n-icon>
            <CaretUp v-if="isProfileHover" /> <CaretDown v-else />
          </n-icon>
        </n-button>
      </n-dropdown>
    </div>
  </n-layout-header>
</template>
<script setup lang="ts">
import {
  CaretDown,
  CaretUp,
  LogOutOutline,
  Pencil,
  PersonCircleOutline,
} from "@vicons/ionicons5";
import { useRoute } from "vue-router";
import { Routes, type Route } from "~/constants/routes";
import { useDarkThemeStore } from "~/store/theme";

const route = useRoute();
const themeStore = useDarkThemeStore();
// const isDark = computed(() => themeStore.isDarkTheme);

const routeObject: Route | undefined = Routes.find(
  (r) => r.key === route.name
);

const isProfileHover: Ref<boolean> = ref(false);

const options = [
  {
    label: "Profile",
    key: "profile",
    icon: renderIcon(PersonCircleOutline),
  },
  {
    label: "Edit Profile",
    key: "editProfile",
    icon: renderIcon(Pencil),
  },
  {
    label: "Logout",
    key: "logout",
    icon: renderIcon(LogOutOutline),
  },
];

const switchTheme = (value: boolean) => themeStore.setDarkTheme(value);

const onProfileHover = () => (isProfileHover.value = true);

const onProfileOut = () => (isProfileHover.value = false);
</script>
