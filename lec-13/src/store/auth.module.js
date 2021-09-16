import AuthService from "../common/auth.service";
import ApiService from "../common/api.service";
import {saveToken, getToken, removeToken} from "../common/jwt.service";
import {CLEAR_AUTH, SET_ERRORS, SET_USER} from "./mutations.type";
import {CHECK_AUTHENTICATE, LOGIN, LOGOUT, REGISTER} from "./actions.type";

const state = {
    isAuthenticated: false,
    errors: null,
    user: {},
};

const getters = {
    errors(state) {
        if (!state.errors) {
            return [];
        }
        return Object.keys(state.errors).map((key) => {
            return `${key} ${state.errors[key].join(" ")}`;
        });
    },
};

const actions = {
    async [LOGIN]({commit}, credentials) {
        try {
            const response = await AuthService.login(credentials);
            const {user} = response.data;

            commit(SET_USER, user);
            return true;
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
            return false;
        }
    },
    async [LOGOUT]({commit}) {
        try {
            await removeToken()
            commit(CLEAR_AUTH)
            return true;
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
            return false;
        }
    },
    async [REGISTER]({commit}, credentials) {
        try {
            const response = await AuthService.register(credentials);
            const {user} = response.data;

            commit(SET_USER, user);
            return true;
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
            return false;
        }
    },
    async [CHECK_AUTHENTICATE]({commit}) {
        if (getToken()) {
            try {
                ApiService.setHeader();
                const {data} = await AuthService.me();
                const {user} = data;

                commit(SET_USER, user);
            } catch (err) {
                commit(SET_ERRORS, err.response.data.errors);
            }
        } else {
            commit(CLEAR_AUTH);
        }
    },
};

const mutations = {
    [SET_USER](state, user) {
        state.user = user;
        state.errors = null;
        state.isAuthenticated = true;
        saveToken(user.token);
    },
    [SET_ERRORS](state, errors) {
        state.errors = errors;
    },
    [CLEAR_AUTH](state) {
        state.user = {};
        state.errors = null;
        state.isAuthenticated = false;
        removeToken();
    },
};

export default {
    state,
    actions,
    mutations,
    getters,
};
