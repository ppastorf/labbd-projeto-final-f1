import Vue from 'vue'
import Router from 'vue-router'
import Meta from 'vue-meta'

import Constructor from './views/constructor'
import Admin from './views/admin'
import Driver from './views/driver'
import Login from './views/login'
import './style.css'

Vue.use(Router)
Vue.use(Meta)
export default new Router({
  mode: 'history',
  routes: [
    {
      name: 'LoginPage',
      path: '/',
      component: Login,
    },
    {
      name: 'AdminPage',
      path: '/admin',
      component: Admin,
    },
    {
      name: 'ConstructorPage',
      path: '/constructor',
      component: Constructor,
    },
    {
      name: 'DriverPage',
      path: '/driver',
      component: Driver,
    },
  ]
})
