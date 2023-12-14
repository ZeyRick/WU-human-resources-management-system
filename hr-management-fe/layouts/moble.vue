<template>
  <n-config-provider class="fullScreen">
    <n-layout has-sider style="height: 100%; display: flex">
      <SideBar />
      <div style="flex-grow: 1; display: flex; flex-direction: column;">
          <slot style="flex-grow: 1;" />
      </div>
    </n-layout>
    </n-config-provider>
</template>

<script setup lang="ts">
</script>

<style>
.fullScreen,
html,
body,
#__nuxt,
#__layout {
  height: 100% !important;
  width: 100% !important;
  max-width:  480px !important;
  overflow: hidden !important;
}
</style>
