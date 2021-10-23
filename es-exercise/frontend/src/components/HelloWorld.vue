<template>
  <div class="hello">
    <input type="text" placeholder="search..." v-model="keyword">
    <button @click="this.getDataES">Search</button>

    <table v-if="dataES != null" class="center">
      <tr>
        <th>Text</th>
        <th>FileName</th>
        <th>MEDShortTitle</th>
        <th>SourceCorpus</th>
        <th>Edn</th>
        <th>MS</th>
        <th>MEDData</th>
        <th>Area</th>
        <th>VP</th>
        <th>Genre</th>
        <th>Words</th>
      </tr>
      <tr v-for="data in dataES" :key="data.ID">
        <td>{{ data.Text }}</td>
        <td>{{ data.FileName }}</td>
        <td>{{ data.MEDShortTitle }}</td>
        <td>{{ data.SourceCorpus }}</td>
        <td>{{ data.Edn }}</td>
        <td>{{ data.MS }}</td>
        <td>{{ data.MEDData }}</td>
        <td>{{ data.Area }}</td>
        <td>{{ data.VP }}</td>
        <td>{{ data.Genre }}</td>
        <td>{{ data.Words }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'HelloWorld',
  data() {
    return {
      from: 0,
      size: 100000,
      keyword: '',
      dataES: null,
    }
  },
  methods: {
    getDataES() {
      console.log("Fetch data...", this.keyword)
      Vue.axios.get(`http://localhost:8088/search?keyword=${this.keyword}`).then((response) => {
        console.log(response)
        this.dataES = response.data;
      }).catch((err) => {
        console.log(err);
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

table, th, td {
  border: 1px solid black;
}

.center {
  margin-left: auto;
  margin-right: auto;
}
</style>
