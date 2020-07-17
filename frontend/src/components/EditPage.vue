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
        <div class="left button"
          v-on:click="savePage()"
          v-bind:class="{'disabled-button':(this.page.Title === '')}">
            save
        </div>
        <div class="left">]</div>
        <div class="left button-gap">[</div>
        <div class="left button" v-on:click="finished($route.params.title)">cancel</div>
        ]
      </div>
      <div id="title">Title</div>
      <input ref="titleInput" class="input" v-model.trim="pageTitle" />
      <div id="title">Content</div>
      <textarea ref="bodyInput" class="input input-body" v-model.trim="pageBody" />
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
      page: null
    }
  },
  created () {
    // Fetch the data when the view is created
    this.fetchData()
  },
  computed: {
    pageTitle: {
      get () { return this.page ? this.page.Title : '' },
      set (value) { this.page.Title = value }
    },
    pageBody: {
      get () { return this.page ? this.page.Body : '' },
      set (value) { this.page.Body = value }
    }
  },
  methods: {
    fetchData () {
      this.loading = true
      this.error = null
      this.page = null
      if (this.$route.params.title === '*') {
        // Special case: new page
        this.page = { Title: '', Body: '' }
        this.loading = false
        this.setFocus('titleInput')
      } else {
        this.$http.get(this.$getConst('Url') + '/pages/' + this.$route.params.title)
          .then(response => {
            this.page = response.data
            this.setFocus('bodyInput')
          })
          .catch((error) => {
            this.error = error.response ? error.response.data : error
          })
          .finally(() => {
            this.loading = false
          })
      }
    },
    savePage () {
      if (this.page.Title === '') {
        return
      }
      this.loading = true
      this.error = null
      this.$http.post(this.$getConst('Url') + '/pages/' + this.$route.params.title, this.page)
        .then(response => {
          this.finished(this.page.Title)
        })
        .catch(error => {
          this.error = error.response ? error.response.data : error
          this.setFocus('titleInput')
        })
        .finally(() => {
          this.loading = false
        })
    },
    setFocus (fieldRef) {
      setTimeout(() => this.$refs[fieldRef].focus(), 0)
    },
    finished (toPage) {
      // Return to where we came from
      if (this.$route.params.title === '*') {
        this.$router.push('/')
      } else {
        this.$router.push('/view/' + toPage)
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.disabled-button {
  color: #a0a0a0;
}
.disabled-button:hover {
  text-decoration: none;
  cursor: default;
}
.input {
  width: 90%;
  max-width: 800px;
}
.input-body {
  height: 400px;
}
</style>
