import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "home",
        component: () => import("../views/Home.vue"),
        children: [
            {
                path: "",
                name: "global-feed",
                component: () => import("../views/MainFeed.vue"),
            },
            {
                path: "/:type",
                name: "my-feed",
                component: () => import("../views/MainFeed.vue"),
            },
        ],
    },
    {
        path: "/login",
        name: "login",
        component: () => import("../views/Login.vue"),
    },
    {
        path: "/register",
        name: "register",
        component: () => import("../views/Register.vue"),
    },
    {
        path: "/:slug",
        name: "article-slug",
        component: () => import("../views/ArticleDetail.vue"),
    }
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});

export default router;
