import ApiService from "./api.service";

const ArticlesService = {
    query(type, params) {
        return ApiService.query("articles" + (type === "feed" ? "/feed" : ""), {
            params: params,
        });
    },
    get(slug) {
        return ApiService.get("articles", slug);
    },
    create(params) {
        return ApiService.post("articles", {article: params});
    },
    update(slug, params) {
        return ApiService.update("articles", slug, {article: params});
    },
    destroy(slug) {
        return ApiService.delete(`articles/${slug}`);
    },

    favorite(slug) {
        return ApiService.post(`articles/${slug}/favorite`)
    },
    unFavorite(slug) {
        console.log(`articles/${slug}/favorite`)
        return ApiService.delete(`articles/${slug}/favorite`)
    }
};

export default ArticlesService;
