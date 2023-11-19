<template>
  <n-layout
    style="flex-grow: 1; display: flex; flex-direction: column"
  >
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
        <n-text type="primary" style="font-size: 18px">User Table</n-text>
        <n-button strong style="background-color: #63E2B7" color="#5cb85c" text-color="#000000">
          <template #icon>
            <n-icon color="#000000">
              <AddCircleOutline />
            </n-icon>
          </template>
         
             Create
         
        </n-button>
      </div></n-card
    >
    <n-data-table :columns="tableColumns" :data="data" />
    <n-card
      content-style="padding: 10px;"
      style="display: flex; align-items: center; height: 50px; overflow: hidden"
    >
      <n-pagination :page="1" :page-count="100" />
    </n-card>
  </n-layout>
</template>

<script setup lang="ts">
import { AddCircleOutline } from "@vicons/ionicons5";
import { tableColumns } from "./table-columns";

const data = Array.from({ length: 10 }).map((_, index) => ({
  name: `Edward King ${index}`,
  age: 32,
  address: `London, Park Lane no. ${index}`,
}));

definePageMeta({
  layout: "main",
});
</script>
