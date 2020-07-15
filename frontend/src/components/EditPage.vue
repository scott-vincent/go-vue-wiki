<template>
  <div>
    <div id="buttons">
      <div class="left">[</div>
      <div class="left button" v-on:click="$router.push('/')">home</div>
      <div class="left">]</div>
      <div class="left button-gap">[</div>
      <div class="left button" v-on:click="savePage()">save</div>
      <div class="left">]</div>
      <div class="left button-gap">[</div>
      <div class="left button" v-on:click="$router.push('/view/' + $route.params.title)">cancel</div>
      ]
    </div>
    <div id="title">
      {{ title }}
    </div>
    <div v-if="loading">
      <b-spinner small variant="secondary"></b-spinner>
    </div>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div v-if="page">
      <div class="page-content">{{ page.Body }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EditPage',
  data () {
    return {
      loading: false,
      error: null,
      page: null,
      title: null
    }
  },
  created () {
    // Fetch the data when the view is created
    this.fetchData()
    this.title = this.$route.params.title
  },
  methods: {
    fetchData () {
      this.loading = true
      this.error = null
      this.page = null
      this.$http.get(this.$getConst('Url') + '/pages/' + this.$route.params.title)
        .then(response => { this.page = response.data })
        .catch(error => { this.error = error })
        .finally(() => { this.loading = false })
    },
    savePage () {
      this.loading = true
      this.error = null
      this.savedPage = this.page
      this.page = null
      this.$http.post(this.$getConst('Url') + '/pages/' + this.$route.params.title)
        .then(response => {
          this.$router.push('/view/' + this.title)
        })
        .catch(error => {
          this.error = error
        })
        .finally(() => {
          this.loading = false
          this.page = this.savedPage
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page-content {
  margin-top: 10px;
}
</style>
