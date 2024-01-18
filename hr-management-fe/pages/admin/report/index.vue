<template>
  <n-layout>
    <n-card>
      <n-date-picker v-model:value="selectedDate" type="date" placeholder="Select a date" />
      <n-select v-model:value="selectedPeriod" :options="periodOptions" placeholder="Select a period" />
    </n-card>
    <n-data-table :columns="Columns" :data="dataSeed" :bordered="false"/>
  </n-layout>
</template>

<script setup lang="ts">
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { NLayout, NCard, NText, type DataTableColumns } from 'naive-ui'

const Columns: DataTableColumns<RowData> = [
  {
    title: 'ID', 
    key: 'id'
  }, 
  {
    title: 'Name', 
    key: 'name'
  }, 
  {
    title: 'Position',
    key: 'position'
  },
  {
    title: 'Salary',
    key: 'salary'
  },
  {
    title: 'Start Date',
    key: 'start_date'
  },
  {
    title: 'Status',
    key: 'status'
  },
  {
    title: 'Working Hour',
    key: 'workingHour'
  },
  {
    title: 'Attendance',
    key: 'attendance'
  }
]

type Employee = {
  id: number,
  name: string,
  position: string,
  salary: number,
  start_date: string,
  status: string,
  workingHour: number,
  attendance : number
}

const generateDataSeed = (): Employee[] => {
  const dataSeed: Employee[] = [];

  for (let i = 0; i < 10; i++) {
    const employee: Employee = {
      id: i + 1,
      name: `Employee${i + 1}`,
      position: `Position${i + 1}`,
      salary: Math.floor(Math.random() * 50000) + 50000, // Random salary between 50000 and 100000
      start_date: "2022-01-01", // You may want to generate a realistic start date
      status: i % 2 === 0 ? "Active" : "Inactive",
      workingHour: Math.floor(Math.random() * 40) + 20, // Random working hours between 20 and 60
      attendance: Math.floor(Math.random() * 100) // Random attendance percentage
    };

    dataSeed.push(employee);
  }

  return dataSeed;
};
const selectedDate = ref(null);
const selectedPeriod = ref(null);
const periodOptions = [
  { label: 'Week', value: 'week' },
  { label: 'Month', value: 'month' }
];

const dataSeed = generateDataSeed();

definePageMeta({
    layout:"main"
})
</script>