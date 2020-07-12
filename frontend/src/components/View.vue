<template>
  <div>
    <div id="title">
      Page Title
    </div>
    <div id="buttons">
      <div class="row">[</div>
      <div class="row button clickable">delete</div>
      <div class="row">]  [</div>
      <div class="row button clickable">edit</div>
      ]
    </div>
    <div class="page">This is a page</div>
  </div>
</template>

<script>
export default {
  name: 'View'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#title {
  font-size: 2em;
  font-weight: bolder;
  color: #404090;
  margin-bottom: 20px;
}
#buttons {
  font-size: 1em;
  font-weight: bolder;
  color: #404090;
  text-decoration: none;
  margin-bottom: 20px;
}
.row {
  float: left;
}
.button {
  padding: 0 3px 0 3px;
}
.page {
  padding: 5px 20px;
}
.clickable {
  color: #3c4e60;
}
.clickable:hover {
  text-decoration: underline;
  cursor: pointer;
}
</style>
