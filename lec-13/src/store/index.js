import Vue from "vue";
import Vuex from "vuex";
import auth from "./auth.module";
import home from "./home.module";

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        auth,
        home,
    },
});
