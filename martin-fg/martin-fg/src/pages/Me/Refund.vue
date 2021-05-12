<template>
 <div class="refund">
  <div class="refund-head">
    <h1 class="money">¥{{userExamination.remainder}}</h1>
  </div>
  <div class="refund-bottom">
    <mt-field label="退款金额" placeholder="请输入退款金额" type="number" v-model=money :attr="{ maxlength: 10 }" @blur.native.capture="checkInputName" :state=state></mt-field>
  </div>
  <div class="refund-button">
    <button @click="RefundMoney">退款</button>
  </div>
 </div>
</template>

<script>
    import {mapState} from 'vuex'
    import {Toast} from 'mint-ui'
    import {refundMoney} from "../../api";

    export default {
      name: "Refund",
      computed : {
        ...mapState(['userExamination'])
      },
      mounted() {
        this.$store.dispatch('reqUserExamination', this.$store.state.userInfo.open_id)
      },
      methods : {
        checkInputName() {
          let regex = /^[0-9]+(.[0-9]{2})?$/
          this.state = regex.test(this.money);
          console.log(this.state);
        },
        async RefundMoney() {
          if (this.money <= 0) {
            Toast("退款金额不能小于等于0")
            return
          }
          if (!this.state) {
            Toast('退款金额支持到小数点后两位')
            return
          }
          if (this.money > this.$store.state.userExamination.remainder) {
            Toast("余额不足")
            return
          }
          let result = await refundMoney({'money':this.money}, this.$store.state.userInfo.open_id)
          if (result.code !== 0) {
            Toast(result.msg)
            return
          }
          Toast('退款成功')
          setTimeout(()=>{
            this.$router.replace('/me')
          }, 1000)
        }
      },
      data() {
        return {
          money : '',
          state : false,
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
