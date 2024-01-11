<template>
    <n-layout>
      <n-card>
        <n-form v-model="employee" label-placement="left">
          <n-form-item label="ID">
            <n-input v-model="employee.id" disabled/>
          </n-form-item>
          <n-form-item label="Name">
            <n-input v-model="employee.name"/>
          </n-form-item>
          <n-form-item label="Position">
            <n-input v-model="employee.position"/>
          </n-form-item>
          <n-form-item label="Salary">
            <n-input-number v-model="employee.salary"/>
          </n-form-item>
          <n-form-item label="Start Date">
            <n-date-picker v-model="employee.start_date" type="date"/>
          </n-form-item>
          <n-form-item label="Status">
            <n-input v-model="employee.status"/>
          </n-form-item>
          <n-form-item label="Working Hour">
            <n-input-number v-model="employee.workingHour"/>
          </n-form-item>
          <n-form-item label="Attendance">
            <n-input-number v-model="employee.attendance"/>
          </n-form-item>
          <n-button type="primary" @click="updateEmployee">Update</n-button>
        </n-form>
      </n-card>
    </n-layout>
  </template>
  
  <script setup lang="ts">
  import { NLayout, NCard, NForm, NFormItem, NInput, NInputNumber, NDatePicker, NButton } from 'naive-ui'
  import { ref } from 'vue'
  
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
  
  const employee = ref<Employee>({
    id: 1,
    name: 'John Doe',
    position: 'Developer',
    salary: 50000,
    start_date: '2022-01-01',
    status: 'Active',
    workingHour: 40,
    attendance: 90
  })
  
  const updateEmployee = () => {
    // Here you can implement the logic to update the employee's information
    console.log('Updated employee:', employee.value)
  }
  
  definePageMeta({
      layout:"main"
  })
  </script>
