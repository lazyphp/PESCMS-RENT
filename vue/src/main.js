import "./assets/main.less";

import { createApp } from "vue";
import App from "./App.vue";
import router from "./route";
import axios from "axios";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import zhCn from "element-plus/dist/locale/zh-cn.mjs";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "./assets/amazeui.min.css";
import "mingcute_icon/font/Mingcute.css";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

const apiUrl = "http://127.0.0.1:8080";

app.provide("apiUrl", apiUrl);

const instance = axios.create({
  baseURL: apiUrl,
});

// 请求拦截器
instance.interceptors.request.use((config) => {
  const Authorization = localStorage.getItem("token") || "";
  config.headers.Authorization = Authorization;
  return config;
});

// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response && error.response.status === 302) {
      // 清空token
      localStorage.removeItem("token");
      // 跳转到登录页面
      router.push({ path: "/login" });
      return new Promise(() => {});
    } else if (error.response && error.response.status === 404) {
      // 跳转到404页面
      router.push({ path: "/404" });
      return new Promise(() => {});
    }
    // 对响应错误做些什么
    return Promise.reject(error);
  }
);

app.config.globalProperties.$axios = instance;

app.provide("axios", instance);

app.use(router);
app.use(ElementPlus, {
  locale: zhCn,
});
app.mount("#app");
