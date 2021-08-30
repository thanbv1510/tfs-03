<template>
  <div id="app">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <img :src="require('/src/assets/images/active-states.jpg')" alt="this is a image">
    <img src="/lec-10/src/assets/images/active-states.jpg" alt="this is a image">
    <OrderSummary></OrderSummary>
  </div>
</template>

<script>
import OrderSummary from './components/OrderSummary'

export default {
  name: 'App',
  components: {
    OrderSummary
  }
}
</script>
