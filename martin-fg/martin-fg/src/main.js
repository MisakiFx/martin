import Vue from 'vue'
import App from './App'
import store from './store'

// 引入路由对象
import router from './router/index'

import "@/common/css/style.css"

import {Actionsheet} from 'mint-ui'
Vue.component(Actionsheet.name, Actionsheet)



new Vue({
  el:'#app',
  router,
  store,
  render: h=>h(App)
})
