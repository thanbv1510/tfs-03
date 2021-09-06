<template>
  <div id="app" :style="{backgroundImage: `url${require('@/assets/images/pattern-background-desktop.svg')}`}">
    <OrderSummary/>
  </div>
</template>

<script>
import OrderSummary from './components/OrderSummary'

export default {
  name: 'App',
  components: {
    OrderSummary,
  }
}
</script>

<style scoped>
#app {
  font-size: 1.6rem;
  height: 100vh;
  width: 100vw;
}
</style>