<template>
    <List style="height: 95%;" :data="state"></List>
</template>

<script setup>
import { reactive } from 'vue';
import List from '../components/list.vue'
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
</script>

<style scoped>
    .talentBox {
        width: 100%;
        height: 95%;
        display: flex;
        flex-direction: column;
        align-items: center;
        font-family: '宋体';
        color:#848484;
        overflow: hidden;
    }
    .header {
        width: 90%;
        height: 5%;
        display: flex;
        justify-content: space-around;
        align-items: center;
        font-size: 18px;
        font-weight: 700;
    }
    .headerItem {
        flex: 1;
        text-align: center;
    }
    .list {
        width: 90%;
        height: 90%;
        display: flex;
        flex-direction: column;
        align-items: center;
        overflow-y: scroll;
    }
    .list::-webkit-scrollbar {
        width: 2px;
        height: 8px;
    }
    .user {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .info {
        width: 50px;
        height: 50px;
        border-radius: 50%;
        margin-right: 10px;
        background-color: white;
    }
    .listItem {
        display: flex;
        justify-content: space-around;
        align-items: center;
        width: 100%;
        min-height: 80px;
        margin-top: 20px;
        margin-bottom: 20px;
        background-color: #202124;
        color: #eeeeee;
    }
    .item {
        min-width: 170px;
        max-width: 170px;
        overflow-wrap: break-word;
        text-align: center;
    }
</style>
