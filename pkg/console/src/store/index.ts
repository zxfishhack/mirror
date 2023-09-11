import Vue from "vue";
import Vuex from "vuex";
import User, { IUserState } from "./modules/user";
Vue.use(Vuex);

export interface IRootState {
  user: IUserState;
}

const store = new Vuex.Store<IRootState>({
  modules: {
    user: User,
  },
});
export default store;
