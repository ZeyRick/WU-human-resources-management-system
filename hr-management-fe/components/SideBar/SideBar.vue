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
                        ? '/img/logo-big.png'
                        : '/img/logo-small.png'
                "
            />
        </div>
        <n-menu icon-size="30" :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions" />
    </n-layout-sider>
</template>

<script setup lang="ts">
import type { MenuOption, NImage } from 'naive-ui'
import { ref } from 'vue'
import { getMenuOptions } from './items'

const isCollapsed = ref(false)
const menuOptions = getMenuOptions();

const collapseUpdated = (collalsed: boolean) => (isCollapsed.value = collalsed)
</script>
