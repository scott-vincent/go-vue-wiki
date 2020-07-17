<template>
  <div>
    <div v-if="loading">
      <b-spinner small variant="secondary"></b-spinner>
    </div>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div v-if="pages && !loading">
      <div id="title">
        <div class="left">Contents [</div>
        <div id="add-page" class="left button clickable" v-on:click="$router.push('/edit/*')">+</div>
        ]
      </div>
      <div v-for="page in pages" :key="page">
        <div id="view-page" class="wiki-link clickable" v-on:click="$router.push('/view/' + page)">{{ page }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Contents',
  data () {
    return {
      loading: false,
      error: null,
      pages: null
    }
  },
  created () {
    // Fetch the data when the view is created
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.loading = true
      this.error = null
      this.pages = null
      this.$http.get(this.$getConst('Url') + '/pages')
        .then(response => {
          this.pages = response.data
        })
        .catch(error => {
          this.error = error
        })
        .finally(() => {
          this.loading = false
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.wiki-link {
  padding: 2px 20px;
}
.clickable {
  color: #3c4e60;
}
.clickable:hover {
  text-decoration: underline;
  cursor: pointer;
}
</style>
