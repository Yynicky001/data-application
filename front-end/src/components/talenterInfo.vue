<template>
    <div class="talenterBox">
        <div class="backBt" @click="back">返回</div>
        <div class="talenterContent">
            <div class="block information">
                <div class="infor">
                    <a href="https://juejin.cn/user/1814658334332320" target="_blank" rel="noopener noreferrer"></a>
                </div>               
                    <div class="Mes basicMes">
                        <a href="https://juejin.cn/user/1814658334332320" target="_blank" rel="noopener noreferrer" class="name">阿符666</a>
                        <div class="sat">这个人很懒，没有写下他的简介</div>
                    </div>
                    <div class="Mes talentMes">
                        <div class="text area">
                            <div class="title">技术领域:</div>
                            <div class="textContent">go</div>
                        </div>
                        <div class="text nation">
                            <div class="title">所属国家:</div>
                            <div class="textContent">德国</div>
                        </div>
                        <div class="text talent">
                            <div class="title">技术评级:</div>
                            <div class="textContent">4.7</div>
                        </div>
                    </div>   
                    <div class="Mes addressMes">
                        <div class="text">博客地址：</div>
                        <div class="address">https://juejin.cn/user/1814658334332320</div>
                        <div class="text">其他博客地址</div>
                        <div class="address">https://juejin.cn/user/1814658334332320</div>
                    </div>             
            </div>
            <div class="rightBox">
                <div class="block chart">
                    <div class="chartBox bigChart"><Chart :option="chartData.lineData"></Chart></div>
                    <div class="chartBox"><Chart :option="chartData.pieData"></Chart></div>
                    <div class="chartBox"><Chart :option="chartData.pieData"></Chart></div> 
                    <div class="chartBox">
                        <div class="title">States</div>
                        <div class="top">
                            <div class="gitLogo"></div>
                        </div>
                        <div class="bottom">
                            <div class="item">
                                <div class="text">Total Stars:</div>
                                <div class="number">5.4k</div></div>
                            <div class="item">
                                <div class="text">2024 Commits:</div>
                                <div class="number">271</div></div>
                            <div class="item">
                                <div class="text">Total PRs:</div>
                                <div class="number">1.1k</div></div>
                            <div class="item">
                                <div class="text">Total Issues:</div>
                                <div class="number">30</div></div>
                            <div class="item">
                                <div class="text">Contributed to:</div>
                                <div class="number">15</div></div>
                        </div>
                    </div>                  
                    <div class="chartBox"><Chart :option="chartData.barData"></Chart></div>
                    <div class="bottomBox"></div>
                    <div class="bottomBox"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
    import { reactive } from 'vue';
    import Chart from '../components/chart.vue'

    const chartData=reactive({
        pieData:{
            title: {
                text: '领域涉及',
                subtext: '语言',
                left: 'start',
                textStyle:{
                    color:'#ffffff',
                    fontWeight:'bold',
                },
            },
            tooltip: {
                trigger: 'item',
            },
            legend: {
                orient: 'horizontal',
                top: 'bottom',
                textStyle:{
                    color:'#ffffff',
                    fontSize:16,
                }
            },
            series: [
                {
                name: '语言',
                type: 'pie',
                radius: ['40%', '60%'],
                data: [
                    { value: 335, name: 'Go' },
                    { value: 310, name: 'Vue' },
                    { value: 234, name: 'Js' },
                    { value: 135, name: 'Java' },
                    { value: 1548, name: 'Python' },
                ],
                emphasis: {
                    itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0, 0, 0, 0.5)',
                    },
                },
                },
            ],
        },
        barData: {
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                }
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: {
                type: 'category',
                data: ['类目1', '类目2', '类目3', '类目4', '类目5'],
                axisLabel: {
                    interval: 0
                }
            },
            yAxis: {
                type: 'value'
            },
            series: [
                {
                    name: '数据1',
                    type: 'bar',
                    data: [10, 22, 35, 47, 58],
                    itemStyle: {
                        color: '#3398DB'
                    },
                    barGap: '0%', // 使条形图之间没有间隙
                    barCategoryGap: '5%' // 使条形图的宽度充满类别区间
                }
            ]
        },
        lineData:{
            title: {
                text: '折线图示例',
                left: 'center',
                top: '20'
            },
            tooltip: {
                trigger: 'axis',
            },
            xAxis: {
                type: 'category',
                data: ['2023-11', '2023-11', '2024-1', '2024-5', '2024-7', '2024-9', '2024-11'],
                axisLine: {
                    lineStyle: {
                        color: '#ffffff' // 红色 x 轴
                    }
                },
                axisTick: {
                    show: true // 隐藏 x 轴刻度线
                },
                axisLabel: {
                    color: '#ffffff' // x 轴标签的颜色
                },
                boundaryGap: false, // 让折线图的起始点与 x 轴对齐
                splitLine:{
                    show:false,
                }
            },
            yAxis: {
                type: 'value',
                position:'right',
                axisLine: {
                    show: true, // 隐藏 y 轴的线
                },
                axisTick: {
                    show: true // 隐藏 y 轴的刻度线
                },
                splitLine: {
                    show: false
                },
                axisLabel: {
                    color: '#ffffff' // y 轴标签的颜色
                }
            },
            series: [
                {
                    name: '销售额',
                    type: 'line',
                    data: [100, 932, 501, 934, 1790, 1330, 920],
                    symbol: 'none',
                    smooth: true, // 使折线平滑
                    itemStyle: {
                        color: 'green', // 红色折线
                    },
                    areaStyle: {
                        color: 'green' // 红色区域填充，透明度为 0.3
                    }
                }
            ]
        }
    })

    const emit=defineEmits();
    const back=()=>{
        emit('back','返回');
    }
</script>

<style scoped>
    *{
        box-sizing: border-box;
    }
    .talenterBox {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        height: 100%;
        color: white;
    }
    .backBt {
        display: flex;
        align-items: center;
        align-self: flex-start;
        width: 90%;
        height: 8%;
        margin-left: 2%;
        font-family: '宋体';
        font-size: 20px;
        cursor: pointer;
    }
    .talenterContent {
        display: flex;
        width: 90%;
        height: 100%;
        overflow: hidden;
    }
    .block {
        margin-bottom: 15px;
    }
    .information {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 30%;
        height: 98%;
        border-radius: 10px;
        margin-right: 20px;
        padding-top: 50px;
        background-color: rgba(0, 0, 0, .4);
        border: 1px solid rgba(255, 255, 255, .3);
        border-bottom: 1px solid #ffffff;
        font-family: '宋体';
    }
    .infor a {
        display: block;
        width: 100%;
        height: 100%;
    }
    .infor {
        width: 120px;
        height: 120px;
        border-radius: 50%;
        background-color: #ffffff;
    }
    .Mes {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 90%;
    }
    .basicMes {
        margin-top: 20px;
        margin-bottom: 40px;
    }
    .name {
        margin-bottom: 20px;
        text-decoration: none;
        font-size: 24px;
        font-weight: 700;
        color: inherit;
    }
    .sat {
        display: flex;
        align-items: center;
        font-size: 18px;
        opacity: .6;
    }
    .talentMes {
        justify-content: space-around;
        height: 20%;
        margin-bottom: 20px;
        padding-top: 20px;
        padding-bottom: 20px;
        border-top: 1px solid #ffffff;
        border-bottom: 1px solid #ffffff;
    }
    .addressMes {
        border-bottom: 1px solid #ffffff;
    }
    .address {   
        width: 90%;
        margin-top: 10px;
        margin-bottom: 10px;    
        overflow-wrap: break-word;
    }
    .text {
        display: flex;
        justify-content: space-between;
        width: 90%;
    }
    .rightBox {
        width: 80%;
        height: 100%;
        border-radius: 10px;
        background-color: rgba(0, 0, 0, .4);
        border: 1px solid rgba(255, 255, 255, .3);
        overflow-y: scroll;
    }
    .rightBox::-webkit-scrollbar {
        width: 2px;
        height: 8px;
    }
    .chart {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-evenly;
        width: 100%;
        height: 100%;
        padding-top: 5%;
        padding-bottom: 5%;
        /* background-color: azure; */
    }
    .chartBox {
        display: flex;
        flex-direction: column;
        width: 45%;
        height: 55%;
        margin-top: 30px;
        border-radius: 10px;
        background-color: #202124;
    }
    .top {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: 40%;
    }
    .gitLogo {
        width: 120px;
        height: 120px;
        border-radius: 50%;
        background-image: url(../assets/GitHub-Symbol.png);
        background-position: center;
        background-size: cover;
        background-repeat: no-repeat;
    }
    .bottom {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        width: 100%;
        height: 60%;
        padding-left: 20px;
        padding-right: 20px;
    }
    .chartBox .title {
        padding-top: 10px;
        padding-left: 10px;
        font-size: 24px;
    }
    .item {
        display: flex;
        align-items: center;
        width: 100%;
        height: 10%;
    }
    
    .bigChart {
        width: 93%;
        margin-top: 0px;
        margin-bottom: 0px;
    }
    .bottomBox {
        width: 45%;
        height: 5%;
    }

</style>
