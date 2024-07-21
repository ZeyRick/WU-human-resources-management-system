<template>
    <div class="dashboard">
        <n-card id="header" style="min-height: 240px; margin-bottom: 20px" shadow="always">
            <n-space align="center">
                <img style="width: 100px" src="~/assets/western_logo.png" alt="" />
                <p style="font-size: 25px">
                    Hello <span style="color: cyan">{{ userInfo?.name }}</span
                    >. Welcome To
                </p>
            </n-space>
            <div class="school-info">
                <p style="font-size: 25px">
                    <span style="color: cyan">Western University</span> Human Resource Management System
                </p>
            </div>
        </n-card>
        <n-space style="width: 100%; margin-bottom: 20px" justify="start" item-style="flex-grow: 1; width: 200px">
            <Card
                v-for="employeeCount in data?.employeeCounts"
                v-on:click="() => cardClick(employeeCount.employeeType)"
                class="card"
                :label="employeeCount.employeeType"
                :value="employeeCount.totalCount"
            >
                <PeopleOutline v-if="employeeCount.employeeType === EMPLOYEE_TYPE.STAFF" />
                <SchoolOutline v-else-if="employeeCount.employeeType === EMPLOYEE_TYPE.TEACHING_STAFF" />
                <BookOutline v-else-if="employeeCount.employeeType === EMPLOYEE_TYPE.LECTURE" />
            </Card>
        </n-space>

        <n-space style="width: 100%" justify="start" item-style="flex-grow: 1; width: 200px">
            <Card
                v-for="key in Object.keys(data || {}).filter((k) => k !== 'employeeCounts')"
                v-on:click="() => cardClick(key)"
                class="card"
                :value="(data as any)[key] as number"
                :label="key"
            >
                <SchoolOutline v-if="key === 'degreeCount'" />
                <ManOutline v-else-if="key === 'userCount'" />
                <GridOutline v-else-if="key === 'courseCount'" />
            </Card>
        </n-space>
        <n-card id="body" class="content" shadow="always">
            <div
                style="
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    flex-direction: column;
                    font-size: 20px;
                    gap: 20px;
                "
            ></div>
        </n-card>
    </div>
</template>

<script setup lang="ts">
import { apiDashboard } from '~/apis/report'
import { useUserInfoStore } from '~/store/userInfo'
import type { DashboardSummary } from '~/types/report'
import { EMPLOYEE_TYPE } from '~/types/employee'
import { BookOutline, GridOutline, ManOutline, PeopleOutline, SchoolOutline } from '@vicons/ionicons5'

const { userInfo } = useUserInfoStore()
const loading = ref<boolean>(false)
const data = ref<DashboardSummary>()

const cardClick = (key: string) => {
    switch (key) {
        case EMPLOYEE_TYPE.LECTURE:
            navigateTo('/admin/employee-lecture')
            break
        case EMPLOYEE_TYPE.STAFF:
            navigateTo('/admin/employee-staff')
            break
        case EMPLOYEE_TYPE.TEACHING_STAFF:
            navigateTo('/admin/employee-teaching')
            break
        case 'degreeCount':
            navigateTo('/admin/degree')
            break
        case 'courseCount':
            navigateTo('/admin/course')
            break
        case 'userCount':
            navigateTo('/admin/user')
            break

        default:
            break
    }
}

const fetchData = async () => {
    try {
        loading.value = true
        data.value = (await apiDashboard()) as DashboardSummary
    } catch (error) {}
}

onMounted(fetchData)

definePageMeta({
    layout: 'main',
})
</script>

<style scoped>
.dashboard {
    display: flex;
    flex-direction: column;
    height: 100vh;
    padding: 20px;
}

.logo-container {
    width: 100px;
    height: 100px;
}

.logo {
    max-width: 100%;
    max-height: 100%;
}

.school-info {
    text-align: right;
    margin-left: 20px;
}

.content {
    flex: 1;
    margin-top: 20px;
}

.data-section {
    padding: 20px;
    border-radius: 5px;
}

.progress-content {
    text-align: center;
}

.video-container iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

#header {
    background-image: url(../../../public/img/header.jpg);
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100% 100%;
    background-blend-mode: overlay;
}

#body {
    background-image: url(../../../assets/login-bg.jpg);
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100% 100%;
    background-blend-mode: overlay;
}

.card {
    transition: all 0.3s ease;
}
.card:hover {
    transform: scale(1.1);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
    z-index: 99;
    cursor: pointer;
    opacity: 0.7;
}
</style>
