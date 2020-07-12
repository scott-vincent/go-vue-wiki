<template>
  <div id="app">
    <router-link id="title" to="/">Scott's Wiki</router-link>
    <router-view id="view"/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  padding: 30px;
}
#title {
  font-size: 2em;
  font-weight: bolder;
  color: #404090;
  text-decoration: none;
}
#title:hover {
  text-decoration: underline;
}
#view {
  margin-top: 30px;
}
</style>
