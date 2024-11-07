<template>
    <div class="tab" ref="tab">
        <div v-for="(item,index) in tabData" class="tabItem" @click="tabChange(index,item.name)">{{ item.name }}</div>
        <div class="lineBlock"></div>
    </div>
    <List style="height: 84%;" :data="state" @handoff="jumpTalenterInfo"></List>
    <TalenterInfo :class="{move:isMoved}" class="talenter" ref="talenter"  @back="backList"></TalenterInfo>
</template>

<script setup>
    import { onBeforeMount, reactive, ref } from 'vue';
    import List from '../components/list.vue';
    import TalenterInfo from './talenterInfo.vue';
    import axios from 'axios';

    const tabData=reactive(['前端','后端','AI','操作系统'])
    const state=reactive({
        headerName:[
            '',
            '名字',
            '领域',
            '所属国家',
            '所属组织',
            'github地址',
            '技术评级',
        ],
        listData:[
            {
                userInfo:'',
                userName:'ccc',
                area:'后端',
                nation:'中国云南',
                organization:'666',
                address:'https://github.com/Vegetable-center',
                talent:'4.7'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
            {
                userInfo:'',
                userName:'czz',
                area:'前端',
                nation:'中国海南',
                organization:'555',
                address:'https://rp.mockplus.cn/',
                talent:'4.8'
            },
        ]
    })

    const reqAxios=async(domain='js')=>{
        await axios({
            method:'post',
            url:'/v1/github/domain/rank',
            data:{
                domain,
                page:1,
                per_page:10,
            },
        })
        .then(res =>{
            const {data}=res.data;
            state.listData=data;
        })
        .catch(error =>{
            console.log(error);
        })
    }

    // onBeforeMount(async()=>{
    //     await axios({
    //         method:'get',
    //         url:'/v1/github/domain',
    //     })
    //     .then(res =>{
    //         const {data}=res.data;
    //         tabData=data;
    //     })
    //     .catch(error=>{
    //         console.log(error)
    //     })
    //     reqAxios();
    // })

    const tab=ref(null);
    const clickJudge=ref(0)


    // 点击请求不同的数据，将数据更新到state中，对应的List组件就会发生变化
    const tabChange=(index,name)=>{
        console.log('进入函数')
        const moveLen=tab.value.children[0].clientWidth;
        const tar=tab.value.querySelector('.lineBlock');
        const oldVal=tar.offsetLeft;

        // 控制导航栏白条滚动的方法
        if(index>=clickJudge.value){
            const move=moveLen*(index-clickJudge.value)+moveLen*(1/4)*(index-clickJudge.value);
            tar.style.left=`${oldVal+move}px`;
        }     
        else{
            const move=moveLen*(clickJudge.value-index)+moveLen*(1/4)*(clickJudge.value-index);
            tar.style.left=`${oldVal-move}px`;
        }
        clickJudge.value=index;

        // reqAxios(name);
    }

    // 技术人员详细信息组件相关数据
    const talenterData=reactive({
        isHidden:false,
        dataUrl:'',
    })
    const talenter=ref(null);
    const isMoved=ref(false);
    const jumpTalenterInfo=(data)=>{
        talenterData.isHidden=true;
        talenterData.dataUrl=data;
        console.log(talenterData.dataUrl);
        isMoved.value=true;
    }
    const backList=()=>{
        isMoved.value=false;
        talenterData.isHidden=false;
    }
</script>

<style scoped>
    .tab {
        position: relative;
        display: flex;
        justify-content: space-around;
        align-items: center;
        width: 30%;
        height: 60px;
        margin-left: 50px;
        margin-bottom: 2%;
        color: #ffffff;
    }
    .tab::after {
        position: absolute;
        left: 5%;
        bottom: 0;
        content: '';
        width: 100%;
        height: 1px;
        background-color: #ffffff;
    }
    .tabItem {
        position: relative;
        width: 80px;
        height: 60px;
        text-align: center;
        line-height: 60px;
        cursor: pointer;
    }
    .lineBlock {
        transition: all .5s;
        position: absolute;
        left: 10px;
        bottom: 0;
        width: 20%;
        height: 5px;
        border-radius: 30%;
        background-color: #ffffff;
    }
    .talenter {
        transition: all .5s;
        position: absolute;
        top: 100%;
        left: 50%;
        transform: translateX(-50%);
        width: 95%;
        height: 90%;
        opacity: 0;
        background-color: #202124;
    }
    .move {
        top:10%;
        opacity: 1;
    }
</style>
