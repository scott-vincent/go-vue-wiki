<template>
  <div>
    <div v-if="loading">
      <b-spinner small variant="secondary"></b-spinner>
    </div>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div v-if="page && !loading">
      <div id="buttons">
        <div class="left">[</div>
        <div class="left button" v-on:click="$router.push('/')">home</div>
        <div class="left">]</div>
        <div class="left button-gap">[</div>
        <div class="left button" v-on:click="$router.push('/edit/' + $route.params.title)">edit</div>
        <div class="left">]</div>
        <div class="left button-gap">[</div>
        <div class="left button" v-on:click="deletePage()">delete</div>
        ]
      </div>
      <div id="title">
        {{ $route.params.title }}
      </div>
      <div class="page-content">{{ page.Body }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ViewPage',
  data () {
    return {
      loading: false,
      error: null,
      page: null
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
      this.page = null
      this.$http.get(this.$getConst('Url') + '/pages/' + this.$route.params.title)
        .then(response => {
          this.page = response.data
        })
        .catch(error => {
          this.error = error.response ? error.response.data : error
        })
        .finally(() => {
          this.loading = false
        })
    },
    deletePage () {
      this.loading = true
      this.error = null
      this.$http.delete(this.$getConst('Url') + '/pages/' + this.$route.params.title)
        .then(response => {
          this.$router.push('/')
        })
        .catch(error => {
          this.error = error.response ? error.response.data : error
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
.page-content {
  margin-top: 10px;
  white-space: pre-wrap;
  width: 90%;
}
</style>
