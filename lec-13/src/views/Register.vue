<template>
  <div class="auth-page">
    <div class="container page">
      <div class="row">
        <div class="col-md-6 offset-md-3 col-xs-12">
          <h1 class="text-xs-center">Sign up</h1>
          <p class="text-xs-center">
            <router-link :to="{ name: 'login' }">Have an account?</router-link>
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
                  placeholder="Your Name"
                  v-model="username"
              />
            </fieldset>
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
                @click.prevent="register"
            >
              Sign up
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import {REGISTER} from "../store/actions.type";

export default {
  data() {
    return {
      username: "",
      email: "",
      password: "",
    };
  },
  computed: {
    ...mapGetters(["errors"]),
  },
  methods: {
    async register() {
      const isSuccess = await this.$store.dispatch(REGISTER, {
        username: this.username,
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
