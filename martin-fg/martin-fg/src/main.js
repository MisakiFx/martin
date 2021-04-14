import Vue from 'vue'
import App from './App'
import store from './store'

// 引入路由对象
import router from './router/index'

new Vue({
  el:'#app',
  router,
  store,
  render: h=>h(App)
})
