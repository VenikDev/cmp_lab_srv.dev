import { createApp } from 'vue'
import './style.css'
import './index.css'
import App from './view/App.vue'
import router from "./router/main.js";

createApp(App)
  .use(router)
  .mount('#app')

