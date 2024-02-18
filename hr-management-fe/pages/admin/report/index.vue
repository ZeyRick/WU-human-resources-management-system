<template>
  <n-layout>
    <n-card>
      <div style="display: flex; align-items: center; gap: 10px;">
        <n-date-picker
    v-model:value="selectedDateRange"
    type="daterange"
    placeholder="Select a date"
    @update:value = "getReportByDate"
  />
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
        <n-button @click="convertToUTC">
    Click Me
  </n-button>
      </div>
      <n-data-table :columns="columns" :loading="loading" :data="reportType" :bordered="false"/>
    </n-card>
  </n-layout>
</template>


<script setup lang="ts">
import type { RowData } from 'naive-ui/es/data-table/src/interface'
import { NLayout,NInput,NSelect, NCard, NText, type DataTableColumns } from 'naive-ui'
import type { Employee, EmployeeParams, CreateEmployeeType } from '~/types/employee'
import type { Department } from '~/types/department'
import { apiAllDepartment } from '~/apis/department'
import { apiGetReport ,apiGetReportByDate, apiGetReportByDepartment,apiGetReportByEmployee} from '~/apis/report'
import './index.css';
import moment from 'moment'
import {reportTableColumns} from './report-table-columns'
import type {report,reportParamFilterByDate,reprtParamFilterByEmpolyee} from '~/types/report'
const searchTerm = ref('')

import { DATE_TIME_FORMAT } from '~/constants/time'
const loading = ref<boolean>(true)
  const pageOption = ref<Pagination>({ page: 1, size: 10 })
const departmentOptions = ref<{ label: string; value: string }[]>([])
const filterForm = reactive<EmployeeParams>({
    employeeName: '',
    departmentId: '',
})
const reportType = ref<report[]>([])

const columns: DataTableColumns<RowData> = [
    ...reportTableColumns,
    
]
const onDepartmentChange = (value: any) => {
    filterForm.departmentId = value
}
const getDepartment = async () => {
    try {
        loading.value = true
        const res: any = await apiAllDepartment()
        const departments = res as Department[]
        departmentOptions.value
        departments.map((e) => {
            departmentOptions.value.push({
                label: `${e.id} - ${e.alias}`,
                value: e.id,
            })
        })
    } catch (error) {
    } finally {
        loading.value = false
    }
}
const selectedDateRange = ref<Date[]>([new Date(), new Date()])
const startDate = selectedDateRange.value[0];
const endDate = selectedDateRange.value[1];
const filterFormDate =reactive<reportParamFilterByDate> ({
    startDate: startDate.toISOString(),
    endDate: endDate.toISOString()
})
const getReportByDate = async () => {
    try {
        reportType.value = []
        loading.value = true
        const res: any = await apiGetReportByDate(pageOption.value,filterFormDate)
        const report = res as report[]
        report.map((e) => {
            reportType.value.push({
                employeeId: e.employeeId,
                employeeName: e.employeeName,
                departmentAlias: e.departmentAlias,
                totalWorkMinute: e.totalWorkMinute,
                attandance: e.attandance
            })
            console.log()
        })
    } catch (error) {
    } finally {
        loading.value = false
    }
  
}

const getReport= async () => {
    try {
        loading.value = true
        const res: any = await apiGetReport()
        const report = res as report[]
        report.map((e) => {
            reportType.value.push({
                employeeId: e.employeeId,
                employeeName: e.employeeName,
                departmentAlias: e.departmentAlias,
                totalWorkMinute: e.totalWorkMinute,
                attandance: e.attandance
            })
            console.log()
        })
    } catch (error) {
    } finally {
        loading.value = false
    }
  
}



const convertToUTC = (value: Date[]) => {
  const startDate = selectedDateRange.value[0];
const endDate = selectedDateRange.value[1];
  console .log("Start Date in +7: " + startDate.toISOString());
  console .log("Start Date in +0: " + startDate.toUTCString());
  const startDate2 = moment(startDate.toISOString(), 'YYYY-MM-DDTHH:mm:ss.sssZ').utc().format(DATE_TIME_FORMAT)
  console .log("Start Date in UTC: " + startDate2);

}



const  handleDateRangeChange =(newDateRange : any)=>{
  const [startDateString, endDateString] = newDateRange.split();
    const startDate = moment(startDateString);
    const endDate = moment(endDateString);

    // Convert to UTC or Phnom Penh local time based on your requirements
    // (adjust time zone offset and formatting accordingly)
    const startDateInUtc = startDate.utc();
    const endDateInUtc = endDate.utc();
    console.log
}
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
    title: 'Department',
    key: 'department'
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

// Helper function to generate a random date in the format YYYY-MM-DD
function getRandomDate() {
  const start = new Date(2010, 0, 1).getTime()
  const end = new Date().getTime()
  const randomDate = new Date(start + Math.random() * (end - start))
  return randomDate.toISOString().split('T')[0] // Get the date part of the ISO string
}






onMounted(async () => {
    await getDepartment()
    await getReport()
    
}),
console.log(reportType.value)

definePageMeta({
    layout:"main"
})
</script>