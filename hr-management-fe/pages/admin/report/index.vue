<template>
  <n-layout>
    <n-card>
      <div style="display: flex; align-items: center; gap: 10px;">
        <n-date-picker v-model:value="selectedDateRange" type="daterange" placeholder="Select a date" />
        <div style="font-size: 16px; display: flex; align-items: center;white-space: nowrap;">
          Department:
          <n-select
              @update:value="onDepartmentChange"
              style="margin-left: 10px; "
              v-model:value="filterForm.departmentId"
              filterable
              size="small"
              :placeholder="i18n.global.t('department')"
              :options="[{ label: 'All', value: '' }, ...departmentOptions]"
          />
        </div>
        <div style="font-size: 16px; display: flex; align-items: center; white-space: nowrap;">
          Employee Name:
          <n-input
              style="margin-left: 10px;"
              v-model:value="searchTerm"
              size="small"
              placeholder="Search by name"
          />
        </div>
      </div>
      <n-data-table :columns="Columns" :data="filteredData" :bordered="false"/>
    </n-card>
  </n-layout>
</template>


<script setup lang="ts">
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { NLayout,NInput,NSelect, NCard, NText, type DataTableColumns } from 'naive-ui'
import type { Employee, EmployeeParams, CreateEmployeeType } from '~/types/employee'
import type { Department } from '~/types/department'
import { apiAllDepartment } from '~/apis/department'
import './index.css';
const searchTerm = ref('')
const filterForm = reactive<EmployeeParams>({
    employeeName: '',
    departmentId: '',
})
const onDepartmentChange = (value: any) => {
    filterForm.departmentId = value
}
const departmentOptions = ref([
  { label: 'Department 1', value: '1' },
  { label: 'Department 2', value: '2' },
  // Add more departments as needed
]);
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

type EmployeeSeed = {
  id: number,
  name: string,
  position: string,
  salary: number,
  start_date: string,
  status: string,
  workingHour: number,
  attendance : number
}
function getRandomName() {
  const names = ['John', 'Jane', 'Mary', 'James', 'Emma', 'Noah', 'Olivia', 'Liam', 'Ava', 'Isabella']
  const randomIndex = Math.floor(Math.random() * names.length)
  return names[randomIndex]
}

// Helper function to generate a random date in the format YYYY-MM-DD
function getRandomDate() {
  const start = new Date(2010, 0, 1).getTime()
  const end = new Date().getTime()
  const randomDate = new Date(start + Math.random() * (end - start))
  return randomDate.toISOString().split('T')[0] // Get the date part of the ISO string
}
const generateDataSeed = (): EmployeeSeed[] => {
  const dataSeed: EmployeeSeed[] = [];

  for (let i = 0; i < 20; i++) {
    const employee: EmployeeSeed = {
      id: i + 1,
      name: getRandomName(),
      position: `Position${i + 1}`,
      salary: Math.floor(Math.random() * 50000) + 50000, // Random salary between 50000 and 100000
      start_date: getRandomDate(), // You may want to generate a realistic start date
      status: i % 2 === 0 ? "Active" : "Inactive",
      workingHour: Math.floor(Math.random() * 40) + 20, // Random working hours between 20 and 60
      attendance: Math.floor(Math.random() * 100) // Random attendance percentage
    };

    dataSeed.push(employee);
  }

  return dataSeed;
};
const selectedDateRange = ref<Date[] | null>(null);
const selectedPeriod = ref(null);


const dataSeed = generateDataSeed();

const filteredData = computed(() => {
  let result = dataSeed

  // Filter by date range
  if (Array.isArray(selectedDateRange.value) && selectedDateRange.value.length > 0)  {
    const [startDate, endDate] = selectedDateRange.value.map(date => new Date(date))
    result = result.filter(employee => {
      const employeeDate = new Date(employee.start_date)
      return (
        employeeDate >= startDate &&
        employeeDate <= endDate
      )
    })
  }

  // Filter by search term
  if (searchTerm.value) {
    result = result.filter(employee =>
      employee.name.toLowerCase().includes(searchTerm.value.toLowerCase())
    )
  }

  return result
})

definePageMeta({
    layout:"main"
})
</script>