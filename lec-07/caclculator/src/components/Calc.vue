<template>
  <div class="calc">
    <input id="result" class="data-input brown-color" type="text" disabled v-model="input"/>
    <table class="calc-table">
      <tr>
        <td><input class="button-input" type="button" :value="7"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="8"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="9"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="'+'"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" value="Del" @click="del()"/></td>
      </tr>
      <tr>
        <td><input class="button-input" type="button" :value="4"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="5"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="6"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="'-'"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" value="AC" @click="reset"/></td>
      </tr>
      <tr>
        <td><input class="button-input" type="button" :value="1"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="2"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="3"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button" :value="'*'"
                   @click="input = input + $event.target.value"/></td>
        <td><input class="button-input" type="button"/></td>
      </tr>
      <tr>
        <td><input class="button-input" type="button" :value="0"
                   @click="input = input + $event.target.value"></td>
        <td><input class="button-input" type="button" value="."/></td>
        <td><input class="button-input" type="button" :value="'/'"
                   @click="input = input + $event.target.value"></td>
        <td><input class="button-input brown-color" type="button" value="=" @click="calc()"/></td>
        <td><input class="button-input" type="button"/></td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "Calc",
  data() {
    return {
      input: '',
    }
  },
  methods: {
    reset() {
      this.input = ''
    },
    del() {
      this.input = this.input.slice(0, -1)
    },
    calc() {
      fetch(`http://localhost:8088/calc/${this.input}`
        , {
          method: 'GET',
          headers: {
            'Content-type': 'application/json'
          }
        })
        .then((response) => {
          if (!response.ok) {
            this.input = 'Unknown'
            return
          }

          return response.json()
        }).then((dataJson) => {
        console.log(dataJson)
        this.input = dataJson.value
      }).catch(function (error) {
        console.log(error)
        this.input = 'Unknown'
      })
    }
  }
}

</script>

<style scoped>
* {
  box-sizing: border-box;
}

html {
  font-size: 62.5%;
}

.calc {
  width: 300px;
  height: 500px;
  margin: 20px auto;
}

.data-input {
  width: 100%;
  height: 15%;
  text-align: right;
  font-size: 5rem;
  margin-bottom: 10px;
  color: black;
}

table {
  width: 100%;
  height: 75%;
  text-align: center;
}

.button-input {
  height: 50px;
  width: 50px;
  border-radius: 50%;
  border: 1px solid black;
  background-color: silver;
  font-size: 2rem;
}

.brown-color {
  background-color: brown;
}
</style>
