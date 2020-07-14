<template>
  <div id="app">
    <router-link id="heading" to="/">Scott's Wiki</router-link>
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
#heading {
  font-size: 2em;
  font-weight: bolder;
  color: #404090;
}
#heading:hover {
  text-decoration: none;
}
#view {
  margin-top: 20px;
}
#title {
  font-size: 1em;
  font-weight: bolder;
  color: #404090;
  text-decoration: none;
  margin-bottom: 12px;
}
#load {
  margin-left: 20px;
}
.error {
  color: red;
}
.left {
  float: left;
}
#buttons {
  font-size: 14px;
  font-weight: bolder;
  color: #409040;
  margin-bottom: 12px;
}
.button {
  padding: 0 1px 0 1px;
  text-decoration: none;
}
.button:hover {
  text-decoration: underline;
  cursor: pointer;
}
.button-gap {
  margin-left: 10px;
}
</style>
