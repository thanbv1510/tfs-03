<template>
  <div>
    <div class="article-meta">
      <a href="profile.html"><img :src="article.author.image"/></a>
      <div class="info">
        <a href="" class="author">{{ article.author.username }}</a>
        <span class="date">{{ article.createdAt }}</span>
      </div>
      <button class="btn btn-outline-primary btn-sm pull-xs-right" :class="{favarited: isFavorite}"
              @click="favorite(article.slug)">
        <i class="ion-heart"></i> {{ favoriteCount }}
      </button>
    </div>
    <router-link class="preview-link" :to="{ name: 'article-slug', query: {slug: article.slug}}">
      <h1> {{ article.title }}</h1>
      <p>{{ article.description }}</p>
      <span>Read more...</span>
      <ul class="tag-list">
        <li class="tag-default tag-pill tag-outline ng-binding ng-scope"
            v-for="tag in article.tagList" :key="tag">
          {{ tag }}
        </li>
      </ul>
    </router-link>
  </div>
</template>

<script>
import {FAVORITE, UN_FAVORITE} from "../store/actions.type";

export default {
  name: "CardFeed",
  data() {
    return {
      isFavorite: this.article.favorited,
      favoriteCount: this.article.favoritesCount

    }
  },
  methods: {
    favorite(slug) {
      if (this.isFavorite) {
        // un_favorite
        this.$store.dispatch(UN_FAVORITE, slug)
        this.isFavorite = !this.favoriteCount;
        this.favoriteCount--;
        return
      }

      // favorite
      this.$store.dispatch(FAVORITE, slug);
      this.isFavorite = !this.favoriteCount;
      this.favoriteCount++;
    }
  },
  props: {
    article: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
.favarited {
  background: #5cb85c;
  color: #F8F9FC;
}
</style>