import { RouteConfig } from "vue-router";

export const backgroundRouters: RouteConfig[] = [
  {
    path: "/",
    redirect: "/config",
  },
];

export const displayRouters: RouteConfig[] = [
  {
    path: "/config",
    name: "config",
    meta: { display: "镜像回源配置端" },
    component: () => import("@/config/Config.vue"),
  },
];
