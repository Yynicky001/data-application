<template>
  <div class="box">
    <div class="left" ref="leftNav">
      <div class="hiddenBt" @click="hidden" ref="hiddenBt"><</div>
      <div class="logo"></div>
      <div class="btList" ref="btList">
        <HomeBt v-for="(item,index) in state.btData" :btData="item" :class="{active:index==0}" @click="containChange(index)"></HomeBt>
      </div>
      <div class="picture"></div>
      <div class="setter"><i class="iconfont icon-shezhi"></i>设置</div>
    </div>
    <div class="right">
      <div class="contain">
        <div class="navBox">
          <div class="title">{{ state.title }}</div>
          <div class="input">
            <input type="text">           
            <i class="iconfont icon-sousuo_2"></i>
          </div>
        </div>
        <HomeGeneral :component="state.thisCom"></HomeGeneral>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import HomeBt from '../components/home-botton.vue';
import HomeGeneral from '../components/general.vue';
import HomeTalent from '../components/home-talent.vue';
import HomeArea from '../components/home-area.vue';

    const state=reactive({
        btData:[
            {
              text:'技术排行',
              iconfont:'icon-gerenpaihang'
            },
            {
              text:'领域热榜',
              iconfont:'icon-rebang'
            },
            // {
            //   text:'开发者国籍',
            //   iconfont:'icon-bangdan1'
            // },
            {
              text:'数据详情',
              iconfont:'icon-shuju1'
            },
            {
              text:'关于我们',
              iconfont:'icon-guanyuwomen'
            }
        ],
        components:[
          HomeTalent,
          HomeArea
        ],
        thisCom:HomeTalent,
        title:'技术排行'
    });

    const btList=ref(null);
    // 导航栏跳转的功能函数
    const containChange=(index)=>{
      const activeEle=btList.value.querySelector('.active');
      if(activeEle){
        activeEle.classList.remove('active');
        btList.value.children[index].classList.add('active');
      }
      state.thisCom=state.components[index];
      state.title=state.btData[index].text;
    }

    // 左侧导航栏点击隐藏的功能函数
    const leftNav=ref(null);
    const hiddenBt=ref(null);
    const isHidden=ref(false);
    
    const hidden=()=>{
      const setter=leftNav.value.querySelector('.setter');
      if(!isHidden.value){
        leftNav.value.style.width='0px';
        hiddenBt.value.textContent='>';
        btList.value.style.opacity='0'
        setter.style.opacity='0';
      }
      else {
        leftNav.value.style.width='300px';
        hiddenBt.value.textContent='<';
        btList.value.style.opacity='1'
        setter.style.opacity='1';
      }
      isHidden.value=!isHidden.value;
    }

</script>

<style scoped>
  *{
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }
  .box {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100vw;
    height: 100vh;
    background-color: #000000;
  }
  .left {
    position: relative;
    transition: all .5s;
    width: 300px;
    height: 95%;
    margin-left: -5px;
    border-radius: 10px;
    background-color: #131514;
  }
  .active {
    background-color: #2b2b2b;
  }
  /* .left:hover .hiddenBt {
    opacity: 1;
  }  */
  .hiddenBt {
    transition: all .5s;
    position: absolute;
    top: 50%;
    right: 0;
    transform: translate(50%,-50%);
    opacity: 0;
    width: 30px;
    height: 80px;
    background-color: #797979;
    line-height: 80px;
    font-size: 32px;
    color: white;
    z-index: 4;
    cursor: pointer;
  }
  .hiddenBt:hover{
    opacity: 1;
  }
  .logo {
    width: 85%;
    height: 180px;
    margin-top: 10px;
    margin-bottom: 50px;
    margin-left: auto;
    margin-right: auto;
    background-image: url('../assets/home-logo.png');
    background-position: center;
    background-size: contain;
    background-repeat: no-repeat;
  }
  .setter {
    position: absolute;
    bottom: 10px;
    left: 20px;
    color: white;
    font-size: 18px;
    cursor: pointer;
    z-index: 5;
  }
  .setter i {
    margin-right: 5px;
  }
  .right {
    position: relative;
    flex: 1;
    height: 100%;
  }
  .navBox {
    position: relative;
    display: flex;
    align-items: center;
    width: 100%;
    height: 10%;
    padding-left: 30px;
  }
  .title {
    text-align: center;
    font-size: 24px;
    font-family: '宋体';
    font-weight: 700;
    color: white;
  }
  .input {
    display: flex;
    align-items: center;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%,-50%);
    border-radius: 20px;
    border: 1px solid #ffffff;
    overflow: hidden;
  }
  .input input {
    width: 300px;
    height: 30px;
    padding-left: 10px;
    outline: none;
    border: none;
  }
  .icon-sousuo_2 {
    padding-left: 3px;
    padding-right: 5px;
    font-size: 28px;
    color: #ffffff;
  }
  .contain {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    width: 95%;
    height: 95%;
    border-radius: 20px;
    background-color: #131514;
  }
</style>
