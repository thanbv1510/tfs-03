import ApiService from "./api.service";

const AuthService = {
    login(credentials) {
        return ApiService.post("users/login", {user: credentials});
    },
    register(credentials) {
        return ApiService.post("users", {user: credentials});
    },
    me() {
        return ApiService.get("user");
    },
    updateUser(user) {
        return ApiService.put("user", user);
    },
};

export default AuthService;
