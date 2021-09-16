<template>
  <div class="auth-page">
    <div class="container page">
      <div class="row">
        <div class="col-md-6 offset-md-3 col-xs-12">
          <h1 class="text-xs-center">Sign in</h1>
          <p class="text-xs-center">
            <router-link :to="{name: 'register'}">Have an account?</router-link>
            <a href="">Have an account?</a>
          </p>

          <ul class="error-messages">
            <!-- <li v-for="(v, k) in errors" :key="k">{{ k }} {{ v[0] }}</li> -->
            <li v-for="error in errors" :key="error">
              {{ error }}
            </li>
          </ul>

          <form>
            <fieldset class="form-group">
              <input
                  class="form-control form-control-lg"
                  type="text"
                  placeholder="Email"
                  v-model="email"
              />
            </fieldset>
            <fieldset class="form-group">
              <input
                  class="form-control form-control-lg"
                  type="password"
                  placeholder="Password"
                  v-model="password"
              />
            </fieldset>
            <button
                class="btn btn-lg btn-primary pull-xs-right"
                @click.prevent="login"
            >
              Sign in
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import {LOGIN} from "../store/actions.type";

export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  computed: {
    ...mapGetters(["errors"]),
  },
  methods: {
    async login() {
      const isSuccess = await this.$store.dispatch(LOGIN, {
        email: this.email,
        password: this.password,
      });
      if (isSuccess) {
        this.$router.push({name: "home"});
      }
    },
  },
};
</script>
