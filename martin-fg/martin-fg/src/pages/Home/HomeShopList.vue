<template>
  <div class="shop-container" v-if="shopList.length > 0">
    <ul class="shop-list">
      <li class="shop-list-item" v-for="(shop, index) in shopList" :key="index">
        <img :src="shop.image_url" alt="" width="100%">
        <h4 class="list-item-title">{{shop.title}}</h4>
        <div class="list-item-bottom">
          <span class="item-price">¥{{shop.price}}</span>
          <span class="item-buy">
            <button>购买</button>
          </span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
    import {mapState} from 'vuex'
    export default {
        name: "HomeShopList",
        computed:{
          ...mapState(['shopList'])
        }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .shop-container
    padding-bottom 40px
    background-color: #f5f5f5
    .shop-list
      .shop-list-item
        margin-bottom 10px
        background-color: #fff
        display flex
        flex-direction column
        .list-item-title
          line-height 22px
          width 94%
          margin-left 3%
          height 22px
          overflow hidden
        .list-item-bottom
          margin 10px 0
          display flex
          flex-direction row
          justify-content space-around
          align-items center
          .item-price
            font-size 18px
            text-align center
            font-weight bolder
            color red
            flex 1
          .item-buy
            flex 5
            button
              width 80%
              height 34px
              font-size 15px
              border none
              color #fff
              display flex
              justify-content center
              align-items center
              margin-left 30px
              background-color: red
              border-radius 8px
</style>
