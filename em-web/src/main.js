import * as ElIcons from '@element-plus/icons'
import axios from 'axios'
import ElementPlus, {
  ElMessage,
  ElMessageBox
} from 'element-plus'
import 'element-plus/dist/index.css'
import qs from 'qs'
import {
  createApp
} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
const app = createApp(App)
for (const name in ElIcons) {
  app.component(name, ElIcons[name])
}
//响应时间
axios.defaults.timeout = 120000;
//配置请求头
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
//配置接口地址
axios.defaults.baseURL = '/api';
//POST传参序列化(添加请求拦截器)
axios.interceptors.request.use((config) => {
  //在发送这之前做某件事
  if (config.method === 'post') {
    config.data = qs.stringify(config.data);
  }
  return config;
}), (error) => {
  ElMessageBox.alert('系统提示', 'Title', {
    confirmButtonText: '确认',
    callback: (action) => {
      ElMessage({
        type: 'info',
        message: `action: ${action}`,
      })
    },
  })
  return Promise.reject(error);
}
axios.interceptors.response.use(response => {
  // if(response.data.status !== 200) {
  //   ElMessageBox.alert('系统提示', 'Title', {
  //     confirmButtonText: '确认',
  //     callback: (action) => {
  //       ElMessage({
  //         type: 'info',
  //         message: `action: ${action}`,
  //       })
  //     },
  //   })
  //   return Promise.reject(response.data)
  // }
  return Promise.resolve(response.data)
}), error => {
  return Promise.reject(error)
}


app.use(ElementPlus).use(router).use(store)
app.config.globalProperties.$http = axios
app.config.globalProperties.$ElMessage = ElMessage
app.mount('#app')