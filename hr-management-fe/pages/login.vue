<template>
  <div class="flex items-center justify-center h-screen">
    <n-space justify="center" align="center" vertical inline>
      <n-input
        v-model:value="value"
        type="text"
        placeholder="Basic Input"
        size="large"
        round
        style="width: 250px"
      />
      <br />
      <n-input
        type="password"
        show-password-on="mousedown"
        placeholder="password"
        size="large"
        round
        style="width: 250px"
      />
    </n-space>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { GlassesOutline, Glasses } from "@vicons/ionicons5";
import { NSpace, NInput, NIcon, darkTheme } from "naive-ui";

export default defineComponent({
  setup() {
    return {
      GlassesOutline,
      Glasses,
      value: ref(null),
      darkTheme,
    };
  },
});
</script>
