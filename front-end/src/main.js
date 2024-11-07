import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

// 添加axios请求拦截器
axios.interceptors.request.use(
    (config)=>{
        config.url=`${config.url}`;
        return config;
    },
    (error)=>{
        return Promise.reject(error);
    }
)

createApp(App).use(router).mount('#app')
