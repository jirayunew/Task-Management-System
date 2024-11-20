import './assets/main.css'

import { createApp } from "vue";
import router from "./router";
import App from "./App.vue";
import axios from "axios";

// เพิ่ม Default Header สำหรับ JWT
const token = localStorage.getItem("token");
if (token) {
    axios.defaults.headers.common["Authorization"] = token;
}

// สร้างแอปและใช้ Router
const app = createApp(App);
app.use(router);
app.mount("#app");
