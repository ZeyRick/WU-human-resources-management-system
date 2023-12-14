<template>
  <n-layout-sider
    bordered
    show-trigger="bar"
    collapse-mode="width"
    :collapsed-width="64"
    :width="240"
    :native-scrollbar="false"
    :on-update:collapsed="collapseUpdated"
  >
    <div style="height: 53px; margin: 0px; padding: 0px">
      <NImage
        preview-disabled
        :src="
          !isCollapsed
            ? 'https://flyclipart.com/thumb2/lol-logo-920003.png'
            : 'https://flyclipart.com/thumb2/ljl-spring-919994.png'
        "
      />
    </div>
    <n-menu
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :options="menuOptions"
    />
  </n-layout-sider>
</template>

<script setup lang="ts">
import type { NImage } from "naive-ui";
import { menuOptions } from "./items";

const isCollapsed = ref(false);

const collapseUpdated = (collalsed: boolean) => (isCollapsed.value = collalsed);
</script>
