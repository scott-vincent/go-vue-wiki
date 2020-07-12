<template>
  <div id="app">
    <h1>Scott's Wiki</h1>
    <p>
      <router-link to="/">Home</router-link>
      <router-link to="/blah">Go to Blah</router-link>
    </p>
    <router-view></router-view>
  </div>
</template>

<script>
import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

import Home from './components/Home.vue'

const router = new VueRouter({
  routes: [
    { path: '/', component: Home },
    { path: '/blah', component: Home }
  ]
})

const app = new Vue({ router }).$mount('#app')
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
