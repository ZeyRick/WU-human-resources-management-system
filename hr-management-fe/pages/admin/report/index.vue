<template>
  <<n-data-table :columns="Columns" :data=""/>
</template>

<script setup lang="ts">
import type { RowData } from 'naive-ui/es/data-table/src/interface';
import {ReportTableColumns} from './report-table'
import { NLayout, NCard, NText, type DataTableColumns } from 'naive-ui'

const Columns: DataTableColumns<RowData> = [
  {
    title: 'Helo', // 
    key: 'heloData'
  }, {
    title: 'GG',
    key: 'gg'
  }
]

type Employee {
  name: string,
  age: number,
  pasw: string
}
const Data: Employee[] = [
  {
    name: "123",
    age: 2
  },
  {
    name: "123",
    age: 2
  },
  {
    name: "123",
    age: 2
  },
  {
    name: "123",
    age: 2
  }
]

definePageMeta({
    layout:"main"
})
</script>