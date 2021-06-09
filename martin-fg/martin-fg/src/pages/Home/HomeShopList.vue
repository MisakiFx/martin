<template>
  <div class="expense-container" v-if="shopList.length > 0">
    <ul class="check-list">
      <li class="check-list-item" v-for="(shop, index) in shopList" :key="index">
        <img :src="shop.image_url" alt="" width="100%">
        <h4 class="item-time-label">{{shop.title}}</h4>
        <div class="list-item-bottom">
          <span class="item-money">¥{{shop.price}}</span>
          <span class="item-buy">
            <button @click="clickBuy(index)">购买</button>
          </span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
    import {mapState} from 'vuex'
    import { MessageBox, Toast, Indicator} from 'mint-ui';
    import {postBuyExamination} from '../../api'
    export default {
      name: "HomeShopList",
      computed:{
        ...mapState(['shopList'])
      },
      methods: {
          async clickBuy(index = 0) {
            MessageBox.confirm("是否确认购买?").then(async action=>{
              Indicator.open()
              let result = await postBuyExamination({'examination_id':index + 1},{'open_id' : this.$store.state.userInfo.open_id})
              Indicator.close()
              if (result.code !== 0) {
                Toast(result.msg)
                return
              }
              Toast("购买成功")
            })
          }
        }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .expense-container
    padding-bottom 40px
    background-color: #f5f5f5
    .check-list
      .check-list-item
        width 100%
        margin-bottom 10px
        background-color: #fff
        display flex
        flex-direction column
        .item-time-label
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
          .item-money
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
