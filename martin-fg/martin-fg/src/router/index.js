// 1. 引入
import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../pages/Home/Home'
import Me from '../pages/Me/Me'

// 2. 声明使用
Vue.use(VueRouter);


// 3. 输入路由对象
export default new VueRouter({
  // 3.1 配置一级路由
  routes: [
    {
      path: '/home',
      component: Home
      //children: [
      //  {path: 'hot', component: Hot, meta: {showBottomTabBar: true}},
      //  {path: 'box', component: Box, meta: {showBottomTabBar: true}},
      //  {path: 'dress', component: Dress, meta: {showBottomTabBar: true}},
      //  {path: 'ele', component: Ele},
      //  {path: 'food', component: Food},
      //  {path: 'general', component: General},
      //  {path: 'man', component: Man},
      //  {path: 'mbaby', component: Mbaby},
      //  {path: 'shirt', component: Shirt},
      //  {path: '/home', redirect: '/home/hot'}
      //]
    },
    {path: '/me', component: Me, meta: {showBottomTabBar: true}},
    //{path: '/login', component: Login},
    {path: '/', redirect: '/home'}
  ]
});
