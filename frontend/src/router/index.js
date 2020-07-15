import Vue from 'vue'
import VueRouter from 'vue-router'
import axios from 'axios'

import { SpinnerPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// My components
import constants from '../components/constants'
import Contents from '../components/Contents.vue'

// My lazy loaded components (to improve load times)
const ViewPage = () => import('../components/ViewPage.vue')
const EditPage = () => import('../components/EditPage.vue')

// Use plugins
Vue.use(VueRouter)
Vue.use(SpinnerPlugin)
Vue.use(constants)
Vue.prototype.$http = axios

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Contents
  },
  {
    path: '/view/:title',
    name: 'View',
    component: ViewPage
  },
  {
    path: '/edit/:title',
    name: 'Edit',
    component: EditPage
  }
]

const router = new VueRouter({
  routes
})

export default router
