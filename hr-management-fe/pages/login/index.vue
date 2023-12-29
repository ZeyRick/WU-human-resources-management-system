<template>
  <div class="flex items-center justify-center h-screen">
    <n-space justify="center" align="center" vertical inline>
      <n-message-provider>
        <content />
      </n-message-provider>
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
      <n-button strong secondary round type ="success"  @click="HandleClick" >
        Success
      </n-button>
      
    </n-space>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { GlassesOutline, Glasses } from "@vicons/ionicons5";
import { NSpace, NInput, NIcon, darkTheme  } from "naive-ui";
import {useMessage} from 'naive-ui';
import { LoginFnc} from "./method";

export default defineComponent({
  setup() {
    return {
      GlassesOutline,
      Glasses,
      value: ref(null),
      darkTheme,
      LoginFnc,
      message,
      HandleClick
    };

  },
});
const message = useMessage();
            const HandleClick =() => {
                message.warning('Login Success');
            };

</script>

