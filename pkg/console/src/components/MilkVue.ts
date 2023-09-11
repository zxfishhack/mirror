import Vue, { DirectiveOptions } from "vue";
import * as filters from "@/filters";
import * as directives from "@/directives";
import * as Util from "@/assets/utils";
import Service from "@/service";

import AppLayout from "./Layout/common.vue";

Vue.component("app-layout", AppLayout);

Vue.prototype.$utils = Util;
Vue.prototype.service = Service;

Object.keys(directives).forEach((key) => {
  Vue.directive(key, (directives as { [key: string]: DirectiveOptions })[key]);
});
Object.keys(filters).forEach((key) => {
  Vue.filter(key, (filters as { [key: string]: Function })[key]);
});
