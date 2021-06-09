// 1. 引入
import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../pages/Home/Home'
import Me from '../pages/Me/Me'
import Login from '../pages/Login/Login'
import MeDetail from "../pages/Me/MeDetail";
import MeExpense from "../pages/Me/MeExpense";
import CheckBooking from "../pages/Check/CheckBooking";
import CheckList from "../pages/Check/CheckList";
import CheckResult from "../pages/Check/CheckResult";
import Refund from "../pages/Me/Refund";
import AdminTop from "../pages/Admin/AdminTop";
import CheckStart from "../pages/Admin/CheckStart";
import CheckFinish from "../pages/Admin/CheckFinish";
import AdminCheckResult from "../pages/Admin/CheckResult";
import Introduction from "../pages/Introduce/Introduction";

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

// 2. 声明使用
Vue.use(VueRouter);


// 3. 输入路由对象
export default new VueRouter({
  // 3.1 配置一级路由
  mode:'history',
  routes: [
    {
      path: '/home',
      component: Home,
      meta: {showBottomTabBar: true}
    },
    {path: '/me', component: Me, meta: {showBottomTabBar: true}},
    {path: '/login', component: Login},
    {path: '/detail', component: MeDetail},
    {path: '/expense', component: MeExpense},
    {path: '/check_booking', component: CheckBooking},
    {path: '/checks', component: CheckList},
    {path: '/check_result/:id', component: CheckResult},
    {path: '/refund', component: Refund},
    {path: '/admin/top', component: AdminTop},
    {path: '/admin/check_start', component: CheckStart},
    {path: '/admin/check_finish', component: CheckFinish},
    {path: '/admin/check_result', component: AdminCheckResult},
    {path: '/introduce', component: Introduction},
    {path: '/', redirect: '/home'}
  ]
});
