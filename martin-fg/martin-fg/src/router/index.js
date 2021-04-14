// 1. 引入
import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../pages/Home/Home'
import Me from '../pages/Me/Me'
import Login from '../pages/Login/Login'

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

// 2. 声明使用
Vue.use(VueRouter);


// 3. 输入路由对象
export default new VueRouter({
  // 3.1 配置一级路由
  routes: [
    {
      path: '/home',
      component: Home,
      meta: {showBottomTabBar: true}
    },
    {path: '/me', component: Me, meta: {showBottomTabBar: true}},
    {path: '/login', component: Login},
    {path: '/', redirect: '/home'}
  ]
});
