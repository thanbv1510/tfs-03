import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
import API_URL from "./config";
import JwtService from "./jwt.service";

const ApiService = {
    init() {
        Vue.use(VueAxios, axios);
        Vue.axios.defaults.baseURL = API_URL;
    },

    setHeader() {
        Vue.axios.defaults.headers.common[
            "Authorization"
            ] = `Token ${JwtService.getToken()}`;
    },

    query(resource, params) {
        return Vue.axios.get(resource, params);
    },

    getAll(resource) {
        return Vue.axios.all(resource);
    },

    get(resource, slug = "") {
        return Vue.axios.get(`${resource}/${slug}`);
    },

    post(resource, params) {
        return Vue.axios.post(`${resource}`, params);
    },

    update(resource, slug, params) {
        return Vue.axios.put(`${resource}/${slug}`, params);
    },

    put(resource, params) {
        return Vue.axios.put(`${resource}`, params);
    },

    delete(resource) {
        return Vue.axios.delete(resource);
    },
};

export default ApiService;
