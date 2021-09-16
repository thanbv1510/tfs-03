<template>
  <nav class="navbar navbar-light">
    <div class="container">
      <router-link
          class="nav-link"
          active-class="active"
          :to="{ name: 'home' }"
      >
        conduit
      </router-link>
      <ul class="nav navbar-nav pull-xs-right" v-if="!isAuthenticated">
        <li class="nav-item">
          <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'home' }"
          >
            Home
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'login' }"
          >
            Login
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'register' }"
          >
            Register
          </router-link>
        </li>
      </ul>
      <ul class="nav navbar-nav pull-xs-right" v-else>
        <li class="nav-item">
          <!-- Add "active" class when you're on that page" -->
          <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'home' }"
          >
            Home
          </router-link>
        <li class="nav-item" @click="logout">
          <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'home' }"
          > Logout
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import {mapState} from "vuex";
import {LOGOUT} from "../store/actions.type";

export default {
  name: "Header",
  methods: {
    async logout() {
      await this.$store.dispatch(LOGOUT)
    }
  },
  computed: {
    ...mapState({
      isAuthenticated: (state) => state.auth.isAuthenticated,
    }),
  },
};
</script>
