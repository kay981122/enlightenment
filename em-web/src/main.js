import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElIcons from '@element-plus/icons'
import axios from 'axios'
import router from './router/index'
const app = createApp(App)
for(const name in ElIcons) {
    app.component(name,ElIcons[name])
}
app.use(ElementPlus).use(router)
app.config.globalProperties.$http = axios
app.mount('#app')

