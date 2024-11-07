<template>
    <div class="aboutBox">
        <div class="top">
            <div class="logo"></div>
        </div>
        <div class="information">这是一个人员简单的团队制作的对github技术人员进行技术评级的小产品。评级仅代表个人团队排序，无任何其他用途。</div>
        <div class="bottom">
            <div class="worker" v-for="item in state.workerData">
                <a :href="item.htmlurl"></a>
                <div class="avatar"><img :src="item.avatar_url"></div>
                <div class="name">{{ item.login }}</div>
            </div>
        </div>
    </div>
</template>

<script setup>
import axios from 'axios';
import { onBeforeMount, reactive } from 'vue';

    const state=reactive({
        workerData:[
            {
                login:'Vegetable_center',
                htmlurl:'',
                avatar_url:''
            },
            {
                login:'Vegetable-center',
                htmlurl:'',
                avatar_url:''
            },
            {
                login:'Vegetable-center',
                htmlurl:'',
                avatar_url:''
            }
        ]
    })

    // onBeforeMount(async()=>{
    //     await axios({
    //         method:'post',
    //         url:'/v1/github/about_us',
    //     })
    //     .then(res =>{
    //         const {data}=res.data;
    //         state.workerData=data;
    //     })
    //     .catch(error =>{
    //         console.log(error);
    //     })
    // })
</script>

<style scoped>
    .aboutBox {
        width: 100%;
        height: 90%;
        display: flex;
        flex-direction: column;
        align-items: center;
        font-family: '宋体';
        color:#848484;
        overflow: hidden;
    }
    .top {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;
        width: 100%;
        height: 50%;
    }
    .logo {
        width: 60%;
        height: 90%;
        background-image: url(../assets/about-logo.png);
        background-position: center;
        background-repeat: no-repeat;
        background-size: contain;
    }
    .information {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 18%;
        font-size: 22px;
        opacity: .8;
        color: white;
    }
    .bottom {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        width: 100%;
        height: 30%;
    }
    .worker {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;
        width: 25%;
        height: 100%;
    }
    .avatar {
        width: 200px;
        height: 200px;
        border-radius: 50%;
        background-color: aquamarine;
    }
    .name {
        font-size: 18px;
        color: white;
    }
</style>
