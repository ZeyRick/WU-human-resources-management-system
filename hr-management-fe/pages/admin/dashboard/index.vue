<template>
    <div class="dashboard">
        <n-card id="header" style="min-height: 240px" shadow="always">
            <n-space align="center">
                <img style="width: 100px" src="~/assets/western_logo.png" alt="" />
                <p style="font-size: 25px">
                    Hello <span style="color: cyan">{{ userInfo.name }}</span
                    >. Welcome To
                </p>
            </n-space>
            <div class="school-info">
                <p style="font-size: 25px">
                    <span style="color: cyan">Western University</span> Human Resource Management System
                </p>
            </div>
        </n-card>
        <n-card id="body" class="content" shadow="always">
            <n-title style="font-size: 25px">Introduction</n-title>
            <div style="display: flex; justify-content: center; align-items: center; flex-direction:column; font-size: 20px; gap: 20px;">
                <div>Project Name</div>
                <div style="color: aqua;">Human Resource Management System</div>
                <div>Thesis Advisor</div>
                <div style="color: aqua;">Professer Yeun Sophearith</div>
                <div>Team Member</div>
                <div style="display: flex; justify-content: space-between; min-width: 300px; align-items: center;">
                    <div style="color: aqua;">Hea PengVeng</div>
                    <div style="color: aqua;">San Rachy</div>
                </div>
                <div style="display: flex; justify-content: space-between; min-width: 300px; align-items: center;">
                    <div style="color: aqua;">Sam Chanbot</div>
                    <div style="color: aqua;">Chhour Seang</div>
                </div>
            </div>


            <!-- <n-list class="data-section" item-border v-if="employeeCount !== undefined">
                <n-list-item>
                    <n-text>Total Employees:</n-text>
                    <n-progress :percent="calculateProgress(employeeCount)" type="circle" color="#007BFF" size="100px">
                    </n-progress>
                </n-list-item>
                <n-list-item>
                    <n-text>Total Courses:</n-text>
                    <n-text>{{ courseCount }}</n-text>
                </n-list-item>
            </n-list>
            <div v-else>Loading...</div> -->
            <!-- <ClientOnly>
                <apexchart
                    :key="series"
                    height="100%"
                    width="100%"
                    :options="chartOptions"
                    :series="series"
                ></apexchart>
            </ClientOnly> -->
        </n-card>
    </div>
</template>

<script setup>
import { defineProps } from 'vue'
import { useUserInfoStore } from '~/store/userInfo'

const { userInfo } = useUserInfoStore()

// const chartOptions = {
//     chart: {
//         type: 'area',
//         stacked: false,
//         height: 350,
//         zoom: {
//             type: 'x',
//             enabled: true,
//             autoScaleYaxis: true,
//         },
//         toolbar: {
//             autoSelected: 'zoom',
//         },
//     },
//     dataLabels: {
//         enabled: false,
//     },
//     markers: {
//         size: 0,
//     },

//     fill: {
//         type: 'gradient',
//         gradient: {
//             shadeIntensity: 1,
//             inverseColors: false,
//             opacityFrom: 0.5,
//             opacityTo: 0,
//             stops: [0, 90, 100],
//         },
//     },
//     yaxis: {
//         labels: {
//             formatter: function (val) {
//                 return (val / 1000000).toFixed(0)
//             },
//         },
//         title: {
//             text: 'Late',
//         },
//         // show: false
//     },
//     xaxis: {
//         type: 'datetime',
//     },
//     tooltip: {
//         shared: false,
//         y: {
//             formatter: function (val) {
//                 return (val / 1000000).toFixed(0)
//             },
//         },
//     },
// }

// const series = [
//     {
//         colors:['#FF0000', '#FF0000', '#FF0000'],
//         name: 'XYZ MOTORS',
//         data: [
//             {
//                 x: new Date(),
//                 y: 20
//             }
//         ],
//     },
// ]
// const series = [
//     {
//         name: 'TEAM 1',
//         data: generateDayWiseTimeSeries(new Date('11 Feb 2017 GMT').getTime(), 20, {
//             min: 10,
//             max: 60,
//         }),
//     },
//     {
//         name: 'TEAM 2',
//         data: generateDayWiseTimeSeries(new Date('11 Feb 2017 GMT').getTime(), 20, {
//             min: 10,
//             max: 60,
//         }),
//     },
//     {
//         name: 'TEAM 3',
//         data: generateDayWiseTimeSeries(new Date('11 Feb 2017 GMT').getTime(), 30, {
//             min: 10,
//             max: 60,
//         }),
//     },
//     {
//         name: 'TEAM 4',
//         data: generateDayWiseTimeSeries(new Date('11 Feb 2017 GMT').getTime(), 10, {
//             min: 10,
//             max: 60,
//         }),
//     },
//     {
//         name: 'TEAM 5',
//         data: generateDayWiseTimeSeries(new Date('11 Feb 2017 GMT').getTime(), 30, {
//             min: 10,
//             max: 60,
//         }),
//     },
// ]

const props = defineProps({
    schoolName: {
        type: String,
        default: 'Western University',
    },
    systemName: {
        type: String,
        default: 'Human Resource Management System',
    },
    employeeCount: {
        type: Number,
        default: 150,
    },
    courseCount: {
        type: Number,
        default: 10,
    },
})

const calculateProgress = (value) => {
    return (value / 300) * 100 // Assuming max value is 300 for demonstration
}

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
</style>
