<template>
    <div class="bt">
        <i class="iconfont" :class="data.btData.iconfont"></i>
        {{ data.btData.text }}
    </div>
</template>

<script setup>
    const data=defineProps(['btData']);
</script>

<style scoped>
    .icon {
        color: #ffffff;
    }
    .bt {
        width: 250px;
        height: 70px;
        margin: 0 auto;
        margin-top: 2px;
        margin-bottom: 2px;
        border-radius: 10px;
        text-align: center;
        line-height: 70px;
        /* background-color: #FF5B13; */
        letter-spacing: 10px;
        font-weight: 700;
        font-family: '宋体';
        font-size: 24px;
        color:#ffffff;
        cursor: pointer;
    }
    .bt:hover {
        background-color: #2b2b2b;
    }
</style>
