<template>
  <div>
    <div id="load">
      <div v-if="loading">
        <b-spinner small variant="secondary"></b-spinner>
      </div>
      <div v-if="error" class="error">
        {{ error }}
      </div>
    </div>
    <div v-if="pages">
      <div id="title">
        <div class="left">Contents [</div>
        <div id="add-page" class="left button clickable" v-on:click="addPage">+</div>
        ]
      </div>
      <div v-for="page in pages" :key="page">
        <div id="view-page" class="wiki-link clickable" v-on:click="viewPage(page)">{{ page }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
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
      this.$http.get('/pages')
        .then(response => {
          this.pages = response.data
        })
        .catch(error => { this.error = error })
        .finally(() => { this.loading = false })
    },
    addPage (event) {
      alert('Add page')
    },
    viewPage (title) {
      alert('View page ' + title)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#load {
  margin-left: 20px;
}
.error {
  color: red;
}
#title {
  font-size: 1em;
  font-weight: bolder;
  color: #404090;
  text-decoration: none;
  margin-bottom: 20px;
}
.left {
  float: left;
}
.button {
  padding: 0 3px 0 3px;
}
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
