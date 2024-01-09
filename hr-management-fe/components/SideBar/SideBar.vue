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
        <n-menu :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions" />
    </n-layout-sider>
</template>

<script setup lang="ts">
import type { MenuOption, NImage } from 'naive-ui'
import { h, ref } from 'vue';
import { RouterLink } from 'vue-router';

const menuOptions: MenuOption[] = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: '/clock',
                },
                { default: () => '123' },
            ),
        key: '/clock',
    },
    {
        label: 'Hear the Wind Sing',
        key: 'hear-the-wind-sing1',
        href: 'https://en.wikipedia.org/wiki/Hear_the_Wind_Sing',
    },
]
const isCollapsed = ref(false)

const collapseUpdated = (collalsed: boolean) => (isCollapsed.value = collalsed)
</script>
