import { createRouter, createWebHistory } from "vue-router";
import axios from "axios";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: "login",
      component: () => import("./views/login.vue"),
      meta: { title: "登录账号", layout: "login" },
    },
    {
      path: "/findpw",
      name: "findpw",
      component: () => import("./views/findpw.vue"),
      meta: { title: "找回密码", layout: "findpw" },
    },
    {
      path: "/install",
      name: "install",
      component: () => import("./views/Install.vue"),
      meta: { title: "软件安装", layout: "install" },
    },
    {
      path: "/model/list",
      name: "modelList",
      component: () => import("./views/Model/Model_list.vue"),
      meta: {
        title: "模型列表",
        layout: "layout",
        model: "model",
        action: "list",
      },
    },
    {
      path: "/field/list",
      name: "fieldList",
      component: () => import("./views/Field/Field_list.vue"),
      meta: {
        title: "模型字段管理",
        layout: "layout",
        model: "field",
        action: "list",
      },
    },
    {
      path: "/house/list",
      name: "houseList",
      component: () => import("./views/House/House_list.vue"),
      meta: {
        title: "房屋管理",
        layout: "layout",
        model: "house",
        action: "list",
      },
    },
    {
      path: "/room/list",
      name: "roomList",
      component: () => import("./views/Room/Room_list.vue"),
      meta: {
        title: "房间管理",
        layout: "layout",
        model: "room",
        action: "list",
      },
    },
    {
      path: "/room/action",
      name: "roomAction",
      component: () => import("./views/Room/Room_action.vue"),
      meta: {
        title: "房间信息管理",
        layout: "layout",
        model: "room",
        action: "action",
      },
    },
    {
      path: "/room/fee",
      name: "roomFee",
      component: () => import("./views/Room/Room_fee.vue"),
      meta: {
        title: "房间租金管理",
        layout: "layout",
        model: "room",
        action: "fee",
      },
    },
    {
      path: "/field/action",
      name: "fieldAction",
      component: () => import("./views/Field/Field_action.vue"),
      meta: {
        title: "模型字段新增/编辑",
        layout: "layout",
        model: "field",
        action: "action",
      },
    },
    {
      path: "/:model/list",
      name: "contentList",
      component: () => import("./views/Content/Content_list.vue"),
      meta: { title: "模型列表", layout: "layout" },
    },
    {
      path: "/:model/action",
      name: "contentAction",
      component: () => import("./views/Content/Content_action.vue"),
      meta: { title: "模型新增/编辑", layout: "layout" },
    },
    {
      path: "/404",
      name: "404notfound",
      component: () => import("./views/404.vue"),
      meta: { title: "404页面丢失了", layout: "notfound" },
    },
  ],
});

router.beforeEach((to, from, next) => {
  NProgress.start();
  var token = localStorage.getItem("token");
  if (to.name === "install") {
    next();
  } else if (to.name === "findpw") {
    next();
  } else if (!token && to.name !== "login") {
    next({ name: "login" });
  } else if (to.path === "/") {
    next("/house/list");
  } else {
    next();
  }
});

router.afterEach((to) => {
  NProgress.done();
});

export default router;
