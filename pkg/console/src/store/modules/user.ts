import {
  VuexModule,
  Module,
} from "vuex-module-decorators";
import store from "@/store";
import { RouteConfig } from "vue-router";

export interface IUserState {
  menuData: RouteConfig[];
  userInfo: any;
  spinning: boolean;
  isAuto: boolean;
}

@Module({ namespaced: true, store, name: "user" })
export default class User extends VuexModule implements IUserState {
  public userInfo: any = {};
  public isAuto: boolean = false;
  public spinning: boolean = false;
  public menuData: RouteConfig[] = [];
}
