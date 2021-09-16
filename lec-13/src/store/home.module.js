import {FAVORITE, FETCH_ARTICLES, UN_FAVORITE} from "./actions.type";
import ArticlesService from "../common/articles.service";
import {SET_ERRORS, SET_ARTICLES, SET_ARTICLES_COUNT} from "./mutations.type";

const state = {
    articles: [],
    articlesCount: 0,
    tags: [],
    errors: null,
};

const getters = {
    getArticleBySlug: (state) => (slug) => {
        let result = state.articles.filter(item => item.slug === slug);
        return result ? result[0] : null;
    }
};

const actions = {
    async [FETCH_ARTICLES]({commit}, {type, ...payload}) {
        try {
            const {data} = await ArticlesService.query(type, payload);
            const {articles, articlesCount} = data;
            commit(SET_ARTICLES, articles);
            commit(SET_ARTICLES_COUNT, articlesCount);
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
        }
    },

    async [FAVORITE]({commit}, slug) {
        try {
            const response = await ArticlesService.favorite(slug);
            console.log("response", response.data.article.favorited)
            if (response.data.article.favorited) {
                commit(SET_ARTICLES_COUNT, response.data.article.favoritesCount)
            }
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
        }
    },

    async [UN_FAVORITE]({commit}, slug) {
        try {
            const response = await ArticlesService.unFavorite(slug);
            if (!response.favorited) {
                commit(SET_ARTICLES_COUNT, response.data.article.favoritesCount)
            }
        } catch (err) {
            commit(SET_ERRORS, err.response.data.errors);
        }
    }
};

const mutations = {
    [SET_ARTICLES](state, articles) {
        state.articles = articles;
    },
    [SET_ARTICLES_COUNT](state, articlesCount) {
        state.articlesCount = articlesCount;
    },
};

export default {
    state,
    getters,
    actions,
    mutations,
};
