<template>
  <div>
    <div class="article-preview" v-for="article in articles" :key="article.slug">
      <CardFeed :article="article"/>
    </div>
  </div>
</template>

<script>
import {mapState} from "vuex";
import {FETCH_ARTICLES} from "../store/actions.type";
import CardFeed from "../components/CardFeed";

export default {
  name: "Feed",
  components: {CardFeed},
  data() {
    return {
      offset: 0,
      limit: 10,
      type: '',
    };
  },
  watch: {
    '$route.params.type'() {
      this.type = this.$route.params.type;
      this.$store.dispatch(FETCH_ARTICLES, {
        type: this.type,
        offset: this.offset,
        limit: this.limit
      })
    }
  },
  computed: {
    ...mapState({
      articles: (state) => state.home.articles,
      articlesCount: (state) => state.home.articlesCount,
    }),
  },
  beforeCreate() {
    this.$store.dispatch(FETCH_ARTICLES, {
      type: this.$route.params.type,
      offset: this.offset,
      limit: this.limit
    })
  }
}
</script>