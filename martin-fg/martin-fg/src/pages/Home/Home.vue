<template>
    <div class="home">
      <!--广告位-->
      <div class="hot-ad">
        <img src="../../../static/img/banner.jpeg" alt="" width="100%">
      </div>
      <div class="introduce">
        <img src="../../../static/img/introduce.png" alt="" width="100%" @click="$router.push('/introduce')">
      </div>
      <!--商品列表-->
      <home-shop-list />
    </div>
</template>

<script>
    import HomeShopList from "./HomeShopList";
    export default {
      name: "Home",
      components:{
        HomeShopList
      },
      mounted() {
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .home
    background: #f5f5f5
    .hot-ad
      background-color: #fff
    .introduce
      text-align center
      background-color: #fff
</style>
