<template>
    <div class="home">
      <!--广告位-->
      <div class="hot-ad">
        <img src="../../../static/img/banner.jpeg" alt="" width="100%">
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
      margin-bottom 8px
</style>
