import Vue from 'vue'
import VueRouter from 'vue-router'
import axios from 'axios'

import { SpinnerPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// My components
import constants from '../components/constants'
import Home from '../components/Home.vue'

// My lazy loaded components (to improve load times)
const View = () => import('../components/View.vue')

// Use plugins
Vue.use(VueRouter)
Vue.use(SpinnerPlugin)
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
