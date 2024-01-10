
<template>
    <h1>Report</h1>
  <n-layout style="flex-grow: 1; display: flex; flex-direction: column">
    <n-card
      content-style="padding: 10px;"
      style="height: 50px; overflow: hidden"
    >
      <div
        style="
          flex-direction: row;
          display: flex;
          align-items: center;
          justify-content: space-between;
          overflow: hidden;
        "
      >
        <n-text type="primary" style="font-size: 18px">Report</n-text>
      </div></n-card
    >
  </n-layout>
</template>

<script setup lang="ts">
definePageMeta({
    layout:"main"
})
</script>