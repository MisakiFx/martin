<template>
 <div class="refund">
  <div class="refund-head">
    <h1 class="money">¥123</h1>
  </div>
  <div class="refund-bottom">
    <mt-field label="退款金额" placeholder="Input number" type="number" v-model="refundMoney"></mt-field>
  </div>
  <div class="refund-button">
    <button>退款</button>
  </div>
 </div>
</template>

<script>
    export default {
      name: "Refund",
      data() {
        return {
          refundMoney : 0
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .refund
    height 100%
    weight 100%
    .refund-head
      padding-top 40px
      width 100%
      height 100px
      background-color: #18bc18
      text-align center
      font-size 16px
      border-radius 10px
      .money
        color white
        padding-top 10px
        font-size 30px
        text-align center
        font-weight border
    .refund-bottom
      margin-top 10px
      height 40px
    .refund-button
      button
        width 90%
        height 40px
        line-height 40px
        background-color: #e9232c
        text-align center
        margin 60px 5%
        border none
        font-size 16px
        color #fff
        border-radius 10px
</style>
