import Vue from 'vue'
import VueRouter from 'vue-router'
import axios from 'axios'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// My components
import constants from '../components/constants'
import Home from '../components/Home.vue'
import View from '../components/View.vue'

// Use plugins
Vue.use(VueRouter)
Vue.use(BootstrapVue)
Vue.use(constants)
Vue.prototype.$http = axios

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/view',
    name: 'View',
    component: View
  }
]

const router = new VueRouter({
  routes
})

export default router
