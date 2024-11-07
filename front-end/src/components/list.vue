<template>
    <div class="talentBox">
        <div class="header">
            <div class="headerItem" v-for="item in props.data.headerName">{{ item }}</div>
        </div>
        <div class="list">
            <div class="listItem" v-for="(item,index) in props.data.listData" @click="jumpTalenterInfo">
                <div class="item top">{{ index+1 }}</div>
                <div class="item user">
                    <!-- <div class="info"><img :src="item.avatar_url"></img></div> -->
                    <div class="info"><img src="https://tse3-mm.cn.bing.net/th/id/OIP-C.yp-D-KHI3e2nN4eMBJcEVAAAAA?w=203&h=203&c=7&r=0&o=5&dpr=1.5&pid=1.7"></img></div>
                    <div class="name">{{ item.login }}</div>
                </div>
                <div class="item area">{{ item.domain }}</div>
                <div class="item nation">{{ item.nation }}</div>
                <!-- <div class="item organization">{{ item.organization }}</div> -->
                <div class="item address">{{ item.html_url }}</div>
                <div class="item talent">{{ item.talent }}</div>
            </div>
        </div>
    </div>
</template>

<script setup>    
    const props=defineProps(['data']);
    const emit=defineEmits();
    const jumpTalenterInfo=()=>{
        emit('handoff','获取数据相关字段')
    }
</script>

<style scoped>
    .talentBox {
        width: 100%;
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
        overflow: hidden;
    }
    .info img {
        width: 100%;
        height: 100%;
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
        cursor: pointer;
    }
    .item {
        min-width: 170px;
        max-width: 170px;
        overflow-wrap: break-word;
        text-align: center;
    }
</style>
