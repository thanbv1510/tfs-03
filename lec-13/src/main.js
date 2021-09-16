import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import ApiService from "./common/api.service";
import {CHECK_AUTHENTICATE} from "./store/actions.type";

Vue.config.productionTip = false;

ApiService.init();

router.beforeEach(async (to, from, next) => {
    await store.dispatch(CHECK_AUTHENTICATE);
    next();
});

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app");
