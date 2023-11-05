<template>
  
    <n-layout has-sider style="height: 100%; display: flex">
      <SideBar />
      <div style="flex-grow: 1; display: flex; flex-direction: column;">
        <n-layout-header bordered style="height: 10%">
          Header Header Header
          <n-menu mode="horizontal" />
        </n-layout-header>
        <div style="flex-grow: 1">
          <slot />
        </div>
      </div>
    </n-layout>
  
</template>

<script lang="ts"></script>

<style>
html,
body,
#__nuxt,
#__layout {
  height: 100% !important;
  width: 100% !important;
  overflow: hidden !important;
}
</style>
