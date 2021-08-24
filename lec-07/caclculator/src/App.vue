<template>
  <main>
    <h1> Calculator </h1>
    <Calc/>
  </main>
</template>

<script>
import Calc from "./components/Calc.vue";

export default {
  name: "App",
  components: {Calc},
}
</script>

<style scoped>
  h1 {
    color: red;
    text-align: center;
  }
</style>
