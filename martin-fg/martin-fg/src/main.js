import Vue from 'vue'
import App from './App'
import store from './store'

// 引入路由对象
import router from './router/index'

import "@/common/css/style.css"

import {Actionsheet} from 'mint-ui'
Vue.component(Actionsheet.name, Actionsheet)

import { Checklist } from 'mint-ui';
Vue.component(Checklist.name, Checklist);

import { Radio } from 'mint-ui';
Vue.component(Radio.name, Radio);


import { InfiniteScroll } from 'mint-ui';
Vue.use(InfiniteScroll);

import { Spinner } from 'mint-ui';
Vue.component(Spinner.name, Spinner);

import { Field } from 'mint-ui';
Vue.component(Field.name, Field);

import VueDirectiveImagePreviewer from 'vue-directive-image-previewer'
import 'vue-directive-image-previewer/dist/assets/style.css'
Vue.use(VueDirectiveImagePreviewer)

new Vue({
  el:'#app',
  router,
  store,
  render: h=>h(App)
})
