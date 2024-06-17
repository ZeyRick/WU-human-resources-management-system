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
        <div v-if="isCollapsed" style="margin: 0px; padding: 10px">
            <NImage style="width: 40px;" preview-disabled src="/img/logo-big.png" />
        </div>
        <div v-else style="margin: 0px; padding-left: 20px; display: flex; align-items: center;">
            <NImage style="width: 43px;" preview-disabled src="/img/logo-big.png" />
            <p style="margin-left: 20px; font-size: 15px;">Western University</p>
        </div>
        <n-menu style="margin-top: 20px;" icon-size="30" :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions" />
    </n-layout-sider>
</template>

<script setup lang="ts">
import type { MenuOption, NImage } from 'naive-ui'
import { ref } from 'vue'
import { getMenuOptions } from './items'

const isCollapsed = ref(false)
const menuOptions = getMenuOptions()

const collapseUpdated = (collalsed: boolean) => (isCollapsed.value = collalsed)
</script>
