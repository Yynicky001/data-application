<template>
    <div class="talentBox">
        <div class="header">
            <div class="headerItem" v-for="item in state.data.headerName">{{ item }}</div>
        </div>
        <div class="list">
            <div class="listItem" v-for="(item,index) in state.data.listData">
                <div class="item top">{{ index+1 }}</div>
                <div class="item user">
                    <div class="info"></div>
                    <div class="name">{{ item.userName }}</div>
                </div>
                <div class="item area">{{ item.area }}</div>
                <div class="item nation">{{ item.nation }}</div>
                <div class="item organization">{{ item.organization }}</div>
                <div class="item address">{{ item.address }}</div>
                <div class="item talent">{{ item.talent }}</div>
            </div>
        </div>
    </div>
</template>

<script setup>    
    const state=defineProps(['data']);
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
