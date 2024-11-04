<template>
    <div class="tab" ref="tab">
        <div v-for="(item,index) in tabData" class="tabItem" @click="tabChange(index)">{{ item }}</div>
        <div class="lineBlock"></div>
    </div>
    <List style="height: 84%;" :data="state"></List>
</template>

<script setup>
    import { reactive, ref } from 'vue';
    import List from '../components/list.vue';
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

    const tab=ref(null);
    const clickJudge=ref(0)
    const tabChange=(index)=>{
        console.log('进入函数')
        const moveLen=tab.value.children[0].clientWidth;
        const tar=tab.value.querySelector('.lineBlock');
        
        const oldVal=tar.offsetLeft;
        console.log(oldVal)
        if(index>=clickJudge.value){
            const move=moveLen*(index-clickJudge.value)+moveLen*(1/4)*(index-clickJudge.value);
            console.log(move)
            tar.style.left=`${oldVal+move}px`;
        }     
        else{
            const move=moveLen*(clickJudge.value-index)+moveLen*(1/4)*(clickJudge.value-index);
            console.log(move)
            tar.style.left=`${oldVal-move}px`;
        }
        clickJudge.value=index;
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
        left: 0px;
        bottom: 0;
        width: 20%;
        height: 5px;
        border-radius: 30%;
        background-color: #ffffff;
    }
</style>
