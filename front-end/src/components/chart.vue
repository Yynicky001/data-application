<template>
    <div class="chartContent" ref="chart"></div>
</template>

<script setup>
    import axios from 'axios';
import *as echarts from 'echarts';
    import { onMounted, ref, defineProps, onBeforeUnmount, onBeforeMount, } from 'vue'

    const props=defineProps(['option']);
    const chart=ref(null);
    

    onBeforeMount(async()=>{
        await axios({
            method:'get',
            url:'https://github-profile-summary-cards.vercel.app/api/cards/profile-details?username=0125nia&theme=zenburn',
        })
        .then(res=>{
            console.log(res)
        })
        .catch(error =>{
            console.log(error)
        })
    })

    // 图表的配置项
    const options=props.option;

    onMounted(()=>{
        const chartInstance=echarts.init(chart.value);
        chartInstance.setOption(options);

        const observer=new ResizeObserver(()=>{
            chartChange(chartInstance);
        })

        const chartChange=(chart)=>{
            chart.resize();
        }
        observer.observe(chart.value);
    })

</script>

<style scoped>
    .chartContent {
        width: 100%;
        height: 100%;
        padding: 10px;
    }
</style>
