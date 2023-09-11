import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { backgroundRouters, displayRouters } from "@/router/routers";
import PageNotFound from "@/components/PageNotFound/PageNotFound.vue";

Vue.use(VueRouter);

const routers: RouteConfig[] = [
  ...backgroundRouters,
  ...displayRouters,
  { path: "*", component: PageNotFound },
];

const router = new VueRouter({
  routes: routers,
});

export default router;
